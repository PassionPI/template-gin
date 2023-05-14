package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Listen listens and serves HTTP requests with graceful shutdown.
func Listen(
	handler http.Handler,
	addr string,
	timeout time.Duration,
) {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
