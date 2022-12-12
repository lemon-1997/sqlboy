package generator

type Mode int8

const (
	ModeGorm = iota + 1
	ModeSqlx
)

//type genOption struct {
//	mode Mode
//}
