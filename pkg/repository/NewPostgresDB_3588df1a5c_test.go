package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewPostgresDB_3588df1a5c(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.DB
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				cfg: Config{
					Host:     "localhost",
					Port:     "5432",
					Username: "postgres",
					Password: "my-secret-password",
					DBName:   "my-database",
					SSLMode:  "disable",
				},
			},
			want: &sqlx.DB{},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				cfg: Config{
					Host:     "localhost",
					Port:     "5432",
					Username: "postgres",
					Password: "wrong-password",
					DBName:   "my-database",
					SSLMode:  "disable",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresDB(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewPostgresDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}
