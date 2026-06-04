```markdown
# Archiver

Консольный архиватор с использованием RLE (Run‑Length Encoding) на Go.
Работает как скрипт, можно запускать через `go run` или как собранный исполняемый файл.

## Возможности
- Сжатие файлов (.bmp, .pcx, .tga, .ico, .cur, .txt (с повторами), .log, .csv (с повторами), .raw, .pbm, .pgm, .ppm)
- Распаковка обратно без потерь
- Создание тестовых файлов для проверки
- Собственный формат архива с заголовком

## Формат архива
[Заголовок] [Сжатые данные]

Заголовок:
- `ARC4` (4 байта) — сигнатура
- Исходный размер (8 байт)
- Сжатый размер (8 байт)

Все данные записываются в формате Little‑Endian.

## Установка

```bash
git clone https://github.com/moond0wner/archiver-go.git
cd archiver
```

Windows:
```bash
go build -o archiver.exe
```

Linux:
```bash
go build -o archiver
```

## Использование

### Через собранный бинарник (после `go build`)
```bash
# Справка
archiver help

# Создать тестовые файлы
archiver gen-txt test         # создаст test.txt
archiver gen-bmp test         # создаст test.bmp

# Сжатие
archiver compress input.txt output.arc
archiver compress image.bmp output.arc

# Распаковка
archiver decompress output.arc restored.txt
archiver decompress output.arc restored.bmp
```

## Через go run (без сборки)

```bash
go run main.go help
go run main.go gen-txt test
go run main.go compress input.txt output.arc
go run main.go decompress output.arc restored.txt
```

## Структура проекта

```
archiver-go/
├── compress
│   └── compress.go
├── decompress
│   └── decompress.go
├── header
│   └── header.go
├── help
│   └── help.go
├── rle
│   └── rle.go
├── utils
│   └── test_files.go
├── .gitignore
├── archive.exe
├── main.go
├── README.md
└── go.mod
```
## Технологии
```
- Go 1.21+
- Только стандартная библиотека
```
## Выполнил
```
- Озернов Максим Б-ПИ-101
```
