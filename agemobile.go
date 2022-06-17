package agemobile

import (
	"filippo.io/age"
	_ "golang.org/x/mobile/bind"
)

// GenerateX25519Identity randomly generates a new X25519Identity.
func GenerateX25519Identity() (*age.X25519Identity, error) {
	return age.GenerateX25519Identity()
}

// ParseX25519Identity returns a new X25519Identity from a Bech32 private key
// encoding with the "AGE-SECRET-KEY-1" prefix.
func ParseX25519Identity(s string) (*age.X25519Identity, error) {
	return age.ParseX25519Identity(s)
}

// ParseX25519Recipient returns a new X25519Recipient from a Bech32 public key
// encoding with the "age1" prefix.
func ParseX25519Recipient(s string) (*age.X25519Recipient, error) {
	return age.ParseX25519Recipient(s)
}
