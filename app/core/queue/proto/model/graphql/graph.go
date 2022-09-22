package graphql

import "encoding/json"

// graphql query: 用于接收 graphql 查询结果
type Query struct {
	Markets []struct {
		ID          string
		Description string
	}
}

////////////////////////////////////////////////////////////////////////////////////////////

// query: func
type QueryCreateMarketEvents struct {
	// array:
	CreateMarketEvents []CreateMarketEvent `graphql:"createMarketEvents(orderBy: timestamp, first: 100, orderDirection: desc)"`
}

type CreateMarketEvent struct {
	ID                string
	Timestamp         string
	MarketCreator     string
	Outcomes          []string
	MaxPrice          string
	MinPrice          string
	Topic             string
	MarketCreationFee string
	Block             string
	Address           string
	Universe          string

	Market struct {
		ID           string
		Description  string
		MarketType   string
		ExtraInfoRaw string
		Status       string
	}
}

func (m *QueryCreateMarketEvents) FromJson(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *QueryCreateMarketEvents) ToJsonString() (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

////////////////////////////////////////////////////////////////////////////////////////////

//
type QueryFinalizeMarketEvents struct {
	// array: top 10
	FinalizeMarketEvents []FinalizeMarketEvent `graphql:"finalizeMarketEvents(orderBy: timestamp, first: 10, orderDirection: desc)"`
}

type FinalizeMarketEvent struct {
	ID        string
	Timestamp string
	Block     string
	Market    struct {
		ID     string
		Status string
	}
}

func (m *QueryFinalizeMarketEvents) FromJson(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *QueryFinalizeMarketEvents) ToJsonString() (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
