package integrationtest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/johnmcdnl/auth/auth"
	"github.com/johnmcdnl/darts-api/darts"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

type targets struct {
	User         auth.User
	ResponseCode int
}

func (t *targets) theApplicationIsRunning() error {
	go darts.StartServer()
	return nil
}

func (t *targets) iHaveAValidUser() error {

	var u auth.User
	u.Username = uuid.NewV4().String()
	u.Password = uuid.NewV4().String()

	j, _ := json.Marshal(t.User)

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:4500/darts/api/auth/register", bytes.NewReader([]byte(j)))
	client := &http.Client{}
	client.Do(req)

	req, _ = http.NewRequest(http.MethodPost, "http://localhost:4500/darts/api/auth/login", bytes.NewReader([]byte(j)))
	client = &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	var token auth.JWTAccessToken
	if err := json.Unmarshal(content, &token); err != nil {
		return err
	}

	t.User = u
	t.User.Password = string(content)

	return nil
}

func (t *targets) iMakeARequestPOSTDartsapitargets(body *gherkin.DocString) error {

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:4500/darts/api/targets", bytes.NewReader([]byte(body.Content)))
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", t.User.Password))
	client := &http.Client{}
	resp, _ := client.Do(req)

	t.ResponseCode = resp.StatusCode

	return nil
}

func (t *targets) iGetAResponse(expectedCode int) error {
	if t.ResponseCode != expectedCode {
		return errors.New(fmt.Sprint("Expected: ", expectedCode, "\t", "Actual: ", t.ResponseCode))
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	t := &targets{}

	s.Step(`^the application is running$`, t.theApplicationIsRunning)
	s.Step(`^I have a valid user$`, t.iHaveAValidUser)
	s.Step(`^I make a request POST \/darts\/api\/targets$`, t.iMakeARequestPOSTDartsapitargets)
	s.Step(`^I get a (\d+) response$`, t.iGetAResponse)
}
