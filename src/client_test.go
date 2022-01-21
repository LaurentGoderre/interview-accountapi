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
    gotUrl = r.URL.String();
    gotMethod = r.Method;
    body, _ := io.ReadAll(r.Body);
    gotBody = string(body);
    w.WriteHeader(http.StatusCreated);
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

func TestFetch(t *testing.T) {
  wantUrl := "/v1/organisation/accounts/foo"
  wantMethod := "GET"
  wantAccount := &models.AccountData {ID: "foo", Type: "test"};
  gotUrl := "";
  gotMethod := "";
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    gotUrl = r.URL.String();
    gotMethod = r.Method;
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("{\"data\":{\"id\":\"foo\",\"type\":\"test\"}}"));
  }))

  host = srv.URL;

  gotAccount, err := Fetch(&models.AccountData {ID: "foo"});
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
  if gotAccount.Type != wantAccount.Type {
    t.Errorf("want http request body %s, got %s", wantAccount.Type, gotAccount.Type);
  }
}

func TestFetchFail(t *testing.T) {
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    switch r.URL.String() {
    case "/v1/organisation/accounts/foo":
      http.Error(w, "Not Found", http.StatusNotFound);
    case "/v1/organisation/accounts/bar":
      w.WriteHeader(http.StatusOK);
      w.Write([]byte("not json!"));
    }
    r.Body.Close();
    return;
  }))

  host = srv.URL;

  _, httpErr := Fetch(&models.AccountData {ID: "foo"});
  _, jsonErr := Fetch(&models.AccountData {ID: "bar"});
  srv.Close();

  if httpErr == nil {
    t.Errorf("didn't return the http error");
  }
  if jsonErr == nil {
    t.Errorf("didn't return the json error");
  }
}

func TestDelete(t *testing.T) {
  wantUrl := "/v1/organisation/accounts/foo?version=0"
  wantMethod := "DELETE"
  gotUrl := "";
  gotMethod := "";
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    gotUrl = r.URL.String();
    gotMethod = r.Method;
    w.WriteHeader(http.StatusOK);
  }))

  host = srv.URL;

  err := Delete(&models.AccountData {ID: "foo"});
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
}

func TestDeleteFail(t *testing.T) {
  srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Bad Request", http.StatusBadRequest);
    r.Body.Close();
    return;
  }))

  host = srv.URL;

  err := Delete(&models.AccountData {ID: "foo"});
  srv.Close();

  if err == nil {
    t.Errorf("didn't return the error");
  }
}
