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
	coinserver "github.com/UnLess24/coin/client/internal/server/coin_server"
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
	defer func() { _ = db.Close() }()

	coinSrv, err := coinserver.New(cfg.CoinServer.Schema, cfg.CoinServer.Host, cfg.CoinServer.Port, cfg.CoinServer.Type)
	if err != nil {
		slog.Error("failed to create coin server", "error", err)
		return
	}
	defer func() { _ = coinSrv.Close() }()

	srv := server.New(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), db, coinSrv, cfg)
	defer func() { _ = srv.Close() }()

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
