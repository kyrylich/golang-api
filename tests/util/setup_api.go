package util

import (
	"github.com/gavv/httpexpect/v2"
	"golangpet/internal/api/router"
	"golangpet/internal/app"
	"golangpet/internal/models"
	"net/http/httptest"
	"testing"
)

var application app.AppInterface = &app.App{}

func SetupApi(t *testing.T) (*httpexpect.Expect, func()) {
	if err := application.Boot(); err != nil {
		t.Fatal(err)
	}
	CleanUpDatabase(models.DB)

	// server
	handler := router.CreateRouter()
	server := httptest.NewServer(handler)

	e := httpexpect.Default(t, server.URL)

	return e, server.Close
}
