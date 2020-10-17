package app

import (
	"github.com/snehal1112/taxcal/model"
)

var rulesCollection = "rules"

func (a *App) CreateRule(rule *model.Rule) error {
	return getSession(a).CreateDocument(rulesCollection, rule)
}

func (a *App) ListRule() (model.Rules, error) {
	rules, err := getSession(a).Search(rulesCollection, nil,0,0)
	rulesObj := make(model.Rules, 0)
	for _, item := range rules {
		rulesObj = append(rulesObj, model.NewRule().ToBSON(item))
	}
	return rulesObj, err
}
