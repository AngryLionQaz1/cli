package util

import "testing"

func TestDownload(t *testing.T) {
	type args struct {
		Url  string
		Path string
	}
	tests := []struct {
		Url  string
		Path string
	}{
		{
			Url:  "https://download.jetbrains.8686c.com/idea/ideaIU-2018.3.exe",
			Path: "/cli",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Url, func(t *testing.T) {
			Download(tt.Url, tt.Path)
		})
	}
}
