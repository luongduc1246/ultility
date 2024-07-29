package crytype

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	fmt.Println(EncryHash("Natural@Duc1246"))
}

func TestCreateSecrecKey(t *testing.T) {
	k, e := CreateSecrecKey()
	fmt.Println(k, e)
}

func TestCreateKeyWithTime(t *testing.T) {
	k, e := CreateSecrecKeyWithTime()
	fmt.Println(k, e)
}

func TestTime(t *testing.T) {
	b := time.Now().UTC().Unix()
	fmt.Println(len([]byte(fmt.Sprint(b))))

}

func BenchmarkCreateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSecrecKey()
	}
}
func BenchmarkCreateKeyTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSecrecKeyWithTime()
	}
}

func TestEncrypEAS(t *testing.T) {
	ser, err := CreateSecrecKey()
	token, err := EncryAES(ser, "ABC")
	fmt.Println(token, err)

}

func TestDecryptEAS(t *testing.T) {
	ser, err := CreateSecrecKey()
	token, err := EncryAES(ser, "bca")
	result, err := DecryAES(ser, token)
	fmt.Println(result, err)
}

type Claims struct {
	Name string
	Code string
}

func TestEncryGob(t *testing.T) {
	ser, err := CreateSecrecKey()
	to, err := EnCryptGob(ser, Claims{Name: "aldkjfl", Code: "ladjfla"})
	fmt.Println(to, err)
}
func TestDecruptGob(t *testing.T) {
	ser, err := CreateSecrecKey()
	to, err := EnCryptGob(ser, Claims{Name: "Duc", Code: "CBA"})
	var c *Claims
	c = &Claims{}
	err = DecryptGob(ser, to, c)
	fmt.Println(c, err)

}

func BenchmarkT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ser, _ := CreateSecrecKey()
		to, _ := EnCryptGob(ser, Claims{Name: "Duc", Code: "CBA"})
		var c *Claims
		c = &Claims{}
		_ = DecryptGob(ser, to, c)
	}
}

func TestLog(t *testing.T) {

}

func TestChangeBy(t *testing.T) {
	a := []byte(fmt.Sprint(time.Now().UTC().Unix()))
	var wg sync.WaitGroup
	for i, s := range a {
		wg.Add(1)
		go func(i int, s byte) {
			if i%2 == 0 {
				a[i] = s + 49
			} else {
				a[i] = s + 17
			}
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	fmt.Println(string(a))
}

func BenchmarkChangeBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []byte(fmt.Sprint(time.Now().UTC().Unix()))
		var wg sync.WaitGroup
		for i, s := range a {
			wg.Add(1)
			go func(i int, s byte) {
				if i%2 == 0 {
					a[i] = s + 49
				} else {
					a[i] = s + 17
				}
				wg.Done()
			}(i, s)
		}
		wg.Wait()
	}
}

func BenchmarkChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []byte(fmt.Sprint(time.Now().UTC().Unix()))
		for i, s := range a {
			if i%2 == 0 {
				a[i] = s + 49
			} else {
				a[i] = s + 17
			}
		}
	}
}
