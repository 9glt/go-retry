# go-retry
Retry func if err 

```go
package main

import (
    retry "github.com/9glt/go-retry"
)

func main() {

    var data string
    err := retry.Run(func(r int) error {
        var err error
        data, err = myfunc()
        return err
    }, 3, time.Second)

    if err != nil {
        panic(err)
    }
    println(data)
}

func myfunc() (string, error) {
    return "hi", nil
}
```