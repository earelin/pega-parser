package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_toString(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   string
	}{
		{
			name:   "Empty",
			config: Config{},
			want:   "/",
		},
		{
			name:   "Database",
			config: Config{Database: "pega"},
			want:   "/pega",
		},
		{
			name:   "Host",
			config: Config{Host: "localhost"},
			want:   "localhost/",
		},
		{
			name: "No password",
			config: Config{
				Host:     "localhost",
				User:     "root",
				Database: "pega",
			},
			want: "root@localhost/pega",
		},
		{
			name: "Full",
			config: Config{
				Host:     "localhost",
				User:     "root",
				Password: "secret",
				Database: "pega",
			},
			want: "root:secret@localhost/pega",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.config.toString())
		})
	}
}
