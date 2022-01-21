package main

func main() {
  country := "GB";
  matchingOpt := false;
  classification := "Personal";
  jointAccount := false;
  account := &AccountData {
    ID: "0f0d5bfd-c041-44cc-82b9-7c01214b9354",
    OrganisationID: "be48dd41-9358-4670-97cd-64b93ad42324",
    Type: "accounts",
    Attributes: &AccountAttributes {
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
  Create(account);

  Delete(account);
}
