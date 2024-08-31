# Reverse proxy to proxy by path to multiple hosts according to a configuration file

🇯🇵 設定ファイルに従って複数ホストにパスでプロキシするリバースプロキシ

```yaml
# config.yaml
hosts:
  - path: tv
    host: http://192.168.1.104
```

```
simple-http-reverse-proxy --addr 0.0.0.0:8089 --config ./config.yml --static ./static/
```

When you access `http://localhost:8089/tv/sht31`, it will be proxied to `http://192.168.1.104/sht31`.

🇯🇵

http://localhost:8089/tv/sht31 にアクセスすると http://192.168.1.104/sht31 にプロキシされる

## install

```
go install github.com/74th/simple-http-reverse-proxy@latest
```