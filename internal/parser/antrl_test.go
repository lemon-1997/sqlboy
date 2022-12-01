package parser

import (
	parser "sqlboy/antrl"
	"testing"
)

func Test_parseStmt(t *testing.T) {
	tests := []struct {
		ddl   string
		name  string
		decls []parser.ColumnDecl
	}{
		{
			ddl:   `CREATE TABLE "t1" (c1 INT)`,
			name:  "t1",
			decls: []parser.ColumnDecl{{ColumnName: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   "CREATE TABLE `t1` (c1 INT)",
			name:  "t1",
			decls: []parser.ColumnDecl{{ColumnName: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   `CREATE TABLE db."t1" (c1 INT)`,
			name:  "t1",
			decls: []parser.ColumnDecl{{ColumnName: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   "CREATE TABLE db.t1 (c1 INT)",
			name:  "t1",
			decls: []parser.ColumnDecl{{ColumnName: "c1", GoType: parser.Int32}},
		},
		{
			ddl: `
CREATE TABLE t1 (
    c1 bigint(20) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    c2 bool,
    c3 bit,
    c4 year,
    c5 point)
ENGINE=NDB
COMMENT="NDB_TABLE=READ_BACKUP=0,PARTITION_BALANCE=FOR_RP_BY_NODE";`,
			name: "t1",
			decls: []parser.ColumnDecl{
				{
					ColumnName: "c1",
					GoType:     parser.Uint64,
				},
				{
					ColumnName: "c2",
					GoType:     parser.Bool,
				},
				{
					ColumnName: "c3",
					GoType:     parser.SliceUint8,
				},
				{
					ColumnName: "c4",
					GoType:     parser.Uint16,
				},
				{
					ColumnName: "c5",
					GoType:     parser.Invalid,
				},
			},
		},
	}
	for _, tt := range tests {
		name, decls, errs := parseStmt(tt.ddl)
		if len(errs) > 0 {
			t.Fatal()
		}
		if name != tt.name {
			t.Errorf("ddl(%s) name got (%s) want(%s)", tt.ddl, name, tt.name)
		}
		if len(decls) != len(tt.decls) {
			t.Errorf("ddl(%s) coulumn length got(%d) want(%d)", tt.ddl, len(decls), len(tt.decls))
		}
		for i := range decls {
			if decls[i].ColumnName != tt.decls[i].ColumnName {
				t.Errorf("decl(%s) column got (%s) want(%s)", decls[i].RawText, decls[i].ColumnName, tt.decls[i].ColumnName)
			}
			if decls[i].GoType != tt.decls[i].GoType {
				t.Errorf("decl(%s) type got (%d) want(%d)", decls[i].RawText, decls[i].GoType, tt.decls[i].GoType)
			}
		}
	}
}
