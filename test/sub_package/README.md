##### 1.引入代码包
```shell
go get github.com/OneHundred86/golang/test/sub_package
```

##### 2.使用代码包
```go
import "github.com/OneHundred86/golang/test/sub_package"

sum := sub.Sum(1, 2, 3)
fmt.Println(sum)
```
