package core

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/todo/internal/step"
)

type Plugin struct {
	drone.Base
}

func NewPlugin() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewTodo()).Name("样例").Build(),
	}
}
