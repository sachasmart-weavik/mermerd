package main

import (
	"github.com/spf13/viper"

	"github.com/sachasmart-weavik/mermerd/cmd"
)

// ldflags flags from goreleaser
var (
	version = "dev"
	commit  = "none"
)

func main() {
	viper.Set("version", version)
	viper.Set("commit", commit)

	cmd.Execute()
}
