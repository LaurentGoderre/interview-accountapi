package main

import(
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "strings"
)

var host = os.Getenv("API_HOST");
var jsonMime = "application/json"

type AccountDataPayload struct {
    Data *AccountData `json:"data"`
}

func Create(payload *AccountData) (bool, error) {
  url := fmt.Sprintf("%s/v1/organisation/accounts", host);
  json, _ := json.Marshal(AccountDataPayload{Data: payload});
  _, err := http.Post(url, jsonMime, strings.NewReader(string(json)));

  if err != nil {
    return false, err;
  } else {
    return true, err;
  }
}
