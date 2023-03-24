package consts

import "time"

var (
	TimeOut              = time.Second * 30
	TrackerFreshSchedule = "*/10 * * * * *"
	StorageFreshSchedule = "*/6000 * * * * *"
	ReportSchedule       = "*/5 * * * * *"
)
