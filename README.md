## Log_Monitor
通过 tail 实现监听日志。较 ```fsnotify``` 方法缺少灵活性，但更稳定。    
只支持包含 ```tail``` 工具的平台。


## Usage
```
go build main.go

export MONITOR_SCRIPTS=脚本存放的目录
export MONITOR_TARGET=监听的对象
export MONITOR_HOOK=Golang Plugin 文件 

./main

```
## Script
### 说明
1. vm 有自己的栈，脚本里可以调用 ```stack``` 来保存过往数据。
2. 脚本内需要标示 ```keywords```

### 示例
```javascript
var keywords = "test,username"
if(!stack) stack = {}
    
console.log(line);

var result = "123"; // result 会返回给 hook 进行处理
```

### 脚本热更新
```
1. 拿到该程序 pid
2. kill -USR1 [pid]
```

## Hook
在设置过 ```MONITOR_HOOK``` 变量后，每次脚本返回的数据可以通过 hook 进行处理。    
处理 hook 的程序使用 golang plugin。
