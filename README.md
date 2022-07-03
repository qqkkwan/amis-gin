# amis-gin
一个前端只写json，结合golang后端的快速搭建后台的项目

## 相关技术
ui: [amis](https://github.com/baidu/amis),

后端：[gin](https://github.com/gin-gonic/gin),

资源一键打包：embed,

热加载：[air](https://github.com/cosmtrek/air)

## 使用
```golang
git clone https://github.com/qqkkwan/amis-gin.git 
go run main.go
```

使用热加载：
```golang
go install github.com/cosmtrek/air@latest 
cd /path/to/your_project
air
```


## 打包
静态资源全部打包
```golang
go build main.go
```


