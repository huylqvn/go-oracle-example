package helpers

import "time"

func Watch(f func()) {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}
