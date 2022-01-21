package client

import(
  "encoding/json"
  "io"
  "fmt"
  "net/http"
  "os"
  "strings"
  "form3/account/models"
)

var host = os.Getenv("API_HOST");
var jsonMime = "application/json"
var accountRootUrl = "v1/organisation/accounts";

type AccountDataPayload struct {
    Data *models.AccountData `json:"data"`
}

func Create(payload *models.AccountData) (bool, error) {
  url := fmt.Sprintf("%s/%s", host, accountRootUrl);
  json, _ := json.Marshal(AccountDataPayload{Data: payload});
  _, err := http.Post(url, jsonMime, strings.NewReader(string(json)));

  if err != nil {
    return false, err;
  }

  return true, err;
}

func Fetch(payload *models.AccountData) (*models.AccountData, error) {
  return FetchById(&payload.ID);
}

func FetchById(id *string) (*models.AccountData, error) {
  url := fmt.Sprintf("%s/%s/%s", host, accountRootUrl, *id);
  resp, getErr := http.Get(url);

  if getErr != nil {
    return nil, getErr;
  }

  defer resp.Body.Close();
  payload, readErr := io.ReadAll(resp.Body);

  if readErr != nil {
    return nil, readErr;
  }

  account := &AccountDataPayload{};
  jsonErr := json.Unmarshal(payload, account);

  if jsonErr != nil {
    return nil, jsonErr;
  }

  return account.Data, nil;
}

func Delete(payload *models.AccountData) (bool, error)  {
  return DeleteById(&payload.ID);
}

func DeleteById(id *string) (bool, error) {
  url := fmt.Sprintf("%s/%s/%s?version=0", host, accountRootUrl, *id);
  req, reqErr := http.NewRequest(http.MethodDelete, url, nil);

  if reqErr != nil {
    return false, reqErr;
  }

  client := &http.Client{}
  _, err := client.Do(req)

  if err != nil {
    return false, err;
  }

  return true, err;
}
