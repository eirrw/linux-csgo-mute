package pipewire

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

const (
	PwDumpCmd     = "pw-dump"
	PwCliCmd      = "pw-cli"
	PwCliSetOpt   = "s"
	PwCliPropsOpt = "Props"
	PwCliVolProp  = "{ volume: %f }"
)

// Pipewire is a struct for managing access to a pipewire instance, holding node information and mutex locks for access
// to the system
type Pipewire struct {
	Nodes []node      // A slice containing nodes loaded from the local pipewire server
	mutex *sync.Mutex // Used to lock access to racy system calls
}

// node is the base struct to contain information about pipewire nodes
type node struct {
	Id          int      `json:"id"`
	Type        string   `json:"type"`
	Version     int      `json:"version"`
	Permissions []string `json:"permissions"`
	Info        info     `json:"info"`
}

// info contains basic information about the pipewire node, including the server cookie and property lists
type info struct {
	Cookie     int       `json:"cookie"`
	UserName   string    `json:"user-name"`
	HostName   string    `json:"host-name"`
	Name       string    `json:"name"`
	ChangeMask []string  `json:"change-mask"`
	Props      infoProps `json:"props"`
}

// infoProps holds a series of useful values for managing properties of a pipewire node
type infoProps struct {
	LogLevel                 int    `json:"log.level"`
	NodeName                 string `json:"node.name"`
	ApplicationProcessId     int    `json:"application.process.id"`
	ApplicationProcessUser   string `json:"application.process.user"`
	ApplicationProcessHost   string `json:"application.process.host"`
	ApplicationProcessBinary string `json:"application.process.binary"`
	ApplicationName          string `json:"application.name"`
	MediaClass               string `json:"media.class"`
	NodeDescription          string `json:"node.description"`
	ClientName               string `json:"client.name"`
	MediaType                string `json:"media.type"`
	MediaCategory            string `json:"media.category"`
	MediaRole                string `json:"media.role"`
}

// New creates and loads a new instance of the Pipewire struct.
func New() *Pipewire {
	var pipewire Pipewire
	pipewire.loadPipewire()
	pipewire.mutex = &sync.Mutex{}

	return &pipewire
}

// loadPipewire imports the output of pw-dump into the Pipewire struct to load the state of the local pipewire server.
func (pw *Pipewire) loadPipewire() {
	out, err := exec.Command(PwDumpCmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	pw.Nodes = make([]node, 0)
	// pw-dump struct
	if err := json.Unmarshal(out, &pw.Nodes); err != nil {
		log.Fatal(err)
	}
}

// SetVolume changes the volume property of the given pipewire node using the pw-cli command.
func (pw *Pipewire) SetVolume(nodeId int, volume float32) error {
	pw.mutex.Lock()
	defer pw.mutex.Unlock()
	cmd := exec.Command(
		PwCliCmd,
		PwCliSetOpt,
		fmt.Sprintf("%d", nodeId),
		PwCliPropsOpt,
		fmt.Sprintf(PwCliVolProp, volume),
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// GetNodeIdByName returns the pipewire node id of the node with a matching name. If there are multiple matching nodes,
// which one is returned is undetermined. Returns -1 if no matching node is found.
func (pw *Pipewire) GetNodeIdByName(nodeName string) int {
	for _, node := range pw.Nodes {
		if nodeName == node.Info.Props.NodeName && node.Type == "PipeWire:Interface:Node" {
			return node.Id
		}
	}

	return -1
}
