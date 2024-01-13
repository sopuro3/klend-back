package argon2

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/sopuro3/klend-back/pkg/password"
)

//nolint:revive
type Argon2Encoder struct{}

const (
	algorithm        = "argon2id"
	version          = argon2.Version
	time      uint32 = 1
	memory    uint32 = 64 * 1024 // 64MB
	threads   uint8  = 4
	keyLen    uint32 = 32
	saltLen   uint32 = 32
)

func NewArgon2Encoder() *Argon2Encoder {
	return &Argon2Encoder{}
}

func createHashPassword(rawPassword string, salt []byte) []byte {
	return argon2.IDKey([]byte(rawPassword), salt, time, memory, threads, keyLen)
}

func createEncodedPassword(hashedPassword, salt []byte) password.EncodedPassword {
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64HashedPassword := base64.RawStdEncoding.EncodeToString(hashedPassword)
	// $<algorithm name>$v=<version>$m=<memory size>,t=<time>,p=<threads>$<b64 salt>$<b64 hash value>
	//nolint:lll
	hash := fmt.Sprintf("$%s$v=%d$m=%d,t=%d,p=%d$%s$%s", algorithm, version, memory, time, threads, base64Salt, base64HashedPassword)

	return password.EncodedPassword(hash)
}

func (e *Argon2Encoder) EncodePassword(rawPassword string) (password.EncodedPassword, error) {
	salt := make([]byte, saltLen)

	if _, err := rand.Read(salt); err != nil {
		return password.EncodedPassword(""), fmt.Errorf("could not generate salt, %w", err)
	}
	// to hash
	hashedPassword := createHashPassword(rawPassword, salt)

	return createEncodedPassword(hashedPassword, salt), nil
}

func (e *Argon2Encoder) IsMatchPassword(inputPassword string, storedPassword password.EncodedPassword) (bool, error) {
	hashedStoredPassword, storedSalt, err := decodeHash(storedPassword)
	if err != nil {
		return false, err
	}

	inputHashedPassword := createHashPassword(inputPassword, storedSalt)

	if string(inputHashedPassword) == string(hashedStoredPassword) {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedPassword password.EncodedPassword) ([]byte, []byte, error) {
	split := strings.Split(string(encodedPassword), "$")
	//nolint:gomnd
	if len(split) != 6 {
		//nolint:goerr113
		err := fmt.Errorf("password is invalid format")

		return []byte(""), []byte(""), err
	}

	salt, err := base64.RawStdEncoding.DecodeString(split[4])
	if err != nil {
		return []byte(""), []byte(""), err
	}

	hashedPassword, err := base64.RawStdEncoding.DecodeString(split[5])
	if err != nil {
		return []byte(""), []byte(""), err
	}

	return hashedPassword, salt, nil
}
