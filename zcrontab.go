package zcrontab

import (
	"errors"
	"sync"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/satori/go.uuid"
)

type ZCrontab struct {
	zCrontabInfo map[string]interface{}
	mut          sync.RWMutex
}

type contrabInfo struct {
	conStr string
	args   interface{}
	f      func(interface{})
}

type atInfo struct {
	sec      int
	min      int
	hour     int
	day      int
	mon      time.Month
	year     int
	repeated bool
	args     interface{}
	f        func(interface{})
}

// NewZCrontab ...
func NewZCrontab() *ZCrontab {
	zTimer := new(ZCrontab)
	zTimer.zCrontabInfo = make(map[string]interface{})
	go zTimer.startConsume()
	return zTimer
}

// Crontab add a crontab task,cronStr is a crontab string,f and args is the callback and input
func (z *ZCrontab) Crontab(cronStr string, f func(interface{}), args interface{}) (string, error) {
	_, erro := cronexpr.Parse(cronStr)
	if erro != nil {
		return "", errors.New("bad format string")
	}

	uuid := z.getUUID()
	z.mut.Lock()
	z.zCrontabInfo[uuid] = contrabInfo{
		conStr: cronStr,
		args:   args,
		f:      f,
	}
	z.mut.Unlock()
	return uuid, nil
}

// At add a at task, t is the task time, f and args is the callback and input, repeat indicate wheather
// the task will repeat every day（will ignore the date info）
func (z *ZCrontab) At(t time.Time, f func(interface{}), args interface{}, repeated bool) (string, error) {
	uuid := z.getUUID()
	z.mut.Lock()
	z.zCrontabInfo[uuid] = atInfo{
		sec:      t.Second(),
		min:      t.Minute(),
		hour:     t.Hour(),
		day:      t.Day(),
		mon:      t.Month(),
		year:     t.Year(),
		repeated: repeated,
		args:     args,
		f:        f,
	}
	z.mut.Unlock()
	return uuid, nil
}

// RmAt remove the task
func (z *ZCrontab) RmAt(id string) {
	z.mut.Lock()
	delete(z.zCrontabInfo, id)
	z.mut.Unlock()
}

func (z *ZCrontab) getUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

func (z *ZCrontab) startConsume() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		cur := time.Now()
		z.mut.RLock()
		for _, value := range z.zCrontabInfo {
			atT, isAt := value.(atInfo)
			if isAt {
				if cur.Second() == atT.sec && cur.Minute() == atT.min && cur.Hour() == atT.hour {
					if cur.Day() == atT.day && cur.Month() == atT.mon && cur.Year() == atT.year {
						go atT.f(atT.args)
					} else if atT.repeated == true {
						go atT.f(atT.args)
					}
				}
			} else {
				if cur.Second() == 0 {
					conInfo, isCon := value.(contrabInfo)
					if isCon {
						preMin := cur.Add(-1 * time.Minute)
						nextTime := cronexpr.MustParse(conInfo.conStr).Next(preMin)

						if nextTime.Weekday() == cur.Weekday() && nextTime.Month() == cur.Month() && nextTime.Day() == cur.Day() && nextTime.Hour() == cur.Hour() && nextTime.Minute() == cur.Minute() {
							conT := value.(contrabInfo)
							go conT.f(conT.args)
						}
					}
				}
			}
		}
		z.mut.RUnlock()
	}
}
