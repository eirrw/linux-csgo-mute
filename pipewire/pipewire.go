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

type Pipewire struct {
	Nodes []Node
	Mutex *sync.Mutex
}

func New() *Pipewire {
	var pipewire Pipewire
	pipewire.loadPipewire()
	pipewire.Mutex = &sync.Mutex{}

	return &pipewire
}

func (pw *Pipewire) loadPipewire() {
	out, err := exec.Command(PwDumpCmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	pw.Nodes = make([]Node, 10)
	// pw-dump struct
	if err := json.Unmarshal(out, &pw.Nodes); err != nil {
		log.Fatal(err)
	}
}

func (pw *Pipewire) SetVolume(nodeId int, volume float32) error {
	pw.Mutex.Lock()
	defer pw.Mutex.Unlock()
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

func (pw *Pipewire) GetNodeIdByName(nodeName string) int {
	for _, node := range pw.Nodes {
		if nodeName == node.Info.Props.NodeName {
			return node.Id
		}
	}

	return -1
}
