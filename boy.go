package sqlboy

import (
	"fmt"
	parserAntlr "github.com/lemon-1997/sqlboy/antlr"
	"github.com/lemon-1997/sqlboy/bus"
	"github.com/lemon-1997/sqlboy/inter"
	"github.com/lemon-1997/sqlboy/internal/generator"
	"github.com/lemon-1997/sqlboy/internal/generator/ast"
	"github.com/lemon-1997/sqlboy/internal/generator/render"
	"github.com/lemon-1997/sqlboy/internal/parser"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
)

type GenMode string

const (
	ModeGorm GenMode = "gorm"
	ModeSqlx         = "sqlx"
)

const (
	TopicAstParse       bus.Topic = "topicAstParse"
	TopicAntlrParse               = "topicAntlrParse"
	TopicAssertGenerate           = "topicAssertGenerate"
	TopicModelGenerate            = "topicModelGenerate"
	TopicDaoGenerate              = "topicDaoGenerate"
	TopicTxGenerate               = "topicTxGenerate"
	TopicQueryGenerate            = "topicQueryGenerate"
)

type Option func(*Boy)

func Mode(mode GenMode) Option {
	return func(boy *Boy) {
		boy.mode = mode
	}
}

type Boy struct {
	file       string
	mode       GenMode
	genPackage string

	bus  *bus.AsyncEventBus
	err  chan error
	done chan struct{}
	data chan interface{}
}

func NewBoy(filePath string, opts ...Option) *Boy {
	boy := &Boy{
		file: filePath,
		mode: ModeGorm,

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
	_ = eventBus.Subscribe(TopicAntlrParse, b.eventAntlrParse)
	_ = eventBus.Subscribe(TopicAssertGenerate, b.eventAssertGenerate)
	_ = eventBus.Subscribe(TopicModelGenerate, b.eventModelGenerate)
	_ = eventBus.Subscribe(TopicDaoGenerate, b.eventDaoGenerate)
	_ = eventBus.Subscribe(TopicTxGenerate, b.eventTxGenerate)
	_ = eventBus.Subscribe(TopicQueryGenerate, b.eventQueryGenerate)
	b.bus = eventBus
}

func (b *Boy) Do() error {
	b.bus.Publish(TopicAstParse)
	var genTables, genCount int
	tables := make(map[string][]parserAntlr.ColumnDecl)
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
					s, err := strconv.Unquote(stmt)
					if err != nil {
						return err
					}
					b.bus.Publish(TopicAntlrParse, s)
				}
			case parser.AntlrParseOut:
				res := data.(parser.AntlrParseOut)
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
				return nil
			}
		case err := <-b.err:
			return err
		default:
			continue
		}
	}
}

func (b *Boy) eventAstParse() {
	b.parse(&parser.AstParser{}, parser.AstParseIn{Path: b.file})
}

func (b *Boy) eventAntlrParse(stmt string) {
	b.parse(&parser.AntlrParser{}, parser.AntlrParseIn{Stmt: stmt})
}

func (b *Boy) eventAssertGenerate(asserts map[string]string) {
	b.generate(&ast.AssertGenerator{}, ast.AssertGenIn{
		PackageName: b.genPackage,
		Asserts:     asserts,
	}, "assert.go")
}

func (b *Boy) eventModelGenerate(tables map[string][]parserAntlr.ColumnDecl) {
	b.generate(&ast.ModelGenerator{}, ast.ModelGenIn{
		PackageName: b.genPackage,
		Tables:      tables,
		Mode:        b.genMode(),
	}, "model.go")
}

func (b *Boy) eventDaoGenerate() {
	b.generate(&render.DaoGenerator{}, render.DaoGenIn{
		PackageName: b.genPackage,
		Mode:        b.genMode(),
	}, "dao.go")
}

func (b *Boy) eventTxGenerate() {
	b.generate(&render.TxGenerator{}, render.TxGenIn{
		PackageName: b.genPackage,
		Mode:        b.genMode(),
	}, "transaction.go")
}

func (b *Boy) eventQueryGenerate(data render.QueryData) {
	b.generate(&render.QueryGenerator{}, render.QueryGenIn{
		Data: data,
		Mode: b.genMode(),
	}, fmt.Sprintf("query_%s.go", data.Table))
}

func (b *Boy) parse(parser inter.Parser, in interface{}) {
	res, err := parser.Parse(in)
	if err != nil {
		b.err <- fmt.Errorf("%s:%w", parser.Name(), err)
		return
	}
	b.data <- res
}

func (b *Boy) generate(gen inter.Generator, in interface{}, file string) {
	buf, err := gen.Generate(in)
	if err != nil {
		b.err <- fmt.Errorf("%s:%w", gen.Name(), err)
		return
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		b.err <- fmt.Errorf("%s:%w", gen.Name(), err)
		return
	}
	if err = os.WriteFile(b.genPath(file), source, 0664); err != nil {
		b.err <- fmt.Errorf("%s:%w", gen.Name(), err)
		return
	}
	b.done <- struct{}{}
}

func (b *Boy) transRenderData(res parser.AntlrParseOut) render.QueryData {
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
		uk = append(uk, []render.Column{})
		for _, column := range item.Columns {
			uk[len(uk)-1] = append(uk[len(uk)-1], render.Column{
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

func (b *Boy) genPath(name string) string {
	return filepath.Join(filepath.Dir(b.file), name)
}
