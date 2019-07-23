package retry

import (
	"errors"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	var cnt int

	var data string
	Run(func(r int) error {
		cnt++
		var err error
		data, err = testfn(r)
		return err
	}, 3, 300*time.Millisecond)

	if data != "hi" {
		t.Fatal()
	}

	if cnt != 3 {
		t.Fatal()
	}
}

func testfn(i int) (string, error) {
	if i == 1 {
		return "hi", nil
	}
	return "", errors.New("error")
}
