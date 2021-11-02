FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main