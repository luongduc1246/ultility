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

const HASH_SERECT_PASSWORD = "N@Tu12@lBui1d3"

const AES_NONCE = "N@Tu12@lBui1"

func CreateSecrecKeyWithTime() (secrecKey string, err error) {
	tb := CreateTimeToArrayByte()
	l := 32 - len(tb)
	r := []byte(random.CreateCodeRamdomNumerals(l))
	key := append(r, tb...)
	secrecKey = string(key)
	return
}

func CreateTimeToArrayByte() []byte {
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

func CreateSecrecKey() (secrecKey string, err error) {
	key := make([]byte, 16)
	if _, err = rand.Read(key); err != nil {
		return
	}
	secrecKey = hex.EncodeToString(key)
	return
}

// EncryHash is a hash function with sha256 and hmac
func EncryHash(str string) (result string, err error) {
	hash := sha256.New()
	_, err = hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf("%x", hash.Sum(nil))
	hash = hmac.New(sha256.New, []byte(HASH_SERECT_PASSWORD))
	_, err = hash.Write([]byte(result))
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf("%x", hash.Sum(nil))
	return result, nil
}

func EnCryptGob(secrecKey string, obj interface{}) (result string, err error) {
	var gobResult bytes.Buffer
	end := gob.NewEncoder(&gobResult)
	err = end.Encode(obj)
	if err != nil {
		return
	}
	result, err = EncryAES(secrecKey, string(gobResult.Bytes()))
	return
}

func DecryptGob(secrettKey string, token string, obj interface{}) (err error) {
	de, err := DecryAES(secrettKey, token)
	if err != nil {
		return
	}
	byteBuffer := bytes.NewBuffer([]byte(de))
	dec := gob.NewDecoder(byteBuffer) // Will read from byteBuffer
	err = dec.Decode(obj)
	return
}

// EncryAES is encode function with AES
func EncryAES(secretKey string, str string) (result string, err error) {
	plaintext := []byte(str)
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	ciphertext := aesgcm.Seal(nil, []byte(AES_NONCE), plaintext, nil)
	result = fmt.Sprintf("%x", ciphertext)
	return result, err
}

// DecryAES decode function with AES
func DecryAES(secretKey string, str string) (result string, err error) {
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
	plaintext, err := aesgcm.Open(nil, []byte(AES_NONCE), ciphertext, nil)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("%s", plaintext)
	return result, err
}
