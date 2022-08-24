package bd

import "golang.org/x/crypto/bcrypt"

func encriptarPassword(pass string) (string, error) {
	//es el numero de 2 elevado al costo, por ende a mayor costo mas seguridad pero mas lento
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
