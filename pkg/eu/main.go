// Error Utils

package eu

import "log"

// if (E)rror (L)og
func El(err error) {
	if err != nil {
		log.Println(err)
	}
}

// if (E)rror (L)og (F)atal
func Elf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// if (E)rror (L)og (F)atal (M)sg
func Elfm(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// if (E)rror (L)og (F)atal (D)o
func Elfd(err error, do func()) {
	if err != nil {
		do()
	}
}
