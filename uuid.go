package faker

import (
	"crypto/rand"
	"fmt"
	"io"
)

// UUID is struct for UUID
type UUID struct {
	Faker *Faker
}

// V4 returns UUID V4 as string
func (u UUID) V4() string {
	var uuid [16]byte
	io.ReadFull(rand.Reader, uuid[:])
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
