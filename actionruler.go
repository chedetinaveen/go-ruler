package goruler

import (
	"encoding/json"
)

// ActionRuler returns action when query satisfies the ruler
type ActionRuler struct {
	Ruler  *Ruler      `json:"ruler"`
	Action interface{} `json:"action"`
}

// NewActionRuler ...
func NewActionRuler(ruler *Ruler, action interface{}) *ActionRuler {
	return &ActionRuler{
		Ruler:  ruler,
		Action: action,
	}
}

// NewActionRulerWithJSON ...
func NewActionRulerWithJSON(jsonstr []byte) (*ActionRuler, error) {
	ruler := ActionRuler{}
	err := json.Unmarshal(jsonstr, &ruler)
	if err != nil {
		return nil, err
	}
	return &ruler, nil
}

// Test tests the actionruler
func (r *ActionRuler) Test(o map[string]interface{}) interface{} {
	if r.Ruler == nil {
		return nil
	}

	if r.Ruler.Test(o) {
		return r.Action
	}
	return nil
}
