package dto

type HealCheckDataRes struct {
	Uptime string `json:"uptime"`
}

type HealthCheckRes struct {
	BaseRes
	Data HealCheckDataRes `json:"data"`
}
