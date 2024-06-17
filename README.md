




# 测试
```sh

curl localhost/long/16
curl localhost/long/32
curl localhost
```

# 测试镜像 swr.cn-north-4.myhuaweicloud.com/hfbbg4/longconn:v0.1


# 构建


```sh

cd src

Regin=swr.cn-north-4.myhuaweicloud.com
Org=hfbbg4
AppName=longconn
Version=v0.1

docker build --build-arg TARGETARCH=amd64 --platform linux/amd64 -t ${Regin}/${Org}/${AppName}:${Version} --no-cache .

docker push ${Regin}/${Org}/${AppName}:${Version}



```


