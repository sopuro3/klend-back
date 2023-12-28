package argon2

import (
	"encoding/base64"
	"github.com/sopuro3/klend-back/pkg/password"
	"reflect"
	"testing"
)

// $argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY
// input text password
// salt 12345678901234567890123456789012
// m 64x1024
// t 1
// p 4

func Test_createHashPassword(t *testing.T) {
	type args struct {
		rawPassword string
		salt        []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success", args{"password", []byte("12345678901234567890123456789012")}, "WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY"},
		{"success", args{"test", []byte("12345678901234567890123456789012")}, "Dh3MvyBffzwDsnqMVErNbtLMHJeuI69CvRi3OLE2Of4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := createHashPassword(tt.args.rawPassword, tt.args.salt)
			b64Got := base64.RawStdEncoding.EncodeToString(got)
			if !reflect.DeepEqual(b64Got, tt.want) {
				t.Errorf("createHashPassword() = %v, want %v", b64Got, tt.want)
			}
		})
	}
}

func Test_createEncodedPassword(t *testing.T) {
	type args struct {
		hashedPassword []byte
		salt           []byte
	}
	tests := []struct {
		name string
		args args
		want password.EncodedPassword
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createEncodedPassword(tt.args.hashedPassword, tt.args.salt); got != tt.want {
				t.Errorf("createEncodedPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2Encoder_EncodePassword(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    password.EncodedPassword
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Argon2Encoder{}
			got, err := e.EncodePassword(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2Encoder_IsMatchPassword(t *testing.T) {
	type args struct {
		inputPassword  string
		storedPassword password.EncodedPassword
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Argon2Encoder{}
			got, err := e.IsMatchPassword(tt.args.inputPassword, tt.args.storedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsMatchPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsMatchPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeHash(t *testing.T) {
	type args struct {
		encodedPassword password.EncodedPassword
	}
	tests := []struct {
		name               string
		args               args
		wantHashedPassword []byte
		wantSalt           []byte
		wantErr            bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHashedPassword, gotSalt, err := decodeHash(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHashedPassword, tt.wantHashedPassword) {
				t.Errorf("decodeHash() gotHashedPassword = %v, want %v", gotHashedPassword, tt.wantHashedPassword)
			}
			if !reflect.DeepEqual(gotSalt, tt.wantSalt) {
				t.Errorf("decodeHash() gotSalt = %v, want %v", gotSalt, tt.wantSalt)
			}
		})
	}
}
