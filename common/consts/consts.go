package consts

import "time"

var (
	TimeOut        = time.Second * 30
	FreshSchedule  = "*/6000 * * * * *"
	ReportSchedule = "*/5 * * * * *"
)
