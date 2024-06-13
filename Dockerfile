# Étape de construction
FROM golang:1.22-alpine

WORKDIR /test_http

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

# Exposer le port
EXPOSE 8080

CMD ["./main"]