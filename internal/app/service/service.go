package service

import (
	// "fmt"
	// // "net/http"
	"time"
)

type Service struct{
}

func New() *Service {
	return &Service{}
}

func (s *Service) DayLeft () int64 {
	d := time.Date(2025, time.January, 1,0,0,0,0,time.UTC)
	dure := time.Until(d)

	
	return int64(dure.Hours())/24
}