# zassert
###### assert风格框架

## 获得zassert
`go get -u github.com/zlyuancn/zassert`

## 使用zassert

```go
package main
import "github.com/zlyuancn/zassert"

func main() {
    zassert.True(true, "断言表达式结果为真")
    zassert.False(false, "断言表达式结果为假")
    zassert.Equal(1, 1, "断言两个表达式结果相等")
    zassert.NotEqual(1, 2, "断言两个表达式结果不相等")
    zassert.Nil(nil, "断言表达式结果为nil")
    zassert.NotNil(1, "断言表达式结果不为nil")
    zassert.Zero(0, "断言表达式结果为0")
    zassert.NotZero(1, "断言表达式结果不为0")
    zassert.EmptyString("", "断言表达式结果为空字符串")
    zassert.NotEmptyString("1", "断言表达式结果不为空字符串")
}
```

### 捕获zassert.AssertError类型的错误

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zassert"
)

func main() {
    defer func() {
        err := recover()
        switch err.(type) {
        case zassert.AssertError:
            fmt.Println("zassert.AssertError类型错误", err)
        default:
            fmt.Println("其他类型错误", err)
        }
    }()

    zassert.PanicAssertError("直接产生zassert.AssertError类型错误")
}
```
