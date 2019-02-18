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

type SearchResponseCommons struct {
	StatusCode   int
	Type         string  `json:"type"`
	Query        string  `json:"query"`
	Limit        int64   `json:"limit"`
	StartIndex   int64   `json:"startIndex"`
	Metric       string  `json:"metric"`
	Time         float64 `json:"time"`
	TotalResults int64   `json:"totalResults,omitempty"`
}

type TSMetaDataLookupResponse struct {
	SearchResponseCommons
	Tags    []TagKeyValue     `json:"tags"`
	Results []TerseTSMetaData `json:"results"`
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
	SearchResponseCommons
	Results []VerboseTSMetaData `json:"results,omitempty"`
}

type VerboseTSMetaData struct {
	TSMetaData
	Custom          map[string]string `json:"custom,omitempty"`
	Units           string            `json:"units,omitempty"`
	Metric          UIDMetaData       `json:"metric,omitempty"`
	Tags            []UIDMetaData     `json:"tags,omitempty"`
	Created         int64             `json:"created,omitempty"`
	LastReceived    int64             `json:"lastReceived,omitempty"`
	TotalDatapoints int64             `json:"totalDatapoints,omitempty"`
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
	searchTSMetaEndpoint := fmt.Sprintf("%s%s", c.tsdbEndpoint, TSMetaSearchPath)
	tsMetadDataSearchResp := TSMetaDataSearchResponse{}
	if err := c.sendRequest(PostMethod,
		searchTSMetaEndpoint,
		string(body),
		&tsMetadDataSearchResp); err != nil {
		return nil, err
	}
	return &tsMetadDataSearchResp, nil
}

type TSUIDSearchResponse struct {
	SearchResponseCommons
	Results []string `json:"results,omitempty"`
}

func (r *TSUIDSearchResponse) SetStatus(code int) {
	r.StatusCode = code
}

func (r *TSUIDSearchResponse) GetCustomParser() func([]byte) error {
	return nil
}

func (r *TSUIDSearchResponse) String() string {
	content, _ := json.Marshal(r)
	return string(content)
}

func (c *clientImpl) SearchTSUID(
	searchRequest TSMetaDataSearchRequestParams) (*TSUIDSearchResponse, error) {
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
	searchTSMetaEndpoint := fmt.Sprintf("%s%s", c.tsdbEndpoint, TSUIDSearchPath)
	resp := TSUIDSearchResponse{}
	if err := c.sendRequest(PostMethod,
		searchTSMetaEndpoint,
		string(body),
		&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UIDMetaSearchResponse struct {
	SearchResponseCommons
	Results []UIDMetaData `json:"results,omitempty"`
}

func (r *UIDMetaSearchResponse) SetStatus(code int) {
	r.StatusCode = code
}

func (r *UIDMetaSearchResponse) GetCustomParser() func([]byte) error {
	return nil
}

func (r *UIDMetaSearchResponse) String() string {
	content, _ := json.Marshal(r)
	return string(content)
}

func (c *clientImpl) SearchUIDMeta(
	searchRequest TSMetaDataSearchRequestParams) (*TSUIDSearchResponse, error) {
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
	searchTSMetaEndpoint := fmt.Sprintf("%s%s", c.tsdbEndpoint, TSUIDSearchPath)
	resp := TSUIDSearchResponse{}
	if err := c.sendRequest(PostMethod,
		searchTSMetaEndpoint,
		string(body),
		&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
