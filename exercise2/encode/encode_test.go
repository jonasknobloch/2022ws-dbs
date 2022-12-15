package encode

import (
	"fmt"
	"testing"
)

func TestFOREncoding(t *testing.T) {
	input := []int32{1, 3, 7, 12, 13, 13, 14, 17}
	output := FOR(input)

	fmt.Println(countBits(input), input)
	fmt.Println(countBits(output), output)
}

func TestDeltaEncoding(t *testing.T) {
	input := []int32{1, 3, 7, 12, 13, 13, 14, 17}
	output := Delta(input)

	fmt.Println(countBits(input), input)
	fmt.Println(countBits(output), output)
}

func TestDictionaryEncoding(t *testing.T) {
	input := []int32{1, 3, 7, 12, 13, 13, 14, 17}
	output := Dictionary(input)

	fmt.Println(countBits(input), input)
	fmt.Println(countBits(output), output)
}
