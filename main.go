package main

import (
	"github.com/dronestock/drone"
)

func main() {
	drone.New(core.NewPlugin).Boot()
}
