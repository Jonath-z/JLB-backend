FROM golang:latest

WORKDIR /app

COPY . ./

RUN go install github.com/githubnemo/CompileDaemon
RUN go mod download

EXPOSE 4040

CMD go run main.go
#ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
