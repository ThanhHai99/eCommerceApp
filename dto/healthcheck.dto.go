package dto

type HealCheckDataRes struct {
	Uptime string `json:"uptime"`
}

type HealthCheckRes struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    HealCheckDataRes `json:"data"`
}
