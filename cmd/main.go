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
		opts = append(opts, sqlboy.Mode(sqlboy.GenMode(mode)))
	}
	boy := sqlboy.NewBoy(os.Args[1], opts...)
	err := boy.Do()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
