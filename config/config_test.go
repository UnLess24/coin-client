package config

import "testing"

func TestMustRead(t *testing.T) {
	tests := []struct {
		name      string
		fns       []Option
		wantError bool
	}{
		{
			name:      "Valid",
			fns:       []Option{WithConfigPath("..")},
			wantError: false,
		},
		{
			name:      "Invalid",
			fns:       []Option{WithConfigName("invalid"), WithConfigPath("..")},
			wantError: true,
		},
		{
			name:      "Valid config file",
			fns:       []Option{WithConfigPath(".."), WithConfigName("config")},
			wantError: false,
		},
		{
			name:      "InValid",
			fns:       []Option{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			func() {
				defer func() {
					if r := recover(); r != nil && !tt.wantError {
						t.Errorf("MustRead() error = %v, doesn't want panic %v", r, tt.wantError)
					} else if tt.wantError && r == nil {
						t.Errorf("MustRead() want panic %v, but it didn't", tt.wantError)
					}
				}()
				MustRead(tt.fns...)
			}()
		})
	}
}
