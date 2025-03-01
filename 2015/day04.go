package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var (
	hasher  = md5.New()
	INPUT   = "iwrupvqb"
	MAX_INT = int(^uint(0) >> 1)
)

func main() {
	// data := "abcdef609043"
	// data = "abcdef"

	for i := 0; i < MAX_INT; i++ {
		// for i := 0; i < 609045; i++ {
		s := INPUT + strconv.Itoa(i)

		computedHash := hash(s)
		// fmt.Println(computedHash, i, computedHash[0:5])

		if (computedHash[0:6]) == "000000" {
			fmt.Println(computedHash, i)
			break
		}
	}
}

func hash(input string) string {
	hasher.Write([]byte(input))
	result := hex.EncodeToString(hasher.Sum(nil))
	hasher.Reset()
	return result
}
