package client

import (
	"encoding/json"
	"fmt"
	"log"
)

type TSMetaDataSearchRequestParams struct {
	Query      string
	Limit      int64
	UseMeta    bool
	StartIndex int64
	Metric     string
	Tags       []TagKeyValue
}

type TSMetaDataLookupResponse struct {
	StatusCode int
	Type       string            `json:"type"`
	Query      string            `json:"query"`
	Limit      int64             `json:"limit"`
	StartIndex int64             `json:"startIndex"`
	Metric     string            `json:"metric"`
	Time       float64           `json:"time"`
	Tags       []TagKeyValue     `json:"tags"`
	Results    []TerseTSMetaData `json:"results"`
}

func (r *TSMetaDataLookupResponse) SetStatus(code int) {
	r.StatusCode = code
}

func (r *TSMetaDataLookupResponse) GetCustomParser() func([]byte) error {
	return nil
}

func (r *TSMetaDataLookupResponse) String() string {
	content, _ := json.Marshal(r)
	return string(content)
}

// there are different representation of the tsmetadata this is the simplest
// version: name, uid and a tags map
type TerseTSMetaData struct {
	Tags   map[string]string `json:"tags"`
	Metric string            `json:"metric"`
	TSUID  string            `json:"tsuid"`
}

type TagKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c *clientImpl) TSMetaDataLookup(
	searchRequest TSMetaDataSearchRequestParams) (*TSMetaDataLookupResponse, error) {
	body, err := json.Marshal(struct {
		Metric string        `json:"metric"`
		Tags   []TagKeyValue `json:"tags"`
	}{
		Metric: searchRequest.Metric,
		Tags:   searchRequest.Tags,
	})
	if err != nil {
		return nil, err
	}
	log.Print(string(body))
	searchTSMetaEndpoint := fmt.Sprintf("%s%s", c.tsdbEndpoint, TSMetaDataLookup)
	tsMetadDataSearchResp := TSMetaDataLookupResponse{}
	if err := c.sendRequest(PostMethod,
		searchTSMetaEndpoint,
		string(body),
		&tsMetadDataSearchResp); err != nil {
		return nil, err
	}
	return &tsMetadDataSearchResp, nil
}

type TSMetaDataSearchResponse struct {
	StatusCode int
	Type       string        `json:"type"`
	Query      string        `json:"query"`
	Limit      int64         `json:"limit"`
	StartIndex int64         `json:"startIndex"`
	Metric     string        `json:"metric"`
	Time       float64       `json:"time"`
	Tags       []TagKeyValue `json:"tags"`
	Results    []TSMetaData  `json:"results"`
}

func (r *TSMetaDataSearchResponse) SetStatus(code int) {
	r.StatusCode = code
}

func (r *TSMetaDataSearchResponse) GetCustomParser() func([]byte) error {
	return nil
}

func (r *TSMetaDataSearchResponse) String() string {
	content, _ := json.Marshal(r)
	return string(content)
}

func (c *clientImpl) SearchTSMetaData(
	searchRequest TSMetaDataSearchRequestParams) (*TSMetaDataSearchResponse, error) {
	body, err := json.Marshal(struct {
		Query      string `json:"query"`
		Limit      int64  `json:"limit"`
		StartIndex int64  `json:"startIndex"`
	}{
		Query:      searchRequest.Query,
		Limit:      searchRequest.Limit,
		StartIndex: searchRequest.StartIndex,
	})
	if err != nil {
		return nil, err
	}
	log.Print(string(body))

	searchTSMetaEndpoint := fmt.Sprintf("%s%s", c.tsdbEndpoint, TSMetaSearchPath)
	log.Print(searchTSMetaEndpoint)
	tsMetadDataSearchResp := TSMetaDataSearchResponse{}
	if err := c.sendRequest(PostMethod,
		searchTSMetaEndpoint,
		string(body),
		&tsMetadDataSearchResp); err != nil {
		return nil, err
	}
	return &tsMetadDataSearchResp, nil
}
