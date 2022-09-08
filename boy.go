package sqlboy

type Data interface {
	DDL() string
}

type Config struct {
	Stmt string
	// todo opts(default([xxx]) method(curdl) mod(gorm sqlx sql))
	//Opts opts
}

func (c *Config) DDL() string {
	return c.Stmt
}

func Curd(c *Config) Data {
	return c
}
