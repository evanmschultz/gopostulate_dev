package router

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )

// func TestRouter(t *testing.T) {
// 	// Create a temporary directory for test pages
// 	tempDir, err := os.MkdirTemp("", "test-pages")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp directory: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Create test pages
// 	pages := map[string]string{
// 		"index.go": `
// package pages

// import (
// 	"fmt"
// 	"net/http"
// )

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the home page")
// }`,
// 		"about.go": `
// package pages

// import (
// 	"fmt"
// 	"net/http"
// )

// func About(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "About Us")
// }`,
// 		"users/[id].go": `
// package users

// import (
// 	"fmt"
// 	"net/http"
// )

// func UserDetails(w http.ResponseWriter, r *http.Request) {
// 	id := r.Context().Value(ParamKey("id")).(string)
// 	fmt.Fprintf(w, "User Details for ID: %s", id)
// }`,
// 	}

// 	for filename, content := range pages {
// 		fullPath := filepath.Join(tempDir, filename)
// 		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
// 		if err != nil {
// 			t.Fatalf("Failed to create directory for %s: %v", filename, err)
// 		}
// 		err = os.WriteFile(fullPath, []byte(content), 0644)
// 		if err != nil {
// 			t.Fatalf("Failed to create test page %s: %v", filename, err)
// 		}
// 	}

// 	router, err := GenerateRoutes(tempDir)
// 	if err != nil {
// 		t.Fatalf("Failed to generate routes: %v", err)
// 	}

// 	tests := []struct {
// 		name           string
// 		path           string
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{"Home Page", "/", http.StatusOK, "Handler Index for /"},
// 		{"About Page", "/about", http.StatusOK, "Handler About for /about"},
// 		{"User Details", "/users/123", http.StatusOK, "Handler UserDetails for /users/[id]"},
// 		{"Not Found", "/nonexistent", http.StatusNotFound, "404 page not found"},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req, err := http.NewRequest("GET", tt.path, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()
// 			router.ServeHTTP(rr, req)

// 			if status := rr.Code; status != tt.expectedStatus {
// 				t.Errorf("handler returned wrong status code: got %v want %v",
// 					status, tt.expectedStatus)
// 			}

// 			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
// 				t.Errorf("handler returned unexpected body: got %v want %v",
// 					rr.Body.String(), tt.expectedBody)
// 			}
// 		})
// 	}
// }