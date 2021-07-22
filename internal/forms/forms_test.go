package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	req := httptest.NewRequest("POST", "/endpoint", nil)
	form := New(req.PostForm)

	valid := form.Valid()
	if !valid {
		t.Error("Error, got invalid form when should be valid")
	}

}

func TestForm_Required(t *testing.T) {
	req := httptest.NewRequest("POST", "/endpoint", nil)
	form := New(req.PostForm)
	form.Required("input1", "input2", "input3")
	if form.Valid() {
		t.Error("Error, Form says valid when required fields are missing")
	}

	postData := url.Values{}
	postData.Add("input1", "testing1")
	postData.Add("input2", "testing2")
	postData.Add("input3", "testing3")

	req, _ = http.NewRequest("POST", "/endpoint", nil)
	req.PostForm = postData

	form = New(req.PostForm)
	form.Required("input1", "input2", "input3")
	if !form.Valid() {
		t.Error("Shows does not have required fields when it does.")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(url.Values{})

	has := form.Has("any-val")
	if has {
		t.Error("Error, form shows it has field when it does not have field")
	}

	postData = url.Values{}
	postData.Add("a", "a-input")
	form = New(postData)

	has = form.Has("a")
	if !has {
		t.Error("Error, shows form does not have this field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postValues := url.Values{}
	form := New(postValues)
	field := "some-string"
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("Error, form shows minlength for a field that does not exist")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Error, the field should have an error but did not get one")
	}

	postValues = url.Values{}
	postValues.Add("some-field", field)
	form = New(postValues)
	form.MinLength("some-field", 100)
	if form.Valid() {
		t.Error("Error, shows minLength of 100 met when data is shorter")
	}

	postValues = url.Values{}
	postValues.Add("some-field", field)
	form = New(postValues)
	form.MinLength("some-field", 1)
	if !form.Valid() {
		t.Error("Error, shows minLength of 1 is not met when it is")
	}

	isError = form.Errors.Get("some-field")
	if isError != "" {
		t.Error("Error, the field should not have an error but got one")
	}
}

func TestForm_ValidateEmail(t *testing.T) {
	postValues := url.Values{}
	form := New(postValues)

	form.ValidateEmail("abc.com")
	if form.Valid() {
		t.Error("Error, form shows valid email for non-existent field.")
	}

	postValues = url.Values{}
	postValues.Add("email", "abc@gmail.com")
	form = New(postValues)
	form.ValidateEmail("email")
	if !form.Valid() {
		t.Error("Error, form shows invalid email when it is indeed valid.")
	}

	postValues = url.Values{}
	postValues.Add("email", "abc.com")
	form = New(postValues)
	form.ValidateEmail("email")
	if form.Valid() {
		t.Error("Error, form shows valid email when it is not valid.")
	}

}
