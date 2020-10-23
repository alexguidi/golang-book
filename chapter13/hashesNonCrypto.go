package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main() {

	//non-cryptographic hash
	//we can use to see if the file are the same, if the Sum32 is the same they are the same file (not 100% certain)
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h2 := crc32.NewIEEE()
	h2.Write([]byte("Test"))
	v2 := h2.Sum32()
	fmt.Println(v2)

	h3, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h4, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h3, h4, h3 == h4)
}
