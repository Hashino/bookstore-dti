package ui

import (
	"bookstore/model"
	"bookstore/repo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func RunMenu() {
	for {
		fmt.Println()
		fmt.Println("Sistema de Gerenciamento de Livros")
		fmt.Println("1. Adicionar Livro")
		fmt.Println("2. Listar Livros")
		fmt.Println("3. Buscar Livro")
		fmt.Println("4. Atualizar Livro")
		fmt.Println("5. Excluir Livro")
		fmt.Println("6. Sair")
		fmt.Print("Digite a opção: ")

		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			book := readBookInput()
			err := repo.AddBook(book)
			if err != nil {
				fmt.Println("Erro ao adicionar livro:", err)
			} else {
				fmt.Println("Livro adicionado com sucesso.")
			}
		case 2:
			books, err := repo.ListBooks()
			if err != nil {
				fmt.Println("Erro:", err)
				break
			}
			fmt.Println()
			fmt.Println("Lista de Livros:")
			for _, b := range books {
				fmt.Printf("%d. %s de %s, lançado em %d\n", b.ID, b.Title, b.Author, b.ReleaseDate.Year())
			}
		case 3:
			id := readIDInput("Digite o ID do livro: ")
			book, err := repo.GetBookByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println()
				fmt.Printf("ID: %d\n", book.ID)
				fmt.Printf("Título: %s\n", book.Title)
				fmt.Printf("Autor: %s\n", book.Author)
				fmt.Printf("Data de Lançamento: %s\n", book.ReleaseDate.Format("2006-01-02"))
				fmt.Printf("Nota: %.1f\n", book.Score)
				fmt.Printf("Resumo: %s\n", book.Summary)
			}
		case 4:
			id := readIDInput("Digite o ID do livro: ")
			book, err := repo.GetBookByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
				break
			}
			updated := readUpdatedBookInput(book)
			repo.UpdateBook(updated)
			fmt.Println("Livro atualizado.")
		case 5:
			id := readIDInput("ID do livro: ")
			ok, _ := repo.DeleteBook(id)
			if ok {
				fmt.Println("Livro excluído.")
			} else {
				fmt.Println("Livro não encontrado.")
			}
		case 6:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func readBookInput() model.Book {
	fmt.Println()
	fmt.Println("Adicionar Novo Livro")
	fmt.Print("Título: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Autor: ")
	scanner.Scan()
	author := scanner.Text()

	fmt.Print("Data de Lançamento (AAAA-MM-DD): ")
	scanner.Scan()
	dateStr := scanner.Text()
	releaseDate, _ := time.Parse("2006-01-02", dateStr)

	fmt.Print("Nota (0.0 - 10.0): ")
	scanner.Scan()
	score, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Resumo: ")
	scanner.Scan()
	summary := scanner.Text()

	return model.Book{
		Title:       title,
		Author:      author,
		ReleaseDate: releaseDate,
		Score:       score,
		Summary:     summary,
	}
}

func readUpdatedBookInput(old model.Book) model.Book {
	fmt.Printf("Título [%s]: ", old.Title)
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = old.Title
	}
	fmt.Printf("Autor [%s]: ", old.Author)
	scanner.Scan()
	author := scanner.Text()
	if author == "" {
		author = old.Author
	}
	fmt.Printf("Data [%s]: ", old.ReleaseDate.Format("2006-01-02"))
	scanner.Scan()
	dateStr := scanner.Text()
	var releaseDate time.Time
	if dateStr == "" {
		releaseDate = old.ReleaseDate
	} else {
		releaseDate, _ = time.Parse("2006-01-02", dateStr)
	}
	fmt.Printf("Nota [%.1f]: ", old.Score)
	scanner.Scan()
	scoreStr := scanner.Text()
	score := old.Score
	if scoreStr != "" {
		score, _ = strconv.ParseFloat(scoreStr, 64)
	}
	fmt.Printf("Resumo [%s]: ", old.Summary)
	scanner.Scan()
	summary := scanner.Text()
	if summary == "" {
		summary = old.Summary
	}
	old.Title = title
	old.Author = author
	old.ReleaseDate = releaseDate
	old.Score = score
	old.Summary = summary
	return old
}

func readIDInput(prompt string) int {
	fmt.Print(prompt)
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	return id
}
