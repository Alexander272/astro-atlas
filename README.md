## Запуск проекта:

Создайте файл .env в корневом каталоге и добавьте следующие значения:

    APP_ENV="dev"

    MONGO_URI="mongodb://127.0.0.1:27017"
    MONGO_USER="<you username>"
    MONGO_PASSWORD="<your password>"

    REDIS_PASSWORD="<your password>"

    HTTP_HOST="http://localhost"
    HTTP_DOMAIN="localhost"

    JWT_KEY="<secret key>"

    BCRYPT_MIN_COST=6
    BCRYPT_DEFAULT_COST=14
    BCRYPT_MAX_COST=30

Запуск докер командой `docker-compose up`
запуск dev версии `air -d`
