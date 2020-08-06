package test

import (
	"net/http"
	"strings"
)

type UsersSuite struct {
	AbstractSuite
}

// SetupSuite setup before all tests
func (s *UsersSuite) SetupSuite() {
	s.cleanupTables()
	s.seed("users")
}

func (s *UsersSuite) TestLoginFail() {
	req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"username":"Admin","password":"123456789"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusUnauthorized, code)
	s.Require().JSONEq(`{"message":"bad login or password"}`, response)
}

func (s *UsersSuite) TestLoginSuccess() {
	req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"username":"Admin","password":"12345678"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, "token")
}

func (s *UsersSuite) TestRegisterSuccess() {
	req, _ := http.NewRequest(http.MethodPost, "/api/register", strings.NewReader(`{"username":"alexander","password":"11112222"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusCreated, code)
	s.Require().Contains(response, "token")
}
