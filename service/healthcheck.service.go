package service

import (
	"eCommerce/dto"
	"eCommerce/util"
	"fmt"
	gosysinfo "github.com/elastic/go-sysinfo"
	"time"
)

func HealthCheck() dto.HealthCheckRes {
	res := dto.HealthCheckRes{}
	res.Code = util.SUCCESS_CODE
	res.Message = "Service is ok"

	process, _ := gosysinfo.Self()
	processInfo, _ := process.Info()
	startAt := processInfo.StartTime
	uptime := time.Now().Sub(startAt).String()
	fmt.Println(process.CPUTime())
	res.Data.Uptime = uptime

	return res
}
