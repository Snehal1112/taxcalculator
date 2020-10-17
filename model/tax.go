package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Tax struct {
	Name string `json:"name" bson:"name"`
	Age int `json:"age" bson:"age"`
	Year int `json:"year" bson:"year"`
	Income int `json:"income" bson:"income"`
	Investment int `json:"investment" bson:"investment"`
	TaxPayable int `json:"tax_payable" bson:"tax_payable"`
	Cess int `json:"cess" bson:"cess"`
	TotalTax int `json:"total_tax" bson:"total_tax"`
}

type TaxCalculations []*Tax

func NewTax() *Tax {
	return &Tax{}
}

// ToJson convert a Tax to a json string
func (t *Tax) ToJson() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJson will decode the input and return a Tax
func (t *Tax)FromJson(data io.Reader) *Tax {
	decoder := json.NewDecoder(data)

	err := decoder.Decode(&t)
	if err == nil {
		return t
	}
	return nil
}

// ToBSON will decode the input and return a Tax
func (t *Tax)ToBSON(data bson.M) *Tax {
	uProps, _ := bson.Marshal(data)
	bson.Unmarshal(uProps, &t)
	return t
}
