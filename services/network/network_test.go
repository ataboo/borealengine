package network

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/satori/go.uuid"
	"github.com/ataboo/gowssrv/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func TestTokenMiddleWare(t *testing.T) {
	fakeToken := uuid.Must(uuid.NewV4())
	fakeUser := models.User{
		ID: bson.NewObjectId(),
		Username: "test-user",
		AuthHashed: nil,
		SessionId: fakeToken.String(),
	}
	service := NewService()

	// Mock method for getting user via token/sessionId
	(accounts.(*FakeAccountDao)).FakeTTU = func(token string) (models.User, error) {
		if token != fakeToken.String() {
			return models.User{}, fmt.Errorf("invalid token")
		}

		return fakeUser, nil
	}

	// Expect the fake user model to be put into the request context
	expectUser := func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.UserPublic)
		if user.ID != fakeUser.ID {
			t.Errorf("unnexpected user model in context: %+v", user)
		}

		if user.Username != fakeUser.Username {
			t.Errorf("unnexpected username in context: %+v", user)
		}
	}

	// Expect no user model to be present in the request context
	expectNoUser := func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		if user != nil {
			t.Errorf("unnexpected user in context: %+v", user)
		}
	}

	// Rejects when Authorization absent
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ws", nil)
	service.tokenMiddleware(expectNoUser).ServeHTTP(recorder, req)

	if recorder.Code != 401 {
		t.Errorf("expected 401, got %d", recorder.Code)
	}

	// Rejects when Authorization token is invalid uuid
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ws", nil)
	req.Header.Add("Authorization", "this-is-no-token")
	service.tokenMiddleware(expectNoUser).ServeHTTP(recorder, req)

	if recorder.Code != 401 {
		t.Errorf("expected 401, got %d", recorder.Code)
	}

	// Rejects when Authorization token doesn't match uuid
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ws", nil)
	req.Header.Add("Authorization", uuid.Must(uuid.NewV4()).String())
	service.tokenMiddleware(expectNoUser).ServeHTTP(recorder, req)

	if recorder.Code != 401 {
		t.Errorf("expected 401, got %d", recorder.Code)
	}

	// Passes when token matches
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ws", nil)
	req.Header.Add("Authorization", fakeToken.String())
	service.tokenMiddleware(expectUser).ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("expected 200, got %d", recorder.Code)
	}
}
