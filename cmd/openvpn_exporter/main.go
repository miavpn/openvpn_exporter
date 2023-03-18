package main

import (
	"os"

	"github.com/ricardoalonso/openvpn_exporter/pkg/command"
)

func main() {
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
