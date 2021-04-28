package filepathx

import (
	"reflect"
	"testing"
)

// before run, create directory
// mkdir -p /tmp/dd/{aa,bb}/cc
// mkdir -p /tmp/dd/cc
// mkdir -p /tmp/dd/dd
func TestGlob(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test for glob",
			args: args{
				pattern: "/tmp/dd/**/cc",
			},
			want:    []string{"/tmp/dd/cc", "/tmp/dd/aa/cc", "/tmp/dd/bb/cc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Glob(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Glob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Glob() got = %v, want %v", got, tt.want)
			}
		})
	}
}
