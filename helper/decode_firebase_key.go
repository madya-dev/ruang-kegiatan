package helper

import (
	"encoding/base64"
	"madyasantosa/ruangkegiatan/config"
)

func GetDecodedFireBaseKey(c config.Config) ([]byte, error) {

	fireBaseAuthKey := c.FirebaseAuthKey

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}
