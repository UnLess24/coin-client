package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/UnLess24/coin/client/config"
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

	srv := server.New(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port))

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
