package todoist

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	getTasksPath = "%s/rest/v2/tasks"

	accessTokenPath = "%s/oauth/access_token"
)

// Client is a http client for Todoist API
type Client struct {
	client       http.Client
	devBaseURL   string
	authzBaseURL string
}

// NewClient is a constructor for Client
func NewClient(client http.Client, devBaseURL, authzBaseURL string) *Client {
	return &Client{
		client:       client,
		devBaseURL:   devBaseURL,
		authzBaseURL: authzBaseURL,
	}
}

// GetTasks returns user's tasks
func (c *Client) GetTasks(ctx context.Context, token string) ([]*Task, error) {
	address := fmt.Sprintf(getTasksPath, c.devBaseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tasks []*Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("decode response: %v", err)
	}

	return tasks, nil
}

// RetrieveAccessToken returns user's access token
func (c *Client) RetrieveAccessToken(ctx context.Context, clientID, clientSecret, code, redirectURL string) (string, error) {
	address := fmt.Sprintf(accessTokenPath, c.authzBaseURL)

	form := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURL},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, address, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var token struct {
		AccessToken string `json:"access_token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", fmt.Errorf("decode response: %v", err)
	}

	return token.AccessToken, nil
}
