package consts

import "time"

var (
	TimeOut        = time.Second * 30
	FreshSchedule  = "*/10 * * * * *"
	ReportSchedule = "*/5 * * * * *"
)
