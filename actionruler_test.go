package goruler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActionRulerWithJSON(t *testing.T) {
	jsonstr := []byte(`{
    "action" : 100.0,
    "ruler" : {"rules" : [{"comparator": "eq", "path": "name", "value": "Thomas"}]}}
    `)
	actionruler, err := NewActionRulerWithJSON(jsonstr)

	assert.Nil(t, err)
	assert.Equal(t, actionruler, &ActionRuler{
		Action: 100.0,
		Ruler: &Ruler{
			Rules: []*Rule{
				{
					Comparator: "eq",
					Path:       "name",
					Value:      "Thomas",
				},
			},
		}})
}

func TestActionResult(t *testing.T) {
	cases := []struct {
		desc        string
		actionRuler *ActionRuler
		testcase    map[string]interface{}
		expectedRes interface{}
	}{
		{
			desc: "Happy path",
			actionRuler: &ActionRuler{
				Ruler: &Ruler{
					Rules: []*Rule{
						{
							Value:      2,
							Path:       "x",
							Comparator: "gt",
						},
					},
				},
				Action: "Passed",
			},
			testcase: map[string]interface{}{
				"x": 3,
			},
			expectedRes: "Passed",
		},
		{
			desc: "Sad path",
			actionRuler: &ActionRuler{
				Ruler: &Ruler{
					Rules: []*Rule{
						{
							Value:      2,
							Path:       "x",
							Comparator: "gt",
						},
					},
				},
				Action: "Passed",
			},
			testcase: map[string]interface{}{
				"x": 1,
			},
			expectedRes: nil,
		},
		{
			desc: "Happy path",
			actionRuler: &ActionRuler{
				Ruler: &Ruler{
					Rules: []*Rule{
						{
							Value:      3,
							Path:       "surge",
							Comparator: "gte",
						},
						{
							Value:      4.0,
							Path:       "farePerKM",
							Comparator: "gt",
						},
					},
				},
				Action: "Fares are high due to high demand",
			},
			testcase: map[string]interface{}{
				"surge":     3.2,
				"farePerKM": 4.3,
			},
			expectedRes: "Fares are high due to high demand",
		},
		{
			desc: "Sad path",
			actionRuler: &ActionRuler{
				Ruler: &Ruler{
					Rules: []*Rule{
						{
							Value:      3,
							Path:       "surge",
							Comparator: "gte",
						},
						{
							Value:      4.0,
							Path:       "farePerKM",
							Comparator: "gt",
						},
					},
				},
				Action: "Fares are high due to high demand",
			},
			testcase: map[string]interface{}{
				"surge":     2.2,
				"farePerKM": 2.3,
			},
			expectedRes: nil,
		},
	}

	for _, scenario := range cases {
		result := scenario.actionRuler.Test(scenario.testcase)
		assert.Equal(t, result, scenario.expectedRes)
	}
}

func TestActionResultWithJSON(t *testing.T) {
	scenarios := []struct {
		desc           string
		rules          []byte
		query          map[string]interface{}
		expectedAction interface{}
	}{
		{
			desc: "Happy path",
			rules: []byte(`{
		    "action" : "Fares are high due to high demand",
		    "ruler" : {"rules" : [{"comparator": "gt", "path": "surge", "value": 2.3},{"comparator": "gte", "path": "farePerKM", "value": 4.3}]}}
				`),
			query: map[string]interface{}{
				"surge":     2.4,
				"farePerKM": 5.2,
			},
			expectedAction: "Fares are high due to high demand",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			actionRuler, err := NewActionRulerWithJSON(scenario.rules)
			assert.Nil(t, err)
			action := actionRuler.Test(scenario.query)
			assert.Equal(t, action, scenario.expectedAction)
		})
	}
}
