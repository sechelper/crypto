package symmetric

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDes3ECBEncrypt(t *testing.T) {
	key := []byte("012345678901234567890123") // 24-byte key for 3DES

	plaintext := []byte("Hello, Secself!") // Plaintext to encrypt
	encrypt, err := Des3ECBEncrypt(key, plaintext)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))

}
func TestDes3ECBDecrypt(t *testing.T) {
	key := []byte("012345678901234567890123") // 24-byte key for 3DES
	encrypt, _ := base64.StdEncoding.DecodeString("6Dxl7vNNWD4G2sGt6JN1RA==")
	decrypt, err := Des3ECBDecrypt(key, encrypt)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(decrypt))
}
