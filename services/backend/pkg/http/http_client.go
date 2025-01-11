package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SendRequest makes a GET or POST request to the a endpoint with optional headers and body
func SendRequest(
	method, baseURL string,
	headers map[string]string,
	body string,
	params map[string][]string,
) (string, error) {
	// Validate HTTP method
	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		return "", errors.New("invalid HTTP method: only GET and POST are supported")
	}

	// Build the full URL with query parameters 
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Add query parameters from the params map
	query := reqURL.Query()
	for key, values := range params {
		query.Set(key, strings.Join(values, ","))
	}
	reqURL.RawQuery = query.Encode()

	// Create request
	var req *http.Request
	if method == "POST" {
		req, err = http.NewRequest(method, reqURL.String(), bytes.NewBuffer([]byte(body)))
	} else {
		req, err = http.NewRequest(method, reqURL.String(), nil)
	}
	if err != nil {
		return "", err
	}

	// Add headers if provided
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
