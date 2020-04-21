package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// test entrypoint
func TestMain(m *testing.M) {
	//fmt.Println("Initing tests..")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL:        "https://example.com/users/login",
		// expected body
		ReqBody: `{"email":"email@gmail.com", "password":"the-password"}`,
		// response planned
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	//assert.EqualValues(t, "", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
