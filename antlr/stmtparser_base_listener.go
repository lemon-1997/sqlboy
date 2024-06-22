// Code generated from ./antlr/StmtParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // StmtParser

import "github.com/antlr4-go/antlr/v4"

// BaseStmtParserListener is a complete listener for a parse tree produced by StmtParser.
type BaseStmtParserListener struct{}

var _ StmtParserListener = &BaseStmtParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStmtParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStmtParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStmtParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStmtParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseStmtParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseStmtParserListener) ExitProg(ctx *ProgContext) {}

// EnterSqlStatement is called when production sqlStatement is entered.
func (s *BaseStmtParserListener) EnterSqlStatement(ctx *SqlStatementContext) {}

// ExitSqlStatement is called when production sqlStatement is exited.
func (s *BaseStmtParserListener) ExitSqlStatement(ctx *SqlStatementContext) {}

// EnterCharSet is called when production charSet is entered.
func (s *BaseStmtParserListener) EnterCharSet(ctx *CharSetContext) {}

// ExitCharSet is called when production charSet is exited.
func (s *BaseStmtParserListener) ExitCharSet(ctx *CharSetContext) {}

// EnterIntervalType is called when production intervalType is entered.
func (s *BaseStmtParserListener) EnterIntervalType(ctx *IntervalTypeContext) {}

// ExitIntervalType is called when production intervalType is exited.
func (s *BaseStmtParserListener) ExitIntervalType(ctx *IntervalTypeContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseStmtParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseStmtParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseStmtParserListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseStmtParserListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterCreateDefinitions is called when production createDefinitions is entered.
func (s *BaseStmtParserListener) EnterCreateDefinitions(ctx *CreateDefinitionsContext) {}

// ExitCreateDefinitions is called when production createDefinitions is exited.
func (s *BaseStmtParserListener) ExitCreateDefinitions(ctx *CreateDefinitionsContext) {}

// EnterColumnDeclaration is called when production columnDeclaration is entered.
func (s *BaseStmtParserListener) EnterColumnDeclaration(ctx *ColumnDeclarationContext) {}

// ExitColumnDeclaration is called when production columnDeclaration is exited.
func (s *BaseStmtParserListener) ExitColumnDeclaration(ctx *ColumnDeclarationContext) {}

// EnterConstraintDeclaration is called when production constraintDeclaration is entered.
func (s *BaseStmtParserListener) EnterConstraintDeclaration(ctx *ConstraintDeclarationContext) {}

// ExitConstraintDeclaration is called when production constraintDeclaration is exited.
func (s *BaseStmtParserListener) ExitConstraintDeclaration(ctx *ConstraintDeclarationContext) {}

// EnterIndexDeclaration is called when production indexDeclaration is entered.
func (s *BaseStmtParserListener) EnterIndexDeclaration(ctx *IndexDeclarationContext) {}

// ExitIndexDeclaration is called when production indexDeclaration is exited.
func (s *BaseStmtParserListener) ExitIndexDeclaration(ctx *IndexDeclarationContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseStmtParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseStmtParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterNullColumnConstraint is called when production nullColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterNullColumnConstraint(ctx *NullColumnConstraintContext) {}

// ExitNullColumnConstraint is called when production nullColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitNullColumnConstraint(ctx *NullColumnConstraintContext) {}

// EnterDefaultColumnConstraint is called when production defaultColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) {}

// ExitDefaultColumnConstraint is called when production defaultColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) {}

// EnterVisibilityColumnConstraint is called when production visibilityColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterVisibilityColumnConstraint(ctx *VisibilityColumnConstraintContext) {
}

// ExitVisibilityColumnConstraint is called when production visibilityColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitVisibilityColumnConstraint(ctx *VisibilityColumnConstraintContext) {
}

// EnterAutoIncrementColumnConstraint is called when production autoIncrementColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) {
}

// ExitAutoIncrementColumnConstraint is called when production autoIncrementColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) {
}

// EnterPrimaryKeyColumnConstraint is called when production primaryKeyColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) {
}

// ExitPrimaryKeyColumnConstraint is called when production primaryKeyColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) {
}

// EnterUniqueKeyColumnConstraint is called when production uniqueKeyColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) {
}

// ExitUniqueKeyColumnConstraint is called when production uniqueKeyColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) {
}

// EnterCommentColumnConstraint is called when production commentColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterCommentColumnConstraint(ctx *CommentColumnConstraintContext) {}

// ExitCommentColumnConstraint is called when production commentColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitCommentColumnConstraint(ctx *CommentColumnConstraintContext) {}

// EnterFormatColumnConstraint is called when production formatColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterFormatColumnConstraint(ctx *FormatColumnConstraintContext) {}

// ExitFormatColumnConstraint is called when production formatColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitFormatColumnConstraint(ctx *FormatColumnConstraintContext) {}

// EnterStorageColumnConstraint is called when production storageColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterStorageColumnConstraint(ctx *StorageColumnConstraintContext) {}

// ExitStorageColumnConstraint is called when production storageColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitStorageColumnConstraint(ctx *StorageColumnConstraintContext) {}

// EnterReferenceColumnConstraint is called when production referenceColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) {
}

// ExitReferenceColumnConstraint is called when production referenceColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) {
}

// EnterCollateColumnConstraint is called when production collateColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterCollateColumnConstraint(ctx *CollateColumnConstraintContext) {}

// ExitCollateColumnConstraint is called when production collateColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitCollateColumnConstraint(ctx *CollateColumnConstraintContext) {}

// EnterGeneratedColumnConstraint is called when production generatedColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) {
}

// ExitGeneratedColumnConstraint is called when production generatedColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) {
}

// EnterSerialDefaultColumnConstraint is called when production serialDefaultColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) {
}

// ExitSerialDefaultColumnConstraint is called when production serialDefaultColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) {
}

// EnterCheckColumnConstraint is called when production checkColumnConstraint is entered.
func (s *BaseStmtParserListener) EnterCheckColumnConstraint(ctx *CheckColumnConstraintContext) {}

// ExitCheckColumnConstraint is called when production checkColumnConstraint is exited.
func (s *BaseStmtParserListener) ExitCheckColumnConstraint(ctx *CheckColumnConstraintContext) {}

// EnterPrimaryKeyTableConstraint is called when production primaryKeyTableConstraint is entered.
func (s *BaseStmtParserListener) EnterPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) {
}

// ExitPrimaryKeyTableConstraint is called when production primaryKeyTableConstraint is exited.
func (s *BaseStmtParserListener) ExitPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) {
}

// EnterUniqueKeyTableConstraint is called when production uniqueKeyTableConstraint is entered.
func (s *BaseStmtParserListener) EnterUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) {
}

// ExitUniqueKeyTableConstraint is called when production uniqueKeyTableConstraint is exited.
func (s *BaseStmtParserListener) ExitUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) {}

// EnterForeignKeyTableConstraint is called when production foreignKeyTableConstraint is entered.
func (s *BaseStmtParserListener) EnterForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) {
}

// ExitForeignKeyTableConstraint is called when production foreignKeyTableConstraint is exited.
func (s *BaseStmtParserListener) ExitForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) {
}

// EnterCheckTableConstraint is called when production checkTableConstraint is entered.
func (s *BaseStmtParserListener) EnterCheckTableConstraint(ctx *CheckTableConstraintContext) {}

// ExitCheckTableConstraint is called when production checkTableConstraint is exited.
func (s *BaseStmtParserListener) ExitCheckTableConstraint(ctx *CheckTableConstraintContext) {}

// EnterReferenceDefinition is called when production referenceDefinition is entered.
func (s *BaseStmtParserListener) EnterReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// ExitReferenceDefinition is called when production referenceDefinition is exited.
func (s *BaseStmtParserListener) ExitReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// EnterReferenceAction is called when production referenceAction is entered.
func (s *BaseStmtParserListener) EnterReferenceAction(ctx *ReferenceActionContext) {}

// ExitReferenceAction is called when production referenceAction is exited.
func (s *BaseStmtParserListener) ExitReferenceAction(ctx *ReferenceActionContext) {}

// EnterReferenceControlType is called when production referenceControlType is entered.
func (s *BaseStmtParserListener) EnterReferenceControlType(ctx *ReferenceControlTypeContext) {}

// ExitReferenceControlType is called when production referenceControlType is exited.
func (s *BaseStmtParserListener) ExitReferenceControlType(ctx *ReferenceControlTypeContext) {}

// EnterSimpleIndexDeclaration is called when production simpleIndexDeclaration is entered.
func (s *BaseStmtParserListener) EnterSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) {}

// ExitSimpleIndexDeclaration is called when production simpleIndexDeclaration is exited.
func (s *BaseStmtParserListener) ExitSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) {}

// EnterSpecialIndexDeclaration is called when production specialIndexDeclaration is entered.
func (s *BaseStmtParserListener) EnterSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) {}

// ExitSpecialIndexDeclaration is called when production specialIndexDeclaration is exited.
func (s *BaseStmtParserListener) ExitSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) {}

// EnterTableOptionEngine is called when production tableOptionEngine is entered.
func (s *BaseStmtParserListener) EnterTableOptionEngine(ctx *TableOptionEngineContext) {}

// ExitTableOptionEngine is called when production tableOptionEngine is exited.
func (s *BaseStmtParserListener) ExitTableOptionEngine(ctx *TableOptionEngineContext) {}

// EnterTableOptionEngineAttribute is called when production tableOptionEngineAttribute is entered.
func (s *BaseStmtParserListener) EnterTableOptionEngineAttribute(ctx *TableOptionEngineAttributeContext) {
}

// ExitTableOptionEngineAttribute is called when production tableOptionEngineAttribute is exited.
func (s *BaseStmtParserListener) ExitTableOptionEngineAttribute(ctx *TableOptionEngineAttributeContext) {
}

// EnterTableOptionAutoextendSize is called when production tableOptionAutoextendSize is entered.
func (s *BaseStmtParserListener) EnterTableOptionAutoextendSize(ctx *TableOptionAutoextendSizeContext) {
}

// ExitTableOptionAutoextendSize is called when production tableOptionAutoextendSize is exited.
func (s *BaseStmtParserListener) ExitTableOptionAutoextendSize(ctx *TableOptionAutoextendSizeContext) {
}

// EnterTableOptionAutoIncrement is called when production tableOptionAutoIncrement is entered.
func (s *BaseStmtParserListener) EnterTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) {
}

// ExitTableOptionAutoIncrement is called when production tableOptionAutoIncrement is exited.
func (s *BaseStmtParserListener) ExitTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) {}

// EnterTableOptionAverage is called when production tableOptionAverage is entered.
func (s *BaseStmtParserListener) EnterTableOptionAverage(ctx *TableOptionAverageContext) {}

// ExitTableOptionAverage is called when production tableOptionAverage is exited.
func (s *BaseStmtParserListener) ExitTableOptionAverage(ctx *TableOptionAverageContext) {}

// EnterTableOptionCharset is called when production tableOptionCharset is entered.
func (s *BaseStmtParserListener) EnterTableOptionCharset(ctx *TableOptionCharsetContext) {}

// ExitTableOptionCharset is called when production tableOptionCharset is exited.
func (s *BaseStmtParserListener) ExitTableOptionCharset(ctx *TableOptionCharsetContext) {}

// EnterTableOptionChecksum is called when production tableOptionChecksum is entered.
func (s *BaseStmtParserListener) EnterTableOptionChecksum(ctx *TableOptionChecksumContext) {}

// ExitTableOptionChecksum is called when production tableOptionChecksum is exited.
func (s *BaseStmtParserListener) ExitTableOptionChecksum(ctx *TableOptionChecksumContext) {}

// EnterTableOptionCollate is called when production tableOptionCollate is entered.
func (s *BaseStmtParserListener) EnterTableOptionCollate(ctx *TableOptionCollateContext) {}

// ExitTableOptionCollate is called when production tableOptionCollate is exited.
func (s *BaseStmtParserListener) ExitTableOptionCollate(ctx *TableOptionCollateContext) {}

// EnterTableOptionComment is called when production tableOptionComment is entered.
func (s *BaseStmtParserListener) EnterTableOptionComment(ctx *TableOptionCommentContext) {}

// ExitTableOptionComment is called when production tableOptionComment is exited.
func (s *BaseStmtParserListener) ExitTableOptionComment(ctx *TableOptionCommentContext) {}

// EnterTableOptionCompression is called when production tableOptionCompression is entered.
func (s *BaseStmtParserListener) EnterTableOptionCompression(ctx *TableOptionCompressionContext) {}

// ExitTableOptionCompression is called when production tableOptionCompression is exited.
func (s *BaseStmtParserListener) ExitTableOptionCompression(ctx *TableOptionCompressionContext) {}

// EnterTableOptionConnection is called when production tableOptionConnection is entered.
func (s *BaseStmtParserListener) EnterTableOptionConnection(ctx *TableOptionConnectionContext) {}

// ExitTableOptionConnection is called when production tableOptionConnection is exited.
func (s *BaseStmtParserListener) ExitTableOptionConnection(ctx *TableOptionConnectionContext) {}

// EnterTableOptionDataDirectory is called when production tableOptionDataDirectory is entered.
func (s *BaseStmtParserListener) EnterTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) {
}

// ExitTableOptionDataDirectory is called when production tableOptionDataDirectory is exited.
func (s *BaseStmtParserListener) ExitTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) {}

// EnterTableOptionDelay is called when production tableOptionDelay is entered.
func (s *BaseStmtParserListener) EnterTableOptionDelay(ctx *TableOptionDelayContext) {}

// ExitTableOptionDelay is called when production tableOptionDelay is exited.
func (s *BaseStmtParserListener) ExitTableOptionDelay(ctx *TableOptionDelayContext) {}

// EnterTableOptionEncryption is called when production tableOptionEncryption is entered.
func (s *BaseStmtParserListener) EnterTableOptionEncryption(ctx *TableOptionEncryptionContext) {}

// ExitTableOptionEncryption is called when production tableOptionEncryption is exited.
func (s *BaseStmtParserListener) ExitTableOptionEncryption(ctx *TableOptionEncryptionContext) {}

// EnterTableOptionIndexDirectory is called when production tableOptionIndexDirectory is entered.
func (s *BaseStmtParserListener) EnterTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) {
}

// ExitTableOptionIndexDirectory is called when production tableOptionIndexDirectory is exited.
func (s *BaseStmtParserListener) ExitTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) {
}

// EnterTableOptionInsertMethod is called when production tableOptionInsertMethod is entered.
func (s *BaseStmtParserListener) EnterTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) {}

// ExitTableOptionInsertMethod is called when production tableOptionInsertMethod is exited.
func (s *BaseStmtParserListener) ExitTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) {}

// EnterTableOptionKeyBlockSize is called when production tableOptionKeyBlockSize is entered.
func (s *BaseStmtParserListener) EnterTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) {}

// ExitTableOptionKeyBlockSize is called when production tableOptionKeyBlockSize is exited.
func (s *BaseStmtParserListener) ExitTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) {}

// EnterTableOptionMaxRows is called when production tableOptionMaxRows is entered.
func (s *BaseStmtParserListener) EnterTableOptionMaxRows(ctx *TableOptionMaxRowsContext) {}

// ExitTableOptionMaxRows is called when production tableOptionMaxRows is exited.
func (s *BaseStmtParserListener) ExitTableOptionMaxRows(ctx *TableOptionMaxRowsContext) {}

// EnterTableOptionMinRows is called when production tableOptionMinRows is entered.
func (s *BaseStmtParserListener) EnterTableOptionMinRows(ctx *TableOptionMinRowsContext) {}

// ExitTableOptionMinRows is called when production tableOptionMinRows is exited.
func (s *BaseStmtParserListener) ExitTableOptionMinRows(ctx *TableOptionMinRowsContext) {}

// EnterTableOptionPackKeys is called when production tableOptionPackKeys is entered.
func (s *BaseStmtParserListener) EnterTableOptionPackKeys(ctx *TableOptionPackKeysContext) {}

// ExitTableOptionPackKeys is called when production tableOptionPackKeys is exited.
func (s *BaseStmtParserListener) ExitTableOptionPackKeys(ctx *TableOptionPackKeysContext) {}

// EnterTableOptionPassword is called when production tableOptionPassword is entered.
func (s *BaseStmtParserListener) EnterTableOptionPassword(ctx *TableOptionPasswordContext) {}

// ExitTableOptionPassword is called when production tableOptionPassword is exited.
func (s *BaseStmtParserListener) ExitTableOptionPassword(ctx *TableOptionPasswordContext) {}

// EnterTableOptionRowFormat is called when production tableOptionRowFormat is entered.
func (s *BaseStmtParserListener) EnterTableOptionRowFormat(ctx *TableOptionRowFormatContext) {}

// ExitTableOptionRowFormat is called when production tableOptionRowFormat is exited.
func (s *BaseStmtParserListener) ExitTableOptionRowFormat(ctx *TableOptionRowFormatContext) {}

// EnterTableOptionStartTransaction is called when production tableOptionStartTransaction is entered.
func (s *BaseStmtParserListener) EnterTableOptionStartTransaction(ctx *TableOptionStartTransactionContext) {
}

// ExitTableOptionStartTransaction is called when production tableOptionStartTransaction is exited.
func (s *BaseStmtParserListener) ExitTableOptionStartTransaction(ctx *TableOptionStartTransactionContext) {
}

// EnterTableOptionSecondaryEngineAttribute is called when production tableOptionSecondaryEngineAttribute is entered.
func (s *BaseStmtParserListener) EnterTableOptionSecondaryEngineAttribute(ctx *TableOptionSecondaryEngineAttributeContext) {
}

// ExitTableOptionSecondaryEngineAttribute is called when production tableOptionSecondaryEngineAttribute is exited.
func (s *BaseStmtParserListener) ExitTableOptionSecondaryEngineAttribute(ctx *TableOptionSecondaryEngineAttributeContext) {
}

// EnterTableOptionRecalculation is called when production tableOptionRecalculation is entered.
func (s *BaseStmtParserListener) EnterTableOptionRecalculation(ctx *TableOptionRecalculationContext) {
}

// ExitTableOptionRecalculation is called when production tableOptionRecalculation is exited.
func (s *BaseStmtParserListener) ExitTableOptionRecalculation(ctx *TableOptionRecalculationContext) {}

// EnterTableOptionPersistent is called when production tableOptionPersistent is entered.
func (s *BaseStmtParserListener) EnterTableOptionPersistent(ctx *TableOptionPersistentContext) {}

// ExitTableOptionPersistent is called when production tableOptionPersistent is exited.
func (s *BaseStmtParserListener) ExitTableOptionPersistent(ctx *TableOptionPersistentContext) {}

// EnterTableOptionSamplePage is called when production tableOptionSamplePage is entered.
func (s *BaseStmtParserListener) EnterTableOptionSamplePage(ctx *TableOptionSamplePageContext) {}

// ExitTableOptionSamplePage is called when production tableOptionSamplePage is exited.
func (s *BaseStmtParserListener) ExitTableOptionSamplePage(ctx *TableOptionSamplePageContext) {}

// EnterTableOptionTablespace is called when production tableOptionTablespace is entered.
func (s *BaseStmtParserListener) EnterTableOptionTablespace(ctx *TableOptionTablespaceContext) {}

// ExitTableOptionTablespace is called when production tableOptionTablespace is exited.
func (s *BaseStmtParserListener) ExitTableOptionTablespace(ctx *TableOptionTablespaceContext) {}

// EnterTableOptionTableType is called when production tableOptionTableType is entered.
func (s *BaseStmtParserListener) EnterTableOptionTableType(ctx *TableOptionTableTypeContext) {}

// ExitTableOptionTableType is called when production tableOptionTableType is exited.
func (s *BaseStmtParserListener) ExitTableOptionTableType(ctx *TableOptionTableTypeContext) {}

// EnterTableOptionUnion is called when production tableOptionUnion is entered.
func (s *BaseStmtParserListener) EnterTableOptionUnion(ctx *TableOptionUnionContext) {}

// ExitTableOptionUnion is called when production tableOptionUnion is exited.
func (s *BaseStmtParserListener) ExitTableOptionUnion(ctx *TableOptionUnionContext) {}

// EnterTableType is called when production tableType is entered.
func (s *BaseStmtParserListener) EnterTableType(ctx *TableTypeContext) {}

// ExitTableType is called when production tableType is exited.
func (s *BaseStmtParserListener) ExitTableType(ctx *TableTypeContext) {}

// EnterTablespaceStorage is called when production tablespaceStorage is entered.
func (s *BaseStmtParserListener) EnterTablespaceStorage(ctx *TablespaceStorageContext) {}

// ExitTablespaceStorage is called when production tablespaceStorage is exited.
func (s *BaseStmtParserListener) ExitTablespaceStorage(ctx *TablespaceStorageContext) {}

// EnterPartitionDefinitions is called when production partitionDefinitions is entered.
func (s *BaseStmtParserListener) EnterPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// ExitPartitionDefinitions is called when production partitionDefinitions is exited.
func (s *BaseStmtParserListener) ExitPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// EnterPartitionFunctionHash is called when production partitionFunctionHash is entered.
func (s *BaseStmtParserListener) EnterPartitionFunctionHash(ctx *PartitionFunctionHashContext) {}

// ExitPartitionFunctionHash is called when production partitionFunctionHash is exited.
func (s *BaseStmtParserListener) ExitPartitionFunctionHash(ctx *PartitionFunctionHashContext) {}

// EnterPartitionFunctionKey is called when production partitionFunctionKey is entered.
func (s *BaseStmtParserListener) EnterPartitionFunctionKey(ctx *PartitionFunctionKeyContext) {}

// ExitPartitionFunctionKey is called when production partitionFunctionKey is exited.
func (s *BaseStmtParserListener) ExitPartitionFunctionKey(ctx *PartitionFunctionKeyContext) {}

// EnterPartitionFunctionRange is called when production partitionFunctionRange is entered.
func (s *BaseStmtParserListener) EnterPartitionFunctionRange(ctx *PartitionFunctionRangeContext) {}

// ExitPartitionFunctionRange is called when production partitionFunctionRange is exited.
func (s *BaseStmtParserListener) ExitPartitionFunctionRange(ctx *PartitionFunctionRangeContext) {}

// EnterPartitionFunctionList is called when production partitionFunctionList is entered.
func (s *BaseStmtParserListener) EnterPartitionFunctionList(ctx *PartitionFunctionListContext) {}

// ExitPartitionFunctionList is called when production partitionFunctionList is exited.
func (s *BaseStmtParserListener) ExitPartitionFunctionList(ctx *PartitionFunctionListContext) {}

// EnterSubPartitionFunctionHash is called when production subPartitionFunctionHash is entered.
func (s *BaseStmtParserListener) EnterSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) {
}

// ExitSubPartitionFunctionHash is called when production subPartitionFunctionHash is exited.
func (s *BaseStmtParserListener) ExitSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) {}

// EnterSubPartitionFunctionKey is called when production subPartitionFunctionKey is entered.
func (s *BaseStmtParserListener) EnterSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) {}

// ExitSubPartitionFunctionKey is called when production subPartitionFunctionKey is exited.
func (s *BaseStmtParserListener) ExitSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) {}

// EnterPartitionComparison is called when production partitionComparison is entered.
func (s *BaseStmtParserListener) EnterPartitionComparison(ctx *PartitionComparisonContext) {}

// ExitPartitionComparison is called when production partitionComparison is exited.
func (s *BaseStmtParserListener) ExitPartitionComparison(ctx *PartitionComparisonContext) {}

// EnterPartitionListAtom is called when production partitionListAtom is entered.
func (s *BaseStmtParserListener) EnterPartitionListAtom(ctx *PartitionListAtomContext) {}

// ExitPartitionListAtom is called when production partitionListAtom is exited.
func (s *BaseStmtParserListener) ExitPartitionListAtom(ctx *PartitionListAtomContext) {}

// EnterPartitionListVector is called when production partitionListVector is entered.
func (s *BaseStmtParserListener) EnterPartitionListVector(ctx *PartitionListVectorContext) {}

// ExitPartitionListVector is called when production partitionListVector is exited.
func (s *BaseStmtParserListener) ExitPartitionListVector(ctx *PartitionListVectorContext) {}

// EnterPartitionSimple is called when production partitionSimple is entered.
func (s *BaseStmtParserListener) EnterPartitionSimple(ctx *PartitionSimpleContext) {}

// ExitPartitionSimple is called when production partitionSimple is exited.
func (s *BaseStmtParserListener) ExitPartitionSimple(ctx *PartitionSimpleContext) {}

// EnterPartitionDefinerAtom is called when production partitionDefinerAtom is entered.
func (s *BaseStmtParserListener) EnterPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) {}

// ExitPartitionDefinerAtom is called when production partitionDefinerAtom is exited.
func (s *BaseStmtParserListener) ExitPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) {}

// EnterPartitionDefinerVector is called when production partitionDefinerVector is entered.
func (s *BaseStmtParserListener) EnterPartitionDefinerVector(ctx *PartitionDefinerVectorContext) {}

// ExitPartitionDefinerVector is called when production partitionDefinerVector is exited.
func (s *BaseStmtParserListener) ExitPartitionDefinerVector(ctx *PartitionDefinerVectorContext) {}

// EnterSubpartitionDefinition is called when production subpartitionDefinition is entered.
func (s *BaseStmtParserListener) EnterSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// ExitSubpartitionDefinition is called when production subpartitionDefinition is exited.
func (s *BaseStmtParserListener) ExitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// EnterPartitionOptionEngine is called when production partitionOptionEngine is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionEngine(ctx *PartitionOptionEngineContext) {}

// ExitPartitionOptionEngine is called when production partitionOptionEngine is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionEngine(ctx *PartitionOptionEngineContext) {}

// EnterPartitionOptionComment is called when production partitionOptionComment is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionComment(ctx *PartitionOptionCommentContext) {}

// ExitPartitionOptionComment is called when production partitionOptionComment is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionComment(ctx *PartitionOptionCommentContext) {}

// EnterPartitionOptionDataDirectory is called when production partitionOptionDataDirectory is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) {
}

// ExitPartitionOptionDataDirectory is called when production partitionOptionDataDirectory is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) {
}

// EnterPartitionOptionIndexDirectory is called when production partitionOptionIndexDirectory is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) {
}

// ExitPartitionOptionIndexDirectory is called when production partitionOptionIndexDirectory is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) {
}

// EnterPartitionOptionMaxRows is called when production partitionOptionMaxRows is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) {}

// ExitPartitionOptionMaxRows is called when production partitionOptionMaxRows is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) {}

// EnterPartitionOptionMinRows is called when production partitionOptionMinRows is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) {}

// ExitPartitionOptionMinRows is called when production partitionOptionMinRows is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) {}

// EnterPartitionOptionTablespace is called when production partitionOptionTablespace is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) {
}

// ExitPartitionOptionTablespace is called when production partitionOptionTablespace is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) {
}

// EnterPartitionOptionNodeGroup is called when production partitionOptionNodeGroup is entered.
func (s *BaseStmtParserListener) EnterPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) {
}

// ExitPartitionOptionNodeGroup is called when production partitionOptionNodeGroup is exited.
func (s *BaseStmtParserListener) ExitPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) {}

// EnterFullId is called when production fullId is entered.
func (s *BaseStmtParserListener) EnterFullId(ctx *FullIdContext) {}

// ExitFullId is called when production fullId is exited.
func (s *BaseStmtParserListener) ExitFullId(ctx *FullIdContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseStmtParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseStmtParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterFullColumnName is called when production fullColumnName is entered.
func (s *BaseStmtParserListener) EnterFullColumnName(ctx *FullColumnNameContext) {}

// ExitFullColumnName is called when production fullColumnName is exited.
func (s *BaseStmtParserListener) ExitFullColumnName(ctx *FullColumnNameContext) {}

// EnterIndexColumnName is called when production indexColumnName is entered.
func (s *BaseStmtParserListener) EnterIndexColumnName(ctx *IndexColumnNameContext) {}

// ExitIndexColumnName is called when production indexColumnName is exited.
func (s *BaseStmtParserListener) ExitIndexColumnName(ctx *IndexColumnNameContext) {}

// EnterMysqlVariable is called when production mysqlVariable is entered.
func (s *BaseStmtParserListener) EnterMysqlVariable(ctx *MysqlVariableContext) {}

// ExitMysqlVariable is called when production mysqlVariable is exited.
func (s *BaseStmtParserListener) ExitMysqlVariable(ctx *MysqlVariableContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseStmtParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseStmtParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseStmtParserListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseStmtParserListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterEngineName is called when production engineName is entered.
func (s *BaseStmtParserListener) EnterEngineName(ctx *EngineNameContext) {}

// ExitEngineName is called when production engineName is exited.
func (s *BaseStmtParserListener) ExitEngineName(ctx *EngineNameContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseStmtParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseStmtParserListener) ExitUid(ctx *UidContext) {}

// EnterSimpleId is called when production simpleId is entered.
func (s *BaseStmtParserListener) EnterSimpleId(ctx *SimpleIdContext) {}

// ExitSimpleId is called when production simpleId is exited.
func (s *BaseStmtParserListener) ExitSimpleId(ctx *SimpleIdContext) {}

// EnterDottedId is called when production dottedId is entered.
func (s *BaseStmtParserListener) EnterDottedId(ctx *DottedIdContext) {}

// ExitDottedId is called when production dottedId is exited.
func (s *BaseStmtParserListener) ExitDottedId(ctx *DottedIdContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseStmtParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseStmtParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterFileSizeLiteral is called when production fileSizeLiteral is entered.
func (s *BaseStmtParserListener) EnterFileSizeLiteral(ctx *FileSizeLiteralContext) {}

// ExitFileSizeLiteral is called when production fileSizeLiteral is exited.
func (s *BaseStmtParserListener) ExitFileSizeLiteral(ctx *FileSizeLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseStmtParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseStmtParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseStmtParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseStmtParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseStmtParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseStmtParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterNullNotnull is called when production nullNotnull is entered.
func (s *BaseStmtParserListener) EnterNullNotnull(ctx *NullNotnullContext) {}

// ExitNullNotnull is called when production nullNotnull is exited.
func (s *BaseStmtParserListener) ExitNullNotnull(ctx *NullNotnullContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseStmtParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseStmtParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterIntegerDataType is called when production integerDataType is entered.
func (s *BaseStmtParserListener) EnterIntegerDataType(ctx *IntegerDataTypeContext) {}

// ExitIntegerDataType is called when production integerDataType is exited.
func (s *BaseStmtParserListener) ExitIntegerDataType(ctx *IntegerDataTypeContext) {}

// EnterFloatDataType is called when production floatDataType is entered.
func (s *BaseStmtParserListener) EnterFloatDataType(ctx *FloatDataTypeContext) {}

// ExitFloatDataType is called when production floatDataType is exited.
func (s *BaseStmtParserListener) ExitFloatDataType(ctx *FloatDataTypeContext) {}

// EnterTimeDataType is called when production timeDataType is entered.
func (s *BaseStmtParserListener) EnterTimeDataType(ctx *TimeDataTypeContext) {}

// ExitTimeDataType is called when production timeDataType is exited.
func (s *BaseStmtParserListener) ExitTimeDataType(ctx *TimeDataTypeContext) {}

// EnterStringDataType is called when production stringDataType is entered.
func (s *BaseStmtParserListener) EnterStringDataType(ctx *StringDataTypeContext) {}

// ExitStringDataType is called when production stringDataType is exited.
func (s *BaseStmtParserListener) ExitStringDataType(ctx *StringDataTypeContext) {}

// EnterBinaryDataType is called when production binaryDataType is entered.
func (s *BaseStmtParserListener) EnterBinaryDataType(ctx *BinaryDataTypeContext) {}

// ExitBinaryDataType is called when production binaryDataType is exited.
func (s *BaseStmtParserListener) ExitBinaryDataType(ctx *BinaryDataTypeContext) {}

// EnterSpecialDataType is called when production specialDataType is entered.
func (s *BaseStmtParserListener) EnterSpecialDataType(ctx *SpecialDataTypeContext) {}

// ExitSpecialDataType is called when production specialDataType is exited.
func (s *BaseStmtParserListener) ExitSpecialDataType(ctx *SpecialDataTypeContext) {}

// EnterSpatialDataType is called when production spatialDataType is entered.
func (s *BaseStmtParserListener) EnterSpatialDataType(ctx *SpatialDataTypeContext) {}

// ExitSpatialDataType is called when production spatialDataType is exited.
func (s *BaseStmtParserListener) ExitSpatialDataType(ctx *SpatialDataTypeContext) {}

// EnterCollectionOptions is called when production collectionOptions is entered.
func (s *BaseStmtParserListener) EnterCollectionOptions(ctx *CollectionOptionsContext) {}

// ExitCollectionOptions is called when production collectionOptions is exited.
func (s *BaseStmtParserListener) ExitCollectionOptions(ctx *CollectionOptionsContext) {}

// EnterConvertedDataType is called when production convertedDataType is entered.
func (s *BaseStmtParserListener) EnterConvertedDataType(ctx *ConvertedDataTypeContext) {}

// ExitConvertedDataType is called when production convertedDataType is exited.
func (s *BaseStmtParserListener) ExitConvertedDataType(ctx *ConvertedDataTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseStmtParserListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseStmtParserListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoDimension is called when production lengthTwoDimension is entered.
func (s *BaseStmtParserListener) EnterLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// ExitLengthTwoDimension is called when production lengthTwoDimension is exited.
func (s *BaseStmtParserListener) ExitLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseStmtParserListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseStmtParserListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterUidList is called when production uidList is entered.
func (s *BaseStmtParserListener) EnterUidList(ctx *UidListContext) {}

// ExitUidList is called when production uidList is exited.
func (s *BaseStmtParserListener) ExitUidList(ctx *UidListContext) {}

// EnterTables is called when production tables is entered.
func (s *BaseStmtParserListener) EnterTables(ctx *TablesContext) {}

// ExitTables is called when production tables is exited.
func (s *BaseStmtParserListener) ExitTables(ctx *TablesContext) {}

// EnterIndexColumnNames is called when production indexColumnNames is entered.
func (s *BaseStmtParserListener) EnterIndexColumnNames(ctx *IndexColumnNamesContext) {}

// ExitIndexColumnNames is called when production indexColumnNames is exited.
func (s *BaseStmtParserListener) ExitIndexColumnNames(ctx *IndexColumnNamesContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseStmtParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseStmtParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterCurrentTimestamp is called when production currentTimestamp is entered.
func (s *BaseStmtParserListener) EnterCurrentTimestamp(ctx *CurrentTimestampContext) {}

// ExitCurrentTimestamp is called when production currentTimestamp is exited.
func (s *BaseStmtParserListener) ExitCurrentTimestamp(ctx *CurrentTimestampContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseStmtParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseStmtParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterScalarFunctionName is called when production scalarFunctionName is entered.
func (s *BaseStmtParserListener) EnterScalarFunctionName(ctx *ScalarFunctionNameContext) {}

// ExitScalarFunctionName is called when production scalarFunctionName is exited.
func (s *BaseStmtParserListener) ExitScalarFunctionName(ctx *ScalarFunctionNameContext) {}

// EnterIsExpression is called when production isExpression is entered.
func (s *BaseStmtParserListener) EnterIsExpression(ctx *IsExpressionContext) {}

// ExitIsExpression is called when production isExpression is exited.
func (s *BaseStmtParserListener) ExitIsExpression(ctx *IsExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseStmtParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseStmtParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseStmtParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseStmtParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseStmtParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseStmtParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterSoundsLikePredicate is called when production soundsLikePredicate is entered.
func (s *BaseStmtParserListener) EnterSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// ExitSoundsLikePredicate is called when production soundsLikePredicate is exited.
func (s *BaseStmtParserListener) ExitSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseStmtParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseStmtParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterJsonMemberOfPredicate is called when production jsonMemberOfPredicate is entered.
func (s *BaseStmtParserListener) EnterJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) {}

// ExitJsonMemberOfPredicate is called when production jsonMemberOfPredicate is exited.
func (s *BaseStmtParserListener) ExitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) {}

// EnterBinaryComparisonPredicate is called when production binaryComparisonPredicate is entered.
func (s *BaseStmtParserListener) EnterBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// ExitBinaryComparisonPredicate is called when production binaryComparisonPredicate is exited.
func (s *BaseStmtParserListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseStmtParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseStmtParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseStmtParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseStmtParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseStmtParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseStmtParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterRegexpPredicate is called when production regexpPredicate is entered.
func (s *BaseStmtParserListener) EnterRegexpPredicate(ctx *RegexpPredicateContext) {}

// ExitRegexpPredicate is called when production regexpPredicate is exited.
func (s *BaseStmtParserListener) ExitRegexpPredicate(ctx *RegexpPredicateContext) {}

// EnterJsonExpressionAtom is called when production jsonExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterJsonExpressionAtom(ctx *JsonExpressionAtomContext) {}

// ExitJsonExpressionAtom is called when production jsonExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitJsonExpressionAtom(ctx *JsonExpressionAtomContext) {}

// EnterUnaryExpressionAtom is called when production unaryExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// ExitUnaryExpressionAtom is called when production unaryExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// EnterCollateExpressionAtom is called when production collateExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterCollateExpressionAtom(ctx *CollateExpressionAtomContext) {}

// ExitCollateExpressionAtom is called when production collateExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitCollateExpressionAtom(ctx *CollateExpressionAtomContext) {}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// ExitMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// EnterBinaryExpressionAtom is called when production binaryExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// ExitBinaryExpressionAtom is called when production binaryExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// EnterFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// ExitFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// EnterBitExpressionAtom is called when production bitExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// ExitBitExpressionAtom is called when production bitExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// EnterNestedExpressionAtom is called when production nestedExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// ExitNestedExpressionAtom is called when production nestedExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// EnterNestedRowExpressionAtom is called when production nestedRowExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) {}

// ExitNestedRowExpressionAtom is called when production nestedRowExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) {}

// EnterMathExpressionAtom is called when production mathExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// ExitMathExpressionAtom is called when production mathExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// EnterIntervalExpressionAtom is called when production intervalExpressionAtom is entered.
func (s *BaseStmtParserListener) EnterIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// ExitIntervalExpressionAtom is called when production intervalExpressionAtom is exited.
func (s *BaseStmtParserListener) ExitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseStmtParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseStmtParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseStmtParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseStmtParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseStmtParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseStmtParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterBitOperator is called when production bitOperator is entered.
func (s *BaseStmtParserListener) EnterBitOperator(ctx *BitOperatorContext) {}

// ExitBitOperator is called when production bitOperator is exited.
func (s *BaseStmtParserListener) ExitBitOperator(ctx *BitOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseStmtParserListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseStmtParserListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterJsonOperator is called when production jsonOperator is entered.
func (s *BaseStmtParserListener) EnterJsonOperator(ctx *JsonOperatorContext) {}

// ExitJsonOperator is called when production jsonOperator is exited.
func (s *BaseStmtParserListener) ExitJsonOperator(ctx *JsonOperatorContext) {}

// EnterCharsetNameBase is called when production charsetNameBase is entered.
func (s *BaseStmtParserListener) EnterCharsetNameBase(ctx *CharsetNameBaseContext) {}

// ExitCharsetNameBase is called when production charsetNameBase is exited.
func (s *BaseStmtParserListener) ExitCharsetNameBase(ctx *CharsetNameBaseContext) {}

// EnterTransactionLevelBase is called when production transactionLevelBase is entered.
func (s *BaseStmtParserListener) EnterTransactionLevelBase(ctx *TransactionLevelBaseContext) {}

// ExitTransactionLevelBase is called when production transactionLevelBase is exited.
func (s *BaseStmtParserListener) ExitTransactionLevelBase(ctx *TransactionLevelBaseContext) {}

// EnterPrivilegesBase is called when production privilegesBase is entered.
func (s *BaseStmtParserListener) EnterPrivilegesBase(ctx *PrivilegesBaseContext) {}

// ExitPrivilegesBase is called when production privilegesBase is exited.
func (s *BaseStmtParserListener) ExitPrivilegesBase(ctx *PrivilegesBaseContext) {}

// EnterIntervalTypeBase is called when production intervalTypeBase is entered.
func (s *BaseStmtParserListener) EnterIntervalTypeBase(ctx *IntervalTypeBaseContext) {}

// ExitIntervalTypeBase is called when production intervalTypeBase is exited.
func (s *BaseStmtParserListener) ExitIntervalTypeBase(ctx *IntervalTypeBaseContext) {}

// EnterDataTypeBase is called when production dataTypeBase is entered.
func (s *BaseStmtParserListener) EnterDataTypeBase(ctx *DataTypeBaseContext) {}

// ExitDataTypeBase is called when production dataTypeBase is exited.
func (s *BaseStmtParserListener) ExitDataTypeBase(ctx *DataTypeBaseContext) {}

// EnterKeywordsCanBeId is called when production keywordsCanBeId is entered.
func (s *BaseStmtParserListener) EnterKeywordsCanBeId(ctx *KeywordsCanBeIdContext) {}

// ExitKeywordsCanBeId is called when production keywordsCanBeId is exited.
func (s *BaseStmtParserListener) ExitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) {}

// EnterFunctionNameBase is called when production functionNameBase is entered.
func (s *BaseStmtParserListener) EnterFunctionNameBase(ctx *FunctionNameBaseContext) {}

// ExitFunctionNameBase is called when production functionNameBase is exited.
func (s *BaseStmtParserListener) ExitFunctionNameBase(ctx *FunctionNameBaseContext) {}
