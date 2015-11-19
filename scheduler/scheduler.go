package scheduler

import (
	"errors"
	"time"
)

// A schedule function for schedule function what for a time in one day
// what is the function to be scheduled
// clock is a variable length param which likes 0, 0, 0, 12, 23, 11
// the clock suggests that the what func with be scheduled at 12:23:11 only once
// if the task should be scheduled repeatedly at a interval of year, month or day
// which coresponding to the first three params
// the task should be guarenteed that it will be finished
func Schedule(what func(), clock ...int) (chan bool, error) {
	specificTime := []int(clock)
	length := len(specificTime)
	if length > 7 || length < 6 {
		return nil, errors.New("illegal parameters")
	}
	if length == 6 {
		specificTime = append(specificTime, 0)
	}
	stop := make(chan bool)
	go func() {
		now := time.Now()
		
		for {
			today := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				specificTime[3], specificTime[4],
				specificTime[5], specificTime[6],
				time.Local,
			)
			delay := today.Sub(now)
			if delay <= 0 {
				next := now.AddDate(specificTime[0],
					specificTime[1],
					specificTime[2])
				tomorrow := time.Date(
				next.Year(),
				next.Month(),
				next.Day(),
				specificTime[3], specificTime[4],
				specificTime[5], specificTime[6],
				time.Local)
				delay = tomorrow.Sub(now)
				if delay <= 0 {
					<-stop
					return
				}
				now = tomorrow
			} else {
				now = today
			}
			select {
			case <-time.After(delay):
				go what()
			case <-stop:
				return
			}
		}
	}()
	return stop, nil
}
