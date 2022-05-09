package gorngseed

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	mathRand "math/rand"
)

// Register is a wrapper for Seed. It exits the application on error.
func Register() {
	err := Seed()
	if err != nil {
		log.Fatalln(err)
	}
}

// Seed seeds the math/rand RNG according to: https://stackoverflow.com/a/54491783.
// It should be initialized as soon as possible in the application.
func Seed() error {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		return fmt.Errorf("cannot seed math/rand: %w", err)
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return nil
}
