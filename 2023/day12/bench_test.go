package main

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
	"strings"
	"testing"
)

func keyStrconv(cond string, nums []int) string {
	strNums := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strNums[i] = strconv.Itoa(nums[i])
	}
	return cond + strings.Join(strNums, ",")
}

type Key struct {
	cond string
	nums [32]byte
}

func keySHA256(cond string, nums []int) Key {
	sha := sha256.New()

	buf := make([]byte, 8)
	for _, num := range nums {
		binary.LittleEndian.PutUint64(buf, uint64(num))
		sha.Write(buf)
	}

	return Key{cond, [32]byte(sha.Sum(nil))}
}

func BenchmarkKeyStrconv(b *testing.B) {
	cond := "#?????#????#?#?#..#"
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		keyStrconv(cond, nums)
	}
}

func BenchmarkKeySHA256(b *testing.B) {
	cond := "#?????#????#?#?#..#"
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		keySHA256(cond, nums)
	}
}
