package sbapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Get sends a get request to the silverbullet API using the custom client
// Takes a page parameter
// Returns the content of the page and error
func (client *SBClient) Get(page string) (body string, err error) {
	// Create a GET request to the endpoint without an io.Reader
	request, err := client.CreateRequest("GET", page, nil)
	if err != nil {
		return "", err
	}

	// Send the request using the HttpClient
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// Reads the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), err
}

// Put sets the page contents to the given data
// Takes a page and a data parameter
// Returns the response body and error
func (client *SBClient) Put(page string, data string) (body string, err error) {
	// Create a PUT request with an io.Reader
	request, err := client.CreateRequest("PUT", page, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}

	// add timestamps to request
	request.Header.Set("X-Created", fmt.Sprintf("%d", time.Now().UnixMilli()))
	request.Header.Set("X-Last-Modified", fmt.Sprintf("%d", time.Now().UnixMilli()))

	// Sends the request using the HttpClient
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// Reads the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), err
}

// Delete removes the given page
// Takes a page parameter
// Returns the response body and error
func (client *SBClient) Delete(page string) (body string, err error) {
	// Creates a DELETE request with no io.Reader
	request, err := client.CreateRequest("DELETE", page, nil)
	if err != nil {
		return "", err
	}

	// Sends the request using the HttpClient
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// Reads the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), err
}

// Append uses GET and PUT methods. First gets the current contents of the file
// After that, appends the new data after the separator. current + separator + new data
// Takes page, data and separator strings
// Returns the response body and error
func (client *SBClient) Append(page string, data string, separator string) (body string, err error) {
	// Get the current data
	currentData, err := client.Get(page)
	if err != nil {
		return body, err
	}

	// Update the data
	updatedData := fmt.Sprintf("%s%s%s", currentData, separator, data)

	// PUT the updated data
	body, err = client.Put(page, updatedData)

	return body, err
}

// CreateRequest creates a http.Request with the given parameters
// and sets the necessary headers for the silverbullet API
// Takes url, token, method strings and a body io.Reader
// returns *http.Request, error
func (client *SBClient) CreateRequest(method string, page string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", client.Endpoint, page), body)
	if err != nil {
		return req, nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.Token))
	req.Header.Set("content-type", "text/markdown")
	req.Header.Set("X-Sync-Mode", "true")

	return req, err
}
