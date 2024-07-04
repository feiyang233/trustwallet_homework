# for local test API
```shell
go run main.go handler.go rpcClient.go
```

```shell
curl --location 'http://localhost:8080/getBlockNumber' \
--header 'Content-Type: application/json' \
--data '{
    "id": 28
}'

{"id":28,"jsonrpc":"2.0","result":"0x38389df"}


 curl --location --request GET 'http://localhost:8080/health_check' \
--header 'Content-Type: application/json' \
--data '{
    "id": 2
    "blockNumber":
}'

ok

curl --location 'http://localhost:8080/getBlockByNumber' \
--header 'Content-Type: application/json' \
--data '{
    "id": 2,
    "blockNumber": "0x134e82a"
}'
```

# for unit test
```shell
➜  src go test -v

=== RUN   TestGetBlockNumber
--- PASS: TestGetBlockNumber (0.00s)
=== RUN   TestGetBlockByNumber
--- PASS: TestGetBlockByNumber (0.00s)
=== RUN   TestHealthCheck
--- PASS: TestHealthCheck (0.00s)
PASS
ok      blockchain-client       0.315s
```

# for docker build
You can build and upload to your docker hub space
```shell
docker build --platform linux/arm64 -t feiyang233/proxy-client:arm64 .
docker build --platform linux/amd64 -t feiyang233/proxy-client:amd64 .

docker login

docker push feiyang233/proxy-client:arm64
docker push feiyang233/proxy-client:amd64
```
You can get this image from https://hub.docker.com/r/feiyang233/proxy-client/tags