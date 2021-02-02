# Log_Monitor
暂未完成，稳定性未测试。缺少与脚本的交互。

# Usage
```
go build main.go

./main [script dir]
```
# Script
## 说明
1. vm 有自己的栈，脚本里可以调用 ```stack``` 来保存过往数据。
2. 脚本内需要标示 keywords

## 示例
```javascript
var keywords = "test,username"
if(!stack) stack = {}
    
console.log(line);

```
