package template

const QueryGorm = `
{{- $tableExport := caseExport .Table -}}
{{- $tableInternal := caseInternal .Table -}}

{{- define "keyParams" -}}
{{range .}}, {{caseInternal .Name}} {{.Type}}{{end}}
{{- end -}}

{{- define "keyFunSign" -}} 
{{range .}}{{caseExport .Name}}{{end}} 
{{- end -}}

{{- define "whereQuery" -}}
{{range .}} AND {{.Name}} = ?{{end}}
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

func (d *{{$tableExport}}Impl) Create{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error {
	return d.dao.DB(ctx).Create({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) BatchCreate{{$tableExport}}(ctx context.Context, list []*{{$tableExport}}, batchSize int) error {
	return d.dao.DB(ctx).CreateInBatches(list, batchSize).Error
}


{{- if len .PrimaryKey}}

func (d *{{$tableExport}}Impl) Find{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).Where("1=1{{template "whereQuery" .PrimaryKey}}"{{template "whereValues" .PrimaryKey}}).First(&{{$tableInternal}}).Error; err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) Update{{$tableExport}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	return d.dao.DB(ctx).Where("1=1{{template "whereQuery" .PrimaryKey}}"{{range .PrimaryKey}}, {{$tableInternal}}.{{caseExport .Name}}{{end}}).Updates({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) Delete{{$tableExport}}(ctx context.Context{{template "keyParams" .PrimaryKey}}) error{
	return d.dao.DB(ctx).Where("1=1{{template "whereQuery" .PrimaryKey}}"{{template "whereValues" .PrimaryKey}}).Delete(&{{$tableExport}}{}).Error
}

{{- end}}


{{- range .UniqueKeys}}
	
func (d *{{$tableExport}}Impl) FindBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) (*{{$tableExport}}, error){
	var {{$tableInternal}} {{$tableExport}}
	if err := d.dao.DB(ctx).Where("1=1{{template "whereQuery" .}}"{{template "whereValues" .}}).First(&{{$tableInternal}}).Error; err != nil {
		return nil, err
	}
	return &{{$tableInternal}}, nil
}

func (d *{{$tableExport}}Impl) UpdateBy{{template "keyFunSign" .}}(ctx context.Context, {{$tableInternal}} *{{$tableExport}}) error{
	return d.dao.DB(ctx).Where("1=1{{template "whereQuery" .}}"{{range .}}, {{$tableInternal}}.{{caseExport .Name}}{{end}}).Updates({{$tableInternal}}).Error
}

func (d *{{$tableExport}}Impl) DeleteBy{{template "keyFunSign" .}}(ctx context.Context{{template "keyParams" .}}) error{
	return d.dao.DB(ctx).Where("1=1{{template "whereQuery" .}}"{{template "whereValues" .}}).Delete(&{{$tableExport}}{}).Error
}
	
{{- end}}
`
