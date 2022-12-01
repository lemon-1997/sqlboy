package parser

type ColumnDecl struct {
	RawText    string
	ColumnName string
	GoType     GoType
}

type StmtListener struct {
	*BaseStmtParserListener
	walkDecl   string
	walkColumn string

	TableName string
	Columns   []ColumnDecl
}

func NewStmtListener() *StmtListener {
	return new(StmtListener)
}

// EnterTableName db_name.tbl_name
func (l *StmtListener) EnterTableName(ctx *TableNameContext) {
	var tableName string
	switch ctx.GetStop().GetTokenType() {
	case StmtParserREVERSE_QUOTE_ID, StmtParserCHARSET_REVERSE_QOUTE_STRING, StmtParserSTRING_LITERAL:
		name := ctx.GetStop().GetText()
		if len(name) <= 2 {
			return
		}
		tableName = name[1 : len(name)-1]
	case StmtParserDOT_ID:
		name := ctx.GetStop().GetText()
		if len(name) <= 1 {
			return
		}
		tableName = name[1:]
	default:
		tableName = ctx.GetText()
	}
	l.TableName = tableName
}

// EnterColumnDeclaration remove quote
func (l *StmtListener) EnterColumnDeclaration(ctx *ColumnDeclarationContext) {
	switch ctx.Uid().GetStop().GetTokenType() {
	case StmtParserREVERSE_QUOTE_ID, StmtParserCHARSET_REVERSE_QOUTE_STRING, StmtParserSTRING_LITERAL:
		column := ctx.Uid().GetText()
		if len(column) <= 2 {
			return
		}
		l.walkColumn = column[1 : len(column)-1]
	default:
		l.walkColumn = ctx.Uid().GetText()
	}
	l.walkDecl = ctx.GetText()
}

// EnterIntegerDataType
// INT1 => TINYINT INT2 => SMALLINT INT3 => MEDIUMINT INT4 => INT INT8 => BIGINT
func (l *StmtListener) EnterIntegerDataType(ctx *IntegerDataTypeContext) {
	unsigned := ctx.UNSIGNED(0)
	var goType GoType
	switch ctx.GetStart().GetTokenType() {
	case StmtParserTINYINT, StmtParserINT1:
		goType = Int8
		if unsigned != nil {
			goType = Uint8
		}
	case StmtParserSMALLINT, StmtParserINT2:
		goType = Int16
		if unsigned != nil {
			goType = Uint16
		}
	case StmtParserMEDIUMINT, StmtParserMIDDLEINT, StmtParserINT, StmtParserINTEGER, StmtParserINT3, StmtParserINT4:
		goType = Int32
		if unsigned != nil {
			goType = Uint32
		}
	case StmtParserBIGINT, StmtParserINT8:
		goType = Int64
		if unsigned != nil {
			goType = Uint64
		}
	default:
		goType = Invalid
	}
	l.Columns = append(l.Columns, l.TransColumnDecl(goType))
}

// EnterFloatDataType
// FLOAT | FLOAT4
// DOUBLE | FLOAT8 | DOUBLE PRECISION | REAL
// DECIMAL | DEC | FIXED | NUMERIC
func (l *StmtListener) EnterFloatDataType(ctx *FloatDataTypeContext) {
	switch ctx.GetStart().GetTokenType() {
	case StmtParserFLOAT, StmtParserFLOAT4:
		l.Columns = append(l.Columns, l.TransColumnDecl(Float32))
	default:
		l.Columns = append(l.Columns, l.TransColumnDecl(Float64))
	}
}

func (l *StmtListener) EnterTimeDataType(ctx *TimeDataTypeContext) {
	l.Columns = append(l.Columns, l.TransColumnDecl(Time))
}

func (l *StmtListener) EnterStringDataType(ctx *StringDataTypeContext) {
	l.Columns = append(l.Columns, l.TransColumnDecl(String))
}

func (l *StmtListener) EnterBinaryDataType(ctx *BinaryDataTypeContext) {
	l.Columns = append(l.Columns, l.TransColumnDecl(SliceByte))
}

func (l *StmtListener) EnterSpecialDataType(ctx *SpecialDataTypeContext) {
	switch ctx.GetStart().GetTokenType() {
	case StmtParserBOOL, StmtParserBOOLEAN:
		l.Columns = append(l.Columns, l.TransColumnDecl(Bool))
	case StmtParserSERIAL:
		// SERIAL is an alias for BIGINT UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE
		l.Columns = append(l.Columns, l.TransColumnDecl(Uint64))
	case StmtParserBIT:
		l.Columns = append(l.Columns, l.TransColumnDecl(SliceUint8))
	case StmtParserYEAR:
		// MySQL displays YEAR values in YYYY format, with a range of 1901 to 2155, and 0000.
		l.Columns = append(l.Columns, l.TransColumnDecl(Uint16))
	default:
		l.Columns = append(l.Columns, l.TransColumnDecl(Invalid))
	}
}

func (l *StmtListener) EnterSpatialDataType(ctx *SpatialDataTypeContext) {
	l.Columns = append(l.Columns, l.TransColumnDecl(Invalid))
}

func (l *StmtListener) TransColumnDecl(t GoType) ColumnDecl {
	return ColumnDecl{
		RawText:    l.walkDecl,
		ColumnName: l.walkColumn,
		GoType:     t,
	}
}
