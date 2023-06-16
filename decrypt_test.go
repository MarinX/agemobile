package agemobile

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"

	"filippo.io/age"
)

func TestDecryptPass(t *testing.T) {
	b64key := "YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IHNjcnlwdCBwT29hWWJxSUF6a0FxMnpiSXVJVGxBIDE4CkkrdXpIZk1uQlhqSWxscCtiSDdGOHlZdkNUTkNzUEMyOEx4aXVET2w0ZzQKLS0tIDlYREdmZUJMT2w4Y2Y3M3dsUmJDZlgweTZ6TjExMitUVVdSV1VnM3U5azgKvUblkUHUQXjEyG1kAKL5XmJHKqUuIeClo9b18JN7y4AMo+6qUjZ1IerK6RkRvRoMzG75NKfsJ/YiSFcPC61JW95o73+IXDAvVVLTp44ihm6ESkJeEEOfFVKDstA8jMTSxmxmP32pxkpLav8bAP65mQL/L5JzJ4vJkXUw5Pj/4y/bhtKRiBo86cl8ZAZDJAnHiwCg7JkUm/dxrCCW0Gpqj/lyXaaq+4LjqSo/eAinV5a7vTAWv4jQqE3240hPqVlP7GWqqouS9TMpLmjixhN6sh7kNdYdiUZ/2CPmnJE="
	key, _ := b64.StdEncoding.DecodeString(b64key)
	enc, err := DecryptPass(
		"adjust-sample-fault-museum-grid-uncle-comic-gold-impact-myth",
		string(key),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b64.StdEncoding.EncodeToString([]byte(enc)))
}

func TestDecrypt(t *testing.T) {
	txt := "Hello World"
	id, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	out, err := Encrypt(id.Recipient().String(), txt, false)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := Decrypt(id.String(), out)
	if err != nil {
		t.Fatal(err)
	}
	if dec != txt {
		t.Fatalf("Expected %v but got %v\n", txt, dec)
	}
}

func TestDecryptArmor(t *testing.T) {
	txt := "Hello World"
	id, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	out, err := Encrypt(id.Recipient().String(), txt, true)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := Decrypt(id.String(), out)
	if err != nil {
		t.Fatal(err)
	}
	if dec != txt {
		t.Fatalf("Expected %v but got %v\n", txt, dec)
	}
}

func TestDecryptFile(t *testing.T) {
	txt := "Hello World"
	id, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	out, err := Encrypt(id.Recipient().String(), txt, false)
	if err != nil {
		t.Fatal(err)
	}
	fdin, err := os.CreateTemp("", "age_input_*")
	if err != nil {
		t.Fatal(err)
	}
	defer fdin.Close()

	io.WriteString(fdin, out)

	fdout, err := os.CreateTemp("", "age_*")
	if err != nil {
		t.Fatal(err)
	}
	defer fdout.Close()

	err = DecryptFile(id.String(), fdin.Name(), fdout.Name())
	if err != nil {
		t.Fatal(err)
	}

	buff, err := io.ReadAll(fdout)
	if err != nil {
		t.Fatal(err)
	}
	if string(buff) != txt {
		t.Fatalf("expected decrypted %v but got %v\n", txt, string(buff))
	}
}

func TestDecryptFileArmor(t *testing.T) {
	txt := "Hello World"
	id, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	out, err := Encrypt(id.Recipient().String(), txt, true)
	if err != nil {
		t.Fatal(err)
	}
	fdin, err := os.CreateTemp("", "age_input_*")
	if err != nil {
		t.Fatal(err)
	}
	defer fdin.Close()

	io.WriteString(fdin, out)

	fdout, err := os.CreateTemp("", "age_*")
	if err != nil {
		t.Fatal(err)
	}
	defer fdout.Close()

	err = DecryptFile(id.String(), fdin.Name(), fdout.Name())
	if err != nil {
		t.Fatal(err)
	}

	buff, err := io.ReadAll(fdout)
	if err != nil {
		t.Fatal(err)
	}
	if string(buff) != txt {
		t.Fatalf("expected decrypted %v but got %v\n", txt, string(buff))
	}
}
