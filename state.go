package main

import (
	"fmt"
	"sync"
)

type DemonStatus struct {
	Connected  bool   `json:"connected"`
	Online     bool   `json:"online"`
	FFmpeg     bool   `json:"ffmpeg"`
	Rendering  bool   `json:"rendering"`
	CurrentJob string `json:"currentJob"`
	AgentID    string `json:"agentId"`
	Hostname   string `json:"hostname"`
	Version    string `json:"version"`
	OS         string `json:"os"`
}

var demonState DemonStatus

var stateMutex sync.RWMutex

func InitDemonState() {
	identity := GetIdentity()
	stateMutex.Lock()
	defer stateMutex.Unlock()
	demonState = DemonStatus{
		Connected:  true,
		Online:     true,
		FFmpeg:     CheckFFmpeg(),
		Rendering:  false,
		CurrentJob: "",
		AgentID:    identity.ID,
		Hostname:   identity.Hostname,
		Version:    version,
		OS: 				identity.OS,
	}
	fmt.Println("Demon state initialized")
	fmt.Printf("%+v\n", demonState)
}

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
		fmt.Println("STATUS UPDATE")
		fmt.Printf("%+v\n", demonState)
		broadcastStatus(demonState)
	}
}