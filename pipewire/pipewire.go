package pipewire

import (
	"encoding/json"
	"log"
	"os/exec"
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
}

func New() *Pipewire {
	var pipewire Pipewire
	pipewire.loadPipewire()

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
