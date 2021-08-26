# Devops

### cmd and mian.go

cobra project

### example

` Algorithm` algorithm example

#### coroutine example

` Algorithm` algorithm example

### NetPing

project instruction ping written in go

```bash
Ping once per minute (ping 10 packets, interval 0.5, recording delay and packet loss)
 ```

#### txt

Defines a map of type interface{}

```go
var ipInfo = make(map[string]interface{})
```

把数据保存在txt文件中，

##### txt_v2

Defines a structure

```go
type Packets struct {
Time     string        `json:"time"`
Addr     string        `json:"addr"`
Ip       *net.IPAddr   `json:"ip""`
Sent     int           `json:"sent"`
Received int           `json:"received"`
Loss     float64       `json:"loss"`
Avg      time.Duration `json:"avg""`
}
```

