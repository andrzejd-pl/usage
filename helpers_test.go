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
				bufferValue: "error: \n",
				panic:       true,
			},
		},
		{
			name: "Error with test data",
			got:  errors.New("test"),
			want: wants{
				bufferValue: "error: test\n",
				panic:       true,
			},
		},
	}

	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if tt.want.panic {
					if r == nil {
						t.Error("code did not panic, but want it")
					} else if r != tt.want.bufferValue {
						t.Errorf("got writer value %s want %s", r, tt.want.bufferValue)
					}
				} else if r != nil {
					t.Error("code panic, but did not want it")
				}
			}()

			writer := bytes.NewBufferString("")
			CheckErrorWithPanic(writer, tt.got)
		})
	}
}

func TestCheckErrorWithOnlyLogging(t *testing.T) {
	testSet := []struct {
		name string
		got  error
		want string
	}{
		{
			name: "Empty test",
			got:  nil,
			want: "",
		},
		{
			name: "Empty error",
			got:  errors.New(""),
			want: "error: \n",
		},
		{
			name: "Error with test data",
			got:  errors.New("test"),
			want: "error: test\n",
		},
	}

	for _, tt := range testSet {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("code panic, but did not want it")
				}
			}()

			writer := bytes.NewBufferString("")
			CheckErrorWithOnlyLogging(writer, tt.got)

			if buff := string(writer.Bytes()); buff != tt.want {
				t.Errorf("got writer value %s want %s", buff, tt.want)
			}
		})
	}
}
