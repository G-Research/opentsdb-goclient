package client

import (
	"fmt"
	"testing"

	"github.com/bluebreezecf/opentsdb-goclient/config"
)

func TestTSMetaDataLookup(t *testing.T) {
	client, err := NewClient(config.OpenTSDBConfig{OpentsdbHost: "ec2-3-88-64-48.compute-1.amazonaws.com:4242"})
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.TSMetaDataLookup(TSMetaDataSearchRequestParams{
		Metric: "*",
		Tags:   []TagKeyValue{{Key: "type", Value: "tcp"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range rsp.Results {
		fmt.Println(v.Metric, v.TSUID, v.Tags)
	}
	t.Fail()
}

func TestSearchTSMetaData(t *testing.T) {
	client, err := NewClient(config.OpenTSDBConfig{OpentsdbHost: "ec2-3-88-64-48.compute-1.amazonaws.com:4242"})
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.SearchTSMetaData(TSMetaDataSearchRequestParams{
		Query: "tsuid:*",
		//		Tags:  []TagKeyValue{{Key: "type", Value: "tcp"}},
		Limit: 100,
	})
	if err != nil {
		//t.Fatal(err)
	}
	for _, v := range rsp.Results {
		fmt.Println(v)
	}
	t.Fail()
}
