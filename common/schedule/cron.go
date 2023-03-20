package schedule

import "github.com/robfig/cron/v3"

func StartCronTask(spec string, fc func()) {
	cr := cron.New(cron.WithSeconds())
	cr.AddFunc(spec, fc)
	cr.Start()
}
