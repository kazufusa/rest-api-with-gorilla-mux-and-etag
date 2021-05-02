# Gorilla mux and ETag

```sh
$ go run main.go &
[1] 3918

$ curl -i localhost:8080/api/clock -H "If-None-Match:84fd5a79b06b85c9e054ab7627150d5fa2086dfd"
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: application/json;charset=UTF-8
Etag: 4d1d31e08c7a31e07c024557311c5cde5b44713e
Date: Mon, 03 May 2021 12:47:12 GMT
Content-Length: 60

{"unix":1620046020,"string":"2021-05-03 21:47:00 +0900 JST"}%

$ curl -i localhost:8080/api/clock -H "If-None-Match:4d1d31e08c7a31e07c024557311c5cde5b44713e"
HTTP/1.1 304 Not Modified
Date: Mon, 03 May 2021 12:47:28 GMT
```
