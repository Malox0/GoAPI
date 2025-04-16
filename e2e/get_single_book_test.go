package e2e_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetSingleBookSuite struct {
	suite.Suite
}

func TestGetSingleBookSuite(t *testing.T) {
	suite.Run(t, new(GetSingleBookSuite))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")

	body, _ := io.ReadAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(`{"Code" : "001", "Msg" : "No book with ISBN 123456789"}`, string(body))
}
