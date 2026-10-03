package solvanta

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// Solve creates a new solve request. The payload should be a struct or map
// matching the task's payload schema (e.g. map[string]interface{} or a typed
// struct). The response is returned synchronously — the API blocks until
// the solver completes or times out.
func (c *Client) Solve(ctx context.Context, task string, payload interface{}) (*SolveResponse, error) {
	return c.SolveWithKey(ctx, task, payload, "")
}

// SolveWithKey creates a solve request with an idempotency key. The key must
// be unique per operation (up to 128 characters). Replaying the same key with
// the same task and payload returns the original result without charging again.
func (c *Client) SolveWithKey(ctx context.Context, task string, payload interface{}, idempotencyKey string) (*SolveResponse, error) {
	body := SolveRequest{Task: task, Payload: payload}

	var headers map[string]string
	if idempotencyKey != "" {
		headers = map[string]string{"Idempotency-Key": idempotencyKey}
	}

	var resp SolveResponse
	if err := c.do(ctx, "POST", "/v1/solve", body, &resp, headers); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSolve retrieves a previously created solve by its ID.
func (c *Client) GetSolve(ctx context.Context, id string) (*SolveResponse, error) {
	var resp SolveResponse
	if err := c.do(ctx, "GET", "/v1/solves/"+url.PathEscape(id), nil, &resp, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSolves returns the most recent solves for the workspace. Pass limit <= 0
// to use the server default.
func (c *Client) ListSolves(ctx context.Context, limit int) ([]SolveResponse, error) {
	path := "/v1/solves"
	if limit > 0 {
		path = fmt.Sprintf("/v1/solves?limit=%s", strconv.Itoa(limit))
	}
	var list []SolveResponse
	if err := c.do(ctx, "GET", path, nil, &list, nil); err != nil {
		return nil, err
	}
	return list, nil
}
