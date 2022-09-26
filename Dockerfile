FROM golang:1.18 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o server .

CMD ["./server"]
