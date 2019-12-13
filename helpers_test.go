package usage

import (
	"bytes"
	"errors"
	"testing"
)

func TestCheckErrorWithPanic(t *testing.T) {
	type wants struct {
		bufferValue string
		panic       bool
	}
	testSet := []struct {
		name string
		got  error
		want wants
	}{
		{
			name: "Empty test",
			got:  nil,
			want: wants{
				bufferValue: "",
				panic:       false,
			},
		},
		{
			name: "Empty error",
			got:  errors.New(""),
			want: wants{
				bufferValue: "",
				panic:       true,
			},
		},
		{
			name: "Error with test data",
			got:  errors.New("Test"),
			want: wants{
				bufferValue: "test",
				panic:       true,
			},
		},
	}

	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("code did not panic, but want it")
					}
				}()
			}

			writer := bytes.NewBufferString("")
			CheckErrorWithPanic(writer, tt.got)

			if buff := string(writer.Bytes()); buff != tt.want.bufferValue {
				t.Errorf("got writer value %s want %s", buff, tt.want.bufferValue)
			}
		})
	}
}
