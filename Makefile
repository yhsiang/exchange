all: docker run-mysql run-exchange

test:
	go test ./...

tools:
	go get github.com/vektra/mockery/...
	go get github.com/gin-gonic/gin
	go get github.com/go-sql-driver/mysql
	go get github.com/google/uuid

docker: docker-mysql docker-exchange

mysql: docker-mysql run-mysql

docker-mysql:
	docker build --pull \
		--file Dockerfile.mysql \
		-t exchange-mysql:latest .

run-mysql:
	docker run --name exchange-mysql -e MYSQL_ROOT_PASSWORD=test -d exchange-mysql:latest


dev: docker-exchange run-exchange

docker-exchange:
	docker build --pull \
		--file Dockerfile \
		-t exchange:latest .

run-exchange:
	docker run --name exchange --link exchange-mysql:db -p 8080:8080 -e MYSQL_ROOT_PASSWORD=test -e DB_HOST=db -d exchange:latest