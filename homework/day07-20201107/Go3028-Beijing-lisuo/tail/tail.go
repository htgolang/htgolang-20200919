package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
	"time"
)

var pathname = "./test.tail"
var mask uint32 = 755

func main() {
	// to implement
	fileToTail := flag.String("f", "foo", "Specify a file to tail.")
	lastLines := flag.Int("n", 10, "Specify to display last n lines.")
	flag.Usage = func() {
		fmt.Println("Usage: go run tail.go -f /path/to/foo -n <NUM>")
	}
	flag.Parse()
	if *fileToTail == "foo" {
		flag.Usage()
		os.Exit(1)
	}

	tail(*fileToTail, *lastLines)
}

// tail
func tail(file string, lastLines int) error {
	f, e := os.Open(file)
	if e != nil {
		log.Fatal(e)
	}
	fd, err := syscall.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	_, errW := syscall.InotifyAddWatch(fd, file, syscall.IN_MODIFY)
	if errW != nil {
		log.Fatal(errW)
	}
	r := bufio.NewReader(f)

	// print last n lines
	if err := showLastLines(f, lastLines); err != nil {
		fmt.Println("err: ", err)
	}
	var rbytes []byte
	for {
		rbytes, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		// print added lines
		fmt.Print(string(rbytes))
		if err == io.EOF {
			time.Sleep(2 * time.Second)
			continue
		}
		if err = watchForModify(fd); err != nil {
			return err
		}

	}

}

// wath a fd's modify event
func watchForModify(fd int) error {
	for {
		var tmpBuf [syscall.SizeofInotifyEvent]byte
		_, err := syscall.Read(fd, tmpBuf[:])
		if err != nil {
			return err
		}
		r := bytes.NewReader(tmpBuf[:])
		var event = syscall.InotifyEvent{}
		errR := binary.Read(r, binary.LittleEndian, &event)
		if errR != nil {
			return errR
		}
		if event.Mask&syscall.IN_MODIFY == syscall.IN_MODIFY {
			return nil
		}
	}

}

//  print last n lines
func showLastLines(f *os.File, n int) error {
	s := bufio.NewScanner(f)
	var lastLines = n
	var cnt int
	var lastLinesChan = make(chan []byte, lastLines+1)

	for {
		s.Scan()
		lastLinesChan <- s.Bytes()
		if s.Text() == "" {
			fmt.Println("done read...")
			fmt.Println("total lines: ", cnt)
			break
		}
		cnt++
		if len(lastLinesChan) == cap(lastLinesChan) {
			<-lastLinesChan
			continue
		}
	}
	for i := 0; i < lastLines; i++ {
		fmt.Println(string(<-lastLinesChan))
	}
	return nil
}
