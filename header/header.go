package header

import (
	"encoding/binary"
	"errors"
	"os"
)

// Header содержит служебную информацию в начале архива.
// Записывается в файл перед сжатыми данными.
type Header struct {
	Magic          [4]byte // сигнатура архива
	OriginalSize   int64
	CompressedSize int64
}

func NewHeader(
	magic [4]byte,
	originalSize int64,
	compressedSize int64,
) *Header {
	return &Header{
		Magic:          magic,
		OriginalSize:   originalSize,
		CompressedSize: compressedSize,
	}
}

func (h *Header) Read(f *os.File) error {
	if err := binary.Read(f, binary.LittleEndian, &h.Magic); err != nil {
		return err
	}
	if err := binary.Read(f, binary.LittleEndian, &h.OriginalSize); err != nil {
		return err
	}
	if err := binary.Read(f, binary.LittleEndian, &h.CompressedSize); err != nil {
		return err
	}
	return nil

}
func (h *Header) Write(f *os.File) error {
	if err := binary.Write(f, binary.LittleEndian, h.Magic); err != nil {
		return err
	}
	if err := binary.Write(f, binary.LittleEndian, h.OriginalSize); err != nil {
		return err
	}
	return binary.Write(f, binary.LittleEndian, h.CompressedSize)
}

func (h *Header) Validate() error {
	if string(h.Magic[:]) != "ARC4" {
		return errors.New("Неизвестный формат архива")
	}

	return nil
}
