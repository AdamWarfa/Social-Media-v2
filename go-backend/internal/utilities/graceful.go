package utilities

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func AwaitSignal(app *fiber.App) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		fmt.Printf("Received signal: %v\n", sig)
		shutdown(app)
	}
}

func shutdown(app *fiber.App) {
	fmt.Println("Shutting down gracefully...")

	// Perform cleanup operations here, if needed

	if err := app.Shutdown(); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)

	}
}
