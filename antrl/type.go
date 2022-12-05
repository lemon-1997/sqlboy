package parser

type GoType string

const (
	Invalid = "invalid"

	Bool   = "bool"
	Int8   = "int8"
	Int16  = "int16"
	Int32  = "int32"
	Int64  = "int64"
	Uint8  = "uint8"
	Uint16 = "uint16"
	Uint32 = "uint32"
	Uint64 = "uint64"

	Float32 = "float32"
	Float64 = "float64"

	String     = "string"
	Time       = "time.Time"
	SliceByte  = "[]byte"
	SliceUint8 = "[]uint8"
)
