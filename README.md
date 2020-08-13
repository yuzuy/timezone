# tzlocation
A package deals with timezone locations

## Installation
```
go get -u github.com/yuzuy/tzlocation
```

## Usage
```
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/yuzuy/tzlocation"
)

func main() {
    locTokyo, err := tzlocation.AsiaTokyo.Load()
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(time.Now().In(locTokyo))
}
```
