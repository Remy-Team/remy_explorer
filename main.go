package remy_explorer

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
	//Init database client
	pool, err := repo.NewClient(context.TODO(), cfg.Storage, 5)
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
		errs <- fmt.Errorf("%s", <-c)
	}()
	level.Info(logger).Log("message", "Service is ready to listen and serve", "type", cfg.Listen.Type, "bind_ip", cfg.Listen.BindIP, "port", cfg.Listen.Port)

	endpoints := handler.MakeEndpoints(fileSvc, folderSvc)

	go func() {
		fmt.Println("Listening on", cfg.Listen)
		httpHandler := handler.NewHTTPServer(endpoints)
		errs <- http.ListenAndServe(cfg.Listen.BindIP+":"+cfg.Listen.Port, httpHandler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
