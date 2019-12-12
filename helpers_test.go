package usage

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestCheckErrorWithPanic(t *testing.T) {
	type args struct {
		writer io.Writer
		err    error
	}
	type wants struct {
		bufferValue string
		panic       bool
	}
	testSet := []struct {
		name string
		got  args
		want wants
	}{
		{
			name: "Empty test",
			got: args{
				writer: bytes.NewBufferString(""),
				err:    nil,
			},
			want: wants{
				bufferValue: "",
				panic:       false,
			},
		},
		{
			name: "Empty error",
			got: args{
				writer: bytes.NewBufferString(""),
				err:    errors.New(""),
			},
			want: wants{
				bufferValue: "",
				panic:       true,
			},
		},
		{
			name: "Error with test data",
			got: args{
				writer: bytes.NewBufferString(""),
				err:    errors.New("Test"),
			},
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

			CheckErrorWithPanic(tt.got.writer, tt.got.err)
		})
	}
}
