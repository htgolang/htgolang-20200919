package funcs

import (
	"fmt"
	"os"
	"sync"
)

const cBufferSize = 200

type chunk struct {
	bufsize int
	offset int64
}

func ConcurrentReads() {

	file, err := os.Open("day02/I have a dream")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := int(fileInfo.Size())
	// Number of go routines we need to spawn.
	concurrency := fileSize / cBufferSize
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunkSizes := make([]chunk, concurrency)

	for i := 0; i < concurrency; i++ {
		chunkSizes[i].bufsize = cBufferSize
		chunkSizes[i].offset = int64(cBufferSize * i)
	}

	if remainder := fileSize % cBufferSize; remainder != 0 {
		c := chunk{bufsize: remainder, offset: int64(concurrency * cBufferSize)}
		concurrency++
		chunkSizes = append(chunkSizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunkSizes []chunk, i int) {
			defer wg.Done()
			chunk := chunkSizes[i]
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := file.ReadAt(buffer, chunk.offset)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("bytes read, string(bytestream):", bytesread)
			fmt.Println("bytestream to string:", string(buffer))
		}(chunkSizes, i)
	}
	wg.Wait()
}
