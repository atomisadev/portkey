package main

import (
	"fmt"

	"github.com/atomisadev/portkey/cmd"
	"github.com/atomisadev/portkey/internal/config"
)

func main() {
	hosts, err := config.LoadHosts()
	if err != nil {
		fmt.Printf("Warning: Could not load SSH config: %v\n", err)
	} else {
		fmt.Printf("DEBUG: Found %d hosts in ~/.ssh/config\n", len(hosts))
	}

	cmd.Execute()
}
