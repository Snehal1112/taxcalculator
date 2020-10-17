package app

import "github.com/snehal1112/taxcal/model"

var taxCollection = "calculation"

func (a *App) SaveTaxCalculation(tax *model.Tax) error {
	return getSession(a).CreateDocument(taxCollection, tax)
}

func (a *App) ListCalculation() (model.TaxCalculations, error) {
	rules, err := getSession(a).Search(taxCollection, nil,0,0)
	rulesObj := make(model.TaxCalculations, 0)
	for _, item := range rules {
		rulesObj = append(rulesObj, model.NewTax().ToBSON(item))
	}
	return rulesObj, err
}