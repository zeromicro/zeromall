package graphql

import (
	"reflect"
	"testing"
)

func TestQueryCreateMarketEvents_ToJsonString(t *testing.T) {

	in := &QueryCreateMarketEvents{
		[]CreateMarketEvent{
			CreateMarketEvent{
				ID:        "111",
				Timestamp: "1343243343",
				Outcomes: []string{
					"aaa",
					"bbb",
				},
			},
		},
	}
	out, _ := in.ToJsonString()
	t.Logf("resp: type=%v, out=%v", reflect.TypeOf(out), out)
}

func TestQueryCreateMarketEvents_FromJson(t *testing.T) {
	raw := `
{
	"CreateMarketEvents":
	[
		{
			"ID":"111",
			"Timestamp":"1343243343",
			"MarketCreator":"Creator",
			"Outcomes":["aaa","bbb"],
			"MaxPrice":"22",
			"MinPrice":"33",
			"Topic":"",
			"MarketCreationFee":"",
			"Block":"",
			"Address":"",
			"Universe":"",
			"Market": {
				"ID":"",
				"Description":"",
				"MarketType":"",
				"ExtraInfoRaw":"",
				"Status":""
			}
		}
	]
}
`

	item := new(QueryCreateMarketEvents)
	err := item.FromJson([]byte(raw))
	t.Logf("FromJson: %+v, err: %v", item, err)
}

func TestQueryFinalizeMarketEvents_ToJsonString(t *testing.T) {

}
