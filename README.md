
# Сервис авторизации и регистрации

## Запуск приложения

Чтобы запустить приложение:
- Создайте `.env` файл, в котором укажите переменные окружения `DB_PASSWORD` и `JWT_KEY`.
- Запустите композицию командами: `docker-compose build`, `docker-compose up`.

## Требования

- go 1.18.2
- docker & docker-compose

## Роутинг
- /auth/Login
- /auth/Register