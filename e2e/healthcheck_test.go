package e2e_test

import (
	"io"

	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func (s *EndToEndSuite) TestHappyHealthcheck() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/healthcheck")
	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := io.ReadAll(r.Body)
	s.JSONEq(`{"status" : "OK", "messages" : []}`, string(b))
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}
