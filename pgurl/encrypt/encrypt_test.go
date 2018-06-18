package encrypt

import (
	"testing"
)

func TestAES(t *testing.T) {
	// write some test
	key := []byte("LKHlhb899Y09olUi")
	encryptMsg, _ := Encrypt(key, "PGURL is GOOD !!!!!")
	msg, _ := Decrypt(key, encryptMsg)
	t.Logf("%s -> %s", msg, encryptMsg) // Hello World
}
