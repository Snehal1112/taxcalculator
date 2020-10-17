package api

import (
	"encoding/json"
	"github.com/snehal1112/taxcal/model"
	"net/http"
)

func (api *API) InitRule() {
	rule := api.BaseRoutes["Rule"]
	rule.Handle("/list", api.ApiHandler(listRules)).Methods("GET")
	rule.Handle("", api.ApiHandler(createRules)).Methods("POST")
}

func createRules(c *Context, w http.ResponseWriter, r *http.Request) {
	rule := model.NewRule()
	rule.FromJson(r.Body)

	err := c.App.CreateRule(rule)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{success: false}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{success: true}`))
}

func listRules(c *Context, w http.ResponseWriter, r *http.Request) {
	rules, err := c.App.ListRule()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{success: false}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(rules)
	w.Write(result)
}
