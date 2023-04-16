package http

import (
	"bytes"
	"net/http"
)

func PutRequest(url string, contentType string, bytesBuffer *bytes.Buffer) (*http.Response, error) {
	// Create request
	request, err := http.NewRequest(http.MethodPut, url, bytesBuffer)
	if err != nil {
		return nil, err
	}

	// Set content type
	request.Header.Set("Content-Type", contentType)

	// Send request
	client := &http.Client{}
	response, err := client.Do(request)

	// Return response and err
	return response, err
}
