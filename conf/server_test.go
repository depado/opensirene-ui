package conf

import "testing"

func TestServer_ListenString(t *testing.T) {
	type fields struct {
		Host  string
		Port  int
		Debug bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{Host: "", Port: 0}, ":0"},
		{"standard", fields{Host: "127.0.0.1", Port: 8080}, "127.0.0.1:8080"},
		{"open", fields{Host: "0.0.0.0", Port: 8080}, "0.0.0.0:8080"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Host:  tt.fields.Host,
				Port:  tt.fields.Port,
				Debug: tt.fields.Debug,
			}
			if got := s.ListenString(); got != tt.want {
				t.Errorf("Server.ListenString() = %v, want %v", got, tt.want)
			}
		})
	}
}
