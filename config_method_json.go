package mainservice

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func ConfigMethodJSON(config interface{}, paths...string) error {
	var (
		err error
		data []byte
	)
	for _, path := range paths {
		if data, err = ioutil.ReadFile(path); err != nil {
			return fmt.Errorf("json config read err: %s", err)
		}
		if err = json.Unmarshal(data, config); err != nil {
			return fmt.Errorf("json config unmarshal err: %s", err)
		}
	}
	return nil
}