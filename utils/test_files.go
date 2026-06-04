package utils

import (
	"fmt"
	"os"
)

func GenerateTXT(filename string) error {
	content := []byte("AAAAABBBBBCCCCCDDDDD")
	if err := os.WriteFile(filename+".txt", content, 0644); err != nil {
		return fmt.Errorf("ошибка генерации тестового файла: %v", err)
	}
	fmt.Println("Тестовый", filename, ".txt успешно создан")
	return nil
}

func GenerateBMP(filename string) error {
	data := make([]byte, 54+4*4*3)
	for i := range data {
		data[i] = 255
	}
	if err := os.WriteFile(filename+".bmp", data, 0644); err != nil {
		return fmt.Errorf("ошибка генерации тестового файла: %v", err)
	}
	fmt.Println("Тестовый", filename, ".bmp успешно создан")
	return nil
}
