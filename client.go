package solvanta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultBaseURL = "https://backend.solvanta.io"

// Client communicates with the Solvanta API.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL overrides the default API origin.
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = strings.TrimRight(url, "/") }
}

// WithHTTPClient provides a custom *http.Client for all requests.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.http = hc }
}

// WithTimeout sets the HTTP client timeout. Ignored when WithHTTPClient is used.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if c.http == http.DefaultClient || c.http == nil {
			c.http = &http.Client{Timeout: d}
		} else {
			c.http.Timeout = d
		}
	}
}

// New creates a Solvanta API client. The apiKey is sent as a Bearer token.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		http:    &http.Client{Timeout: 90 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// do executes an HTTP request and decodes the JSON response into dst.
// If the API returns a non-2xx status, an *Error is returned.
func (c *Client) do(ctx context.Context, method, path string, body interface{}, dst interface{}, headers map[string]string) error {
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("solvanta: marshal request: %w", err)
		}
		reqBody = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("solvanta: build request: %w", err)
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("solvanta: request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return fmt.Errorf("solvanta: read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var env errorEnvelope
		_ = json.Unmarshal(raw, &env)
		return &Error{
			StatusCode: resp.StatusCode,
			Code:       env.Error.Code,
			Message:    env.Error.Message,
		}
	}

	if dst != nil && len(raw) > 0 {
		if err := json.Unmarshal(raw, dst); err != nil {
			return fmt.Errorf("solvanta: decode response: %w", err)
		}
	}
	return nil
}
