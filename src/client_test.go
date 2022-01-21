package client

import (
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
  "form3/account/models"
)

func TestGetUrl(t *testing.T) {
  host = "https://blah";

  wanted := "https://blah/v1/organisation/accounts"
  got := getUrl();

  if got != wanted {
    t.Errorf("got %q, wanted %q", got, wanted)
  }

  wanted = "https://blah/v1/organisation/accounts/blah"
  got = getUrl("blah");

  if got != wanted {
    t.Errorf("got %q, wanted %q", got, wanted)
  }
}

func TestCreate(t *testing.T) {
  wantUrl := "/v1/organisation/accounts"
  wantMethod := "POST"
  wantBody := "{\"data\":{\"id\":\"foo\"}}";
  gotUrl := "";
  gotMethod := "";
  gotBody := "";
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated);
    gotUrl = r.URL.String();
    gotMethod = r.Method;
    body, _ := io.ReadAll(r.Body);
    gotBody = string(body);
  }))

  host = srv.URL;

  err := Create(&models.AccountData {ID: "foo"});
  srv.Close();

  if err != nil {
    t.Errorf("Unexpected error on request: %s", err);
  }
  if gotUrl != wantUrl {
    t.Errorf("want url %s, got %s", wantUrl, gotUrl);
  }
  if gotMethod != wantMethod {
    t.Errorf("want http method %s, got %s", wantMethod, gotMethod);
  }
  if gotBody != wantBody {
    t.Errorf("want http request body %s, got %s", wantBody, gotBody);
  }
}

func TestCreateFail(t *testing.T) {
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Bad Request", http.StatusBadRequest);
    r.Body.Close();
    return;
  }))

  host = srv.URL;

  err := Create(&models.AccountData {ID: "foo"});
  srv.Close();

  if err == nil {
    t.Errorf("didn't return the error");
  }
}
