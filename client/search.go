package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type UIDMetaDataSearchResults struct {
	Results []UIDMetaData `json:"results,omitempty"`
	Query   string        `json:"query"`
	Type    string        `json:"type"`

	StatusCode int
	ErrorInfo  map[string]interface{} `json:"error,omitempty"`
}

func (res *UIDMetaDataSearchResults) SetStatus(code int) {
	res.StatusCode = code
}

func (res *UIDMetaDataSearchResults) GetCustomParser() func(respCnt []byte) error {
	return nil
}

func (res *UIDMetaDataSearchResults) String() string {
	bslice, err := json.Marshal(res)
	if err != nil {
		return "could not stringify"
	}
	return string(bslice)
}

func (c *clientImpl) SearchUIDMeataData(query string) (*UIDMetaDataSearchResults, error) {
	encodedQuery := url.Values{"query": []string{query}}.Encode()
	resp := UIDMetaDataSearchResults{}
	endpoint := fmt.Sprintf("%s%s?%s", c.tsdbEndpoint, SearchUIDMetaPath, encodedQuery)
	if err := c.sendRequest(GetMethod, endpoint, "", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
