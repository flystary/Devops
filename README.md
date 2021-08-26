# Devops

### cmd and mian.go

This is an example of the `cobra` project

Execute the following command in the current path to compile the project,you will appear a ly executable file.

```bash
go build -o ly  main.go
```

Use `ly --help`, you will be prompted how to use the ` ly` command

```bash
Show Current Time. For example:

With this command, you can view current time or customize the time format.

Usage:
  ly [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  show        Display the current time
  version     A brief description of your command

Flags:
  -h, --help   help for ly

Use "ly [command] --help" for more information about a command.
```

How to use `ly`

```bash
ly version     #显示当前ly版本
ly show        #展示当前时间  2021-08-26 13:24:25.3019767 +0800 CST m=+0.018521101
ly show parse -f  "2006/01/02 15:04:05"  ##将把当前时间格式化成"2006/01/02 15:04:05" 2021/08/26 13:26:55
```

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
Save the data in a txt file

##### main.go
Defines a map of type interface{}

```go
var ipInfo = make(map[string]interface{})

ipInfo["Addr"] = pinger.Addr()
ipInfo["Ip"] = pinger.IPAddr()
ipInfo["Sent"] = pinger.PacketsSent
ipInfo["Received"] = pinger.PacketsRecv
ipInfo["Loss"] = stats.PacketLoss
ipInfo["Avg"] = stats.AvgRtt
ipInfo["Time"] = time.Now().Format("2006/01/02 15:04:05")
```

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


ipInfo = Packets{
Time:     time.Now().Format("2006/01/02 15:04:05"),
Addr:     pinger.Addr(),
Ip:       pinger.IPAddr(),
Sent:     pinger.PacketsSent,
Received: pinger.PacketsRecv,
Loss:     stats.PacketLoss,
Avg:      stats.AvgRtt,
}
```

