package advanced

import (
	"bytes"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {

	tests := []struct {
		intput  string
		wantW   string
		wantErr bool
	}{
		{intput: "[1,2,3]", wantW: "6", wantErr: false},
		{intput: "[]", wantW: "0", wantErr: false},
		{intput: "[a,b]", wantW: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.intput, func(t *testing.T) {
			r := strings.NewReader(tt.intput)
			w := &bytes.Buffer{}
			err := sum(r, w)

			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("sum() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
