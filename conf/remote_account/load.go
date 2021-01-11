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

func init() {
	ReadConfig()
}
func GetConfig() *Config {
	return config
}

func ReadConfig() {
	// parse remote_account.json
	file_ptr, _ := os.Open(config_path)
	defer file_ptr.Close()
	data, _ := ioutil.ReadAll(file_ptr)
	config = new(Config)
	_ = json.Unmarshal(data, &config.Account)
}
