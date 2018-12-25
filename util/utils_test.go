package util

import "testing"

func TestNginxTip(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ssss"},
	}
	for range tests {
		t.Run("", func(t *testing.T) {
			NginxTip()
		})
	}
}
