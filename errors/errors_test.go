package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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
				err: errors.New("test"),
			},
			want: Err{
				Err: errors.New("test"),
				Traces: []Err{
					{
						Source: "go-lib/errors/errors_test.go:",
						Func:   "errors.TestAddTrace",
					},
				},
				Source: "go-lib/errors/errors_test.go:",
				Func:   "errors.TestAddTrace",
			},
		},
		{
			name: "when Err error",
			args: args{
				err: Err{
					Err: errors.New("test"),
					Traces: []Err{
						{
							Source: "anotherSource",
							Func:   "anotherFunc",
						},
					},
					Source: "anotherSource",
					Func:   "anotherFunc",
				},
			},
			want: Err{
				Err: errors.New("test"),
				Traces: []Err{
					{
						Source: "anotherSource",
						Func:   "anotherFunc",
					},
					{
						Source: "go-lib/errors/errors_test.go:",
						Func:   "errors.TestAddTrace",
					},
				},
				Source: "anotherSource",
				Func:   "anotherFunc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trace := AddTrace(tt.args.err, tt.args.errs...)
			if !strings.Contains(trace.Source, tt.want.Source) {
				t.Errorf("AddTrace() Source = %s, want %s", trace.Source, tt.want.Source)
			}

			if !strings.Contains(trace.Func, tt.want.Func) {
				t.Errorf("AddTrace() Source = %s, want %s", trace.Func, tt.want.Func)
			}

			if len(trace.Traces) != len(tt.want.Traces) {
				t.Errorf("AddTrace() len(Traces) = %d, want %d", len(trace.Traces), len(tt.want.Traces))
			}

			for i := range tt.want.Traces {
				resultTrace := trace.Traces[i]
				wantTrace := tt.want.Traces[i]
				if !strings.Contains(resultTrace.Source, wantTrace.Source) {
					t.Errorf("AddTrace().Traces[%d] Source = %s, want %s", i, resultTrace.Source, wantTrace.Source)
				}

				if !strings.Contains(resultTrace.Func, wantTrace.Func) {
					t.Errorf("AddTrace().Traces[%d] Func = %s, want %s", i, resultTrace.Func, wantTrace.Func)
				}
			}
		})
	}
}

func TestPrintErr(t *testing.T) {
	err := Err{
		Err: errors.New("test"),
		Traces: []Err{
			{
				Source: "anotherSource",
				Func:   "anotherFunc",
			},
		},
		Source: "anotherSource",
		Func:   "anotherFunc",
	}

	assert.Equal(t, "{\"Error\": \"test\", \"Code\": \"\", \"Sources\": [\"anotherSource\"], \"Funcs\": [\"anotherFunc\"]}", fmt.Sprintf("%v", err))
}
