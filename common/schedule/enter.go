package schedule

import "github.com/robfig/cron/v3"

type Schedule struct {
	ID     cron.EntryID
	Cron   string
	Task   func()
	Status string
}

type ScheduleManage struct {
	Cron      *cron.Cron
	Schedules map[cron.EntryID]*Schedule
}

func NewScheduleManage() *ScheduleManage {
	return &ScheduleManage{
		Cron:      cron.New(cron.WithSeconds()),
		Schedules: make(map[cron.EntryID]*Schedule),
	}
}

func (sm *ScheduleManage) Add(spec string, cmd func()) error {
	id, err := sm.Cron.AddFunc(spec, cmd)
	if err != nil {
		return err
	}
	sm.Schedules[id] = &Schedule{
		ID:     id,
		Cron:   spec,
		Task:   cmd,
		Status: "waiting",
	}
	return nil
}
func (sm *ScheduleManage) GetAll() map[cron.EntryID]*Schedule {
	return sm.Schedules
}
func (sm *ScheduleManage) StartAll() {
	for k := range sm.Schedules {
		sm.Schedules[k].Status = "working"
	}
	sm.Cron.Start()
}
