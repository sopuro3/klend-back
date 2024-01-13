//nolint:lll,wsl,varnamelen
package argon2

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"

	"github.com/sopuro3/klend-back/pkg/password"
)

// $argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY
// input text password
// salt 12345678901234567890123456789012
// m 64x1024
// t 1
// p 4

func Test_createHashPassword(t *testing.T) {
	t.Parallel()
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := createHashPassword(tt.args.rawPassword, tt.args.salt)
			b64Got := base64.RawStdEncoding.EncodeToString(got)
			if !reflect.DeepEqual(b64Got, tt.want) {
				t.Errorf("createHashPassword() = %v, want %v", b64Got, tt.want)
			}
		})
	}
}

func Test_createEncodedPassword(t *testing.T) {
	t.Parallel()
	hashedPassword, err := base64.RawStdEncoding.DecodeString("WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")
	if err != nil {
		fmt.Println("could not decode b64 password")
	}
	type args struct {
		hashedPassword []byte
		salt           []byte
	}
	tests := []struct {
		name string
		args args
		want password.EncodedPassword
	}{
		{"success", args{hashedPassword, []byte("12345678901234567890123456789012")}, password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := createEncodedPassword(tt.args.hashedPassword, tt.args.salt); got != tt.want {
				t.Errorf("createEncodedPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeHash(t *testing.T) {
	t.Parallel()
	wantHashedPassword, err := base64.RawStdEncoding.DecodeString("WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")
	if err != nil {
		t.Error(err)
	}

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
		{"success", args{password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")}, wantHashedPassword, []byte("12345678901234567890123456789012"), false},
		{"invalid format", args{password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY$")}, wantHashedPassword, []byte("12345678901234567890123456789012"), true},
		{"invalid b64 salt", args{password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI==$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")}, wantHashedPassword, []byte("12345678901234567890123456789012"), true},
		{"invalid b64 password", args{password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY==")}, wantHashedPassword, []byte("12345678901234567890123456789012"), true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotHashedPassword, gotSalt, err := decodeHash(tt.args.encodedPassword)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("decodeHash() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(gotHashedPassword, tt.wantHashedPassword) {
				t.Errorf("decodeHash() gotHashedPassword = %v, want %v", gotHashedPassword, tt.wantHashedPassword)

				return
			}
			if !reflect.DeepEqual(gotSalt, tt.wantSalt) {
				t.Errorf("decodeHash() gotSalt = %v, want %v", gotSalt, tt.wantSalt)

				return
			}
		})
	}
}

func TestEncoder_IsMatchPassword(t *testing.T) {
	t.Parallel()
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
		{"success", args{inputPassword: "password", storedPassword: password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")}, true, false},
		{"fail", args{inputPassword: "failed", storedPassword: password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY")}, false, false},
		{"error", args{inputPassword: "password", storedPassword: password.EncodedPassword("$argon2id$v=19$m=65536,t=1,p=4$MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI$WBo5t5PvTcN/kEJbQLWhYcF4d+n+r6hSdLX+6aJymIY$")}, false, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Argon2Encoder{}
			got, err := e.IsMatchPassword(tt.args.inputPassword, tt.args.storedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsMatchPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("IsMatchPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_EncodePassword(t *testing.T) {
	t.Parallel()
	t.Run("check hash differences", func(t *testing.T) {
		t.Parallel()
		enc := &Argon2Encoder{}
		hash1, err := enc.EncodePassword("password")
		if err != nil {
			t.Errorf("EncodePassword() error = %v", err)

			return
		}
		dhash1, salt1, err := decodeHash(hash1)
		if err != nil {
			t.Errorf("decodeHash() error = %v", err)

			return
		}
		hash2, err := enc.EncodePassword("password")
		if err != nil {
			t.Errorf("EncodePassword() error = %v", err)

			return
		}
		dhash2, salt2, err := decodeHash(hash2)
		if err != nil {
			t.Errorf("decodeHash() error = %v", err)

			return
		}
		if hash1 == hash2 {
			t.Errorf("Generate same salt")

			return
		}
		if string(salt1) == string(salt2) {
			t.Errorf("Generate same salt")

			return
		}
		if string(dhash1) == string(dhash2) {
			t.Errorf("Generate same salt")

			return
		}
	})
}

func TestEncoder_Argon2Encoder(t *testing.T) {
	t.Parallel()
	type args struct {
		inputPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"sequence of events", args{"password"}, true, false},
		{
			"sequence of events",
			args{"^@}_RKyKFr4!Y^UCMa6FA,U+kjnURedZkuYhvu-s?rxAc7u6iVLR3]hJKG]yf!Hf2]Qd.D?Qyhy2s,r#!r,3dyij>h]=Hr,NuR*e"},
			true, false,
		},
		{
			"sequence of events",
			//nolint:lll
			args{"7IVQq4L0sBy2zjSgqeLy6wnDci55mYOuMr0uuv1NFuf0Zgv3Y8JP2sl6riqm8U6N30980BUqp3ISdJn9o47Vdfp24xjj9sWP6nnkaPB4Pa75Bq3QEZLmo4IFPJRVLdFB1lek1RbG8FkELNbkmfB5kwUnC1aw2w3x3ty5J165CWP7CdtasoFbpz5RpcNgQlipvhZDQxyhzVpmfkWxgJqOqWnRgOgqQFexTK2qSNHEXidxyrKlSUoGWP5IHvSbl9Y6"},
			true, false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			enc := NewArgon2Encoder()
			hashedPassword, err := enc.EncodePassword(tt.args.inputPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodePassword() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if ok, err := enc.IsMatchPassword(tt.args.inputPassword, hashedPassword); ok != tt.want || err != nil {
				t.Errorf("IsMatchPassword() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
