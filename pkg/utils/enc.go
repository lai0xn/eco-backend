package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func CheckPassword(enc_pass string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(enc_pass), []byte(password))

	if err != nil {
		return err
	}
	return nil
}
