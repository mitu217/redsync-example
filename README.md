# RedSync example

Tried Redsync example.

## Dependencies

- Go 1.12

- [Redsync](https://github.com/go-redsync/redsync)

## Desctiption

ServerAとServerBでは同様のサービスが動作しており、そのサービスは次のような動作をする

1. リクエストの受信
2. ロックの取得
3. Verificationサービスへロック取得成功の通知
4. ランダムなsleep
5. Verificationサービスへロック解除開始の通知
6. ロックの開放
7. レスポンスの返却

ロックにはServerに関係なく同一のキーが使われる想定です

Verificationサービスでは上記の3,5の通知が同一サービスから連続して通知されることを期待する
もし、連続した通知が確認でいない場合エラーを返す

このサンプルではVerificationサービスが期待した動作をするかどうかを検証するとともに、高負荷時のサーバーおよびRedisの動作を確認する

## Reference

### Redis Docs
[English](https://redis.io/topics/distlock)

[Japanese](https://redis-documentasion-japanese.readthedocs.io/ja/latest/topics/distlock.html)

### Others
[Distributed Locks using Golang and Redis](https://kylewbanks.com/blog/distributed-locks-using-golang-and-redis)
