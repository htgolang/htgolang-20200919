package miscs

import (
	"time"
)

//DateCheck ...
func DateCheck(date string) (err error) {
	_, err = time.Parse("2006-01-02", date)
	return
}
