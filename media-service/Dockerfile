FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o media-service 

EXPOSE 3000

CMD ["./media-service"]

