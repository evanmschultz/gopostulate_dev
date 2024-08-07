package server

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"

// 	"github.com/evanmschultz/gopostulate_dev/pkg/router"
// )

// func TestServer(t *testing.T) {
// 	// Create a temporary directory for test pages
// 	tempDir, err := os.MkdirTemp("", "test-pages")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp directory: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)
// // 
// 	// Create a test page
// 	testPage := filepath.Join(tempDir, "index.go")
// 	err = os.WriteFile(testPage, []byte(`
// package pages

// import (
// 	"fmt"
// 	"net/http"
// )

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the home page")
// }
// `), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test page: %v", err)
// 	}

// 	// Create a router with the test pages
// 	router, err := router.GenerateRoutes(tempDir)
// 	if err != nil {
// 		t.Fatalf("Failed to generate routes: %v", err)
// 	}

// 	// Create a server with the test router
// 	server := &Server{
// 		router: router,
// 		// addr:   ":8080",
// 	}

// 	// Test the home page
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	server.router.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := "Handler Index for /"
// 	if !strings.Contains(rr.Body.String(), expected) {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }