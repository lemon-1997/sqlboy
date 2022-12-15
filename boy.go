package sqlboy

import (
	"fmt"
	"go/format"
	"os"
	parserAntrl "sqlboy/antrl"
	"sqlboy/bus"
	"sqlboy/internal"
	"sqlboy/internal/generator"
	"sqlboy/internal/generator/ast"
	"sqlboy/internal/generator/render"
	"sqlboy/internal/parser"
)

type GenMode string

const (
	ModeGorm GenMode = "gorm"
	ModeSqlx GenMode = "sqlx"
)

const (
	TopicAstParse       bus.Topic = "topicAstParse"
	TopicAntrlParse     bus.Topic = "topicAntrlParse"
	TopicAssertGenerate bus.Topic = "topicAssertGenerate"
	TopicModelGenerate  bus.Topic = "topicModelGenerate"
	TopicDaoGenerate    bus.Topic = "topicDaoGenerate"
	TopicTxGenerate     bus.Topic = "topicTxGenerate"
	TopicQueryGenerate  bus.Topic = "topicQueryGenerate"
)

type Option func(*Boy)

func Mode(mode GenMode) Option {
	return func(boy *Boy) {
		boy.mode = mode
	}
}

type Boy struct {
	file string
	mode GenMode

	genDir     string
	genPackage string

	bus  *bus.AsyncEventBus
	err  chan error
	done chan struct{}
	data chan interface{}
}

func NewBoy(filePath string, opts ...Option) *Boy {
	// todo 获取文件夹 判断文件路径是否正确
	boy := &Boy{
		file:   filePath,
		mode:   ModeGorm,
		genDir: "./testGorm",

		err:  make(chan error),
		done: make(chan struct{}),
		data: make(chan interface{}, 10),
	}
	for _, opt := range opts {
		opt(boy)
	}
	boy.register()
	return boy
}

func (b *Boy) register() {
	eventBus := bus.NewAsyncEventBus()
	_ = eventBus.Subscribe(TopicAstParse, b.eventAstParse)
	_ = eventBus.Subscribe(TopicAntrlParse, b.eventAntrlParse)
	_ = eventBus.Subscribe(TopicAssertGenerate, b.eventAssertGenerate)
	_ = eventBus.Subscribe(TopicModelGenerate, b.eventModelGenerate)
	_ = eventBus.Subscribe(TopicDaoGenerate, b.eventDaoGenerate)
	_ = eventBus.Subscribe(TopicTxGenerate, b.eventTxGenerate)
	_ = eventBus.Subscribe(TopicQueryGenerate, b.eventQueryGenerate)
	b.bus = eventBus
}

func (b *Boy) Do() {
	b.bus.Publish(TopicAstParse)
	var genTables, genCount int
	tables := make(map[string][]parserAntrl.ColumnDecl)
	for {
		select {
		case data := <-b.data:
			switch data.(type) {
			case parser.AstParseOut:
				res := data.(parser.AstParseOut)
				genTables = len(res.Stmt)
				b.genPackage = res.PackageName
				b.bus.Publish(TopicDaoGenerate)
				b.bus.Publish(TopicTxGenerate)
				b.bus.Publish(TopicAssertGenerate, res.Stmt)
				for _, stmt := range res.Stmt {
					if len(stmt) <= 2 {
						return
					}
					b.bus.Publish(TopicAntrlParse, stmt[1:len(stmt)-1])
				}
			case parser.AntrlParseOut:
				res := data.(parser.AntrlParseOut)
				tables[res.TableName] = res.Columns
				b.bus.Publish(TopicQueryGenerate, b.transRenderData(res))
				if len(tables) == genTables {
					b.bus.Publish(TopicModelGenerate, tables)
				}
			}
		case <-b.done:
			genCount++
			// assert,model,dao,tx is 4 file
			if genCount >= genTables+4 {
				return
			}
		case err := <-b.err:
			panic(err)
		default:
			continue
		}
	}
}

func (b *Boy) eventAstParse() {
	b.parse(&parser.AstParser{}, parser.AstParseIn{Path: b.file})
}

func (b *Boy) eventAntrlParse(stmt string) {
	b.parse(&parser.AntrlParser{}, parser.AntrlParseIn{Stmt: stmt})
}

func (b *Boy) eventAssertGenerate(asserts map[string]string) {
	b.generate(&ast.AssertGenerator{}, ast.AssertGenIn{
		PackageName: b.genPackage,
		Asserts:     asserts,
	}, b.genDir+"/assert.go")
}

func (b *Boy) eventModelGenerate(tables map[string][]parserAntrl.ColumnDecl) {
	b.generate(&ast.ModelGenerator{}, ast.ModelGenIn{
		PackageName: b.genPackage,
		Tables:      tables,
		Mode:        b.genMode(),
	}, b.genDir+"/model.go")
}

func (b *Boy) eventDaoGenerate() {
	b.generate(&render.DaoGenerator{}, render.DaoGenIn{
		PackageName: b.genPackage,
		Mode:        b.genMode(),
	}, b.genDir+"/dao.go")
}

func (b *Boy) eventTxGenerate() {
	b.generate(&render.TxGenerator{}, render.TxGenIn{
		PackageName: b.genPackage,
		Mode:        b.genMode(),
	}, b.genDir+"/transaction.go")
}

func (b *Boy) eventQueryGenerate(data render.QueryData) {
	b.generate(&render.QueryGenerator{}, render.QueryGenIn{
		Data: data,
		Mode: b.genMode(),
	}, b.genDir+fmt.Sprintf("/%s_query.go", data.Table))
}

func (b *Boy) parse(parser internal.Parser, in interface{}) {
	res, err := parser.Parse(in)
	if err != nil {
		b.err <- err
		return
	}
	b.data <- res
}

func (b *Boy) generate(gen internal.Generator, in interface{}, file string) {
	buf, err := gen.Generate(in)
	if err != nil {
		b.err <- err
		return
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		b.err <- err
		return
	}
	if err = os.WriteFile(file, source, 0664); err != nil {
		b.err <- err
		return
	}
	b.done <- struct{}{}
}

func (b *Boy) transRenderData(res parser.AntrlParseOut) render.QueryData {
	var columns []render.Column
	for _, item := range res.Columns {
		columns = append(columns, render.Column{
			Name:      item.Name,
			Type:      string(item.GoType),
			IsNotNull: item.IsNotNull,
		})
	}

	var pk []render.Column
	for _, item := range res.PrimaryKey.Columns {
		pk = append(pk, render.Column{
			Name:      item.Name,
			Type:      string(item.GoType),
			IsNotNull: item.IsNotNull,
		})
	}

	var uk [][]render.Column
	for _, item := range res.UniqueKeys {
		for _, column := range item.Columns {
			pk = append(pk, render.Column{
				Name:      column.Name,
				Type:      string(column.GoType),
				IsNotNull: column.IsNotNull,
			})
		}
	}

	return render.QueryData{
		Package:    b.genPackage,
		Table:      res.TableName,
		Columns:    columns,
		PrimaryKey: pk,
		UniqueKeys: uk,
	}
}

func (b *Boy) genMode() generator.Mode {
	switch b.mode {
	case ModeGorm:
		return generator.ModeGorm
	case ModeSqlx:
		return generator.ModeSqlx
	default:
		return generator.ModeGorm
	}
}
