package useful

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

	testSet := []struct {
		testName  string
		arguments string
		wantError bool
		want      interface{}
	}{
		{
			"empty test",
			"",
			false,
			&yamlConfig{},
		},
		{
			"test with error",
			"test1:1\ntest2:2 test3:3",
			true,
			&yamlConfig{
				Test1: "1",
				Test2: "2",
				Test3: "3",
			},
		},
		{
			"correct test",
			"test1: 1\ntest2: 2\ntest3: 3\n",
			false,
			&yamlConfig{
				Test1: "1",
				Test2: "2",
				Test3: "3",
			},
		},
	}

	for _, testData := range testSet {
		t.Run(testData.testName, func(t *testing.T) {
			stream := bytes.NewBufferString(testData.arguments)
			config, err := NewConfigurationFromReader(stream, &yamlConfig{})

			if err != nil {
				if testData.wantError {
					return
				}

				t.Errorf("code return error, but did not want it %v", err)
			} else if testData.wantError {
				t.Errorf("code did not return error, but want it")
			}

			if !reflect.DeepEqual(config, testData.want) {
				t.Errorf("got %v want %v", config, testData.want)
			}
		})
	}
}
