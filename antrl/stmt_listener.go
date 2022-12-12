package parser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ColumnDecl struct {
	Decl      string
	Name      string
	Comment   string
	SqlType   string
	GoType    GoType
	IsNotNull bool
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

func (c *ColumnDecl) SetIsNotNull(isNotNull bool) {
	c.IsNotNull = isNotNull
}

type ColumnIndex struct {
	Decl    string
	Columns []ColumnDecl
}

type TableAttr struct {
	TableName  string
	Columns    []ColumnDecl
	PrimaryKey ColumnIndex
	UniqueKeys []ColumnIndex
}

type StmtListener struct {
	*BaseStmtParserListener
	column ColumnDecl
	TableAttr
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

func (l *StmtListener) EnterColumnDeclaration(ctx *ColumnDeclarationContext) {
	columnName := l.removeColumnQuote(ctx.Uid().GetStart())
	decl := l.getTextFromTokens(ctx.GetStart(), ctx.GetStop())
	l.column = ColumnDecl{}
	l.column.SetName(columnName)
	l.column.SetDecl(decl)
}

func (l *StmtListener) ExitColumnDeclaration(ctx *ColumnDeclarationContext) {
	l.Columns = append(l.Columns, l.column)
}

func (l *StmtListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {
	start := ctx.DataType().GetStart().GetStart()
	stop := ctx.DataType().GetStop().GetStop()
	l.column.SetSqlType(ctx.GetStart().GetInputStream().GetText(start, stop))
}

func (l *StmtListener) EnterNullColumnConstraint(ctx *NullColumnConstraintContext) {
	if ctx.GetStart().GetTokenType() == StmtParserNOT {
		l.column.SetIsNotNull(true)
	}
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

func (l *StmtListener) EnterPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) {
	l.PrimaryKey = ColumnIndex{
		Decl:    l.column.Decl,
		Columns: []ColumnDecl{l.column},
	}
}

func (l *StmtListener) EnterUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) {
	l.UniqueKeys = append(l.UniqueKeys, ColumnIndex{
		Decl:    l.column.Decl,
		Columns: []ColumnDecl{l.column},
	})
}

func (l *StmtListener) EnterPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) {
	var keys []string
	for _, indexCtx := range ctx.IndexColumnNames().GetChildren() {
		columnCtx, ok := indexCtx.(*IndexColumnNameContext)
		if !ok {
			continue
		}
		columnName := l.removeColumnQuote(columnCtx.GetStart())
		keys = append(keys, columnName)
	}
	l.PrimaryKey.Decl = l.getTextFromTokens(ctx.GetStart(), ctx.GetStop())
	for _, columnName := range keys {
		for _, item := range l.Columns {
			if item.Name == columnName {
				l.PrimaryKey.Columns = append(l.PrimaryKey.Columns, item)
				break
			}
		}
	}
}

func (l *StmtListener) EnterUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) {
	var keys []string
	for _, indexCtx := range ctx.IndexColumnNames().GetChildren() {
		columnCtx, ok := indexCtx.(*IndexColumnNameContext)
		if !ok {
			continue
		}
		columnName := l.removeColumnQuote(columnCtx.GetStart())
		keys = append(keys, columnName)
	}
	var uniqueIndex ColumnIndex
	uniqueIndex.Decl = l.getTextFromTokens(ctx.GetStart(), ctx.GetStop())
	for _, columnName := range keys {
		for _, item := range l.Columns {
			if item.Name == columnName {
				uniqueIndex.Columns = append(uniqueIndex.Columns, item)
				break
			}
		}
	}
	l.UniqueKeys = append(l.UniqueKeys, uniqueIndex)
}

func (l *StmtListener) getTextFromTokens(start, end antlr.Token) string {
	return start.GetInputStream().GetText(start.GetStart(), end.GetStop())
}

func (l *StmtListener) removeColumnQuote(column antlr.Token) string {
	var res string
	switch column.GetTokenType() {
	case StmtParserREVERSE_QUOTE_ID, StmtParserCHARSET_REVERSE_QOUTE_STRING, StmtParserSTRING_LITERAL:
		name := column.GetText()
		if len(name) <= 2 {
			return ""
		}
		res = name[1 : len(name)-1]
	default:
		res = column.GetText()
	}
	return res
}
