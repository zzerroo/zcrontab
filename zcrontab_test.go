package zcrontab

import (
	"fmt"
	"testing"
	"time"
	"zcrontab"
)

func test(args interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), args)
}

func TestMain(t *testing.T) {
	zCrontab := zcrontab.NewZCrontab()
	zCrontab.Crontab("* * * * *", test, "*")
	zCrontab.Crontab("* 9-10 24 6 *", test, "0624,9-10,every minute")
	zCrontab.Crontab("* 9 24 6 *", test, "0624,9,every minute")
	zCrontab.Crontab("* 9-10 24-25 6 *", test, "0624、0625,9-10,every minute")
	zCrontab.Crontab("* 9-10 * * 1", test, "Mon 9-10,every minute")
	zCrontab.Crontab("*/2 * * * *", test, "every 2 minute")
	zCrontab.Crontab("* 15 * * *", test, "every minute 3")
	zCrontab.Crontab("*/2 9,10 * * *", test, "every 2 minute 9,10")
	zCrontab.Crontab("*/2 9-10 * * *", test, "every 2 minute 9-10")
	zCrontab.Crontab("* 9-10/2 * * *", test, "every minute 9-10/2")

	zCrontab.At(time.Now().Add(1*time.Minute), test, "this is test for at repeat", true)
	zCrontab.At(time.Now().Add(1*time.Minute), test, "this is test for at", false)

	select {}
}
