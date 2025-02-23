package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/UnLess24/coin/client/config"
	"github.com/UnLess24/coin/client/internal/database"
	"github.com/UnLess24/coin/client/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.MustRead()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer cancel()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
	}()

	db, err := database.NewPGDB(cfg)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return
	}
	defer db.Close()
	srv := server.New(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port), db, cfg)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return srv.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		slog.Error("exit reason", "message", err)
	}
}
