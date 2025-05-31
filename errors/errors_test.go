package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestErr_ImplementsError(t *testing.T) {
	assert.Implements(t, (*error)(nil), Err{})
}

func TestAddTrace(t *testing.T) {
	type args struct {
		err  error
		errs []Err
	}
	tests := []struct {
		name string
		args args
		want Err
	}{
		{
			name: "when normal error",
			args: args{
				err: errors.New("error1"),
			},
			want: Err{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, _ := os.Getwd()

			assert.Equal(t)
			assert.Equalf(t, tt.want, AddTrace(tt.args.err, tt.args.errs...), "AddTrace(%v, %v)", tt.args.err, tt.args.errs)
		})
	}
}
