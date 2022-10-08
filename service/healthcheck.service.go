package service

import (
	"eCommerce/dto"
	"eCommerce/util"
	gosysinfo "github.com/elastic/go-sysinfo"
	"time"
)

func HealthCheck() dto.HealthCheckRes {
	process, _ := gosysinfo.Self()
	processInfo, _ := process.Info()
	startAt := processInfo.StartTime
	uptime := time.Now().Sub(startAt).String()

	res := dto.HealthCheckRes{}
	res.Code = util.SUCCESS_CODE
	res.Message = "Service is ok"
	res.Data.Uptime = uptime

	return res
}
