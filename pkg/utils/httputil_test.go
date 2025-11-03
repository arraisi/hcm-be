package utils

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arraisi/hcm-be/internal/platform/httpclient"
)

func TestHttpUtil_Get(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Verify token
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got '%s'", auth)
		}

		// Verify custom header
		customHeader := r.Header.Get("X-Custom-Header")
		if customHeader != "custom-value" {
			t.Errorf("Expected X-Custom-Header 'custom-value', got '%s'", customHeader)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()
	customHeaders := map[string]string{
		"X-Custom-Header": "custom-value",
	}

	response, err := httpUtil.Get(ctx, server.URL, "test-token", customHeaders)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	expected := `{"status":"success"}`
	if string(response) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(response))
	}
}

func TestHttpUtil_Post(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		// Verify content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
		}

		// Verify token
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got '%s'", auth)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"123","status":"created"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()
	body := map[string]interface{}{
		"name": "Test User",
	}

	response, err := httpUtil.Post(ctx, server.URL, body, "test-token")
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	expected := `{"id":"123","status":"created"}`
	if string(response) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(response))
	}
}

func TestHttpUtil_Put(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		// Verify token
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got '%s'", auth)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"updated"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()
	body := map[string]interface{}{
		"name": "Updated User",
	}

	response, err := httpUtil.Put(ctx, server.URL, body, "test-token")
	if err != nil {
		t.Fatalf("PUT request failed: %v", err)
	}

	expected := `{"status":"updated"}`
	if string(response) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(response))
	}
}

func TestHttpUtil_Delete(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify method
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		// Verify token
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got '%s'", auth)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()

	response, err := httpUtil.Delete(ctx, server.URL, "test-token")
	if err != nil {
		t.Fatalf("DELETE request failed: %v", err)
	}

	if len(response) != 0 {
		t.Errorf("Expected empty response, got '%s'", string(response))
	}
}

func TestHttpUtil_WithoutToken(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify no token is sent
		auth := r.Header.Get("Authorization")
		if auth != "" {
			t.Errorf("Expected no Authorization header, got '%s'", auth)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()

	response, err := httpUtil.Get(ctx, server.URL, "")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	expected := `{"status":"ok"}`
	if string(response) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(response))
	}
}

func TestHttpUtil_MultipleHeaders(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify multiple headers
		header1 := r.Header.Get("X-Header-1")
		header2 := r.Header.Get("X-Header-2")

		if header1 != "value1" {
			t.Errorf("Expected X-Header-1 'value1', got '%s'", header1)
		}
		if header2 != "value2" {
			t.Errorf("Expected X-Header-2 'value2', got '%s'", header2)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 1,
	})

	ctx := context.Background()
	headers1 := map[string]string{
		"X-Header-1": "value1",
	}
	headers2 := map[string]string{
		"X-Header-2": "value2",
	}

	response, err := httpUtil.Get(ctx, server.URL, "", headers1, headers2)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	expected := `{"status":"ok"}`
	if string(response) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(response))
	}
}

func TestHttpUtil_ErrorHandling(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"bad request"}`))
	}))
	defer server.Close()

	// Create HttpUtil client
	httpUtil := NewHttpUtil(httpclient.Options{
		Timeout: 5 * time.Second,
		Retries: 0, // No retries for 4xx errors
	})

	ctx := context.Background()

	response, err := httpUtil.Get(ctx, server.URL, "")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Should still return response body
	if string(response) != `{"error":"bad request"}` {
		t.Errorf("Expected error response body, got '%s'", string(response))
	}

	// Verify it's an HTTPError
	httpErr, ok := err.(*httpclient.HTTPError)
	if !ok {
		t.Errorf("Expected HTTPError, got %T", err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %d", httpErr.StatusCode)
	}
}
