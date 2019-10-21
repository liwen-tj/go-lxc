## 运行代码
- 执行deps.sh来安装依赖
- 执行build.sh来生成可执行文件

这个可执行文件./bin/olscheduler支持三种操作命令
- start，可以用于开启scheduler
- add，向一个已经运行的scheduler里面加入一个新的worker
- remove，向一个已经运行的scheduler移除worker

格式分别是:
```
olscheduler start [-c|--config=FILEPATH]
olscheduler worker add URL
olscheduler worker remove URL
```

## start命令启动scheduler
Scheduler is an object that can schedule lambda function workloads to a pool of workers.

worker只是一个URL...

对http请求做了一层封装



















type Balancer interface {
	SelectWorker(r *http.Request, l *lambda.Lambda) (url.URL, *httputil.HttpError)
	ReleaseWorker(workerUrl url.URL)
	AddWorker(workerUrl url.URL)
	RemoveWorker(workerUrl url.URL)
	GetAllWorkers() []url.URL
} // balancer只是一个起到选择worker的功能










