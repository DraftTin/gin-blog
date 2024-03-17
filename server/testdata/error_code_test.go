package testdata

import (
	"encoding/json"
	"os"
	"testing"
)

func TestErrorCode(t *testing.T) {
	file := "../models/response/error_code.json"
	content, err := os.ReadFile(file)
	if err != nil {
		t.Errorf("error1: %s", err)
		return
	}
	error_code := make(map[string]string)
	err = json.Unmarshal(content, &error_code)
	if err != nil {
		t.Errorf("error2: %s", err)
	}

	if error_code["1001"] != "系统错误" {
		t.Errorf("error3: %v", error_code)
	}
}
