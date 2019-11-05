# Exchange demo

This is a demo for a simple exchange api, you can post order, and query orderbook.


## Exchange demo with Docker

If you have a docker environment, please use the command below.

```
$ make all
```

If encounter problem `2019/11/05 06:43:17 dial tcp 172.17.0.2:3306: connect: connection refused`

Simply restart exchange container will solve this problem.
For other solution, we can use bash script to make sure db connection and then start the server.

# Development

Recommend to setup golang environment and use commands below to run.

```
$ make tools
$ DB_HOST=localhost go run main.go
```

## Development with Docker

If you have a docker environment, please use the command below.

```
$ make mysql
$ make dev
```

# Testing

```
$ make test
```