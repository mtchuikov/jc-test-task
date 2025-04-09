# Тестовое задание для JavaCode.

## Запуск программы

Программа может быть запущена при помощи команды `docker compose up`. Эта команда запустит в Docker базу данных Postgres, а также экземпляр приложения. Кроме того, перед началом работы также необходимо выполнить создание необходимых схем, функций и т.п. в базе данных. Это можно сделать при помощи команды: `TODO`

## Конфигурация

Программа может быть сконфигурирована при помощи флагов командной строки или переменных окружения

| Flag | Shorthand | Default value | Description |
| `--server-addr`| `-a` | `SERVER_ADDRESS` | `127.0.0.1:8080` | Specify the IP address and port for the server to listen on |
| `--dsn` | `-d` | `DB_CONN_URL` | |  Define the database DSN (Data Source Name) used to connect to the database |

| Flag | Shorthand | Default value | Description |
|------|-----------|---------------|-------------|
| `--server-addr`| `-a` | `SERVER_ADDRESS` | `127.0.0.1:8080` | Specify the IP address and port for the server to listen on |
| `--dsn` | `-d` | `DB_CONN_URL` | `postgres://username:password@postgres:5432/postgres?sslmode=disable` |  Define the database DSN (Data Source Name) used to connect to the database |
| `--help` | `-h` | | Print information about supported flags |
