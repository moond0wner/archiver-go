package rle_algorithm

// Принимает исходный файл в виде слайса байтов и возвращает слайс байт в формате [count, byte, count, byte, ...])
func Compress(data []byte) []byte {
	compressed := []byte{}

	for i := 0; i < len(data); i++ {
		count := 1
		for i+1 < len(data) && data[i] == data[i+1] && count < 255 {
			count++
			i++
		}
		compressed = append(compressed, byte(count), data[i])
	}
	return compressed
}

// Принимает сжатый файл в виде слайса байтов и возвращает слайс байтов с исходными данными
func Decompress(compressedData []byte) []byte {
	decompressed := []byte{}
	for i := 0; i < len(compressedData); i += 2 {
		count := compressedData[i]
		b := compressedData[i+1]
		for j := 0; j < int(count); j++ {
			decompressed = append(decompressed, b)
		}
	}
	return decompressed
}
