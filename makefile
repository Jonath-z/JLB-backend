PKG: github.com/Jonath-z/JLB-backend
BIN: JLB-backend

build: 
	@go build -o $(BIN) $(PKG)

run-prod: build
	@./bin/JLB-backend

run-dev: 
	@go run ./src/main.go

