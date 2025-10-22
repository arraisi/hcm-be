package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	apphttp "hcm-be/internal/http"
	"hcm-be/internal/http/handlers"
	"hcm-be/internal/repository/memory"
	"hcm-be/internal/service"
)

func setup() http.Handler {
	repo := memory.NewUserRepo()
	svc := service.NewUserService(repo)
	h := handlers.NewUserHandler(svc)
	return apphttp.NewRouter(h, apphttp.RouterOptions{})
}

func TestCreateAndListUsers(t *testing.T) {
	server := httptest.NewServer(setup())
	defer server.Close()

	// create
	resp, err := http.Post(server.URL+"/api/v1/users", "application/json", bytes.NewBufferString(`{"email":"a@b.com","name":"Abdul"}`))
	if err != nil { t.Fatal(err) }
	if resp.StatusCode != http.StatusCreated { t.Fatalf("expected 201 got %d", resp.StatusCode) }

	// list
	resp2, err := http.Get(server.URL+"/api/v1/users")
	if err != nil { t.Fatal(err) }
	if resp2.StatusCode != http.StatusOK { t.Fatalf("expected 200 got %d", resp2.StatusCode) }
}
