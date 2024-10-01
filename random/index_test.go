package random

import (
	"fmt"
	"testing"
)

func TestRandomAlpha(t *testing.T) {
	code := CreateCodeRandomAlphabet(6)
	fmt.Println(code)
}
func TestRandomNumerals(t *testing.T) {
	code := CreateCodeRandomNumerals(6)
	fmt.Println(code)
}
func TestRandomDigit(t *testing.T) {
	code := CreateCodeRandomDigit(6)
	fmt.Println(code)
}
