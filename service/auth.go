package service

import "errors"

// Authentication handles authenticate actions
type Authentication interface {
	Login(username, password string) error
	Logout() error
	LoggedIn() bool
}

// FakeAuthentication faking the authentication by using hard-coded user
type fakeAuthentication struct {
	username string
	password string
	loggedIn bool
}

func (auth *fakeAuthentication) Login(username, password string) error {
	if username == auth.username && password == auth.password {
		auth.loggedIn = true
		return nil
	}
	return errors.New("auth failed")
}

func (auth *fakeAuthentication) Logout() error {
	if auth.loggedIn {
		auth.loggedIn = false
		return nil
	}
	return errors.New("not logged in")
}

// TODO: Use session instead
func (auth *fakeAuthentication) LoggedIn() bool {
	return auth.loggedIn
}

// NewFakeAuthentication - fakeAuthentication constructor
func NewFakeAuthentication(username, password string) Authentication {
	return &fakeAuthentication{
		username: username,
		password: password,
		loggedIn: false,
	}
}
