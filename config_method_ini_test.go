package mainservice

import (
	"testing"
)

type testConfigINI struct {
	Network struct {
		Host string
		Port uint16
	}
}

func TestConfigMethodINI(t *testing.T) {
	config := testConfigINI{}
	if err := ConfigMethodINI(&config, "test/config.ini"); err != nil {
		t.Fatal(err)
	}
	if config.Network.Host != "localhost" {
		t.Fatal("field network.host does not match localhost")
	} else if config.Network.Port != 19091 {
		t.Fatal("field network.port does not match 19091")
	}
}