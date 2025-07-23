package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/renniemaharaj/grouplogs/pkg/logger"
)

func main() {

	group := logger.CreateGroup()

	l1 := logger.New().
		Prefix("Primary").
		Debugging(true).
		JSONMode(false).
		Subscribable(true).MaxLines(100).STDOUT(false).Rotate()

	group.Join(l1)

	l2 := logger.New().
		Prefix("Secondary").
		Debugging(true).
		JSONMode(false).
		Subscribable(true).MaxLines(100).STDOUT(false).Rotate()

	group.Join(l2)

	l1.Info("This is an information").
		Success("This is a success").
		Warning("This is a warning").
		Debug("Is debugging enabled").
		Error("Oh, no. This is an error")

	l2.Info("This is an information").
		Success("This is a success").
		Warning("This is a warning").
		Debug("Is debugging enabled").
		Error("Oh, no. This is an error")

	idleLimit := 500 * time.Millisecond
	timer := time.NewTimer(idleLimit)

	for {
		select {
		case l := <-group.Delegate:
			lBytes, _ := json.Marshal(l)
			fmt.Println(string(lBytes))
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(idleLimit)
		case <-timer.C:
			l1.STDOUT(true)
			l1.Success("Removing logger 1 from group")
			group.Remove(l1)

			l2.STDOUT(true)
			l2.Success("Removed logger 2 from group")
			group.Remove(l2)

			return
		}
	}

}
