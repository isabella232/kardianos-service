package service

import "testing"

func TestTemplateFunctionCmd(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			"bad'p$ass",
			"bad'p$ass",
			`'bad'"'"'p$ass'`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tf["cmd"].(func(string) string)(tt.arg); got != tt.want {
				t.Errorf("tf[cmd](%s) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestTemplateFunctionEnvValueSupervisord(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			`DATA_SOURCE_NAME=123`,
			`'123'`,
		},
		{
			`DATA_SOURCE_NAME=root:qwerty123)(*&^%$#@1@tcp(127.0.0.1:3306)/?timeout=5s`,
			`'root:qwerty123)(*&^%%$#@1@tcp(127.0.0.1:3306)/?timeout=5s'`, // double %
		},
		{
			// %40 is URL encoding for @; username is root@azure
			`DATA_SOURCE_NAME=postgres://root%40azure:qwerty123@127.0.0.1:5432/postgres?connect_timeout=5&sslmode=disable`,
			`'postgres://root%%40azure:qwerty123@127.0.0.1:5432/postgres?connect_timeout=5&sslmode=disable'`, // double %
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := tf["envValueSupervisord"].(func(string) string)(tt.arg); got != tt.want {
				t.Errorf("tf[envValueSupervisord](%s) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
