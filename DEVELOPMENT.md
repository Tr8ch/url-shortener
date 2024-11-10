# DEVELOPMENT

Пререквизиты:
- [docker](https://www.docker.com/get-started/)

## Подготовка окружения

Запустите зависимости (Redis, итд) след.команой.

```sh
docker compose -f compose.yaml up -d
```

### Настроить переменные окружения

Скопируйте шаблонный `.env.example`.

```sh
cp ./cmd/api/.env.example ./cmd/api/.env
```

## Запуск Core API

Перед запуском ОБЯЗАТЕЛЬНО экспортните переменные окружения.

```sh
export $(cat ./cmd/api/.env)
```

Запустить API.

```sh
go run ./cmd/api
```