package main

import (
	"bookstore/db"
	"bookstore/logger"
	"bookstore/ui"
	"fmt"
)

func main() {
	if err := logger.Init(); err != nil {
		fmt.Println("Erro ao iniciar logger:", err)
		return
	}
	defer logger.Close()

	if err := db.InitDB(); err != nil {
		logger.LogError(err)
		fmt.Println("Erro ao iniciar banco de dados:", err)
		return
	}
	defer db.DB.Close()

	ui.RunMenu()
}
