package main

import (
	"fmt"
	"os"

	"github.com/moond0wner/archiver-go/compress"
	"github.com/moond0wner/archiver-go/decompress"
	"github.com/moond0wner/archiver-go/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\nУкажите команду. Используйте 'help' для справки.")
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "compress":
		if len(os.Args) < 4 {
			fmt.Println("Использование: archiver compress <входной> <выходной>")
			return
		}
		if err := compress.File(os.Args[2], os.Args[3]); err != nil {
			fmt.Println("Ошибка сжатия:", err)
		}

	case "decompress":
		if len(os.Args) < 4 {
			fmt.Println("Использование: archiver decompress <входной> <выходной>")
			return
		}
		if err := decompress.File(os.Args[2], os.Args[3]); err != nil {
			fmt.Println("Ошибка распаковки:", err)
		}

	case "gen-txt":
		if len(os.Args) < 3 {
			fmt.Println("Использование: archiver gen-txt <имя_файла>")
			return
		}
		if err := utils.GenerateTXT(os.Args[2]); err != nil {
			fmt.Println("Ошибка создания txt:", err)
		} else {
			fmt.Printf("Создан файл: %s.txt\n", os.Args[2])
		}

	case "gen-bmp":
		if len(os.Args) < 3 {
			fmt.Println("Использование: archiver gen-bmp <имя_файла>")
			return
		}
		if err := utils.GenerateBMP(os.Args[2]); err != nil {
			fmt.Println("Ошибка создания bmp:", err)
		} else {
			fmt.Printf("Создан файл: %s.bmp\n", os.Args[2])
		}

	case "help":
		utils.ShowHelp()

	default:
		fmt.Printf("Неизвестная команда: %s\n", cmd)
	}
}
