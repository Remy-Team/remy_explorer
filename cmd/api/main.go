package api

import (
	"github.com/go-kit/log"
	"os"
	"remy_explorer/internal/config"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("message", "Starting explorer service")
	cfg := config.GetConfig()
	logger.Log("message", "Configuration loaded", "config", cfg)

	logger.Log("message", "Connecting to the database")
}
