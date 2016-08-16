package jobs

import (
  "log"
)

type Job interface {
  Run()
}

type StorageJob struct {
}

func (j StorageJob) Run() {
  log.Println("---> Configure jobs in OnAppStart")
}
