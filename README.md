




# 测试
```sh

ssh dd
kubectl get po -owide -nhwx1166232
curl 10.0.230.172/waitmill/111
curl 10.0.230.158/waitmill/3
curl 10.0.230.158/wait/3


curl localhost/long/16
curl localhost/long/32
curl localhost

curl localhost/waitmill/1 # 等100毫秒
curl localhost/wait/1 # 等1秒

```

# 测试镜像 swr.cn-north-4.myhuaweicloud.com/hfbbg4/longconn:v0.1


# 构建


```sh

cd src

Regin=swr.cn-north-4.myhuaweicloud.com
Org=hfbbg4
AppName=longconn
Version=v0.2

docker build --build-arg TARGETARCH=amd64 --platform linux/amd64 -t ${Regin}/${Org}/${AppName}:${Version} --no-cache .

docker push ${Regin}/${Org}/${AppName}:${Version}



```


