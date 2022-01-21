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
  }

  return true, err;
}

func Delete(payload *AccountData) (bool, error)  {
  return DeleteById(&payload.ID);
}

func DeleteById(id *string) (bool, error) {
  url := fmt.Sprintf("%s/v1/organisation/accounts/%s?version=0", host, *id);
  fmt.Println(url);
  req, reqErr := http.NewRequest(http.MethodDelete, url, nil);

  if reqErr != nil {
    fmt.Println(reqErr.Error);
    return false, reqErr;
  }

  client := &http.Client{}
  _, err := client.Do(req)

  if err != nil {
    fmt.Println(err.Error);
    return false, err;
  }

  return true, err;
}
