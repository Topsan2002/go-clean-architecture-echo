##Docker setup

``` bash

docker pull postgres:alpine

docker run --name cockroachdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:alpine


docker exec -it cockroachdb bash


psql -U postgres



CREATE DATABASE cockroachdb;

```