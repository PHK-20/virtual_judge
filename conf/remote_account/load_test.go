package remote_account

import (
	"testing"
)

func TestLoad(t *testing.T) {
	err := ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(GetConfig())
}
