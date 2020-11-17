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

func main() {
	// to implement
	fileToTail := flag.String("f", "foo", "Specify a file to tail.")
	lastLines := flag.Int("n", 10, "Specify to display last n lines.")
	flag.Usage = func() {
		fmt.Println("Usage: go run tail.go -f <FILE> -n <NUM>")
	}
	flag.Parse()
	if *fileToTail == "foo" {
		flag.Usage()
		os.Exit(1)
	}

	f, e := os.Open(*fileToTail)
	defer f.Close()
	if e != nil {
		log.Fatal(e)
	}
	fd, err := syscall.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	_, errW := syscall.InotifyAddWatch(fd, *fileToTail, syscall.IN_MODIFY)
	if errW != nil {
		log.Fatal(errW)
	}

	// print last n lines
	if err := showLastLines(f, *lastLines); err != nil {
		fmt.Println("err: ", err)
	}

	for {
		errWa := watchForModify(fd)
		if errWa == nil {
			tail(f, *lastLines, fd)
		} else {
			return
		}
	}
}

// tail
func tail(f *os.File, lastLines int, fd int) error {

	r := bufio.NewReader(f)

	//var rbytes []byte
	for {
		rbytes, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		// print added lines
		fmt.Print(string(rbytes))
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		errW := watchForModify(fd)
		if errW != nil {
			return errW
		}
	}
}

// watch a fd's modify event
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
	//closelastLinesChan)
	return nil
}
