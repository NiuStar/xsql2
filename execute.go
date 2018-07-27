package xsql2

import (
	"fmt"

)

func (order *XSql2Order) executeNoResult(req string) {
	order.xsql2.db.QueryRow(req)
}

func (order *XSql2Order) execute(req string) (results []map[string]interface{}) { //SQL

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("数据库执行错误：", err)
		}
	}()

	fmt.Println("Execute执行sql语句: " + req)

	//s.ch = 0
	//s.xs.mLock.RLock()

	//go timer(s)

	rows, err := order.xsql2.db.Query(req)

	//s.xs.mLock.RUnlock()
	//s.ch = 1
	if err != nil {
		fmt.Println("error: ", err)
		//s.xs.mLock.RLock()
		/*s.xs.db.Close()
		db := createDB(s.xs.name, s.xs.password, s.xs.ip, s.xs.port, s.xs.sqlName)
		s.xs.db = db
		s.xs.time_last = time.Now().Unix()

		rows, err = order.xsql2.db.Query(req)
		defer rows.Close()
		checkErr(err)*/
		return nil
	}

	defer rows.Close()
	columns, err2 := rows.Columns()
	if err2 != nil {
		fmt.Println(err2) // proper error handling instead of panic in your app
		return nil
	}

	if len(columns) <= 0 {
		return nil
	}


	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		t := make(map[string]interface{})
		for i, col := range values {

			if col == nil {
				t[order.fields[i].Name] = nil
			} else {
				fmt.Println("order.fields[i].Type_:",order.fields[i].Type_)
				switch order.fields[i].Type_ {

				case "int":
					{
						fmt.Println("order.fields[i].Name:",byte2Int(col.([]byte)))

						if t[order.fields[i].Name] != nil {
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(col.([]byte))
						} else {
							t[order.fields[i].Name] = byte2Int(col.([]byte))
						}

					}
					break
				case "float":
					{
						if t[order.fields[i].Name] != nil {
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Float(col.([]byte))
						} else {
							t[order.fields[i].Name] = byte2Float(col.([]byte))
						}

					}
					break
				case "string":
					{

						if t[order.fields[i].Name] != nil {
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2String(col.([]byte))
						} else {
							t[order.fields[i].Name] = byte2String(col.([]byte))
						}
					}
					break
				default:
					{
						if t[order.fields[i].Name] != nil {
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = getInitValue(col.([]byte))
						} else {
							t[order.fields[i].Name] = getInitValue(col.([]byte))
						}
						//t[order.fields[i].Name] = getInitValue(col.([]byte))
					}
					break
				}
			}

		}
		results = append(results, t)

	}
	return results
}