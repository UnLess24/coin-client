package main

import (
	"os"

	"github.com/UnLess24/coin/client/config"
	"github.com/UnLess24/coin/client/pkg/migrateprocess"
)

func main() {
	cfg := config.MustRead()
	migrateprocess.MustProcess(os.Args[1:], cfg)
}
