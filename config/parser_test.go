package config

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "example",
			args: args{
				filepath: "./config-example.toml",
			},
			want: &Config{
				ServerEnv: map[string]*ServerConfig{
					"development": {
						ListenAddress: "127.0.0.1",
						ListenPort:    8080,
					},
					"production": {
						ListenAddress: "127.0.0.1",
						ListenPort:    8080,
					},
				},
				MysqlEnv: map[string]*MysqlConfig{
					"development": {
						Address:  "127.0.0.1",
						Port:     6379,
						User:     "root",
						Password: "123456",
					},
					"production": {
						Address:  "127.0.0.1",
						Port:     6379,
						User:     "root",
						Password: "123456",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
