FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
# COPY main.go ./
# COPY env ./
RUN go mod tidy
COPY . .
EXPOSE 1323
# CMD ["go", "run", "main.go"]