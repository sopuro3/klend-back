package password

type EncodedPassword string

/*
Encoder
Hash format
$<algorithm name>$v=<version>$m=<memory size>$t=<time>$p=<threads>$l=<len>$<b64 salt>$<b64 hash value>
*/
type Encoder interface {
	EncodePassword(rawPassword string) (EncodedPassword, error)
	IsMatchPassword(rawPassword string, encodedPassword EncodedPassword) bool
}
