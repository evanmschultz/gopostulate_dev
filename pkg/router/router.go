// pkg/router/router.go

package router

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pagesDir string
	routes   map[string]gin.HandlerFunc
}

func NewRouter(pagesDir string) (*Router, error) {
	r := &Router{
		pagesDir: pagesDir,
		routes:   make(map[string]gin.HandlerFunc),
	}

	err := r.generateRoutes()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Router) generateRoutes() error {
	return filepath.Walk(r.pagesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		route := filepath.ToSlash(path)
		route = route[len(r.pagesDir):]
		route = route[:len(route)-3] // remove .go extension

		// TODO: Generate handler based on file content
		r.routes[route] = func(c *gin.Context) {
			c.String(200, "Handler for "+route)
		}

		return nil
	})
}

func (r *Router) SetupRoutes(engine *gin.Engine) {
	for route, handler := range r.routes {
		engine.GET(route, handler)
	}
}