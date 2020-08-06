package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/server"
	_ "github.com/lib/pq"
	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type AbstractSuite struct {
	suite.Suite
	app        *server.App
	DBConn     *sql.DB
	sqlPgx     *dbcon.Db
	dbPolluter *polluter.Polluter
}

const adminID = 1000000
const adminPostID = "5005cbe1-7c76-45af-848c-617f90ba79dd"
const nonAdminPostID = "1005cbe1-7c76-45af-848c-617f90ba79dd"

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	a := new(AbstractSuite)
	a.setupOnce()

	users := new(UsersSuite)
	users.setup(a)
	suite.Run(t, users)

	votes := new(VotesSuite)
	votes.setup(a)
	suite.Run(t, votes)

	comments := new(CommentsSuite)
	comments.setup(a)
	suite.Run(t, comments)

	posts := new(PostsSuite)
	posts.setup(a)
	suite.Run(t, posts)

	// suite.Run(t, new(VotesSuite).setup(a))
}

func (s *AbstractSuite) setup(abstractSuite *AbstractSuite) {
	s.app = abstractSuite.app
	s.DBConn = abstractSuite.DBConn
	s.sqlPgx = abstractSuite.sqlPgx
	s.dbPolluter = abstractSuite.dbPolluter
}

func (s *AbstractSuite) cleanupTables() {
	for _, table := range []string{"users", "comments", "posts", "sessions", "votes"} {
		_, err := s.DBConn.Exec("TRUNCATE " + table)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *AbstractSuite) seed(filename string) {
	seed, err := os.Open(fmt.Sprintf("testdata/%s.yaml", filename))
	if err != nil {
		s.T().Fatalf("failed to open seed file: %s", err)
	}
	defer seed.Close()

	err = s.dbPolluter.Pollute(seed)
	if err != nil {
		s.T().Fatalf("failed to pollute with file %s: %s", filename, err)
	}
}

func (s *AbstractSuite) setupOnce() {
	var err error

	cfg := config.New()
	dbConfig := cfg.DB
	s.app = server.New([]byte("some secret key"), cfg)
	dsn := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	s.DBConn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = s.DBConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	s.cleanupTables()

	s.dbPolluter = polluter.New(polluter.PostgresEngine(s.DBConn))
	s.sqlPgx = dbcon.New(context.Background(), dbConfig)
}

func (s *AbstractSuite) executeRequest(req *http.Request, auth *string) (string, int) {
	req.Header.Set("content-type", "application/json")
	if auth != nil {
		req.Header.Set("authorization", *auth)
	}
	rr := httptest.NewRecorder()
	s.app.HttpHandler.ServeHTTP(rr, req)

	return rr.Body.String(), rr.Code
}

func (s *AbstractSuite) checkResponseCode(expected, actual int) {
	if expected != actual {
		s.T().Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (s *AbstractSuite) getAuthToken() string {
	req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"username":"Admin","password":"12345678"}`))
	response, code := s.executeRequest(req, nil)

	s.checkResponseCode(http.StatusOK, code)
	s.Require().Contains(response, "token")

	fd := &struct {
		Token string `json:"token"`
	}{}
	err := json.Unmarshal([]byte(response), fd)
	if err != nil {
		panic(err)
	}

	return "Bearer " + fd.Token
}
