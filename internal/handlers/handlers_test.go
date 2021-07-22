package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key string
	value string
}

var theTests = []struct {
	name string
	url string
	method string
	params []postData
	expectStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start_date", value: "07-22-2021"},
		{key: "end_date", value: "07-25-2021"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start_date", value: "07-22-2021"},
		{key: "end_date", value: "07-25-2021"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Patrick"},
		{key: "last_name", value: "Cockrill"},
		{key: "email", value: "pcock@here.com"},
		{key: "phone", value: "703-555-5555"},
	}, http.StatusOK},
	//{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			response, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != e.expectStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectStatusCode, response.StatusCode)
			}
		} else if e.method == "POST" {
			values := url.Values{}
			for _, val := range e.params {
				values.Add(val.key, val.value)
			}
			response, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != e.expectStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectStatusCode, response.StatusCode)
			}

		}
	}

}
