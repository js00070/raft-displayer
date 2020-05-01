# 说明

启动后，访问localhost:8080/index可看到页面

## 接口使用例子

启动3节点的raft
```bash
curl localhost:8080/api/startnodes?servers=3
```

获取编号为2的节点的状态（编号从0开始计算）
```bash
curl localhost:8080/api/getstate?number=2
```
返回值为
```bash
{"commitIndex":0,"lastApplied":0,"leaderId":-1,"logs":[{"Command":null,"Term":0}],"number":2,"state":0,"term":1,"votedCount":2,"votedFor":2}
```
字段含义具体见raft.go中Raft结构体

向编号为2的节点发送内容为101的command
```bash
curl "localhost:8080/api/startcommand?number=2&command=101"
```

断开编号为2的节点
```bash
curl localhost:8080/api/disconnect?number=2
```

重连编号为2的节点
```bash
curl localhost:8080/api/reconnect?number=2
```
