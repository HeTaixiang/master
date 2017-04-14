package utils

import "testing"

func TestNewDB(t *testing.T) {
	config, err := ReadToConfig("../db.json")
	if err != nil {
		t.Fatal("test failure 1")
	} else if len(config.DB) == 0 || config.DB != "product" {
		t.Log("test faulure !")
	} else {
		t.Log("test success")
	}

}

func TestFailreNewDB(t *testing.T) {
	_, err := ReadToConfig("")
	if err != nil {
		t.Log("test success !")
	}
}
