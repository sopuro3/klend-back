package argon2

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/sopuro3/klend-back/pkg/password"
	"golang.org/x/crypto/argon2"
)

type Argon2Encoder struct{}

var (
	algorithm        = "argon2id"
	version          = argon2.Version
	time      uint32 = 1
	memory    uint32 = 64 //64MB
	threads   uint8  = 4
	keyLen    uint32 = 32
)

func NewArgon2Encoder() password.Encoder {
	return &Argon2Encoder{}
}

func createHashPassword(rawPassword string, salt []byte) []byte {
	return argon2.IDKey([]byte(rawPassword), salt, time, memory*1024, threads, keyLen)
}

func createEncodedPassword(hashedPassword, salt []byte) password.EncodedPassword {

	base64Salt := base64.StdEncoding.EncodeToString(salt)
	base64HashedPassword := base64.StdEncoding.EncodeToString(hashedPassword)
	// $<algorithm name>$v=<version>$m=<memory size>$t=<time>$p=<threads>$l=<len>$<b64 salt>$<b64 hash value>
	return password.EncodedPassword(fmt.Sprintf("$%s$v=%d$m=%d$t=%d$p=%d$l=%d$%s$%s$", algorithm, version, memory, time, threads, keyLen, base64Salt, base64HashedPassword))
}
func (e *Argon2Encoder) EncodePassword(rawPassword string) (password.EncodedPassword, error) {
	saltLen := 32
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// to hash
	hashedPassword := createHashPassword(rawPassword, salt)

	return createEncodedPassword(hashedPassword, salt), nil
}

func (e *Argon2Encoder) IsMatchPassword(rawPassword string, encodedPassword password.EncodedPassword) bool {
	//TODO implement me
	panic("implement me")
}
