package test

import (
	"net/http"
	"strings"
)

type PostsSuite struct {
	AbstractSuite
	authToken string
}

// SetupSuite setup before all tests
func (s *PostsSuite) SetupSuite() {
	s.cleanupTables()
	s.seed("users")
	s.seed("posts")
	s.authToken = s.getAuthToken()
}

func (s *PostsSuite) TestPostsAddForbidden() {
	req, _ := http.NewRequest(http.MethodPost, "/api/posts", strings.NewReader(`{"category":"music","type":"text","title":"dfgdsfgs","text":"sdgfsdfgs"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message":"unauthorized"}`, response)
}

func (s *PostsSuite) TestPostsAddSuccess() {
	const title = "some unique title 8fojwr8w8jr"
	const body = "some unique body safgdasfgii108fojwr8w8jr"
	req, _ := http.NewRequest(http.MethodPost, "/api/posts", strings.NewReader(`{"category":"music","type":"text","title":"`+title+`","text":"`+body+`"}`))
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Regexp(`"category"\s*:\s*"music"`, response)
	s.Require().Contains(response, title)
	s.Require().Contains(response, body)
}

func (s *PostsSuite) TestPostsDeleteSuccess() {
	req, _ := http.NewRequest(http.MethodDelete, "/api/post/"+adminPostID, nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().JSONEq(`{"message":"success"}`, response)
}
