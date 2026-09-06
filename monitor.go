package main

import (
	"time"
)

func StartMonitor(){

	go func(){

		lastStatus:=GetStatus()

		for {

			current:=DemonStatus{
				Connected:lastStatus.Connected,
				Online:true,
				FFmpeg:CheckFFmpeg(),
				Rendering:lastStatus.Rendering,
				CurrentJob:lastStatus.CurrentJob,
				AgentID:lastStatus.AgentID,
				Hostname:lastStatus.Hostname,
				Version: lastStatus.Version,
				OS: lastStatus.OS,
			}

			if current!=lastStatus{

				UpdateStatus(current)

				lastStatus=current
			}

			time.Sleep(
				2*time.Second,
			)
		}

	}()

}