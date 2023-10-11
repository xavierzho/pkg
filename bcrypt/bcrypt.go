package bcrypt

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd string) ([]byte, error) {

	hashed, err := bcrypt.GenerateFromPassword(bytes.NewBufferString(pwd).Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func Verify(pwd string, hashed []byte) bool {
	return bcrypt.CompareHashAndPassword(hashed, bytes.NewBufferString(pwd).Bytes()) == nil
}
