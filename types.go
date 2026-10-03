// Package solvanta provides a Go client for the Solvanta solver API.
package solvanta

import "encoding/json"

// Task name constants accepted by the API.
const (
	ReCaptchaV2Task           = "ReCaptchaV2Task"
	ReCaptchaV2EnterpriseTask = "ReCaptchaV2EnterpriseTask"
	ReCaptchaV3Task           = "ReCaptchaV3Task"
	ReCaptchaV3EnterpriseTask = "ReCaptchaV3EnterpriseTask"
	CloudflareTask            = "CloudflareTask"
	ArkoseTask                = "ArkoseTask"
	ForterTask                = "ForterTask"
	NuDataTask                = "NuDataTask"
	ThreatMetrixTask          = "ThreatMetrixTask"
	PxmTask                   = "PxmTask"
	UeMsmTask                 = "UeMsmTask"
	Metadata1Task             = "Metadata1Task"
)

// Solve status values returned by the API.
const (
	StatusSolved     = "solved"
	StatusProcessing = "processing"
	StatusFailed     = "failed"
)

// SolveRequest is the body sent to POST /v1/solve.
type SolveRequest struct {
	Task    string      `json:"task"`
	Payload interface{} `json:"payload"`
}

// SolveResponse is the envelope returned by the solve and result endpoints.
type SolveResponse struct {
	ID          string          `json:"id"`
	Status      string          `json:"status"`
	Result      json.RawMessage `json:"result,omitempty"`
	CreditsUsed int             `json:"credits_used"`
	Balance     int             `json:"balance,omitempty"`
	ElapsedMs   int             `json:"elapsed_ms"`
}

// Service describes one registered solver task.
type Service struct {
	Task         string   `json:"task"`
	Name         string   `json:"name"`
	Family       string   `json:"family"`
	Platform     string   `json:"platform"`
	Description  string   `json:"description"`
	Capabilities []string `json:"capabilities"`
	Credits      int      `json:"credits"`
	Aliases      []string `json:"aliases,omitempty"`
	Native       bool     `json:"native"`
	Configured   bool     `json:"configured"`
}

// ServicesResponse is the envelope returned by GET /v1/solver/services.
type ServicesResponse struct {
	Services       []Service `json:"services"`
	WorkerReachable bool     `json:"worker_reachable"`
}

// HealthResponse is the body returned by GET /health.
type HealthResponse struct {
	Status string `json:"status"`
}

// DashboardResponse is the body returned by GET /v1/dashboard.
type DashboardResponse struct {
	Credits    int             `json:"credits"`
	UsedToday  int             `json:"used_today"`
	SolvesToday int            `json:"solves_today"`
	Extra      json.RawMessage `json:"-"`
}

// errorEnvelope is the JSON error shape from the API.
type errorEnvelope struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
