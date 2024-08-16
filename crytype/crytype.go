package crytype

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/luongduc1246/ultility/random"
)

type SecretLoader interface {
	SetHash(string)
	GetHash() string
	SetNonce(string)
	GetNonce() string
}

type secret struct {
	keyHash  string
	keyNonce string
}

func NewSecret(hash string, nonce string) *secret {
	return &secret{
		keyHash:  hash,
		keyNonce: nonce,
	}
}

func (s *secret) SetHash(key string) {
	s.keyHash = key
}
func (s secret) GetHash() string {
	return s.keyHash
}
func (s *secret) SetNonce(key string) {
	s.keyNonce = key
}
func (s secret) GetNonce() string {
	return s.keyNonce
}

type cryType struct {
	secret SecretLoader
}

func NewCryType(secret SecretLoader) *cryType {
	return &cryType{
		secret: secret,
	}
}

func (c cryType) CreateSecrecKeyWithTime() (secrecKey string) {
	tb := c.CreateTimeToArrayByte()
	l := 32 - len(tb)
	r := []byte(random.CreateCodeRamdomNumerals(l))
	key := append(r, tb...)
	secrecKey = string(key)
	return
}

func (c cryType) CreateTimeToArrayByte() []byte {
	t := []byte(fmt.Sprint(time.Now().UTC().Unix()))
	for i, s := range t {
		if i%2 == 0 {
			t[i] = s + 49
		} else {
			t[i] = s + 17
		}
	}
	return t
}

func (cryType) CreateSecrecKey() (secrecKey string, err error) {
	key := make([]byte, 16)
	if _, err = rand.Read(key); err != nil {
		return
	}
	secrecKey = hex.EncodeToString(key)
	return
}

// EncryHash is a hash function with sha256 and hmac
func (c cryType) EncryHash(str string) (result string, err error) {
	hash := sha256.New()
	_, err = hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf("%x", hash.Sum(nil))
	hash = hmac.New(sha256.New, []byte(c.secret.GetHash()))
	_, err = hash.Write([]byte(result))
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf("%x", hash.Sum(nil))
	return result, nil
}

func (c cryType) EnCryptGobAes(secrecKey string, obj interface{}) (result string, err error) {
	var gobResult bytes.Buffer
	end := gob.NewEncoder(&gobResult)
	err = end.Encode(obj)
	if err != nil {
		return
	}
	result, err = c.EncryAES(secrecKey, string(gobResult.Bytes()))
	return
}

func (c cryType) DecryptGobAes(secrettKey string, token string, obj interface{}) (err error) {
	de, err := c.DecryAES(secrettKey, token)
	if err != nil {
		return
	}
	byteBuffer := bytes.NewBuffer([]byte(de))
	dec := gob.NewDecoder(byteBuffer) // Will read from byteBuffer
	err = dec.Decode(obj)
	return
}

// EncryAES is encode function with AES
func (c cryType) EncryAES(secretKey string, str string) (result string, err error) {
	plaintext := []byte(str)
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	ciphertext := aesgcm.Seal(nil, []byte(c.secret.GetNonce()), plaintext, nil)
	result = fmt.Sprintf("%x", ciphertext)
	return result, err
}

// DecryAES decode function with AES
func (c cryType) DecryAES(secretKey string, str string) (result string, err error) {
	ciphertext, err := hex.DecodeString(str)
	if err != nil {
		return result, err
	}
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return result, err
	}
	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		return result, err
	}
	plaintext, err := aesgcm.Open(nil, []byte(c.secret.GetNonce()), ciphertext, nil)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("%s", plaintext)
	return result, err
}
