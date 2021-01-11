package remote_account

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Account RemoteAccount
}

const config_path = "/home/phkli/go_project/beego_judge/conf/remote_account/remote_account.json"

var config *Config

func GetConfig() *Config {
	return config
}

func ReadConfig() error {
	// parse remote_account.json
	if config != nil {
		return nil
	}
	file_ptr, err := os.Open(config_path)
	defer file_ptr.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file_ptr)
	if err != nil {
		return err
	}
	config = new(Config)
	err = json.Unmarshal(data, &config.Account)
	if err != nil {
		return err
	}
	return nil
}
