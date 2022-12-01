package parser

type GoType int

const (
	Invalid GoType = iota

	Bool
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64

	Float32
	Float64

	String
	Time
	SliceByte
	SliceUint8
)
