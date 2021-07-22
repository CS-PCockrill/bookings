package render

import (
	"github.com/CS-PCockrill/bookings-app/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	req, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(req.Context(), "flash", "123")

	result := AddDefaultData(&td, req)
	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)

	}
	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter
	err = RenderTemplate(&ww, "home.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&ww, "non-existent.page.tmpl", &models.TemplateData{}, r)
	if err == nil {
		t.Error("Error, should not have rendered template that doesn't exist")
	}

}

func getSession() (*http.Request, error) {
	req, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := req.Context()
	ctx, _ = session.Load(ctx, req.Header.Get("X-Session"))
	req = req.WithContext(ctx)
	return req, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

}