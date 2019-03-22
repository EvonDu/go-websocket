package core

import (
	"flag"
	"strconv"
)

//定义结构体
type Config struct {
	Port	int
	Test    bool
	Swagger	bool
}

func (t *Config) Load(){
	// 命令行参数
	flag.IntVar(&t.Port, "p", 8080, "port")
	flag.BoolVar(&t.Test, "t", false, "debug page")
	flag.BoolVar(&t.Swagger, "s", false, "swagger api document")
	flag.Parse()
}

func (t *Config) GetPortConsoleString() string{
	s := strconv.Itoa(t.Port)
	for i:=0; i<=4-len(s) ;i++ {
		s += " "
	}
	return s
}