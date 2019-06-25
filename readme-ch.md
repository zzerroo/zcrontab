# ZCrontab

基于go的crontab任务处理器

- 基于ticker的任务处理机制，不依赖crond
- 支持任务的删除、回调

- crontab任务
  - crontab全任务支持*,-,/及,等类型任务处理
- At 任务
  - 支持特定时间的任务执行
  - 支持重复任务

# 如何使用

- ## install 

  git clone https://github.com/zzerroo/zcrontab.git && cd zcrontab && glide install

- ## example

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
  
  

# 第三方库

github.com/gorhill/cronexpr

github.com/satori/go.uuid
