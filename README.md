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
