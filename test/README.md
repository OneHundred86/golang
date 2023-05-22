##### 1.引入代码包
```shell
go get -u github.com/OneHundred86/golang/test
```

##### 2.使用代码包
```go
import (
    "github.com/OneHundred86/golang/test"
    "github.com/OneHundred86/golang/test/aa"
)

test.Print("hello world")
aa.Show("abcd")
// 不能使用bb模块，因为bb模块变成了独立的子模块，需要独立go get引入
// bb.Sum(1,2)
```
