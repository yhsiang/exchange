# Exchange demo


## Exchange demo with Docker

```
$ make all
```

If encounter problem `2019/11/05 06:43:17 dial tcp 172.17.0.2:3306: connect: connection refused`

Simply restart exchange container will solve this problem.
For other solution, we can use bash script to make sure db connection and then start the server.

# Development

$ make tools
$ DB_HOST=localhost go run main.go

## Development with Docker

```
$ make mysql
$ make dev
```