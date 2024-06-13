package crypto

import (
	"encoding/hex"
	"testing"
)

func TestAES(t *testing.T) {
	key, err := GenerateAesKey()
	println(hex.EncodeToString(key))

	if err != nil {
		t.Fatal(err)
	}

	text := "this is a test,包含中文"

	cipher, err := EncryptAES(key, []byte(text))
	if err != nil {
		t.Fatal(err)
	}

	text2, err := DecryptAES(key, cipher)
	if err != nil {
		t.Fatal(err)
	}

	if text != string(text2) {
		t.Fatal("decrypted text is not the same")
	}
}
