package main

import (
	"time"
)

func StartMonitor() {

	go func() {

		lastStatus := GetStatus()

		for {

			current := DemonStatus{
				Online:     true,
				FFmpeg:     CheckFFmpeg(),
				Rendering:  lastStatus.Rendering,
				CurrentJob: lastStatus.CurrentJob,
			}

			if current != lastStatus {

				UpdateStatus(current)

				lastStatus = current
			}

			time.Sleep(
				2 * time.Second,
			)

		}

	}()

}
