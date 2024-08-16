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

type CryType struct {
	secret SecretLoader
}

func NewCryType(secret SecretLoader) *CryType {
	return &CryType{
		secret: secret,
	}
}

func (c CryType) CreateSecrecKeyWithTime() (secrecKey string) {
	tb := c.CreateTimeToArrayByte()
	l := 32 - len(tb)
	r := []byte(random.CreateCodeRamdomNumerals(l))
	key := append(r, tb...)
	secrecKey = string(key)
	return
}

func (c CryType) CreateTimeToArrayByte() []byte {
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

func (CryType) CreateSecrecKey() (secrecKey string, err error) {
	key := make([]byte, 16)
	if _, err = rand.Read(key); err != nil {
		return
	}
	secrecKey = hex.EncodeToString(key)
	return
}

// EncryHash is a hash function with sha256 and hmac
func (c CryType) EncryHash(str string) (result string, err error) {
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

func (c CryType) EnCryptGobAes(secrecKey string, obj interface{}) (result string, err error) {
	var gobResult bytes.Buffer
	end := gob.NewEncoder(&gobResult)
	err = end.Encode(obj)
	if err != nil {
		return
	}
	result, err = c.EncryAES(secrecKey, gobResult.String())
	return
}

func (c CryType) DecryptGobAes(secrettKey string, token string, obj interface{}) (err error) {
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
func (c CryType) EncryAES(secretKey string, str string) (result string, err error) {
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
func (c CryType) DecryAES(secretKey string, str string) (result string, err error) {
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
	result = string(plaintext)
	return result, err
}
