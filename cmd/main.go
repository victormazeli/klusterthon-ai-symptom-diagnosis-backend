package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/api/middlewares"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/api/routes"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "8080"

func main() {
	// Load config
	cfg := config.LoadEnvironmentConfig()

	// Initialize Database Connection
	err := mgm.SetDefaultConfig(nil, cfg.Database.Name, options.Client().ApplyURI(cfg.Database.URL))

	if err != nil {
		log.Fatal("Error occurred connecting to database!")
	}
	log.Print("Connected to database!")

	port := cfg.Server.Port
	if port == "" {
		port = defaultPort
	}

	// Initialize Gin router
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(middlewares.CORS())
	// Health
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is Running!")
	})
	// Not Found
	r.NoRoute(middlewares.NotFound())
	// Method Not Allowed
	r.NoMethod(middlewares.MethodNotAllowed())
	// Setup Route
	rootPath := r.Group("")
	routes.SetupRoute(cfg, nil, rootPath)

	// Setup server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gin in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)

		}
	}()

	quit := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
