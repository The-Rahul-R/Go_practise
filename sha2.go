package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("enter abs filepath")
	var filepath string
	fmt.Scanln(&filepath)

	var ch int = 1
	var input_hash string

	for ch != 3 {
		fmt.Println("select 1 to calc hash and 2 to check file integrity and 3 to exit")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			var sha256 string
			sha256 = calc_sha(filepath)
			fmt.Println(sha256)

		case 2:
			fmt.Println("enter hash")
			fmt.Scanln(&input_hash)
			check_integrity(filepath, input_hash)

		}
	}
}

func calc_sha(fpath string) string {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	hash := h.Sum(nil)
	return hex.EncodeToString(hash[:])
}

func check_integrity(fpath string, hash string) {
	h := calc_sha(fpath)
	if h == hash {
		fmt.Println("file integrity maintained")
	} else {
		fmt.Println("file integrity compromised")
	}

}
