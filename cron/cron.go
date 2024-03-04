package cron

import (
	"time"

	"powerbot/core"

	"github.com/go-co-op/gocron/v2"
)

type Cron struct {
	Scheduler gocron.Scheduler
	core      *core.CoreService
}

func NewCron(c *core.CoreService) *Cron {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	cron := Cron{Scheduler: scheduler, core: c}

	_, err = scheduler.NewJob(gocron.CronJob("0 * * * *", false), gocron.NewTask(func() {
		time.Sleep(time.Millisecond * 20)
		cron.core.TimedQuery(&core.TimedQueryRequest{Hour: int64(time.Now().Hour())})
	}))
	if err != nil {
		panic(err)
	}

	return &cron
}
