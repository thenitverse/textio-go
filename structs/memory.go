package main

import (
	"fmt"
	"unsafe"
)

type contact struct {
	userID       string //16 bytes
	sendingLimit int32  // 4bytes
	age          int32
}

type perms struct {
	permissionLevel int  // 8bytes
	canSend         bool // 1 byte
	canReceive      bool
	canManage       bool
}

func main() {
	c := contact{}
	p := perms{}

	fmt.Println("--- Contact Struct ---")
	fmt.Printf("Total size: %d bytes\n", unsafe.Sizeof(c))
	// On a 64-bit system:
	// string (16) + int32 (4) + int32 (4) = 24 bytes

	fmt.Println("\n--- Perms Struct ---")
	fmt.Printf("Total size: %d bytes\n", unsafe.Sizeof(p))
	// On a 64-bit system:
	// int (8) + bool (1) + bool (1) + bool (1) = 11 bytes
	// However, Go pads this to 16 bytes to align with the 8-byte 'int' boundary.
}

/*output:
--- Contact Struct ---
Total size: 24 bytes

--- Perms Struct ---
Total size: 16 bytes
*/
