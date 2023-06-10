PKG: github.com/Jonath-z/JLB-backend
BIN: JLB-backend

run-dev: 
	@docker-compose up

run-docker-build: 
	docker build -t jlb-backend	.
