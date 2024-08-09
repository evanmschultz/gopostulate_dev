// pkg/server/server.go

package server

// import (
// 	"github.com/evanmschultz/gopostulate_dev/pkg/router"
// 	"github.com/evanmschultz/gopostulate_dev/pkg/template"
// 	"github.com/gin-gonic/gin"
// )

// type Server struct {
// 	Engine *gin.Engine
// 	Router *router.Router
// 	TemplateEngine *template.Engine
// }

// type Config struct {
// 	PagesDir     string
// 	TemplatesDir string
// 	StaticDir    string
// }

// func NewServer(config Config) (*Server, error) {
// 	r := gin.Default()

// 	// Setup custom router
// 	customRouter, err := router.CreateRouter(config.PagesDir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Setup template engine
// 	templateEngine, err := template.NewEngine(config.TemplatesDir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Setup static file serving
// 	r.Static("/public", config.StaticDir)

// 	// Setup routes
// 	customRouter.SetupRoutes(r)

// 	server := &Server{
// 		Engine: r,
// 		Router: customRouter,
// 		TemplateEngine: templateEngine,
// 	}

// 	return server, nil
// }

// func (s *Server) Start(addr string) error {
// 	return s.Engine.Run(addr)
// }