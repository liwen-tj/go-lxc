## Registry
所有注册的代码的集合，代码仓库。
当Worker节点没有需要的handler时，需要从这取

## worker是一个process，只能有一个worker，默认监听5000端口
curl -X POST localhost:5000/run/echo -d '{"hello": "world"}'
响应处理代码在src/server/lambdaServer.go里面，（LambdaServer is a worker server that listens to run lambda requests and forward these requests to its sandboxes.）
通过解析URL获取lambda的名字以后，
借助/src/lambda/lambda.go里面的lambdaMgr（LambdaMgr provides thread-safe getting of lambda functions and collects all
lambda subsystems (resource pullers and sandbox pools) in one place）数据结构

Get在src/lambda/lambda.go line 144
Invoke在src/lambda/lambda.go line 207

## python代码
应该在要执行的python代码的开始说明需要哪些包，如：
- ol-install: parso,jedi,idna,chardet,certifi,requests
- ol-import: parso,jedi,idna,chardet,certifi,requests,urllib3

** python代码的event是dict类型，返回值需要能够被JSON serializable **

## python代码
目前没有依赖包的sock可以跑起来，但是sock没法安装依赖，docker模式下可以安装，但是不太对 （numpy说没有array选项。。）



