package secret_test

import (
	"testing"

	"github.com/Nigel2392/secret"
)

var key = secret.New("0123456789")
var data = []byte("Hello World")

func TestEncryptDecryptBytes(t *testing.T) {
	encrypted, err := key.Bytes().Encrypt(data)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := key.Bytes().Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if string(decrypted) != string(data) {
		t.Fatalf("Decrypted data does not match original data.")
	}
}

func TestEncryptDecryptHTMLSafe(t *testing.T) {
	encrypted, err := key.HTMLSafe().Encrypt(string(data))
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := key.HTMLSafe().Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != string(data) {
		t.Fatalf("Decrypted data does not match original data.")
	}
}

func TestSignVerify(t *testing.T) {
	signature := key.Sign(string(data))
	if !key.Verify(string(data), signature) {
		t.Fatal("Signature does not match data.")
	} else if key.Verify(string(data)+"!", signature) {
		t.Fatal("Bad signature matched data.")
	}
}
