package encryption

import (
	"crypto/rsa"
	"testing"
)

func TestAes(t *testing.T) {
	key1 := []byte("AES256Key-32Characters1234567890")    // 32 bytes for AES-256
	key2 := []byte("AES256Key-32Characters1234567891")    // 32 bytes for AES-256
	keyWrong := []byte("AES256Key-32Characters123456789") // 31 bytes for AES-256

	testAes := []struct {
		message          string
		key1             []byte
		key2             []byte
		wantError        string
		inputEqualOutbut bool
	}{
		{"Hello, RSA!", key1, key1, "", true},
		{"ПРИВЕТ ВСЕМ!", key1, key1, "", true},
		{"привет всем!", key2, key2, "", true},
		{"Hello, RSA!", key1, key2, "", false},
		{"Hello, RSA!", key2, key1, "", false},
		{"Hello, RSA!", key1, key2, "", false},
		{"Hello, RSA!", key2, key1, "", false},
		{"Hello, RSA!", keyWrong, key1, "crypto/aes: invalid key size 31", false},
		{"Hello, RSA!", key1, []byte{}, "crypto/aes: invalid key size 0", false},
	}
	for i, r := range testAes {
		encrypted, err := aesEncrypt(r.key1, []byte(r.message))
		if err != nil {
			if err.Error() != r.wantError {
				t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.wantError, err)
			} else {
				t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.wantError, err)
			}
			continue
		}
		decrypted, err := aesDecrypt(r.key2, encrypted)
		if err != nil {
			if err.Error() != r.wantError {
				t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.wantError, err)
			} else {
				t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.wantError, err)
			}
			continue
		}
		if string(decrypted) != r.message {
			if r.inputEqualOutbut {
				t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.message, decrypted)
			} else {
				t.Logf("Test %d: OK, expected not %s, result %s: OK", i+1, r.message, decrypted)
			}
		} else {
			if r.inputEqualOutbut {
				t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.message, decrypted)
			} else {
				t.Errorf("Test %d: FAIL, expected not %s, result %s", i+1, r.message, decrypted)

			}
		}
	}
}

func TestRsa(t *testing.T) {
	privateKey, _ := generateRSAKeys()
	privateKey2, _ := generateRSAKeys()
	publicKey := &privateKey.PublicKey
	publicKey2 := &privateKey2.PublicKey

	testRsa := []struct {
		message   string
		key1      *rsa.PublicKey
		key2      *rsa.PrivateKey
		wantError string
	}{
		{"Hello, RSA!", publicKey, privateKey, ""},
		{"ПРИВЕТ ВСЕМ!", publicKey, privateKey, ""},
		{"привет всем!", publicKey2, privateKey2, ""},
		{"Hello, RSA!", publicKey, privateKey2, "crypto/rsa: decryption error"},
		{"Hello, RSA!", publicKey2, privateKey, "crypto/rsa: decryption error"},
		{"Hello, RSA!", &rsa.PublicKey{}, privateKey, "crypto/rsa: missing public modulus"},
		{"Hello, RSA!", publicKey, &rsa.PrivateKey{}, "crypto/rsa: missing public modulus"},
	}
	for i, r := range testRsa {
		encrypted, err := rsaEncrypt(r.key1, []byte(r.message))
		if err != nil {
			if err.Error() != r.wantError {
				t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.wantError, err)
			} else {
				t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.wantError, err)
			}
			continue
		}
		decrypted, err := rsaDecrypt(r.key2, encrypted)
		if err != nil {
			if err.Error() != r.wantError {
				t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.wantError, err)
			} else {
				t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.wantError, err)
			}
			continue
		}
		if string(decrypted) != r.message {
			t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.message, decrypted)
		} else {
			t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.message, decrypted)
		}
	}
}
