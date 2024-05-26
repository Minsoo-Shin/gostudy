package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
	"sort"
)

func main() {
	// 입력 이미지 열기
	file, err := os.Open("cmd/image_block/gopher.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	data, err := io.ReadAll(rd)
	if err != nil {
		fmt.Println("Error reading rc")
	}
	// image를 4kb로 나눠보자.
	hasLen := 4096 // 4kb
	numBlocks := len(data) / hasLen

	var blocks = make([]block, hasLen)

	for i := 0; i < numBlocks; i++ {
		blocks[i].order = i
		copy(blocks[i].data[:], data[i*hasLen:(i+1)*hasLen])
	}
	ReadBlocks(blocks)
}

type block struct {
	data  [4096]byte
	order int
}

/*
hashLen := 20 // Length of SHA-1 hash
	buf := []byte(i.Pieces)
	if len(buf)%hashLen != 0 {
		err := fmt.Errorf("Received malformed pieces of length %d", len(buf))
		return nil, err
	}
	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)

	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes, nil
*/

func ReadBlocks(blocks []block) {
	// blocks를 order 필드를 기준으로 정렬
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].order < blocks[j].order
	})

	var data []byte
	for _, block := range blocks {
		data = append(data, block.data[:]...)
	}
	file, _ := os.Create("output_1.png")
	wr := bufio.NewWriter(file)
	nn, err := wr.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("just check", nn)

	reader, err := os.Open("output_1.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	im, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(im.Bounds())

}
