package main

import (
	"sync"
)

type DemonStatus struct {
	Online     bool   `json:"online"`
	FFmpeg     bool   `json:"ffmpeg"`
	Rendering  bool   `json:"rendering"`
	CurrentJob string `json:"currentJob"`
}

var demonState = DemonStatus{
	Online:     true,
	FFmpeg:     false,
	Rendering:  false,
	CurrentJob: "",
}

var stateMutex sync.RWMutex

func GetStatus() DemonStatus {

	stateMutex.RLock()
	defer stateMutex.RUnlock()

	return demonState
}

func UpdateStatus(newStatus DemonStatus) {

	stateMutex.Lock()

	changed := demonState != newStatus

	demonState = newStatus

	stateMutex.Unlock()

	if changed {
		broadcastStatus(demonState)
	}
}
