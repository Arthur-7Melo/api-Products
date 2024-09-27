FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

RUN CGO_ENABLED=0 GOOS=linux go build -o apiExecutavel main.go

CMD ["./apiExecutavel"]