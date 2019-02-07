package client

import (
	"fmt"
	"testing"

	"github.com/bluebreezecf/opentsdb-goclient/config"
)

func TestSearchUIDMeataData(t *testing.T) {
	client, err := NewClient(config.OpenTSDBConfig{OpentsdbHost: "ec2-3-88-64-48.compute-1.amazonaws.com:4242"})
	if err != nil {
		t.Error(err)
	}

	rs, err := client.SearchUIDMeataData("name:*")
	if err != nil {
		t.Error(err)
	}
	for _, k := range rs.Results {
		fmt.Println(k)
	}
	t.Fail()
}
