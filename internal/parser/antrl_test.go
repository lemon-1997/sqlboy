package parser

import (
	parser "sqlboy/antrl"
	"testing"
)

func Test_parseStmt(t *testing.T) {
	tests := []struct {
		ddl  string
		want parser.TableAttr
	}{
		{
			ddl: `CREATE TABLE "t1" (c1 INT PRIMARY KEY)`,
			want: parser.TableAttr{
				TableName: "t1",
				Columns:   []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
				PrimaryKey: parser.ColumnIndex{
					Columns: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
				},
			},
		},
		{
			ddl: "CREATE TABLE `t1` (c1 INT,PRIMARY KEY (`c1`))",
			want: parser.TableAttr{
				TableName: "t1",
				Columns:   []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
				PrimaryKey: parser.ColumnIndex{
					Columns: []parser.ColumnDecl{{Name: "c1", GoType: parser.Int32}},
				},
			},
		},
		{
			ddl: `CREATE TABLE db."t1" (c1 INT,c2 INT,PRIMARY KEY(c1,c2))`,
			want: parser.TableAttr{
				TableName: "t1",
				Columns: []parser.ColumnDecl{
					{Name: "c1", GoType: parser.Int32},
					{Name: "c2", GoType: parser.Int32},
				},
				PrimaryKey: parser.ColumnIndex{
					Columns: []parser.ColumnDecl{
						{Name: "c1", GoType: parser.Int32},
						{Name: "c2", GoType: parser.Int32},
					},
				},
			},
		},
		{
			ddl: "CREATE TABLE db.t1 (`c1` int(11),`c2` int(11),`c3` int(11),UNIQUE KEY `c1k` (`c1`),UNIQUE KEY `c2c3k` (`c2`,`c3`))",
			want: parser.TableAttr{
				TableName: "t1",
				Columns: []parser.ColumnDecl{
					{Name: "c1", GoType: parser.Int32},
					{Name: "c2", GoType: parser.Int32},
					{Name: "c3", GoType: parser.Int32},
				},
				UniqueKeys: []parser.ColumnIndex{
					{
						Columns: []parser.ColumnDecl{
							{Name: "c1", GoType: parser.Int32},
						},
					},
					{
						Columns: []parser.ColumnDecl{
							{Name: "c2", GoType: parser.Int32},
							{Name: "c3", GoType: parser.Int32},
						},
					},
				},
			},
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
			want: parser.TableAttr{
				TableName: "t1",
				Columns: []parser.ColumnDecl{
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
				PrimaryKey: parser.ColumnIndex{
					Columns: []parser.ColumnDecl{
						{Name: "c1", GoType: parser.Uint64},
					},
				},
			},
		},
	}

	checkColumn := func(t *testing.T, got, want []parser.ColumnDecl) {
		for i := range got {
			if got[i].Name != want[i].Name {
				t.Errorf("decl(%s) name got (%s) want(%s)", got[i].Decl, got[i].Name, want[i].Name)
			}
			if got[i].GoType != want[i].GoType {
				t.Errorf("decl(%s) type got (%s) want(%s)", got[i].Decl, got[i].GoType, want[i].GoType)
			}
			if got[i].Comment != want[i].Comment {
				t.Errorf("decl(%s) comment got (%s) want(%s)", got[i].Decl, got[i].Comment, want[i].Comment)
			}
		}
	}

	for _, tt := range tests {
		table, errs := parseStmt(tt.ddl)
		if len(errs) > 0 {
			t.Fatal(tt.ddl)
		}
		if table.TableName != tt.want.TableName {
			t.Errorf("ddl(%s) tableName got (%s) want(%s)", tt.ddl, table.TableName, tt.want.TableName)
		}
		if len(table.Columns) == len(tt.want.Columns) {
			checkColumn(t, table.Columns, tt.want.Columns)
		} else {
			t.Errorf("ddl(%s) coulumn length got(%d) want(%d)", tt.ddl, len(table.Columns), len(tt.want.Columns))
		}
		if len(table.PrimaryKey.Columns) == len(tt.want.PrimaryKey.Columns) {
			checkColumn(t, table.PrimaryKey.Columns, tt.want.PrimaryKey.Columns)
		} else {
			t.Errorf("ddl(%s) primaryKey length got(%d) want(%d)", tt.ddl, len(table.PrimaryKey.Columns), len(tt.want.PrimaryKey.Columns))
		}
		if len(table.UniqueKeys) == len(tt.want.UniqueKeys) {
			for i := range table.UniqueKeys {
				checkColumn(t, table.UniqueKeys[i].Columns, tt.want.UniqueKeys[i].Columns)
			}
		} else {
			t.Errorf("ddl(%s) uniqueKey length got(%d) want(%d)", tt.ddl, len(table.UniqueKeys), len(tt.want.UniqueKeys))
		}
	}
}
