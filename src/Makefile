Regin=swr.cn-north-4.myhuaweicloud.com
Org=hfbbg4
AppName=longconn
Version=v0.5

all:
	docker build --build-arg TARGETARCH=amd64 --platform linux/amd64 -t ${Regin}/${Org}/${AppName}:${Version} --no-cache .
	echo "docker push ${Regin}/${Org}/${AppName}:${Version}"
	-kubectl delete -f test1.yaml -n hwx1166232
	sleep 5;
	-kubectl apply -f test1.yaml -n hwx1166232

test1:
	echo "docker push ${Regin}/${Org}/${AppName}:${Version}"
