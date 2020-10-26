package schedule

import (
	"github.com/robfig/cron/v3"
)

func ScheduleInit() {
	// groutine job

	// cron job
	c := cron.New()
	c.AddFunc("@every 1h", StartHeapCheck)

	c.Start()
}
