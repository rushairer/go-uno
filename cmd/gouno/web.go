package gouno

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/go-uno/internal/core"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

func init() {
	rootCmd.AddCommand(webCmd)

	webCmd.Flags().StringP("config", "c", "./config/config.yaml", "config file path")
	webCmd.Flags().StringP("address", "a", "0.0.0.0", "address to listen on")
	webCmd.Flags().StringP("port", "p", "8080", "port to listen on")
	webCmd.Flags().BoolP("debug", "d", false, "debug mode")
}

func startWebServer(cmd *cobra.Command, args []string) {
	log.Printf("starting web server...")

	err := core.InitConfig(cmd.Flag("config").Value.String())
	if err != nil {
		log.Fatalf("init config failed, err: %v", err)
	}

	core.GlobalConfig.WebServerConfig.Address = cmd.Flag("address").Value.String()
	core.GlobalConfig.WebServerConfig.Port = cmd.Flag("port").Value.String()
	core.GlobalConfig.WebServerConfig.Debug, _ = cmd.Flags().GetBool("debug")

	if core.GlobalConfig.WebServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := gin.Default()

	testGroup := server.Group("/test")
	{
		testGroup.GET(
			"/alive",
			func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", core.GlobalConfig.WebServerConfig.Address, core.GlobalConfig.WebServerConfig.Port),
		Handler: server,
	}

	log.Printf("listening on %s", httpServer.Addr)
	log.Printf("press Ctrl+C to exit")

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Printf("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	// Close

	log.Printf("server exiting")
}
