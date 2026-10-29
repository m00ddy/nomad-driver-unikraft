package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"
	"github.com/m00ddy/nomad-driver-unikraft/internal/driver"
)

func main() {
	plugins.Serve(factory)
}

func factory(log hclog.Logger) interface{} {
	return driver.New(log)
}
