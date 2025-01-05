package structs

import (
	"fmt"
	"log"
)

// MySql bit(1) type implemented as boolean field for DAOs
// the mysql driver returns the field value as []uint8 for some reason
type BitBool bool

func (b *BitBool) Scan(src interface{}) error {
	log.Printf("BitBool.Scan() type of src %T\n", src)
	srcByte, ok := src.([]uint8)
	if !ok {
		return fmt.Errorf("Unexpected type for BitBool: %T", src)
	}
	if len(srcByte) == 0 || srcByte[0] == 0 {
		*b = false
	} else if srcByte[0] == 1 {
		*b = true
	}
	return nil
}
