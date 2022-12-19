package main

import (
	"flag"
	"fmt"
	"github.com/lemon-1997/sqlboy"
	"log"
	"os"
)

const (
	usage = `sqlboy [packages]
    sqlboy $path -mode gorm
    Find more information at: https://github.com/lemon-1997/sqlboy
`
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("sqlboy:")

	if len(os.Args) < 2 {
		log.Fatal("no specify file")
	}

	flag.Usage = func() {
		fmt.Print(usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	var mode string
	fs := flag.NewFlagSet("sqlboy", flag.ExitOnError)
	fs.StringVar(&mode, "mode", "", "gorm or sqlx")
	_ = fs.Parse(os.Args[2:])

	var opts []sqlboy.Option
	if mode != "" {
		genMode := sqlboy.GenMode(mode)
		if genMode != sqlboy.ModeGorm && genMode != sqlboy.ModeSqlx {
			log.Fatalf("mode %s is not gorm or sqlx", mode)
		}
		opts = append(opts, sqlboy.Mode(genMode))
	}
	boy := sqlboy.NewBoy(os.Args[1], opts...)
	err := boy.Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("generate success")
	os.Exit(0)
}
