package miscs

import (
	"io"
	"os"
	"time"
)

//DateCheck ...
func DateCheck(date string) (err error) {
	_, err = time.Parse("2006-01-02", date)
	return
}

//IsDirEmpty ...
func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
