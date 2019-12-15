package usage

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewConfigurationFromReader(t *testing.T) {
	type yamlConfig struct {
		Test1 string `yaml:"test1"`
		Test2 string `yaml:"test2"`
		Test3 string `yaml:"test3"`
	}
	stream := bytes.NewBufferString("")

	config, err := NewConfigurationFromReader(stream, &yamlConfig{})

	if err != nil {
		t.Errorf("code return error, but did not want")
	}

	if !reflect.DeepEqual(config, &yamlConfig{}) {
		t.Errorf("got %v want %v", config, yamlConfig{})
	}
}
