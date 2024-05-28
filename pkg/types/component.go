package types

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Component struct {
	Name      string
	Path      string
	CompType  ComponentType
	Framework Framework
}

// UnmarshalYAML unmarshals a YAML file into a Component struct.
func (comp *Component) UnmarshalYAML(value *yaml.Node) (err error) {
	for i := 0; i < len(value.Content); i++ {
		yamlVal := value.Content[i]
		switch yamlVal.Value {
		case "name":
			comp.Name = value.Content[i+1].Value
		case "path":
			comp.Path = value.Content[i+1].Value
		case "type":
			comp.CompType = ComponentType(value.Content[i+1].Value)
			err = comp.CompType.validateComponentType()
		case "framework":
			framework, ok := frameworks.GetFramework(value.Content[i+1].Value)
			if !ok {
				return fmt.Errorf("framework with name: %s not found", value.Content[i+1].Value)
			}
			comp.Framework = framework
		}
	}

	return
}

func (comp *Component) validateComponentType() error {
	return comp.CompType.validateComponentType()
}

type Components map[string]Component

var components Components = make(Components)

func GetComponent(key string) (comp Component, ok bool) {
	comp, ok = components[key]

	return
}

func RegisterComponent(comp Component) (err error) {
	if _, ok := components[comp.Name]; ok {
		return fmt.Errorf("component with key: %s already registered", comp.Name)
	}

	if err = comp.validateComponentType(); err != nil {
		return
	}

	components[comp.Name] = comp
	return
}
