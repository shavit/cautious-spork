package jobs

import (
  "log"
)

type DecodeSegmentsJob struct {
}

func (j DecodeSegmentsJob) Run() {
  log.Println("---> Decoding...")
}
