package decompress

import (
	"fmt"
	"os"

	"github.com/moond0wner/archiver-go/header"
	rle_algorithm "github.com/moond0wner/archiver-go/rle"
)

func File(inputPath, outputPath string) error {
	in, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("открытие: %w", err)
	}
	defer in.Close()

	var h header.Header
	if err := h.Read(in); err != nil {
		return fmt.Errorf("чтение заголовка: %w", err)
	}

	if err := h.Validate(); err != nil {
		return fmt.Errorf("валидация: %w", err)
	}

	compressed := make([]byte, h.CompressedSize)
	if _, err := in.Read(compressed); err != nil {
		return fmt.Errorf("чтение данных: %w", err)
	}

	decompressed := rle_algorithm.Decompress(compressed)

	if int64(len(decompressed)) != h.OriginalSize {
		return fmt.Errorf("размер восстановленных данных не совпадает")
	}

	if err := os.WriteFile(outputPath, decompressed, 0644); err != nil {
		return fmt.Errorf("запись: %w", err)
	}

	fmt.Printf("Распаковано: %d -> %d байт\n", len(compressed), len(decompressed))
	return nil
}
