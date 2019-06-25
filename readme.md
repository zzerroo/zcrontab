# ZCrontab

zcrontab is a crontab、at  task processor based on golang:

- based on time.ticker，no need of crond service
- support task deletion and callback
- crontab
  - the crontab task support all crontab symbol including *,-,/,,

- at
  - the at task can be repeated at the same time(every day)

# Usage

## Install 

  glide get github.com/zzerroo/zcrontab

## example

```go
import (
	"fmt"
	"time"
	"ztimer/zcrontab"
)

func test(args interface{}) {
	fmt.Println(args)
}
 
func main() {
	zCrontab := zcrontab.NewZCrontab()
	zCrontab.Crontab("* * * * *", test, "*")
	zCrontab.Crontab("*/2 * * * *", test, "every 2 minute")
	zCrontab.Crontab("* 15 * * *", test, "every minute 3")
	zCrontab.Crontab("*/2 15,16 * * *", test, "every 2 minute 3,4")
	zCrontab.Crontab("*/2 15-16 * * *", test, "every 2 minute 3-4")
	zCrontab.Crontab("* 15-16/2 * * *", test, "every minute 3-4/3")
	zCrontab.At(time.Now().Add(1*time.Minute), test, "this is test for at repeat", true)
	zCrontab.At(time.Now().Add(1*time.Minute), test, "this is test for at", false)
	select {}
}
```

#   Third

github.com/gorhill/cronexpr

github.com/satori/go.uuid