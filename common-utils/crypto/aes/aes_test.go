package aes

import (
	"bytes"
	"testing"
)

func Test_AES_Encrypt_Decrypt(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{
			name: "Test 1",
			data: []byte("Hello, World!"),
		},
		{
			name: "Test 2",
			data: []byte("Go is Awesome. It's built-in concurrency capabilities are amazing"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate a new key
			key, err := GenerateSymmetricKey()
			if err != nil {
				t.Fatalf("Failed to generate key: %v", err)
			}

			// Encrypt the data
			encrypted := EncryptData(key, tt.data)
			decryptedData := DecryptData(key, encrypted)

			if !bytes.Equal(decryptedData, tt.data) {
				t.Errorf("decrypted data does not match original.\nExpected: %s\nGot: %s", string(tt.data), string(decryptedData))
			}
		})
	}
}
