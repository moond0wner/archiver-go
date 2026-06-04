```markdown
# Archiver

Консольный архиватор с использованием RLE (Run‑Length Encoding) на Go.

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
go build -o archiver.exe
```

## Использование

```bash
# Справка
archiver.exe help

# Создать тестовые файлы
archiver.exe gen-txt test         # создаст test.txt
archiver.exe gen-bmp test         # создаст test.bmp

# Сжатие
archiver.exe compress input.txt output.arc
archiver.exe compress image.bmp output.arc

# Распаковка
archiver.exe decompress output.arc restored.txt
archiver.exe decompress output.arc restored.bmp
```

## Пример

```bash
archiver.exe gen-txt sample
archiver.exe compress sample.txt sample.arc
archiver.exe decompress sample.arc restored.txt
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
