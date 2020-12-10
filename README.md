# godemo

>用go加速数据处理工作：
1. 协程

---

### 快速上手
  ```
  1. 下载go
  2. tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz
  3. 设置环境变量 ~/.bash_profile
  export GOROOT=/usr/local/go
  PATH="$GOROOT/bin:${PATH}"
  export PATH
  export GOPATH=/Users/roger/go
  export GOPROXY=https://mirrors.aliyun.com/goproxy/
  source ~/.bash_profile

  go mod init  # 初始化go.mod
  go mod tidy  # 更新依赖文件
  go mod download  # 下载依赖文件
  go mod vendor  # 将依赖转移至本地的vendor文件
  go mod edit  # 手动修改依赖文件
  go mod graph  # 打印依赖图
  go mod verify  # 校验依赖

  4. 安装依赖包
  sudo go get -u github.com/go-sql-driver/mysql
  sudo go get go.mongodb.org/mongo-driver/mongo
  ```

### 项目说明


### Reference
- [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)
   中文参考文档。
- [awesome-go](https://github.com/avelino/awesome-go)
  
- [mongodb driver](http://github.com/mongodb/mongo-go-driver)
  The MongoDB supported driver for Go.

- [TheAlgorithms for Go](https://github.com/TheAlgorithms/Go)
- [json-iterator for Go](https://github.com/json-iterator/go)
  go get github.com/json-iterator/go

