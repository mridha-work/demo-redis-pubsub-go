# Demo Redis Pub-Sub in Golang

This repository is a demo on how to implement a very simple [redis pub/sub](https://redis.io/docs/latest/develop/pubsub/) using in golang.<br/>
based on: <https://medium.com/@schalla024/implementing-and-understanding-redis-pub-sub-with-go-3bfdaa7ff1a4>

## Prerequisites

- this repo uses go 1.27.1, so go version >= 1.27.1 need to be installed
- docker & docker-compose
- redis-cli

## Setup

1. set up redis in docker using this docker-compose command

```bash
docker-compose up -d
```

2. run the publisher in one terminal

```bash
make run-pub
```

3. open another terminal and run the subscriber

```bash
make run-sub
```

4. validate, input text in the publisher terminal and check the sent message in the subscriber terminal.<br/>
we can also check the channels using `redis-cli`

```bash
redis-cli -h localhost -p 6379

localhost:6379> PUBSUB CHANNELS *
```

5. stop docker when done testing

```bash
docker-compose down
```
