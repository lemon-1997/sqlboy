// Code generated from ./antlr/StmtParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // StmtParser

import "github.com/antlr4-go/antlr/v4"

// StmtParserListener is a complete listener for a parse tree produced by StmtParser.
type StmtParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterSqlStatement is called when entering the sqlStatement production.
	EnterSqlStatement(c *SqlStatementContext)

	// EnterCharSet is called when entering the charSet production.
	EnterCharSet(c *CharSetContext)

	// EnterIntervalType is called when entering the intervalType production.
	EnterIntervalType(c *IntervalTypeContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterCreateDefinitions is called when entering the createDefinitions production.
	EnterCreateDefinitions(c *CreateDefinitionsContext)

	// EnterColumnDeclaration is called when entering the columnDeclaration production.
	EnterColumnDeclaration(c *ColumnDeclarationContext)

	// EnterConstraintDeclaration is called when entering the constraintDeclaration production.
	EnterConstraintDeclaration(c *ConstraintDeclarationContext)

	// EnterIndexDeclaration is called when entering the indexDeclaration production.
	EnterIndexDeclaration(c *IndexDeclarationContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterNullColumnConstraint is called when entering the nullColumnConstraint production.
	EnterNullColumnConstraint(c *NullColumnConstraintContext)

	// EnterDefaultColumnConstraint is called when entering the defaultColumnConstraint production.
	EnterDefaultColumnConstraint(c *DefaultColumnConstraintContext)

	// EnterVisibilityColumnConstraint is called when entering the visibilityColumnConstraint production.
	EnterVisibilityColumnConstraint(c *VisibilityColumnConstraintContext)

	// EnterAutoIncrementColumnConstraint is called when entering the autoIncrementColumnConstraint production.
	EnterAutoIncrementColumnConstraint(c *AutoIncrementColumnConstraintContext)

	// EnterPrimaryKeyColumnConstraint is called when entering the primaryKeyColumnConstraint production.
	EnterPrimaryKeyColumnConstraint(c *PrimaryKeyColumnConstraintContext)

	// EnterUniqueKeyColumnConstraint is called when entering the uniqueKeyColumnConstraint production.
	EnterUniqueKeyColumnConstraint(c *UniqueKeyColumnConstraintContext)

	// EnterCommentColumnConstraint is called when entering the commentColumnConstraint production.
	EnterCommentColumnConstraint(c *CommentColumnConstraintContext)

	// EnterFormatColumnConstraint is called when entering the formatColumnConstraint production.
	EnterFormatColumnConstraint(c *FormatColumnConstraintContext)

	// EnterStorageColumnConstraint is called when entering the storageColumnConstraint production.
	EnterStorageColumnConstraint(c *StorageColumnConstraintContext)

	// EnterReferenceColumnConstraint is called when entering the referenceColumnConstraint production.
	EnterReferenceColumnConstraint(c *ReferenceColumnConstraintContext)

	// EnterCollateColumnConstraint is called when entering the collateColumnConstraint production.
	EnterCollateColumnConstraint(c *CollateColumnConstraintContext)

	// EnterGeneratedColumnConstraint is called when entering the generatedColumnConstraint production.
	EnterGeneratedColumnConstraint(c *GeneratedColumnConstraintContext)

	// EnterSerialDefaultColumnConstraint is called when entering the serialDefaultColumnConstraint production.
	EnterSerialDefaultColumnConstraint(c *SerialDefaultColumnConstraintContext)

	// EnterCheckColumnConstraint is called when entering the checkColumnConstraint production.
	EnterCheckColumnConstraint(c *CheckColumnConstraintContext)

	// EnterPrimaryKeyTableConstraint is called when entering the primaryKeyTableConstraint production.
	EnterPrimaryKeyTableConstraint(c *PrimaryKeyTableConstraintContext)

	// EnterUniqueKeyTableConstraint is called when entering the uniqueKeyTableConstraint production.
	EnterUniqueKeyTableConstraint(c *UniqueKeyTableConstraintContext)

	// EnterForeignKeyTableConstraint is called when entering the foreignKeyTableConstraint production.
	EnterForeignKeyTableConstraint(c *ForeignKeyTableConstraintContext)

	// EnterCheckTableConstraint is called when entering the checkTableConstraint production.
	EnterCheckTableConstraint(c *CheckTableConstraintContext)

	// EnterReferenceDefinition is called when entering the referenceDefinition production.
	EnterReferenceDefinition(c *ReferenceDefinitionContext)

	// EnterReferenceAction is called when entering the referenceAction production.
	EnterReferenceAction(c *ReferenceActionContext)

	// EnterReferenceControlType is called when entering the referenceControlType production.
	EnterReferenceControlType(c *ReferenceControlTypeContext)

	// EnterSimpleIndexDeclaration is called when entering the simpleIndexDeclaration production.
	EnterSimpleIndexDeclaration(c *SimpleIndexDeclarationContext)

	// EnterSpecialIndexDeclaration is called when entering the specialIndexDeclaration production.
	EnterSpecialIndexDeclaration(c *SpecialIndexDeclarationContext)

	// EnterTableOptionEngine is called when entering the tableOptionEngine production.
	EnterTableOptionEngine(c *TableOptionEngineContext)

	// EnterTableOptionEngineAttribute is called when entering the tableOptionEngineAttribute production.
	EnterTableOptionEngineAttribute(c *TableOptionEngineAttributeContext)

	// EnterTableOptionAutoextendSize is called when entering the tableOptionAutoextendSize production.
	EnterTableOptionAutoextendSize(c *TableOptionAutoextendSizeContext)

	// EnterTableOptionAutoIncrement is called when entering the tableOptionAutoIncrement production.
	EnterTableOptionAutoIncrement(c *TableOptionAutoIncrementContext)

	// EnterTableOptionAverage is called when entering the tableOptionAverage production.
	EnterTableOptionAverage(c *TableOptionAverageContext)

	// EnterTableOptionCharset is called when entering the tableOptionCharset production.
	EnterTableOptionCharset(c *TableOptionCharsetContext)

	// EnterTableOptionChecksum is called when entering the tableOptionChecksum production.
	EnterTableOptionChecksum(c *TableOptionChecksumContext)

	// EnterTableOptionCollate is called when entering the tableOptionCollate production.
	EnterTableOptionCollate(c *TableOptionCollateContext)

	// EnterTableOptionComment is called when entering the tableOptionComment production.
	EnterTableOptionComment(c *TableOptionCommentContext)

	// EnterTableOptionCompression is called when entering the tableOptionCompression production.
	EnterTableOptionCompression(c *TableOptionCompressionContext)

	// EnterTableOptionConnection is called when entering the tableOptionConnection production.
	EnterTableOptionConnection(c *TableOptionConnectionContext)

	// EnterTableOptionDataDirectory is called when entering the tableOptionDataDirectory production.
	EnterTableOptionDataDirectory(c *TableOptionDataDirectoryContext)

	// EnterTableOptionDelay is called when entering the tableOptionDelay production.
	EnterTableOptionDelay(c *TableOptionDelayContext)

	// EnterTableOptionEncryption is called when entering the tableOptionEncryption production.
	EnterTableOptionEncryption(c *TableOptionEncryptionContext)

	// EnterTableOptionIndexDirectory is called when entering the tableOptionIndexDirectory production.
	EnterTableOptionIndexDirectory(c *TableOptionIndexDirectoryContext)

	// EnterTableOptionInsertMethod is called when entering the tableOptionInsertMethod production.
	EnterTableOptionInsertMethod(c *TableOptionInsertMethodContext)

	// EnterTableOptionKeyBlockSize is called when entering the tableOptionKeyBlockSize production.
	EnterTableOptionKeyBlockSize(c *TableOptionKeyBlockSizeContext)

	// EnterTableOptionMaxRows is called when entering the tableOptionMaxRows production.
	EnterTableOptionMaxRows(c *TableOptionMaxRowsContext)

	// EnterTableOptionMinRows is called when entering the tableOptionMinRows production.
	EnterTableOptionMinRows(c *TableOptionMinRowsContext)

	// EnterTableOptionPackKeys is called when entering the tableOptionPackKeys production.
	EnterTableOptionPackKeys(c *TableOptionPackKeysContext)

	// EnterTableOptionPassword is called when entering the tableOptionPassword production.
	EnterTableOptionPassword(c *TableOptionPasswordContext)

	// EnterTableOptionRowFormat is called when entering the tableOptionRowFormat production.
	EnterTableOptionRowFormat(c *TableOptionRowFormatContext)

	// EnterTableOptionStartTransaction is called when entering the tableOptionStartTransaction production.
	EnterTableOptionStartTransaction(c *TableOptionStartTransactionContext)

	// EnterTableOptionSecondaryEngineAttribute is called when entering the tableOptionSecondaryEngineAttribute production.
	EnterTableOptionSecondaryEngineAttribute(c *TableOptionSecondaryEngineAttributeContext)

	// EnterTableOptionRecalculation is called when entering the tableOptionRecalculation production.
	EnterTableOptionRecalculation(c *TableOptionRecalculationContext)

	// EnterTableOptionPersistent is called when entering the tableOptionPersistent production.
	EnterTableOptionPersistent(c *TableOptionPersistentContext)

	// EnterTableOptionSamplePage is called when entering the tableOptionSamplePage production.
	EnterTableOptionSamplePage(c *TableOptionSamplePageContext)

	// EnterTableOptionTablespace is called when entering the tableOptionTablespace production.
	EnterTableOptionTablespace(c *TableOptionTablespaceContext)

	// EnterTableOptionTableType is called when entering the tableOptionTableType production.
	EnterTableOptionTableType(c *TableOptionTableTypeContext)

	// EnterTableOptionUnion is called when entering the tableOptionUnion production.
	EnterTableOptionUnion(c *TableOptionUnionContext)

	// EnterTableType is called when entering the tableType production.
	EnterTableType(c *TableTypeContext)

	// EnterTablespaceStorage is called when entering the tablespaceStorage production.
	EnterTablespaceStorage(c *TablespaceStorageContext)

	// EnterPartitionDefinitions is called when entering the partitionDefinitions production.
	EnterPartitionDefinitions(c *PartitionDefinitionsContext)

	// EnterPartitionFunctionHash is called when entering the partitionFunctionHash production.
	EnterPartitionFunctionHash(c *PartitionFunctionHashContext)

	// EnterPartitionFunctionKey is called when entering the partitionFunctionKey production.
	EnterPartitionFunctionKey(c *PartitionFunctionKeyContext)

	// EnterPartitionFunctionRange is called when entering the partitionFunctionRange production.
	EnterPartitionFunctionRange(c *PartitionFunctionRangeContext)

	// EnterPartitionFunctionList is called when entering the partitionFunctionList production.
	EnterPartitionFunctionList(c *PartitionFunctionListContext)

	// EnterSubPartitionFunctionHash is called when entering the subPartitionFunctionHash production.
	EnterSubPartitionFunctionHash(c *SubPartitionFunctionHashContext)

	// EnterSubPartitionFunctionKey is called when entering the subPartitionFunctionKey production.
	EnterSubPartitionFunctionKey(c *SubPartitionFunctionKeyContext)

	// EnterPartitionComparison is called when entering the partitionComparison production.
	EnterPartitionComparison(c *PartitionComparisonContext)

	// EnterPartitionListAtom is called when entering the partitionListAtom production.
	EnterPartitionListAtom(c *PartitionListAtomContext)

	// EnterPartitionListVector is called when entering the partitionListVector production.
	EnterPartitionListVector(c *PartitionListVectorContext)

	// EnterPartitionSimple is called when entering the partitionSimple production.
	EnterPartitionSimple(c *PartitionSimpleContext)

	// EnterPartitionDefinerAtom is called when entering the partitionDefinerAtom production.
	EnterPartitionDefinerAtom(c *PartitionDefinerAtomContext)

	// EnterPartitionDefinerVector is called when entering the partitionDefinerVector production.
	EnterPartitionDefinerVector(c *PartitionDefinerVectorContext)

	// EnterSubpartitionDefinition is called when entering the subpartitionDefinition production.
	EnterSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// EnterPartitionOptionEngine is called when entering the partitionOptionEngine production.
	EnterPartitionOptionEngine(c *PartitionOptionEngineContext)

	// EnterPartitionOptionComment is called when entering the partitionOptionComment production.
	EnterPartitionOptionComment(c *PartitionOptionCommentContext)

	// EnterPartitionOptionDataDirectory is called when entering the partitionOptionDataDirectory production.
	EnterPartitionOptionDataDirectory(c *PartitionOptionDataDirectoryContext)

	// EnterPartitionOptionIndexDirectory is called when entering the partitionOptionIndexDirectory production.
	EnterPartitionOptionIndexDirectory(c *PartitionOptionIndexDirectoryContext)

	// EnterPartitionOptionMaxRows is called when entering the partitionOptionMaxRows production.
	EnterPartitionOptionMaxRows(c *PartitionOptionMaxRowsContext)

	// EnterPartitionOptionMinRows is called when entering the partitionOptionMinRows production.
	EnterPartitionOptionMinRows(c *PartitionOptionMinRowsContext)

	// EnterPartitionOptionTablespace is called when entering the partitionOptionTablespace production.
	EnterPartitionOptionTablespace(c *PartitionOptionTablespaceContext)

	// EnterPartitionOptionNodeGroup is called when entering the partitionOptionNodeGroup production.
	EnterPartitionOptionNodeGroup(c *PartitionOptionNodeGroupContext)

	// EnterFullId is called when entering the fullId production.
	EnterFullId(c *FullIdContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterFullColumnName is called when entering the fullColumnName production.
	EnterFullColumnName(c *FullColumnNameContext)

	// EnterIndexColumnName is called when entering the indexColumnName production.
	EnterIndexColumnName(c *IndexColumnNameContext)

	// EnterMysqlVariable is called when entering the mysqlVariable production.
	EnterMysqlVariable(c *MysqlVariableContext)

	// EnterCharsetName is called when entering the charsetName production.
	EnterCharsetName(c *CharsetNameContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterEngineName is called when entering the engineName production.
	EnterEngineName(c *EngineNameContext)

	// EnterUid is called when entering the uid production.
	EnterUid(c *UidContext)

	// EnterSimpleId is called when entering the simpleId production.
	EnterSimpleId(c *SimpleIdContext)

	// EnterDottedId is called when entering the dottedId production.
	EnterDottedId(c *DottedIdContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterFileSizeLiteral is called when entering the fileSizeLiteral production.
	EnterFileSizeLiteral(c *FileSizeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *HexadecimalLiteralContext)

	// EnterNullNotnull is called when entering the nullNotnull production.
	EnterNullNotnull(c *NullNotnullContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterIntegerDataType is called when entering the integerDataType production.
	EnterIntegerDataType(c *IntegerDataTypeContext)

	// EnterFloatDataType is called when entering the floatDataType production.
	EnterFloatDataType(c *FloatDataTypeContext)

	// EnterTimeDataType is called when entering the timeDataType production.
	EnterTimeDataType(c *TimeDataTypeContext)

	// EnterStringDataType is called when entering the stringDataType production.
	EnterStringDataType(c *StringDataTypeContext)

	// EnterBinaryDataType is called when entering the binaryDataType production.
	EnterBinaryDataType(c *BinaryDataTypeContext)

	// EnterSpecialDataType is called when entering the specialDataType production.
	EnterSpecialDataType(c *SpecialDataTypeContext)

	// EnterSpatialDataType is called when entering the spatialDataType production.
	EnterSpatialDataType(c *SpatialDataTypeContext)

	// EnterCollectionOptions is called when entering the collectionOptions production.
	EnterCollectionOptions(c *CollectionOptionsContext)

	// EnterConvertedDataType is called when entering the convertedDataType production.
	EnterConvertedDataType(c *ConvertedDataTypeContext)

	// EnterLengthOneDimension is called when entering the lengthOneDimension production.
	EnterLengthOneDimension(c *LengthOneDimensionContext)

	// EnterLengthTwoDimension is called when entering the lengthTwoDimension production.
	EnterLengthTwoDimension(c *LengthTwoDimensionContext)

	// EnterLengthTwoOptionalDimension is called when entering the lengthTwoOptionalDimension production.
	EnterLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// EnterUidList is called when entering the uidList production.
	EnterUidList(c *UidListContext)

	// EnterTables is called when entering the tables production.
	EnterTables(c *TablesContext)

	// EnterIndexColumnNames is called when entering the indexColumnNames production.
	EnterIndexColumnNames(c *IndexColumnNamesContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterCurrentTimestamp is called when entering the currentTimestamp production.
	EnterCurrentTimestamp(c *CurrentTimestampContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterScalarFunctionName is called when entering the scalarFunctionName production.
	EnterScalarFunctionName(c *ScalarFunctionNameContext)

	// EnterIsExpression is called when entering the isExpression production.
	EnterIsExpression(c *IsExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterSoundsLikePredicate is called when entering the soundsLikePredicate production.
	EnterSoundsLikePredicate(c *SoundsLikePredicateContext)

	// EnterExpressionAtomPredicate is called when entering the expressionAtomPredicate production.
	EnterExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// EnterJsonMemberOfPredicate is called when entering the jsonMemberOfPredicate production.
	EnterJsonMemberOfPredicate(c *JsonMemberOfPredicateContext)

	// EnterBinaryComparisonPredicate is called when entering the binaryComparisonPredicate production.
	EnterBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterRegexpPredicate is called when entering the regexpPredicate production.
	EnterRegexpPredicate(c *RegexpPredicateContext)

	// EnterJsonExpressionAtom is called when entering the jsonExpressionAtom production.
	EnterJsonExpressionAtom(c *JsonExpressionAtomContext)

	// EnterUnaryExpressionAtom is called when entering the unaryExpressionAtom production.
	EnterUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// EnterCollateExpressionAtom is called when entering the collateExpressionAtom production.
	EnterCollateExpressionAtom(c *CollateExpressionAtomContext)

	// EnterConstantExpressionAtom is called when entering the constantExpressionAtom production.
	EnterConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// EnterMysqlVariableExpressionAtom is called when entering the mysqlVariableExpressionAtom production.
	EnterMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// EnterBinaryExpressionAtom is called when entering the binaryExpressionAtom production.
	EnterBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// EnterFullColumnNameExpressionAtom is called when entering the fullColumnNameExpressionAtom production.
	EnterFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// EnterBitExpressionAtom is called when entering the bitExpressionAtom production.
	EnterBitExpressionAtom(c *BitExpressionAtomContext)

	// EnterNestedExpressionAtom is called when entering the nestedExpressionAtom production.
	EnterNestedExpressionAtom(c *NestedExpressionAtomContext)

	// EnterNestedRowExpressionAtom is called when entering the nestedRowExpressionAtom production.
	EnterNestedRowExpressionAtom(c *NestedRowExpressionAtomContext)

	// EnterMathExpressionAtom is called when entering the mathExpressionAtom production.
	EnterMathExpressionAtom(c *MathExpressionAtomContext)

	// EnterIntervalExpressionAtom is called when entering the intervalExpressionAtom production.
	EnterIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterBitOperator is called when entering the bitOperator production.
	EnterBitOperator(c *BitOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterJsonOperator is called when entering the jsonOperator production.
	EnterJsonOperator(c *JsonOperatorContext)

	// EnterCharsetNameBase is called when entering the charsetNameBase production.
	EnterCharsetNameBase(c *CharsetNameBaseContext)

	// EnterTransactionLevelBase is called when entering the transactionLevelBase production.
	EnterTransactionLevelBase(c *TransactionLevelBaseContext)

	// EnterPrivilegesBase is called when entering the privilegesBase production.
	EnterPrivilegesBase(c *PrivilegesBaseContext)

	// EnterIntervalTypeBase is called when entering the intervalTypeBase production.
	EnterIntervalTypeBase(c *IntervalTypeBaseContext)

	// EnterDataTypeBase is called when entering the dataTypeBase production.
	EnterDataTypeBase(c *DataTypeBaseContext)

	// EnterKeywordsCanBeId is called when entering the keywordsCanBeId production.
	EnterKeywordsCanBeId(c *KeywordsCanBeIdContext)

	// EnterFunctionNameBase is called when entering the functionNameBase production.
	EnterFunctionNameBase(c *FunctionNameBaseContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitSqlStatement is called when exiting the sqlStatement production.
	ExitSqlStatement(c *SqlStatementContext)

	// ExitCharSet is called when exiting the charSet production.
	ExitCharSet(c *CharSetContext)

	// ExitIntervalType is called when exiting the intervalType production.
	ExitIntervalType(c *IntervalTypeContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitCreateDefinitions is called when exiting the createDefinitions production.
	ExitCreateDefinitions(c *CreateDefinitionsContext)

	// ExitColumnDeclaration is called when exiting the columnDeclaration production.
	ExitColumnDeclaration(c *ColumnDeclarationContext)

	// ExitConstraintDeclaration is called when exiting the constraintDeclaration production.
	ExitConstraintDeclaration(c *ConstraintDeclarationContext)

	// ExitIndexDeclaration is called when exiting the indexDeclaration production.
	ExitIndexDeclaration(c *IndexDeclarationContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitNullColumnConstraint is called when exiting the nullColumnConstraint production.
	ExitNullColumnConstraint(c *NullColumnConstraintContext)

	// ExitDefaultColumnConstraint is called when exiting the defaultColumnConstraint production.
	ExitDefaultColumnConstraint(c *DefaultColumnConstraintContext)

	// ExitVisibilityColumnConstraint is called when exiting the visibilityColumnConstraint production.
	ExitVisibilityColumnConstraint(c *VisibilityColumnConstraintContext)

	// ExitAutoIncrementColumnConstraint is called when exiting the autoIncrementColumnConstraint production.
	ExitAutoIncrementColumnConstraint(c *AutoIncrementColumnConstraintContext)

	// ExitPrimaryKeyColumnConstraint is called when exiting the primaryKeyColumnConstraint production.
	ExitPrimaryKeyColumnConstraint(c *PrimaryKeyColumnConstraintContext)

	// ExitUniqueKeyColumnConstraint is called when exiting the uniqueKeyColumnConstraint production.
	ExitUniqueKeyColumnConstraint(c *UniqueKeyColumnConstraintContext)

	// ExitCommentColumnConstraint is called when exiting the commentColumnConstraint production.
	ExitCommentColumnConstraint(c *CommentColumnConstraintContext)

	// ExitFormatColumnConstraint is called when exiting the formatColumnConstraint production.
	ExitFormatColumnConstraint(c *FormatColumnConstraintContext)

	// ExitStorageColumnConstraint is called when exiting the storageColumnConstraint production.
	ExitStorageColumnConstraint(c *StorageColumnConstraintContext)

	// ExitReferenceColumnConstraint is called when exiting the referenceColumnConstraint production.
	ExitReferenceColumnConstraint(c *ReferenceColumnConstraintContext)

	// ExitCollateColumnConstraint is called when exiting the collateColumnConstraint production.
	ExitCollateColumnConstraint(c *CollateColumnConstraintContext)

	// ExitGeneratedColumnConstraint is called when exiting the generatedColumnConstraint production.
	ExitGeneratedColumnConstraint(c *GeneratedColumnConstraintContext)

	// ExitSerialDefaultColumnConstraint is called when exiting the serialDefaultColumnConstraint production.
	ExitSerialDefaultColumnConstraint(c *SerialDefaultColumnConstraintContext)

	// ExitCheckColumnConstraint is called when exiting the checkColumnConstraint production.
	ExitCheckColumnConstraint(c *CheckColumnConstraintContext)

	// ExitPrimaryKeyTableConstraint is called when exiting the primaryKeyTableConstraint production.
	ExitPrimaryKeyTableConstraint(c *PrimaryKeyTableConstraintContext)

	// ExitUniqueKeyTableConstraint is called when exiting the uniqueKeyTableConstraint production.
	ExitUniqueKeyTableConstraint(c *UniqueKeyTableConstraintContext)

	// ExitForeignKeyTableConstraint is called when exiting the foreignKeyTableConstraint production.
	ExitForeignKeyTableConstraint(c *ForeignKeyTableConstraintContext)

	// ExitCheckTableConstraint is called when exiting the checkTableConstraint production.
	ExitCheckTableConstraint(c *CheckTableConstraintContext)

	// ExitReferenceDefinition is called when exiting the referenceDefinition production.
	ExitReferenceDefinition(c *ReferenceDefinitionContext)

	// ExitReferenceAction is called when exiting the referenceAction production.
	ExitReferenceAction(c *ReferenceActionContext)

	// ExitReferenceControlType is called when exiting the referenceControlType production.
	ExitReferenceControlType(c *ReferenceControlTypeContext)

	// ExitSimpleIndexDeclaration is called when exiting the simpleIndexDeclaration production.
	ExitSimpleIndexDeclaration(c *SimpleIndexDeclarationContext)

	// ExitSpecialIndexDeclaration is called when exiting the specialIndexDeclaration production.
	ExitSpecialIndexDeclaration(c *SpecialIndexDeclarationContext)

	// ExitTableOptionEngine is called when exiting the tableOptionEngine production.
	ExitTableOptionEngine(c *TableOptionEngineContext)

	// ExitTableOptionEngineAttribute is called when exiting the tableOptionEngineAttribute production.
	ExitTableOptionEngineAttribute(c *TableOptionEngineAttributeContext)

	// ExitTableOptionAutoextendSize is called when exiting the tableOptionAutoextendSize production.
	ExitTableOptionAutoextendSize(c *TableOptionAutoextendSizeContext)

	// ExitTableOptionAutoIncrement is called when exiting the tableOptionAutoIncrement production.
	ExitTableOptionAutoIncrement(c *TableOptionAutoIncrementContext)

	// ExitTableOptionAverage is called when exiting the tableOptionAverage production.
	ExitTableOptionAverage(c *TableOptionAverageContext)

	// ExitTableOptionCharset is called when exiting the tableOptionCharset production.
	ExitTableOptionCharset(c *TableOptionCharsetContext)

	// ExitTableOptionChecksum is called when exiting the tableOptionChecksum production.
	ExitTableOptionChecksum(c *TableOptionChecksumContext)

	// ExitTableOptionCollate is called when exiting the tableOptionCollate production.
	ExitTableOptionCollate(c *TableOptionCollateContext)

	// ExitTableOptionComment is called when exiting the tableOptionComment production.
	ExitTableOptionComment(c *TableOptionCommentContext)

	// ExitTableOptionCompression is called when exiting the tableOptionCompression production.
	ExitTableOptionCompression(c *TableOptionCompressionContext)

	// ExitTableOptionConnection is called when exiting the tableOptionConnection production.
	ExitTableOptionConnection(c *TableOptionConnectionContext)

	// ExitTableOptionDataDirectory is called when exiting the tableOptionDataDirectory production.
	ExitTableOptionDataDirectory(c *TableOptionDataDirectoryContext)

	// ExitTableOptionDelay is called when exiting the tableOptionDelay production.
	ExitTableOptionDelay(c *TableOptionDelayContext)

	// ExitTableOptionEncryption is called when exiting the tableOptionEncryption production.
	ExitTableOptionEncryption(c *TableOptionEncryptionContext)

	// ExitTableOptionIndexDirectory is called when exiting the tableOptionIndexDirectory production.
	ExitTableOptionIndexDirectory(c *TableOptionIndexDirectoryContext)

	// ExitTableOptionInsertMethod is called when exiting the tableOptionInsertMethod production.
	ExitTableOptionInsertMethod(c *TableOptionInsertMethodContext)

	// ExitTableOptionKeyBlockSize is called when exiting the tableOptionKeyBlockSize production.
	ExitTableOptionKeyBlockSize(c *TableOptionKeyBlockSizeContext)

	// ExitTableOptionMaxRows is called when exiting the tableOptionMaxRows production.
	ExitTableOptionMaxRows(c *TableOptionMaxRowsContext)

	// ExitTableOptionMinRows is called when exiting the tableOptionMinRows production.
	ExitTableOptionMinRows(c *TableOptionMinRowsContext)

	// ExitTableOptionPackKeys is called when exiting the tableOptionPackKeys production.
	ExitTableOptionPackKeys(c *TableOptionPackKeysContext)

	// ExitTableOptionPassword is called when exiting the tableOptionPassword production.
	ExitTableOptionPassword(c *TableOptionPasswordContext)

	// ExitTableOptionRowFormat is called when exiting the tableOptionRowFormat production.
	ExitTableOptionRowFormat(c *TableOptionRowFormatContext)

	// ExitTableOptionStartTransaction is called when exiting the tableOptionStartTransaction production.
	ExitTableOptionStartTransaction(c *TableOptionStartTransactionContext)

	// ExitTableOptionSecondaryEngineAttribute is called when exiting the tableOptionSecondaryEngineAttribute production.
	ExitTableOptionSecondaryEngineAttribute(c *TableOptionSecondaryEngineAttributeContext)

	// ExitTableOptionRecalculation is called when exiting the tableOptionRecalculation production.
	ExitTableOptionRecalculation(c *TableOptionRecalculationContext)

	// ExitTableOptionPersistent is called when exiting the tableOptionPersistent production.
	ExitTableOptionPersistent(c *TableOptionPersistentContext)

	// ExitTableOptionSamplePage is called when exiting the tableOptionSamplePage production.
	ExitTableOptionSamplePage(c *TableOptionSamplePageContext)

	// ExitTableOptionTablespace is called when exiting the tableOptionTablespace production.
	ExitTableOptionTablespace(c *TableOptionTablespaceContext)

	// ExitTableOptionTableType is called when exiting the tableOptionTableType production.
	ExitTableOptionTableType(c *TableOptionTableTypeContext)

	// ExitTableOptionUnion is called when exiting the tableOptionUnion production.
	ExitTableOptionUnion(c *TableOptionUnionContext)

	// ExitTableType is called when exiting the tableType production.
	ExitTableType(c *TableTypeContext)

	// ExitTablespaceStorage is called when exiting the tablespaceStorage production.
	ExitTablespaceStorage(c *TablespaceStorageContext)

	// ExitPartitionDefinitions is called when exiting the partitionDefinitions production.
	ExitPartitionDefinitions(c *PartitionDefinitionsContext)

	// ExitPartitionFunctionHash is called when exiting the partitionFunctionHash production.
	ExitPartitionFunctionHash(c *PartitionFunctionHashContext)

	// ExitPartitionFunctionKey is called when exiting the partitionFunctionKey production.
	ExitPartitionFunctionKey(c *PartitionFunctionKeyContext)

	// ExitPartitionFunctionRange is called when exiting the partitionFunctionRange production.
	ExitPartitionFunctionRange(c *PartitionFunctionRangeContext)

	// ExitPartitionFunctionList is called when exiting the partitionFunctionList production.
	ExitPartitionFunctionList(c *PartitionFunctionListContext)

	// ExitSubPartitionFunctionHash is called when exiting the subPartitionFunctionHash production.
	ExitSubPartitionFunctionHash(c *SubPartitionFunctionHashContext)

	// ExitSubPartitionFunctionKey is called when exiting the subPartitionFunctionKey production.
	ExitSubPartitionFunctionKey(c *SubPartitionFunctionKeyContext)

	// ExitPartitionComparison is called when exiting the partitionComparison production.
	ExitPartitionComparison(c *PartitionComparisonContext)

	// ExitPartitionListAtom is called when exiting the partitionListAtom production.
	ExitPartitionListAtom(c *PartitionListAtomContext)

	// ExitPartitionListVector is called when exiting the partitionListVector production.
	ExitPartitionListVector(c *PartitionListVectorContext)

	// ExitPartitionSimple is called when exiting the partitionSimple production.
	ExitPartitionSimple(c *PartitionSimpleContext)

	// ExitPartitionDefinerAtom is called when exiting the partitionDefinerAtom production.
	ExitPartitionDefinerAtom(c *PartitionDefinerAtomContext)

	// ExitPartitionDefinerVector is called when exiting the partitionDefinerVector production.
	ExitPartitionDefinerVector(c *PartitionDefinerVectorContext)

	// ExitSubpartitionDefinition is called when exiting the subpartitionDefinition production.
	ExitSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// ExitPartitionOptionEngine is called when exiting the partitionOptionEngine production.
	ExitPartitionOptionEngine(c *PartitionOptionEngineContext)

	// ExitPartitionOptionComment is called when exiting the partitionOptionComment production.
	ExitPartitionOptionComment(c *PartitionOptionCommentContext)

	// ExitPartitionOptionDataDirectory is called when exiting the partitionOptionDataDirectory production.
	ExitPartitionOptionDataDirectory(c *PartitionOptionDataDirectoryContext)

	// ExitPartitionOptionIndexDirectory is called when exiting the partitionOptionIndexDirectory production.
	ExitPartitionOptionIndexDirectory(c *PartitionOptionIndexDirectoryContext)

	// ExitPartitionOptionMaxRows is called when exiting the partitionOptionMaxRows production.
	ExitPartitionOptionMaxRows(c *PartitionOptionMaxRowsContext)

	// ExitPartitionOptionMinRows is called when exiting the partitionOptionMinRows production.
	ExitPartitionOptionMinRows(c *PartitionOptionMinRowsContext)

	// ExitPartitionOptionTablespace is called when exiting the partitionOptionTablespace production.
	ExitPartitionOptionTablespace(c *PartitionOptionTablespaceContext)

	// ExitPartitionOptionNodeGroup is called when exiting the partitionOptionNodeGroup production.
	ExitPartitionOptionNodeGroup(c *PartitionOptionNodeGroupContext)

	// ExitFullId is called when exiting the fullId production.
	ExitFullId(c *FullIdContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitFullColumnName is called when exiting the fullColumnName production.
	ExitFullColumnName(c *FullColumnNameContext)

	// ExitIndexColumnName is called when exiting the indexColumnName production.
	ExitIndexColumnName(c *IndexColumnNameContext)

	// ExitMysqlVariable is called when exiting the mysqlVariable production.
	ExitMysqlVariable(c *MysqlVariableContext)

	// ExitCharsetName is called when exiting the charsetName production.
	ExitCharsetName(c *CharsetNameContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitEngineName is called when exiting the engineName production.
	ExitEngineName(c *EngineNameContext)

	// ExitUid is called when exiting the uid production.
	ExitUid(c *UidContext)

	// ExitSimpleId is called when exiting the simpleId production.
	ExitSimpleId(c *SimpleIdContext)

	// ExitDottedId is called when exiting the dottedId production.
	ExitDottedId(c *DottedIdContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitFileSizeLiteral is called when exiting the fileSizeLiteral production.
	ExitFileSizeLiteral(c *FileSizeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *HexadecimalLiteralContext)

	// ExitNullNotnull is called when exiting the nullNotnull production.
	ExitNullNotnull(c *NullNotnullContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitIntegerDataType is called when exiting the integerDataType production.
	ExitIntegerDataType(c *IntegerDataTypeContext)

	// ExitFloatDataType is called when exiting the floatDataType production.
	ExitFloatDataType(c *FloatDataTypeContext)

	// ExitTimeDataType is called when exiting the timeDataType production.
	ExitTimeDataType(c *TimeDataTypeContext)

	// ExitStringDataType is called when exiting the stringDataType production.
	ExitStringDataType(c *StringDataTypeContext)

	// ExitBinaryDataType is called when exiting the binaryDataType production.
	ExitBinaryDataType(c *BinaryDataTypeContext)

	// ExitSpecialDataType is called when exiting the specialDataType production.
	ExitSpecialDataType(c *SpecialDataTypeContext)

	// ExitSpatialDataType is called when exiting the spatialDataType production.
	ExitSpatialDataType(c *SpatialDataTypeContext)

	// ExitCollectionOptions is called when exiting the collectionOptions production.
	ExitCollectionOptions(c *CollectionOptionsContext)

	// ExitConvertedDataType is called when exiting the convertedDataType production.
	ExitConvertedDataType(c *ConvertedDataTypeContext)

	// ExitLengthOneDimension is called when exiting the lengthOneDimension production.
	ExitLengthOneDimension(c *LengthOneDimensionContext)

	// ExitLengthTwoDimension is called when exiting the lengthTwoDimension production.
	ExitLengthTwoDimension(c *LengthTwoDimensionContext)

	// ExitLengthTwoOptionalDimension is called when exiting the lengthTwoOptionalDimension production.
	ExitLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// ExitUidList is called when exiting the uidList production.
	ExitUidList(c *UidListContext)

	// ExitTables is called when exiting the tables production.
	ExitTables(c *TablesContext)

	// ExitIndexColumnNames is called when exiting the indexColumnNames production.
	ExitIndexColumnNames(c *IndexColumnNamesContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitCurrentTimestamp is called when exiting the currentTimestamp production.
	ExitCurrentTimestamp(c *CurrentTimestampContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitScalarFunctionName is called when exiting the scalarFunctionName production.
	ExitScalarFunctionName(c *ScalarFunctionNameContext)

	// ExitIsExpression is called when exiting the isExpression production.
	ExitIsExpression(c *IsExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitSoundsLikePredicate is called when exiting the soundsLikePredicate production.
	ExitSoundsLikePredicate(c *SoundsLikePredicateContext)

	// ExitExpressionAtomPredicate is called when exiting the expressionAtomPredicate production.
	ExitExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// ExitJsonMemberOfPredicate is called when exiting the jsonMemberOfPredicate production.
	ExitJsonMemberOfPredicate(c *JsonMemberOfPredicateContext)

	// ExitBinaryComparisonPredicate is called when exiting the binaryComparisonPredicate production.
	ExitBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitRegexpPredicate is called when exiting the regexpPredicate production.
	ExitRegexpPredicate(c *RegexpPredicateContext)

	// ExitJsonExpressionAtom is called when exiting the jsonExpressionAtom production.
	ExitJsonExpressionAtom(c *JsonExpressionAtomContext)

	// ExitUnaryExpressionAtom is called when exiting the unaryExpressionAtom production.
	ExitUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// ExitCollateExpressionAtom is called when exiting the collateExpressionAtom production.
	ExitCollateExpressionAtom(c *CollateExpressionAtomContext)

	// ExitConstantExpressionAtom is called when exiting the constantExpressionAtom production.
	ExitConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// ExitMysqlVariableExpressionAtom is called when exiting the mysqlVariableExpressionAtom production.
	ExitMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// ExitBinaryExpressionAtom is called when exiting the binaryExpressionAtom production.
	ExitBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// ExitFullColumnNameExpressionAtom is called when exiting the fullColumnNameExpressionAtom production.
	ExitFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// ExitBitExpressionAtom is called when exiting the bitExpressionAtom production.
	ExitBitExpressionAtom(c *BitExpressionAtomContext)

	// ExitNestedExpressionAtom is called when exiting the nestedExpressionAtom production.
	ExitNestedExpressionAtom(c *NestedExpressionAtomContext)

	// ExitNestedRowExpressionAtom is called when exiting the nestedRowExpressionAtom production.
	ExitNestedRowExpressionAtom(c *NestedRowExpressionAtomContext)

	// ExitMathExpressionAtom is called when exiting the mathExpressionAtom production.
	ExitMathExpressionAtom(c *MathExpressionAtomContext)

	// ExitIntervalExpressionAtom is called when exiting the intervalExpressionAtom production.
	ExitIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitBitOperator is called when exiting the bitOperator production.
	ExitBitOperator(c *BitOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitJsonOperator is called when exiting the jsonOperator production.
	ExitJsonOperator(c *JsonOperatorContext)

	// ExitCharsetNameBase is called when exiting the charsetNameBase production.
	ExitCharsetNameBase(c *CharsetNameBaseContext)

	// ExitTransactionLevelBase is called when exiting the transactionLevelBase production.
	ExitTransactionLevelBase(c *TransactionLevelBaseContext)

	// ExitPrivilegesBase is called when exiting the privilegesBase production.
	ExitPrivilegesBase(c *PrivilegesBaseContext)

	// ExitIntervalTypeBase is called when exiting the intervalTypeBase production.
	ExitIntervalTypeBase(c *IntervalTypeBaseContext)

	// ExitDataTypeBase is called when exiting the dataTypeBase production.
	ExitDataTypeBase(c *DataTypeBaseContext)

	// ExitKeywordsCanBeId is called when exiting the keywordsCanBeId production.
	ExitKeywordsCanBeId(c *KeywordsCanBeIdContext)

	// ExitFunctionNameBase is called when exiting the functionNameBase production.
	ExitFunctionNameBase(c *FunctionNameBaseContext)
}
