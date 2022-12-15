package encode

import (
	"math/bits"
)

func countBits(src []int32) int {
	sum := 0

	for _, v := range src {
		if v < 0 {
			sum += 32
			continue
		}

		sum += 32 - bits.LeadingZeros32(uint32(v))
	}

	return sum
}

// https://doi.org/10.1109/ICDE.1998.655800
func FOR(src []int32) []int32 {
	// min, max := src[0], src[0]
	//
	// for _, v := range src {
	// 	if v < min {
	// 		min = v
	// 	}
	//
	// 	if v > max {
	// 		max = v
	// 	}
	// }
	//
	// offset := max - min

	min := src[0]

	for _, v := range src {
		if v < min {
			min = v
		}
	}

	dst := make([]int32, len(src))

	for i, v := range src {
		dst[i] = v - min
	}

	return dst
}

// https://en.wikipedia.org/wiki/Delta_encoding
func Delta(src []int32) []int32 {
	dst := make([]int32, len(src))

	for i, v := range src {
		if i == 0 {
			dst[i] = v
			continue
		}

		dst[i] = v - src[i-1]
	}

	return dst
}

func Dictionary(src []int32) []int32 {
	dict := make(map[int32]int32)
	dst := make([]int32, len(src))

	var key int32

	for i, v := range src {
		if _, ok := dict[v]; !ok {
			dict[v] = key
			key++
		}

		dst[i] = dict[v]
	}

	return dst
}
