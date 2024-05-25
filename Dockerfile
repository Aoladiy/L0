FROM golang:1.22.3

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Обновляем файлы go.mod и go.sum
RUN go mod tidy

# Запускаем приложение
CMD ["go", "run", "./cmd/main.go"]
