# docker-compose 负载均衡demo

```sh
# 制作镜像(多段构建)
docker build -t go-host .
# 启动服务
docker-compose up --scale go-host=3 -d
```

访问
> http://localhost
