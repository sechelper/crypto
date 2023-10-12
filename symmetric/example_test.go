package symmetric_test

import (
	"encoding/base64"
	"fmt"
	"github.com/sechelper/crypto/symmetric"
)

func ExampleDes3ECBEncrypt() {
	key := []byte("012345678901234567890123") // 24-byte key for 3DES

	plaintext := []byte("Hello, Secself!") // Plaintext to encrypt
	encrypt, err := symmetric.Des3ECBEncrypt(key, plaintext)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))

}
func ExampleDes3ECBDecrypt() {
	key := []byte("012345678901234567890123") // 24-byte key for 3DES
	encrypt, _ := base64.StdEncoding.DecodeString("6Dxl7vNNWD4G2sGt6JN1RA==")
	decrypt, err := symmetric.Des3ECBDecrypt(key, encrypt)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decrypt))
}
