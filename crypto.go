package shell

import (
	"errors"
	"strings"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
	"github.com/ontio/ontology-crypto/keypair"
)

type SymmetricScheme byte

const (
	AES SymmetricScheme = iota
)

var names []string = []string {
	"AES",
}

func GetScheme(name string) (SymmetricScheme, error) {
	for i, v := range names {
		if strings.ToUpper(v) == strings.ToUpper(name) {
			return SymmetricScheme(i), nil
		}
	}
	return 0, errors.New("unknown symmetric scheme " + name)
}

func kdf(pwd []byte, salt []byte) (dKey []byte, err error) {
	param := keypair.GetScryptParameters()
	if param.DKLen < 32 {
		err = errors.New("derived key length too short")
		return nil, err
	}

	// Derive the encryption key
	dKey, err = scrypt.Key([]byte(pwd), salt, param.N, param.R, param.P, param.DKLen)
	return dKey, err
}

func randomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
