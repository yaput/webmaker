package setup

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantConfig Configuration
		wantErr    bool
	}{
		{
			name: "Success Test",
			path: "../../config/config.development.json",
			wantConfig: Configuration{
				Server: ServerConfiguration{
					Port: "8000",
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed File Not Found Test",
			path:       "config/config.nonexist.json",
			wantConfig: Configuration{},
			wantErr:    true,
		},
		{
			name:       "Failed to Decode file Test",
			path:       "test/dummy.config.json",
			wantConfig: Configuration{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := LoadConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("LoadConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
