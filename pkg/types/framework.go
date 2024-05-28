package types

import (
	"fmt"
	"net/url"
	"sync"
	"text/template"
)

type Framework interface {
	GetName() string
	GetURL() url.URL
	GetType() ComponentType
	GetTemplate() template.Template
}

type Frameworks struct {
	mu         sync.Mutex
	frameworks map[string]Framework
}

var frameworks Frameworks = Frameworks{frameworks: make(map[string]Framework)}

// AddFramework adds a new framework to the collection of frameworks.
// It returns an error if a framework with the same name already exists.
func (fws *Frameworks) AddFramework(fw Framework) (err error) {
	fws.mu.Lock()
	if _, ok := fws.frameworks[fw.GetName()]; ok {
		return fmt.Errorf("framework with name: %s already registered", fw.GetName())
	}

	fws.frameworks[fw.GetName()] = fw
	fws.mu.Unlock()
	return
}

// AddOrReplaceFramework adds or replaces a framework in the Frameworks collection.
// If a framework with the same name already exists, it will be replaced.
func (fws *Frameworks) AddOrReplaceFramework(fw Framework) {
	fws.mu.Lock()
	fws.frameworks[fw.GetName()] = fw
	fws.mu.Unlock()
}

// RemoveFramework removes a framework from the Frameworks collection.
// It takes a key parameter specifying the name of the framework to remove.
// If the framework is not found in the collection, it returns an error.
// Otherwise, it removes the framework and returns nil.
func (fws *Frameworks) RemoveFramework(key string) (err error) {
	fws.mu.Lock()
	if _, ok := fws.frameworks[key]; !ok {
		return fmt.Errorf("framework with name: %s not registered", key)
	}

	delete(fws.frameworks, key)
	return
}

func (fws *Frameworks) GetFramework(key string) (fw Framework, ok bool) {
	fws.mu.Lock()
	fw, ok = fws.frameworks[key]
	fws.mu.Unlock()

	return
}
