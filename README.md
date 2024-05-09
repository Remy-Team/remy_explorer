# Remy Explorer

Remy Explorer - это микросервис на языке Go, предназначенный для работы с файлами и папками в облачном хранилище.

## Основные возможности

- Создание папок
- Получение информации о папках
- Получение списка папок по ID родительской папки
- Обновление информации о папках
- Удаление папкок
- Создание файлов
- Получение информации о файлах
- Получение списка файлов по ID папки
- Обновление информации о файлах
- Удаление файлов

## Установка

Для установки проекта вам потребуется Go версии 1.16 или выше. Склонируйте репозиторий и установите зависимости:

```bash
git clone https://github.com/yourusername/remy_explorer.git
cd remy_explorer
go mod download
```

## Запуск

Для запуска проекта используйте команду `go run`:

```bash
go run .
```

## Тестирование

Для запуска тестов используйте команду `go test`:

```bash
go test ./...
```

## Лицензия

Этот проект лицензирован под MIT License - подробности смотрите в файле [LICENSE.md](LICENSE.md).
