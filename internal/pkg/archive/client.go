package archive

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseURL = "https://api.battlefield.rip/archive/bfbc2/players/"
)

var (
	ErrPlayerNotFound = errors.New("player not found")
)

type RequestError struct {
	requestURL *url.URL
	statusCode int
}

func newRequestError(requestURL *url.URL, statusCode int) *RequestError {
	return &RequestError{
		requestURL: requestURL,
		statusCode: statusCode,
	}
}

func (e RequestError) Error() string {
	return fmt.Sprintf("request to %s failed with status code %d", e.requestURL.String(), e.statusCode)
}

func (e RequestError) StatusCode() int {
	return e.statusCode
}

type Client struct {
	baseURL string

	client *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *Client) GetStats(ctx context.Context, platform string, name string) (StatsResponse, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return StatsResponse{}, err
	}

	u = u.JoinPath(platform, name, "stats")

	q := u.Query()
	q.Set("keySet", "all")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return StatsResponse{}, err
	}

	body, err := c.do(req)
	if err != nil {
		if rerr, ok := errors.AsType[*RequestError](err); ok && rerr.statusCode == http.StatusNotFound {
			return StatsResponse{}, ErrPlayerNotFound
		}
		return StatsResponse{}, err
	}

	var resp StatsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return StatsResponse{}, err
	}

	return resp, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if !isSuccessStatusCode(res.StatusCode) {
		return nil, newRequestError(req.URL, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func isSuccessStatusCode(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode <= http.StatusIMUsed
}
