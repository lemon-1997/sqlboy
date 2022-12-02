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
			decls: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   "CREATE TABLE `t1` (c1 INT)",
			name:  "t1",
			decls: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   `CREATE TABLE db."t1" (c1 INT)`,
			name:  "t1",
			decls: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
		},
		{
			ddl:   "CREATE TABLE db.t1 (c1 INT)",
			name:  "t1",
			decls: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
		},
		{
			ddl: `
CREATE TABLE t1 (
    c1 bigint(20) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    c2 bool COMMENT 'this is comment',
    c3 bit,
    c4 year,
    c5 point)
ENGINE=NDB
COMMENT="NDB_TABLE=READ_BACKUP=0,PARTITION_BALANCE=FOR_RP_BY_NODE";`,
			name: "t1",
			decls: []parser.ColumnDecl{
				{
					Name:   "c1",
					GoType: parser.Uint64,
				},
				{
					Name:    "c2",
					GoType:  parser.Bool,
					Comment: "this is comment",
				},
				{
					Name:   "c3",
					GoType: parser.SliceUint8,
				},
				{
					Name:   "c4",
					GoType: parser.Uint16,
				},
				{
					Name:   "c5",
					GoType: parser.Invalid,
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
			if decls[i].Name != tt.decls[i].Name {
				t.Errorf("decl(%s) name got (%s) want(%s)", decls[i].Decl, decls[i].Name, tt.decls[i].Name)
			}
			if decls[i].GoType != tt.decls[i].GoType {
				t.Errorf("decl(%s) type got (%d) want(%d)", decls[i].Decl, decls[i].GoType, tt.decls[i].GoType)
			}
			if decls[i].Comment != tt.decls[i].Comment {
				t.Errorf("decl(%s) comment got (%s) want(%s)", decls[i].Decl, decls[i].Comment, tt.decls[i].Comment)
			}
		}
	}
}
