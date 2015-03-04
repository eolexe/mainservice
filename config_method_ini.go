package mainservice

import (
	"code.google.com/p/gcfg"
)

func ConfigMethodINI(config interface{}, paths...string) error {
	for _, path := range paths {
		err := gcfg.ReadFileInto(config, path)
		if err != nil {
			return err
		}
	}
	return nil
}