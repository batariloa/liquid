FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy only the Go module files first
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Now copy the rest of the source code
COPY . .

RUN go build -o search-service 

EXPOSE 8085

CMD ["./search-service"]

