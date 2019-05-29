### 功能
用`go-kit`实现了最简单的get、post请求

### 目录
```
+-- main.go                 // 程序入口
|
+-- transports.go           // 请求路由层
|
+-- endpoints.go            // 服务路由层
|
+-- service.go              // 服务层
```

### 编译运行
首先进入项目目录，然后执行
```
go build ./
go run gokit-simple-restful-golang
```

### 请求
```
// GET 获取默认结果
curl http://localhost:8080/calculate/result -v

// POST 计算1+1，并返回结果
curl http://localhost:8080/calculate/Add/1/1 -X POST -v
```