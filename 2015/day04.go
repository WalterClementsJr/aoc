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
	for i := 0; i < MAX_INT; i++ {
		s := INPUT + strconv.Itoa(i)

		computedHash := hash(s)

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
