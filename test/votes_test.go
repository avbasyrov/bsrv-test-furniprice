package test

import (
	"net/http"
)

type VotesSuite struct {
	AbstractSuite
	authToken string
}

// SetupSuite setup before all tests
func (s *VotesSuite) SetupSuite() {
	s.cleanupTables()
	s.seed("users")
	s.seed("posts")
	s.authToken = s.getAuthToken()
}

func (s *VotesSuite) TestVoteUpNoAuth() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/upvote", nil)
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message": "unauthorized"}`, response)
}

func (s *VotesSuite) TestUnVoteNoAuth() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/unvote", nil)
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message": "unauthorized"}`, response)
}

func (s *VotesSuite) TestVoteDownNoAuth() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/downvote", nil)
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message": "unauthorized"}`, response)
}

func (s *VotesSuite) TestVoteUpSuccess() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/upvote", nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, adminPostID)
	s.Require().Regexp(`"score"\s*:\s*1[,}]`, response)
}

func (s *VotesSuite) TestVoteDownSuccess() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/downvote", nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, adminPostID)
	s.Require().Regexp(`"score"\s*:\s*-1[,}]`, response)
}

func (s *VotesSuite) TestUnVoteSuccess() {
	req, _ := http.NewRequest(http.MethodGet, "/api/post/"+adminPostID+"/unvote", nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, adminPostID)
	s.Require().Regexp(`"score"\s*:\s*0[,}]`, response)
}
