package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Slab struct {
	LowerBound int `json:"lower_bound" bson:"lower_bound"`
	UpperBound int `json:"upper_bound" bson:"upper_bound"`
	TaxPercentage int `json:"tax_per" bson:"tax_per"`
}

type CESS struct {
	Limit int `json:"limit" bson:"limit"`
	Charge int `json:"charge" bson:"charge"`
}

type SeniorCitizen struct {
	Age       int `json:"age" bson:"age"`
	Reduction int `json:"reduction" bson:"reduction"`
}

type Rule struct {
	Year  int `json:"year" bson:"year"`
	Slabs []Slab `json:"slabs" bson:"slabs"`
	GovBondsLimit int `json:"gov_bonds_limit" bson:"gov_bonds_limit"`
	CESS `json:"cess" bson:"cess"`
	SeniorCitizen `json:"senior_citizen" bson:"senior_citizen"`
}

type Rules []*Rule

func NewRule() *Rule {
	return &Rule{}
}

// ToJson convert a Rule to a json string
func (r *Rule) ToJson() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJson will decode the input and return a Rule
func (r *Rule)FromJson(data io.Reader) *Rule {
	decoder := json.NewDecoder(data)

	err := decoder.Decode(&r)
	if err == nil {
		return r
	}

	return nil
}

// ToBSON will decode the input and return a rules
func (r *Rule)ToBSON(data bson.M) *Rule {
	rProps, _ := bson.Marshal(data)
	bson.Unmarshal(rProps, &r)
	return r
}
