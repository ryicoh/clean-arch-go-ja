package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryicoh/clean-arch/internal/infrastructure/cmd/apiserver"
)

func main() {
	c := apiserver.NewContainer()
	if err := c.Build(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
