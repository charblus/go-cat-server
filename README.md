[go 微服务](https://segmentfault.com/a/1190000014895034)

### 负载测试
> 需要Java的运行环境以及需要安装Apache Maven。

brew cask install java
brew search maven
brew info maven
brew install maven

cd goblog/loadtest   
mvn gatling:execute -Dusers=1000 -Dduration=30 -DbaseUrl=http://localhost:6767

这样就会启动并运行测试。参数如下:

* users: 模拟测试的并发用户数.
* duration: 测试要运行的秒数.
* baseUrl: 我们要测试的服务的基础路径。当我们把它迁移到Docker Swarm后，baseUrl修改修改为Swarm的公共IP. 

> 首次运行，mvn会自动安装一大堆东西。安装完后，测试完成之后，它会将结果写到控制台窗口，同时也会产生一个报告到target/gatling/results中的html中。

[Go微服务 - 嵌入数据库和JSON](https://segmentfault.com/a/1190000014961252)

[author github codes](https://github.com/callistaenterprise/goblog/tree/P3/accountservice)

> p2 已添加读取数据

> p3 docker
### docker

1. 创建Dockerfile
```
FROM iron/base

EXPOSE 6767
ADD catservice-linux-amd64 /
ENTRYPOINT ["./catservice-linux-amd64"]     
```
> FROM: 定义我们将要开始构建我们自己映像来源的基本映像。 iron/base是非常适合运行Go应用程序的映像。
> EXPOSE: 定义我们希望在Docker网络内部暴露的可以到达的端口号。
> ADD: 添加一个文件accountservice-linux-amd64到容器文件系统的根目录/。
> ENTRYPOINT 定义当Docker启动这个映像容器时要运行的可执行文件,这里对应go build 生成的二进制文件

2. 在根目录（main.go所在目录）下 构建catservice-linux-amd64可执行文件。
```
export GOOS=linux
go build -o catservice-linux-amd64
export GOOS=darwin
```

3. 构建Docker映像。 在项目父级目录
```
cd ../
docker build -t oahcoay/catservice go-cat-server

```
> 构建Docker容器映像时，我们通常使用[prefix]/[name]的命名约定来对其名字打标签
> go-cat-server 为当前项目目录名

> p4 负载均衡