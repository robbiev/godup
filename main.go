package main

import (
	"fmt"
	"math"
)

const prime = 1021

func naiveIndexOf(s string, pattern string) int {
Outer:
	for i := 0; i < len(s)-len(pattern); i++ {
		for j := range pattern {
			if s[i+j] != pattern[j] {
				continue Outer
			}
		}
		return i
	}
	return -1
}

func rabinFingerprint(str string) uint32 {
	hash := uint32(0)
	maxPow := len(str) - 1
	for i := 0; i < len(str); i++ {
		hash += uint32(str[i]) * uint32(math.Pow(prime, float64(maxPow-i)))
	}
	return hash
}

func nextRabinFingerprint(pow uint32, oldHash uint32, in, out uint8) uint32 {
	return (prime * (oldHash - (uint32(out) * pow))) + uint32(in)
}

func rabinKarpIndexOf(s string, pattern string) int {
	if len(s) == 0 || len(pattern) == 0 || len(s) < len(pattern) {
		return -1
	}

	if len(s) == len(pattern) {
		if s != pattern {
			return -1
		}
		return 0
	}

	pow := uint32(math.Pow(prime, float64(len(pattern)-1)))

	hpattern := rabinFingerprint(pattern)
	hs := rabinFingerprint(s[0:len(pattern)])

	for i := 0; i <= len(s)-len(pattern); i++ {
		if hs == hpattern {
			if s[i:i+len(pattern)] == pattern {
				return i
			}
		}

		inIndex := i + len(pattern)
		if inIndex < len(s) {
			hs = nextRabinFingerprint(pow, hs, s[inIndex], s[i])
		}
	}
	return -1
}

func main() {
	fmt.Println(naiveIndexOf("blah", "a"))
	fmt.Println(naiveIndexOf("blah", "x"))

	fmt.Println(rabinKarpIndexOf("blah", "a"))
	fmt.Println(rabinKarpIndexOf("blah", "x"))

	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "blah"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blahblah", "blah"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "b"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "l"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "a"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "h"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "bl"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "la"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "ah"))
	fmt.Println("===")
	fmt.Println(rabinKarpIndexOf("blah", "lah"))
}
