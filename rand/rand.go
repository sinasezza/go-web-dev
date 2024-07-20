package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)

	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error reading random bytes: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("read %d bytes, want %d", nRead, n)
	}
	return b, nil
}

// n is number of bytes
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
