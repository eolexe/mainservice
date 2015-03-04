package mainservice

import (
	"testing"
)

type testConfigJSON struct {
	Host string
	Port uint16
}

func TestConfigMethodJSON(t *testing.T) {
	config := testConfigJSON{}
	if err := ConfigMethodJSON(&config, "test/config.json"); err != nil {
		t.Fatal(err)
	}
	if config.Host != "localhost" {
		t.Fatal("field host does not match localhost")
	} else if config.Port != 19091 {
		t.Fatal("fiel port does not match 19091")
	}
}