package utils

import (
	"testing"
)

func TestExternalIP(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExternalIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExternalIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ip:%s", got)
		})
	}
}
