package scheduler

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	test_func := func() {
		fmt.Println("#########")
	}
	done, err := Schedule(test_func, 0, 0, 0,
		17, 54, 0)
	time.Sleep(1 * time.Minute)
	done <- true
	if err != nil {
		t.Error("task executing failed")
	}
}
