package template

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/a-h/templ"
)

func TestTemplateEngine(t *testing.T) {
	engine, err := NewEngine("../../templates")
	if err != nil {
		t.Fatalf("Failed to create template engine: %v", err)
	}

	// Create a test component
	testComponent := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "<h1>Hello, GoPostulate!</h1>")
		return err
	})

	// Add the test component to the engine
	engine.components["test"] = testComponent

	var buf bytes.Buffer
	err = engine.Render(&buf, "test", nil)
	if err != nil {
		t.Fatalf("Failed to render component: %v", err)
	}

	expected := "<h1>Hello, GoPostulate!</h1>"
	if buf.String() != expected {
		t.Errorf("component rendered incorrectly: got %v want %v",
			buf.String(), expected)
	}

	// Test RenderToString
	result, err := engine.RenderToString("test", nil)
	if err != nil {
		t.Fatalf("Failed to render component to string: %v", err)
	}

	if result != expected {
		t.Errorf("RenderToString rendered incorrectly: got %v want %v",
			result, expected)
	}

	// Test rendering a non-existent component
	err = engine.Render(&buf, "nonexistent", nil)
	if err == nil {
		t.Error("Expected an error when rendering a non-existent component, but got nil")
	}
}