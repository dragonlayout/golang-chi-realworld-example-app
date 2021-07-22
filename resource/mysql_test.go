package resource

import (
	"database/sql"
	"github.com/dragonlayout/golang-chi-realworld-example-app/config"
	"testing"
)

func TestInitMysqlPool(t *testing.T) {
	type args struct {
		conf *config.MysqlConfig
	}
	tests := []struct {
		name       string
		args       args
		wantDbConn *sql.DB
		wantErr    bool
	}{
		{
			name:       "example",
			args:       args{
				conf: &config.MysqlConfig{
					Address:         "127.0.0.1",
					Port:            3306,
					User:            "root",
					Password:        "123456",
					Protocol:        "tcp",
					Database:        "real_world",
					ConnMaxLifetime: 10,
					ConnMaxIdleTime: 5,
					MaxIdleConns:    10,
					MaxOpenConns:    5,
				},
			},
			wantDbConn: new(sql.DB),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InitMysqlPool(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitMysqlPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotDbConn, tt.wantDbConn) {
			//	t.Errorf("InitMysqlPool() gotDbConn = %v, want %v", gotDbConn, tt.wantDbConn)
			//}
		})
	}
}
