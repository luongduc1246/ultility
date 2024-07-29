package random

import (
	"fmt"
	"testing"
)

func TestRandomAlpha(t *testing.T) {
	code := CreateCodeRamdomAlphabet(6)
	fmt.Println(code)
}
func TestRandomNumerals(t *testing.T) {
	code := CreateCodeRamdomNumerals(6)
	fmt.Println(code)
}
func TestRandomDigit(t *testing.T) {
	code := CreateCodeRamdomDigit(6)
	fmt.Println(code)
}
