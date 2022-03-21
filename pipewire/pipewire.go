package pipewire

import (
	"encoding/json"
	"log"
	"os/exec"
)

var (
	pwcli      = "pw-cli"
	pwdump     = "pw-dump"
	pw_arg_set = "s"
	pw_arg_ls  = "ls"
)

func loadPipewire() {
	cmd := exec.Command(pwdump)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// pw-dump struct
	var pipewire Node
	if err := json.NewDecoder(stdout).Decode(&pipewire); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
