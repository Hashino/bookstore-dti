package main

import (
	"bookstore/db"
	"bookstore/ui"
	"fmt"
)

func main() {
	if err := db.InitDB(); err != nil {
		fmt.Println("Erro ao iniciar banco:", err)
		return
	}
	defer db.DB.Close()

	ui.RunMenu()
}
