## Log_Monitor
暂未完成，稳定性未测试。

## Usage
```
go build main.go

export MONITOR_SCRIPTS=[scripts dir]
export MONITOR_TARGET=[tail target]
export MONITOR_HOOK=[golang plugin] 

./main

```
## Script
### 说明
1. vm 有自己的栈，脚本里可以调用 ```stack``` 来保存过往数据。
2. 脚本内需要标示 keywords

### 示例
```javascript
var keywords = "test,username"
if(!stack) stack = {}
    
console.log(line);

var result = "123"
```

## Hook
在设置过 ```MONITOR_HOOK``` 变量后，每次脚本返回的数据可以通过 hook 进行处理。    
处理 hook 的程序使用 golang plugin。
