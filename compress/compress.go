package compress

import (
	"fmt"
	"os"

	"github.com/moond0wner/archiver-go/header"
	rle_algorithm "github.com/moond0wner/archiver-go/rle"
)

func File(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("чтение: %w", err)
	}

	compressed := rle_algorithm.Compress(data)

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("создание: %w", err)
	}
	defer out.Close()

	signature := [4]byte{'A', 'R', 'C', '4'}

	h := header.NewHeader(
		signature,
		int64(len(data)),
		int64(len(compressed)),
	)

	if err := h.Write(out); err != nil {
		return fmt.Errorf("запись заголовка: %w", err)
	}
	if _, err := out.Write(compressed); err != nil {
		return fmt.Errorf("запись данных: %w", err)
	}

	fmt.Printf("Сжато: %d -> %d байт\n", len(data), len(compressed))
	return nil
}
