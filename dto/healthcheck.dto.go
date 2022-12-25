package dto

type HealCheckData struct {
	Uptime string `json:"uptime"`
}

type HealthCheckRes struct {
	BaseRes
	Data HealCheckData `json:"data"`
}
