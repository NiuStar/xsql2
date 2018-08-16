package xsql2

import "fmt"

func (x *XSql2)Begin(obj... XSqlObject){
	var err error
	x.tx,err = x.db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	x.txopen=1
}

func (x *XSql2)Commit(obj... XSqlObject){
	var err error
	err = x.tx.Commit()
	x.txopen=0
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (x *XSql2)RollBack(obj... XSqlObject){
	var err error
	err = x.tx.Rollback()
	x.txopen=0
	if err != nil {
		fmt.Println(err)
		return
	}
}