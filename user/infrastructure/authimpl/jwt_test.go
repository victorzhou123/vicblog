package authimpl

import (
	"testing"
	"time"
)

const (
	usernameTest       = "victor"
	signNotExpiredTime = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM3MjI0MTI0MzcsInVzZXJuYW1lIjoidmljdG9yIn0.urI7lJlbe8v_cAZyL6MHSL4Y_-t4CbvQRhahEtS3H_0"
	signExpiredTime    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI0MTI0MzcsInVzZXJuYW1lIjoidmljdG9yIn0.OcjExnd9d5nk0SbqA5HEOM1VTGwmfg_FXBFsugE6bzY"
)

var signJwtExpiredTime, signJwtNotExpiredTime *signJwt

type expiredTimeCreator struct{}

func (t *expiredTimeCreator) AddUnix(add time.Duration) int64 {
	return 1722412437
}

type notExpiredTimeCreator struct{}

func (t *notExpiredTimeCreator) AddUnix(add time.Duration) int64 {
	return 3722412437
}

func init() {
	cfg := &Config{
		SecretKey:  "123456789",
		ExpireTime: 60,
	}

	signJwtExpiredTime = NewSignJwt(&expiredTimeCreator{}, cfg)

	signJwtNotExpiredTime = NewSignJwt(&notExpiredTimeCreator{}, cfg)
}

func Test_signJwt_genSignBasedUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		s       *signJwt
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "gen sign based on notExpired time",
			s:       signJwtNotExpiredTime,
			args:    args{usernameTest},
			want:    signNotExpiredTime,
			wantErr: false,
		}, {
			name:    "gen sign based on expired time",
			s:       signJwtExpiredTime,
			args:    args{usernameTest},
			want:    signExpiredTime,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.genSignBasedUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("signJwt.GenSignBasedUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("signJwt.GenSignBasedUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signJwt_verifyToken(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name    string
		s       *signJwt
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "verify not expired time",
			s:       signJwtNotExpiredTime,
			args:    args{signNotExpiredTime},
			want:    usernameTest,
			wantErr: false,
		}, {
			name:    "verify expired time",
			s:       signJwtExpiredTime,
			args:    args{signExpiredTime},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.verifyToken(tt.args.sign)
			if (err != nil) != tt.wantErr {
				t.Errorf("signJwt.verifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("signJwt.verifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
