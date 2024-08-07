// internal/template/engine.go
package template

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

type Engine struct {
	components map[string]templ.Component
}

func NewEngine(dir string) (*Engine, error) {
	engine := &Engine{
		components: make(map[string]templ.Component),
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".templ" {
			// For simplicity, we're assuming that each .templ file exports
			// a component with the same name as the file (minus extension).
			// In a real implementation, you'd need to parse the file and
			// extract the actual component names.
			name := filepath.Base(path[:len(path)-len(".templ")])
			
			// Here you would actually compile the Templ file and get the component.
			// For this example, we're using a placeholder function.
			component, err := loadTemplComponent(path)
			if err != nil {
				return err
			}
			
			engine.components[name] = component
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return engine, nil
}

func (e *Engine) Render(w io.Writer, name string, data interface{}) error {
	component, ok := e.components[name]
	if !ok {
		return fmt.Errorf("component %s not found", name)
	}

	return component.Render(context.Background(), w)
}

func (e *Engine) RenderToString(name string, data interface{}) (string, error) {
	component, ok := e.components[name]
	if !ok {
		return "", fmt.Errorf("component %s not found", name)
	}

	var buf bytes.Buffer
	err := component.Render(context.Background(), &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// This is a placeholder function. In a real implementation,
// you would need to compile the Templ file and return the actual component.
func loadTemplComponent(path string) (templ.Component, error) {
	// This is where you'd compile the Templ file and return the component.
	// For now, we're returning a placeholder component.
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintf(w, "Placeholder for component from %s", path)
		return err
	}), nil
}