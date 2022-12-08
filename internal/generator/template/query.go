package template

const queryCommon = `
{{- $tableExport := caseExport .Table -}}
{{- $tableInternal := caseInternal .Table -}}

{{- define "keyParams" -}}
{{range .}}, {{caseInternal .Name}} {{.Type}}{{end}}
{{- end -}}

{{- define "keyFunSign" -}} 
{{range .}}{{caseExport .Name}}{{end}} 
{{- end -}}

{{- define "whereValues" -}}
{{range .}}, {{caseInternal .Name}}{{end}}
{{- end -}}

package {{.Package}}

import "context"

type {{$tableExport}}Dao interface {
	Create{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error
	BatchCreate{{$tableExport}}(ctx context.Context, list []*{{$tableExport}}, batchSize int) error

	{{- if len .PrimaryKey}}
	Find{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) (*{{$tableExport}}, error)
	Update{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error
	Delete{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) error
	{{- end}}

	{{- range .UniqueKeys}}
	FindBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) (*{{$tableExport}}, error)
	UpdateBy{{template "keyFunSign" .}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error
	DeleteBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) error
	{{- end}}
}

type {{$tableExport}}Impl struct {
	dao *Dao
}

func New{{$tableExport}}Dao(dao *Dao) {{$tableExport}}Dao {
	return &{{$tableExport}}Impl{
		dao: dao,
	}
}
`

const QueryGorm = queryCommon + `
func (d *{{$tableExport}}Impl) Create{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error {
	return d.dao.DB(ctx).Create({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) BatchCreate{{$tableExport}}(ctx context.Context, list []*{{$tableExport}}, batchSize int) error {
	return d.dao.DB(ctx).CreateInBatches(list, batchSize).Error
}

{{- if len .PrimaryKey}}
{{- $primaryQuery := whereQuery .PrimaryKey}}
func (d *{{$tableExport}}Impl) Find{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).Where("{{$primaryQuery}}"{{template "whereValues" .PrimaryKey}}).First(&{{$tableInternal}}).Error; err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) Update{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	return d.dao.DB(ctx).Where("{{$primaryQuery}}"{{range .PrimaryKey}}, {{$tableInternal}}.{{caseExport .Name}}{{end}}).Updates({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) Delete{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) error{
	return d.dao.DB(ctx).Where("{{$primaryQuery}}"{{template "whereValues" .PrimaryKey}}).Delete(&{{$tableExport}}{}).Error
}

{{- end}}


{{- range .UniqueKeys}}
{{- $uniqueQuery := whereQuery .}}
func (d *{{$tableExport}}Impl) FindBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).Where("{{$uniqueQuery}}"{{template "whereValues" .}}).First(&{{$tableInternal}}).Error; err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) UpdateBy{{template "keyFunSign" .}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	return d.dao.DB(ctx).Where("{{$uniqueQuery}}"{{range .}}, {{$tableInternal}}.{{caseExport .Name}}{{end}}).Updates({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) DeleteBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) error{
	return d.dao.DB(ctx).Where("{{$uniqueQuery}}"{{template "whereValues" .}}).Delete(&{{$tableExport}}{}).Error
}
	
{{- end}}
`

const QuerySqlx = queryCommon + `
{{- $tableQuote := addQuote .Table -}}
{{- $columnsQuote := joinQuote .Columns -}}
{{- $columnsPlace := joinPlace .Columns}}

func (d *{{$tableExport}}Impl) Create{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error {
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "INSERT INTO {{$tableQuote}} ({{$columnsQuote}}) VALUES ({{$columnsPlace}})", {{$tableInternal}})
	return err
}

func (d *{{$tableExport}}Impl) BatchCreate{{$tableExport}}(ctx context.Context, list []*{{$tableExport}}, batchSize int) error {
	return d.dao.InTx(ctx, func(ctx context.Context) error {
		for i := 0; i < len(list); i += batchSize {
			ends := i + batchSize
			if ends > len(list) {
				ends = len(list)
			}
			_, err := d.dao.DB(ctx).NamedExecContext(ctx, "INSERT INTO {{$tableQuote}} ({{$columnsQuote}}) VALUES ({{$columnsPlace}})", list[i:ends])
			if err != nil {
				return err
			}
		}
		return nil
	})
}

{{- if len .PrimaryKey}}

{{- $primaryQuery := whereQuery .PrimaryKey}}
{{- $primaryColumn := sqlxPlaceholder .PrimaryKey}}
{{- $updateColumn := sqlxPlaceholder .Columns}}

func (d *{{$tableExport}}Impl) Find{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).QueryRowxContext(ctx, "SELECT {{$columnsQuote}} FROM {{$tableQuote}} WHERE {{$primaryQuery}}"{{template "whereValues" .PrimaryKey}}).StructScan(&{{$tableInternal}}); err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) Update{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "UPDATE {{$tableQuote}} SET {{$updateColumn}} WHERE {{$primaryColumn}}", {{$tableInternal}})
	return err
}

func (d *{{$tableExport}}Impl) Delete{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) error{
	_, err := d.dao.DB(ctx).ExecContext(ctx, "DELETE FROM {{$tableQuote}} WHERE {{$primaryQuery}}"{{template "whereValues" .PrimaryKey}})
	return err
}

{{- end}}


{{- range .UniqueKeys}}

{{- $uniqueQuery := whereQuery .}}
{{- $uniqueColumn := sqlxPlaceholder .}}
{{- $updateColumn := sqlxPlaceholder $.Columns}}

func (d *{{$tableExport}}Impl) FindBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).QueryRowxContext(ctx, "SELECT {{$columnsQuote}} FROM {{$tableQuote}} WHERE {{$uniqueQuery}}"{{template "whereValues" .}}).StructScan(&{{$tableInternal}}); err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) UpdateBy{{template "keyFunSign" .}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "UPDATE {{$tableQuote}} SET {{$updateColumn}} WHERE {{$uniqueColumn}}", {{$tableInternal}})
	return err
}

func (d *{{$tableExport}}Impl) DeleteBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) error{
	_, err := d.dao.DB(ctx).ExecContext(ctx, "DELETE FROM {{$tableQuote}} WHERE {{$uniqueQuery}}"{{template "whereValues" .}})
	return err
}
	
{{- end}}
`
