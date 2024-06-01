package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"remy_explorer/internal/config"
	handler "remy_explorer/internal/explorer/handler/http"
	repo "remy_explorer/internal/explorer/repository/postgresql"
	"remy_explorer/internal/explorer/service/file"
	"remy_explorer/internal/explorer/service/folder"
	"syscall"
)

//	@title			Remy Explorer API
//	@version		0.0.1
//	@description	This is a file and folder explorer API
//	@BasePath		/
//	@schemes		http
//	@produce		json
//	@consumes		json

//	@contact.name	Remy Team
//	@contact.email	remystorage@yandex.ru
//	@license.name	MIT
//	@license.url	http://opensource.org/licenses/MIT

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "explorer", "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	level.Info(logger).Log("message", "Service started")
	defer level.Info(logger).Log("message", "Service ended")

	cfg := config.GetConfig(logger)

	// Корневой контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Init database client
	pool, err := repo.NewClient(ctx, cfg.Storage, 5)
	if err != nil {
		level.Error(logger).Log("message", "Failed to connect to the database", "error", err)
		return
	}
	defer pool.Close()
	flag.Parse()

	// Create file service
	var fileSvc file.FileService
	{
		rep := repo.NewFileRepo(pool, logger)
		fileSvc = file.NewService(rep, logger)
	}
	var folderSvc folder.FolderService
	{
		rep := repo.NewFolderRepo(pool, logger)
		folderSvc = folder.NewService(rep, logger)
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		level.Info(logger).Log("message", "Received signal", "signal", sig)
		cancel() // Cancel context if needed
		errs <- fmt.Errorf("received signal: %s", sig)
	}()
	level.Info(logger).Log("message", "Service is ready to listen and serve", "type", cfg.Listen.Type, "bind_ip", cfg.Listen.BindIP, "port", cfg.Listen.Port)

	endpoints := handler.MakeEndpoints(fileSvc, folderSvc)

	go func() {
		address := cfg.Listen.BindIP + ":" + cfg.Listen.Port
		level.Info(logger).Log("message", "HTTP server is starting", "address", address)
		httpHandler := handler.NewHTTPServer(logger, endpoints)
		errs <- http.ListenAndServe(address, httpHandler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
