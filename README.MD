## Docker setup

``` bash

$ docker pull postgres:alpine

$ docker run --name cockroachdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:alpine


$ docker exec -it cockroachdb bash


$ psql -U postgres



$ CREATE DATABASE cockroachdb;

```

[ref](https://medium.com/@rayato159/how-to-implement-clean-architecture-in-golang-87e9f2c8c5e4) 