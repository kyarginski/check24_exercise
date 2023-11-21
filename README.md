CHECK24 exercise.
====

The application is a blog rest API for CHECK24.

Apps blog and admin can be run as service.

## Run app as CLI application

```shell
export BLOG_CONFIG_PATH=config/blog/prod.yaml
```

```shell
go run ./cmd/blog
```

```shell
export ADMIN_CONFIG_PATH=config/admin/prod.yaml
```

```shell
go run ./cmd/admin
```


## Run services as docker containers

```shell
docker-compose up -d
```

## Stop server with docker containers

```shell
docker-compose down
```


## Documentation of API (swagger)

See swagger documentation file [here](doc/swagger.json)
