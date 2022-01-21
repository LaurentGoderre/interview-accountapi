package client

import(
  "encoding/json"
  "errors"
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

func getUrl(params ...string) (string) {
  if (len(params) > 0) {
    return fmt.Sprintf("%s/%s/%s", host, accountRootUrl, params[0]);
  }

  return fmt.Sprintf("%s/%s", host, accountRootUrl);
}

func Create(payload *models.AccountData) (error) {
  url := getUrl();
  json, _ := json.Marshal(AccountDataPayload{Data: payload});
  res, err := http.Post(url, jsonMime, strings.NewReader(string(json)));

  if err != nil {
    return err;
  }

  if (res.StatusCode != http.StatusCreated) {
    return errors.New(res.Status)
  }

  return nil;
}

func Fetch(payload *models.AccountData) (*models.AccountData, error) {
  return FetchById(&payload.ID);
}

func FetchById(id *string) (*models.AccountData, error) {
  url := getUrl(*id);
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

func Delete(payload *models.AccountData) (error)  {
  return DeleteById(&payload.ID);
}

func DeleteById(id *string) (error) {
  url := fmt.Sprintf("%s?version=0", getUrl(*id));
  req, reqErr := http.NewRequest(http.MethodDelete, url, nil);

  if reqErr != nil {
    return reqErr;
  }

  client := &http.Client{}
  res, err := client.Do(req)

  if err != nil {
    return err;
  }

  if (res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent) {
    return errors.New(res.Status)
  }

  return nil;
}
