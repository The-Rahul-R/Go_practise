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
	stringvalue := calc_sha(filepath)
	fmt.Println(stringvalue)
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
	fmt.Printf("%x \n", h.Sum(nil))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash[:])

}
