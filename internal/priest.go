package internal

import (
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
)

//go:generate gonectr generate -m=../cmd/server
func Priest(cemetery gone.Cemetery) error {
	_ = goner.XormPriest(cemetery)
	return goner.GinPriest(cemetery)
}
