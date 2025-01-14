package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL string
	client  *http.Client
	token   string
}

func NewAPIClient(baseURL, token string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		token: token,
	}
}

func prettifyJSON(jsonBytes []byte) (string, error) {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, jsonBytes, "", "  ")
	if err != nil {
		return "", fmt.Errorf("JSON error: %w", err)
	}
	return pretty.String(), nil
}

func (api *APIClient) doRequest(method, endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", api.baseURL, endpoint)

	var requestBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize body to JSON: %w", err)
		}
		requestBody = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if api.token != "" {
		req.Header.Set("Authorization", "Bearer "+api.token)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("server error: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse body: %w", err)
	}

	return respBody, nil
}

func (api *APIClient) GetSingleProject(pubKey string) (string, error) {
	endpoint := fmt.Sprintf("/%s/", pubKey)
	bodyBytes, err := api.doRequest(http.MethodGet, endpoint, nil)
	pretty, err := prettifyJSON(bodyBytes)
	return pretty, err
}

func (api *APIClient) GetProjects() (string, error) {
	endpoint := "/"
	bodyBytes, err := api.doRequest(http.MethodGet, endpoint, nil)
	pretty, err := prettifyJSON(bodyBytes)
	return pretty, err
}

func (api *APIClient) CreateProject(body map[string]interface{}) (string, error) {
	endpoint := "/"
	bodyBytes, err := api.doRequest(http.MethodPost, endpoint, body)
	pretty, err := prettifyJSON(bodyBytes)
	return pretty, err
}

func (api *APIClient) DeleteProject(pubKey string) error {
	endpoint := fmt.Sprintf("/%s/", pubKey)
	_, err := api.doRequest(http.MethodDelete, endpoint, nil)
	return err
}

func (api *APIClient) UpdateProject(pubKey string, body map[string]interface{}) (string, error) {
	endpoint := fmt.Sprintf("/%s/", pubKey)
	bodyBytes, err := api.doRequest(http.MethodPost, endpoint, body)
	pretty, err := prettifyJSON(bodyBytes)
	return pretty, err
}
