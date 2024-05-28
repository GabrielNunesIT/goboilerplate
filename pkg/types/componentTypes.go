package types

import "fmt"

type ComponentType string

const (
	Main      ComponentType = "main"
	WebServer ComponentType = "webserver"
	Config    ComponentType = "config"
	DB        ComponentType = "db"
	Log       ComponentType = "log"
)

func (componentType ComponentType) validateComponentType() error {
	switch componentType {
	case Main, WebServer, Config, DB, Log:
		return nil
	default:
		return fmt.Errorf("invalid component type: %s", componentType)
	}
}
