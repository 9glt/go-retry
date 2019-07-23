package retry

import "time"

func Run(fn func(r int) error, count int, delay time.Duration) error {
	var err error
	for i := count; i > 0; i-- {
		err = fn(i)
		time.Sleep(delay)
		if err == nil {
			break
		}
	}
	return err
}
