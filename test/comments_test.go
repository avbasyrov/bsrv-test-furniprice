package test

import (
	"net/http"
	"strings"
)

type CommentsSuite struct {
	AbstractSuite
	authToken string
}

// SetupSuite setup before all tests
func (s *CommentsSuite) SetupSuite() {
	s.cleanupTables()
	s.seed("users")
	s.seed("posts")
	s.authToken = s.getAuthToken()
}

func (s *CommentsSuite) TestAddCommentForbidden() {
	req, _ := http.NewRequest(http.MethodPost, "/api/post/"+adminPostID, strings.NewReader(`{"comment":"afdsasdfsadfas"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message":"unauthorized"}`, response)
}

func (s *CommentsSuite) TestDeleteCommentForbidden() {
	req, _ := http.NewRequest(http.MethodDelete, "/api/post/"+nonAdminPostID, nil)
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusForbidden, code)
	s.Require().JSONEq(`{"message":"unauthorized"}`, response)
}

func (s *CommentsSuite) TestAddCommentToOwnPostSuccess() {
	const comment = "afdsasdfsadxxxxxxxxfas"
	req, _ := http.NewRequest(http.MethodPost, "/api/post/"+adminPostID, strings.NewReader(`{"comment":"`+comment+`"}`))
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, adminPostID)
	s.Require().Contains(response, comment)
}

func (s *CommentsSuite) TestAddCommentToOtherPostSuccess() {
	const comment = "sfgdfzzzzzzzzzzzzxg"
	req, _ := http.NewRequest(http.MethodPost, "/api/post/"+nonAdminPostID, strings.NewReader(`{"comment":"`+comment+`"}`))
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, nonAdminPostID)
	s.Require().Contains(response, comment)
}

func (s *CommentsSuite) TestDeleteCommentToOwnPostSuccess() {
	const commentToOwnPost = "a9d5479b-60b0-4253-8ac5-7e3e635129af"
	req, _ := http.NewRequest(http.MethodDelete, "/api/post/"+adminPostID+"/"+commentToOwnPost, nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, adminPostID)
	s.Require().NotContains(response, commentToOwnPost)
}

func (s *CommentsSuite) TestDeleteCommentToNonOwnPostSuccess() {
	const commentToNotOwnPost = "09d5479b-60b0-4253-8ac5-7e3e635129af"
	req, _ := http.NewRequest(http.MethodDelete, "/api/post/"+nonAdminPostID+"/"+commentToNotOwnPost, nil)
	response, code := s.executeRequest(req, &s.authToken)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, nonAdminPostID)
	s.Require().NotContains(response, commentToNotOwnPost)
}
