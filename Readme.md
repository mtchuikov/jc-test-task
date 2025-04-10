# Тестовое задание для JavaCode.

## Запуск программы

Программа может быть запущена при помощи команды `docker compose up`. Эта команда запустит в Docker базу данных Postgres, создаст в ней необходимые таблицы, функции и т.п. при помощи инструмента [migrate](https://github.com/golang-migrate/migrate).

## Конфигурация

Программа может быть сконфигурирована при помощи флагов командной строки или переменных окружения

| Flag | | Env | Default value | Description |
|------|-----------|---------------|-------------|
| `--server.addr` | ` `| `SERVER_ADDRESS` | `127.0.0.1:8080` | IP address and port for the server to listen on |
| `--db.conn.url` | ` ` | `DB_CONN_URL` | `postgres://username:password@postgres:5432/postgres?sslmode=disable` | URL to connect to the Postgres database |
| `--help` | ` ` | `-h` | Print information about supported flags |
