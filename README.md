# RedSync example

Tried Redsync example.

## Dependencies

- Go 1.12

- [Redsync](https://github.com/go-redsync/redsync)

## Desctiption

![figure1](png/figure1.png)

Same service is running on Server A and Server B

Service operates as follows:

1. Listen request from client
2. Get lock by redsync library
3. Release lock by redsync library
4. Return response to client

The same key is used for locking regardless of Server

## Reference

### Redis Docs
[English](https://redis.io/topics/distlock)

[Japanese](https://redis-documentasion-japanese.readthedocs.io/ja/latest/topics/distlock.html)

### Others
[Distributed Locks using Golang and Redis](https://kylewbanks.com/blog/distributed-locks-using-golang-and-redis)
