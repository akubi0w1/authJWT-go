# auth-jwt

## overview

jwtを発行することでセッションを管理する。発行されたトークンはリクエストヘッダーに付加して利用する。
UIはなく、HTTPリクエストを受け付ける形で確認を行う。

## feature

- ユーザの作成
- 認証
- ログインユーザの確認

## how to start

```
// DBの起動
$ docker-compose up -d

// serverの起動
$ go run main.go
```

サンプルリクエスト
```
// signup
$ curl --request POST \
  --url http://localhost:8080/signup \
  --header 'Content-Type: application/json' \
  --data '{
	"id": "akubi",
	"name": "あくび",
	"password": "password"
}'

// login
$ curl --request POST \
  --url http://localhost:8080/login \
  --header 'Content-Type: application/json' \
  --data '{
	"id": "akubi",
	"password": "password"
}'

// who am i
curl --request GET \
  --url http://localhost:8080/whoami \
  --header 'Authorization: Bearer AUTHORIZED_TOKEN'
```