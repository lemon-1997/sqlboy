package generator

import parser "sqlboy/antrl"

func ToSqlNullType(goType string) string {
	switch goType {
	case parser.Invalid:
		return ""
	case parser.Bool:
		return "sql.NullBool"
	case parser.Int8, parser.Int16, parser.Uint8:
		return "sql.NullInt16"
	case parser.Int32, parser.Uint16:
		return "sql.NullInt32"
	//uint64 out of range https://github.com/golang/go/issues/47953
	case parser.Int64, parser.Uint32, parser.Uint64:
		return "sql.NullInt64"
	case parser.Float32, parser.Float64:
		return "sql.NullFloat64"
	case parser.String:
		return "sql.NullString"
	case parser.Time:
		return "sql.NullTime"
	case parser.SliceByte, parser.SliceUint8:
		return "sql.NullByte"
	default:
		return ""
	}
}
