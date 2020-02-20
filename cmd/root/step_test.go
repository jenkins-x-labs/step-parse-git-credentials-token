package root

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test_options_getToken(t *testing.T) {

	tests := []struct {
		name    string
		filename string
		wantErr bool
	}{
		{name: "test_single_entry", filename:filepath.Join("test_data","test_single_entry"), wantErr: false},
		{name: "test_disabled_entry", filename:filepath.Join("test_data","test_disabled_entry"), wantErr: false},
		{name: "test_multiple_entries", filename:filepath.Join("test_data","test_multiple_entries"), wantErr: true},
		{name: "test_empty_file", filename:filepath.Join("test_data","test_empty_file"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &options{}
			if token, err := o.getToken(tt.filename); (err != nil) != tt.wantErr {
				assert.Equal(t, "wine", token)
				t.Errorf("getToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}