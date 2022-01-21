package client

import (
  "testing"
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

}
