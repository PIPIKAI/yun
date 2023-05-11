// package Schedule
package schedule

import "github.com/robfig/cron/v3"

// Schedule
type Schedule struct {
	ID     cron.EntryID
	Cron   string
	Task   func()
	Status string
}

// ScheduleManage
type ScheduleManage struct {
	Cron      *cron.Cron
	Schedules map[cron.EntryID]*Schedule
}

// NewScheduleManage
//
//	@return *ScheduleManage
func NewScheduleManage() *ScheduleManage {
	return &ScheduleManage{
		Cron:      cron.New(cron.WithSeconds()),
		Schedules: make(map[cron.EntryID]*Schedule),
	}
}

// Add one manage
//
//	@receiver sm
//	@param spec
//	@param cmd
//	@return error
func (sm *ScheduleManage) Add(spec string, cmd func()) error {
	id, err := sm.Cron.AddFunc(spec, cmd)
	if err != nil {
		return err
	}
	sm.Schedules[id] = &Schedule{
		ID:     id,
		Cron:   spec,
		Task:   cmd,
		Status: "waiting start",
	}
	return nil
}

// GetAll Manage
//
//	@receiver sm
//	@return map
func (sm *ScheduleManage) GetAll() map[cron.EntryID]*Schedule {
	return sm.Schedules
}

// StartAll
//
//	@receiver sm
func (sm *ScheduleManage) StartAll() {
	for k := range sm.Schedules {
		sm.Schedules[k].Status = "working"
	}
	sm.Cron.Start()
}
