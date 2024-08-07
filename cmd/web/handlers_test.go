package main

import (
	"net/http"
	"net/url"
	"testing"
	
	"snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()
	
	code, _, body := ts.get(t, "/ping")
	
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	
	ts := newTestServer(t, app.routes())

	defer ts.Close()
	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName = "Bob"
		validPassword = "validPa$$word"
		validEmail = "bob@example.com"
		formTag = "<form action='/user/signup' method='POST' novalidate>"
	)

	tests := []struct {
		name string
		userName string
		userEmail string
		userPassword string
		csrfToken string
		wantCode int
	}{
		{
			name: "Valid submission",
			userName: validName,
			userEmail: validEmail,
			userPassword: validPassword,
			csrfToken: validCSRFToken,
			wantCode: http.StatusSeeOther,
		},
		{
			name: "Invalid CSRF Token",
			userName: validName,
			userEmail: validEmail,
			userPassword: validPassword,
			csrfToken: "wrongToken",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Empty name",
			userName: "",
			userEmail: validEmail,
			userPassword: validPassword,
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Empty email",
			userName: validName,
			userEmail: "",
			userPassword: validPassword,
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Empty password",
			userName: validName,
			userEmail: validEmail,
			userPassword: "",
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Invalid email",
			userName: validName,
			userEmail: "bob@example.",
			userPassword: validPassword,
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Short password",
			userName: validName,
			userEmail: validEmail,
			userPassword: "pa$$",
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Duplicate email",
			userName: validName,
			userEmail: "dupe@example.com",
			userPassword: validPassword,
			csrfToken: validCSRFToken,
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, _ := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.wantCode)
		})
	}
}