package solvanta

import "context"

// Services returns the list of registered solver tasks and their metadata.
func (c *Client) Services(ctx context.Context) (*ServicesResponse, error) {
	var resp ServicesResponse
	if err := c.do(ctx, "GET", "/v1/solver/services", nil, &resp, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Health pings the API and returns its status.
func (c *Client) Health(ctx context.Context) (*HealthResponse, error) {
	var resp HealthResponse
	if err := c.do(ctx, "GET", "/health", nil, &resp, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Dashboard returns credit balance and usage stats for the authenticated workspace.
func (c *Client) Dashboard(ctx context.Context) (*DashboardResponse, error) {
	var resp DashboardResponse
	if err := c.do(ctx, "GET", "/v1/dashboard", nil, &resp, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
