package parser

type ColumnDecl struct {
	Decl    string
	Name    string
	Comment string
	SqlType string
	GoType  GoType
}

func (c *ColumnDecl) SetDecl(decl string) {
	c.Decl = decl
}

func (c *ColumnDecl) SetName(name string) {
	c.Name = name
}

func (c *ColumnDecl) SetComment(comment string) {
	c.Comment = comment
}

func (c *ColumnDecl) SetGoType(gType GoType) {
	c.GoType = gType
}

func (c *ColumnDecl) SetSqlType(sType string) {
	c.SqlType = sType
}

type StmtListener struct {
	*BaseStmtParserListener
	column ColumnDecl

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
	l.column = ColumnDecl{}
	switch ctx.Uid().GetStop().GetTokenType() {
	case StmtParserREVERSE_QUOTE_ID, StmtParserCHARSET_REVERSE_QOUTE_STRING, StmtParserSTRING_LITERAL:
		column := ctx.Uid().GetText()
		if len(column) <= 2 {
			return
		}
		l.column.SetName(column[1 : len(column)-1])
	default:
		l.column.SetName(ctx.Uid().GetText())
	}
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()
	l.column.SetDecl(ctx.GetStart().GetInputStream().GetText(start, stop))
}

func (l *StmtListener) ExitColumnDeclaration(ctx *ColumnDeclarationContext) {
	l.Columns = append(l.Columns, l.column)
}

func (l *StmtListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {
	start := ctx.DataType().GetStart().GetStart()
	stop := ctx.DataType().GetStop().GetStop()
	l.column.SetSqlType(ctx.GetStart().GetInputStream().GetText(start, stop))
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
	l.column.SetGoType(goType)
}

// EnterFloatDataType
// FLOAT | FLOAT4
// DOUBLE | FLOAT8 | DOUBLE PRECISION | REAL
// DECIMAL | DEC | FIXED | NUMERIC
func (l *StmtListener) EnterFloatDataType(ctx *FloatDataTypeContext) {
	switch ctx.GetStart().GetTokenType() {
	case StmtParserFLOAT, StmtParserFLOAT4:
		l.column.SetGoType(Float32)
	default:
		l.column.SetGoType(Float64)
	}
}

func (l *StmtListener) EnterTimeDataType(ctx *TimeDataTypeContext) {
	l.column.SetGoType(Time)
}

func (l *StmtListener) EnterStringDataType(ctx *StringDataTypeContext) {
	l.column.SetGoType(String)
}

func (l *StmtListener) EnterBinaryDataType(ctx *BinaryDataTypeContext) {
	l.column.SetGoType(SliceByte)
}

func (l *StmtListener) EnterSpecialDataType(ctx *SpecialDataTypeContext) {
	switch ctx.GetStart().GetTokenType() {
	case StmtParserBOOL, StmtParserBOOLEAN:
		l.column.SetGoType(Bool)
	case StmtParserSERIAL:
		// SERIAL is an alias for BIGINT UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE
		l.column.SetGoType(Uint64)
	case StmtParserBIT:
		l.column.SetGoType(SliceUint8)
	case StmtParserYEAR:
		// MySQL displays YEAR values in YYYY format, with a range of 1901 to 2155, and 0000.
		l.column.SetGoType(Uint16)
	default:
		l.column.SetGoType(Invalid)
	}
}

func (l *StmtListener) EnterSpatialDataType(ctx *SpatialDataTypeContext) {
	l.column.SetGoType(Invalid)
}

func (l *StmtListener) EnterCommentColumnConstraint(ctx *CommentColumnConstraintContext) {
	comment := ctx.STRING_LITERAL().GetText()
	if len(comment) <= 2 {
		return
	}
	l.column.SetComment(comment[1 : len(comment)-1])
}
