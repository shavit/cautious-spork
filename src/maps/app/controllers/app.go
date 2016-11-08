package controllers

import (
	"github.com/revel/revel"
	appJobs "maps/app/jobs"
	"github.com/revel/modules/jobs/app/jobs"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	jobs.Now(appJobs.DecodeSegmentsJob{})

	return c.Render()
}
