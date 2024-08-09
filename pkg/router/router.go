// pkg/router/router.go

package router

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/a-h/templ"

	// TODO: figure out how to import the pages and components packages from the user's project
	// "myapp/pages"
	// "myapp/components"
)

// Router handles the registration and management of routes in the application.
type Router struct {
	engine *gin.Engine
}

// New creates and returns a new Router instance.
func New() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

// RegisterRoute automatically registers a route based on the file name and location.
// It determines if it's a page or component, imports the corresponding Params struct,
// and sets up the route with the correct handler.
//
// Parameters:
// - filePath: The path to the .templ.go file, relative to the project root.
//
// Returns:
// - error: An error if registration fails, nil otherwise.
func (r *Router) RegisterRoute(filePath string) error {
	// Determine if it's a page or component
	isComponent := strings.HasPrefix(filePath, "components/")
	baseDir := "pages"
	if isComponent {
		baseDir = "components"
	}

	// Extract the route path and file name
	relPath := strings.TrimPrefix(filePath, baseDir+"/")
	relPath = strings.TrimSuffix(relPath, ".templ.go")
	routeParts := strings.Split(relPath, "/")
	fileName := routeParts[len(routeParts)-1]

	// Construct the type names
	paramsTypeName := fmt.Sprintf("%sParams", strings.Title(fileName))
	handlerTypeName := strings.Title(fileName)

	// Get the types
	paramsType, handler, err := getTypes(isComponent, paramsTypeName, handlerTypeName)
	if err != nil {
		return fmt.Errorf("failed to get types: %v", err)
	}

	// Construct the route path
	routePath := "/" + strings.Join(routeParts, "/")
	if fileName == "index" {
		routePath = "/" + strings.Join(routeParts[:len(routeParts)-1], "/")
	}

	// Add route parameters
	routePath = addRouteParams(routePath, paramsType)

	// Register the route
	r.engine.GET(routePath, func(c *gin.Context) {
		// Create and populate the params struct
		params := reflect.New(paramsType).Elem()
		for i := 0; i < paramsType.NumField(); i++ {
			field := paramsType.Field(i)
			tag := field.Tag.Get("route")
			if tag != "" {
				paramName := strings.TrimSuffix(tag, "?")
				params.Field(i).SetString(c.Param(paramName))
			}
		}

		// Call the handler
		result := reflect.ValueOf(handler).Call([]reflect.Value{params})[0].Interface()
		
		// Render the template
		component := result.(templ.Component)
		c.Header("Content-Type", "text/html")
		err := component.Render(c.Request.Context(), c.Writer)
		if err != nil {
			c.String(500, fmt.Sprintf("Failed to render template: %v", err))
		}
	})

	return nil
}

// addRouteParams adds route parameters to the given path based on the struct tags.
func addRouteParams(path string, paramsType reflect.Type) string {
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		tag := field.Tag.Get("route")
		if tag != "" {
			if strings.HasSuffix(tag, "?") {
				path += "/:" + strings.TrimSuffix(tag, "?") + "?"
			} else {
				path += "/:" + tag
			}
		}
	}
	return path
}

// getTypes retrieves the Params struct and handler function from the appropriate package.
//
// Parameters:
// - isComponent: Boolean indicating if this is a component (true) or a page (false).
// - paramsTypeName: The name of the Params struct (e.g., "AboutParams").
// - handlerTypeName: The name of the handler function (e.g., "About").
//
// Returns:
// - reflect.Type: The type of the Params struct.
// - interface{}: The handler function.
// - error: An error if retrieval fails, nil otherwise.
func getTypes(isComponent bool, paramsTypeName, handlerTypeName string) (reflect.Type, interface{}, error) {
	var pkg interface{}
	// if isComponent {
	// 	pkg = components
	// } else {
	// 	pkg = pages
	// }

	// Get the Params struct type
	paramsType := reflect.ValueOf(pkg).Elem().FieldByName(paramsTypeName).Type()
	if paramsType == reflect.TypeOf(nil) {
		return nil, nil, fmt.Errorf("params type %s not found", paramsTypeName)
	}

	// Get the handler function
	handler := reflect.ValueOf(pkg).MethodByName(handlerTypeName)
	if !handler.IsValid() {
		return nil, nil, fmt.Errorf("handler %s not found", handlerTypeName)
	}

	return paramsType, handler.Interface(), nil
}

// DiscoverAndRegisterRoutes automatically discovers and registers all routes
// in the pages and components directories.
//
// Returns:
// - error: An error if discovery or registration fails, nil otherwise.
func (r *Router) DiscoverAndRegisterRoutes() error {
	for _, dir := range []string{"pages", "components"} {
		err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, ".templ.go") {
				return r.RegisterRoute(path)
			}

			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to discover routes in %s: %v", dir, err)
		}
	}

	return nil
}

// Run starts the HTTP server and begins listening for requests.
//
// Parameters:
// - addr: The address and port to listen on (e.g., ":8080").
//
// Returns:
// - error: An error if the server fails to start or encounters an error while running.
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}