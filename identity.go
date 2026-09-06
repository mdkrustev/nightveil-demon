package main
import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
	"github.com/google/uuid"
)
type AgentIdentity struct {
	ID string `json:"id"`
	Hostname string `json:"hostname"`
	OS string `json:"os"`
	Created time.Time `json:"created"`
}
var agentIdentity AgentIdentity
func LoadIdentity() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	path := filepath.Join(
		dir,
		"NightVeilDemon",
		"agent.json",
	)
	os.MkdirAll(
		filepath.Dir(path),
		0755,
	)
	data, err := os.ReadFile(path)
	if err == nil {
		return json.Unmarshal(
			data,
			&agentIdentity,
		)
	}
	hostname, _ := os.Hostname()
	agentIdentity = AgentIdentity{
		ID: uuid.New().String(),
		Hostname: hostname,
		OS: GetOSName(),
		Created: time.Now(),
	}
	jsonData, _ := json.MarshalIndent(
		agentIdentity,
		"",
		"  ",
	)
	return os.WriteFile(
		path,
		jsonData,
		0644,
	)
}
func GetIdentity() AgentIdentity {
	return agentIdentity
}