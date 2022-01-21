package main

import(
  "fmt"
  "form3/account/models"
  "form3/account/client"
)

func main() {
  country := "GB";
  matchingOpt := false;
  classification := "Personal";
  jointAccount := false;
  account := &models.AccountData {
    ID: "0f0d5bfd-c041-44cc-82b9-7c01214b9354",
    OrganisationID: "be48dd41-9358-4670-97cd-64b93ad42324",
    Type: "accounts",
    Attributes: &models.AccountAttributes {
      Country: &country,
      BaseCurrency: "GBP",
      BankID: "400302",
      BankIDCode: "GBDSC",
      Bic: "NWBKGB42",
      Name: []string{"Samantha Holder"},
      AlternativeNames: []string{"Sam Holder"},
      AccountClassification: &classification,
      JointAccount: &jointAccount,
      AccountMatchingOptOut: &matchingOpt,
      SecondaryIdentification: "A1B2C3D4",
    },
  };
  createErr := client.Create(account);

  if (createErr != nil) {
    fmt.Println(createErr.Error());
    return;
  }

  fmt.Println("Created account with id", account.ID);

  data, fetchErr := client.Fetch(&models.AccountData {ID: "0f0d5bfd-c041-44cc-82b9-7c01214b9354"});

  if (fetchErr != nil) {
    fmt.Println(fetchErr.Error());
    return;
  }

  fmt.Println("Account type is", data.Type);

  deleteErr := client.Delete(account);

  if (deleteErr != nil) {
    fmt.Println(deleteErr.Error());
    return;
  }

  fmt.Println("Deleted account with id", account.ID);
}
