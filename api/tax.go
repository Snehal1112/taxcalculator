package api

import (
	"encoding/json"
	"github.com/snehal1112/taxcal/model"
	"net/http"
)

func (api *API) InitTax() {
	rule := api.BaseRoutes["Tax"]
	rule.Handle("", api.ApiHandler(saveCalculation)).Methods("POST")
	rule.Handle("/list", api.ApiHandler(listCalculation)).Methods("GET")
}

func listCalculation(c *Context, w http.ResponseWriter, r *http.Request) {
	rules, err := c.App.ListCalculation()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{success: false}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(rules)
	w.Write(result)
}

func saveCalculation(c *Context, w http.ResponseWriter, r *http.Request) {
	tax := model.NewTax()
	tax.FromJson(r.Body)
	err := c.App.SaveTaxCalculation(tax)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{success: false}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{success: true}`))
}