package useful

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func NewConfigurationFromReader(stream *bytes.Buffer, configuration interface{}) (interface{}, error) {
	yamlData, err := ioutil.ReadAll(stream)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlData, configuration)

	if err != nil {
		return nil, err
	}

	return configuration, err
}
