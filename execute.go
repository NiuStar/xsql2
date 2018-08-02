package xsql2

import (
	"fmt"

)

func (order *XSql2Order) executeNoResult(req string) {
	fmt.Println("ExecuteNoResult执行语句: " , req)
	fmt.Println("ExecuteNoResult执行参数: ",order.args)
	order.xsql2.db.QueryRow(req,order.args...)
}
func (order *XSql2Order) executeForLastInsertId(req string) int64{
	fmt.Println("ExecuteNoResult执行语句: " , req)
	fmt.Println("ExecuteNoResult执行参数: ",order.args)
	r,err :=order.xsql2.db.Exec(req,order.args...)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	n,err := r.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}

func (order *XSql2Order) execute(req string) (results []map[string]interface{}) { //SQL

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("数据库执行错误：", err)
	//	}
	//}()

	fmt.Println("Execute执行语句: " , req)
	fmt.Println("Execute执行参数: ",order.args)

	//s.ch = 0
	//s.xs.mLock.RLock()

	//go timer(s)

	rows, err := order.xsql2.db.Query(req,order.args...)

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
		for i, _ := range values {

			if values[i] == nil {
				if t[order.fields[i].Name] != nil && order.fields[i].AS_ ==""{
					//t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(values[i].([]byte))
					t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = nil
				} else if order.fields[i].AS_ !="" {
					//t[order.fields[i].Name] = byte2Int(values[i].([]byte))
					t[order.fields[i].AS_] = nil
				}else {
					t[order.fields[i].Name] = nil
				}
				//t[order.fields[i].Name] = nil
			} else {
				//fmt.Println("order.fields[i].Type_:",order.fields[i].Type_)
				switch order.fields[i].Type_ {

				case "int":
					{
						//fmt.Println("order.fields[i].Name:",byte2Int(values[i].([]byte)))

						if t[order.fields[i].Name] != nil && order.fields[i].AS_ ==""{
							//t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = values[i].(int64)
						} else if order.fields[i].AS_ !="" {
							//t[order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].AS_] = values[i].(int64)
						}else {
							t[order.fields[i].Name] = values[i].(int64)
						}

					}
					break
				case "float":
					{
						if t[order.fields[i].Name] != nil && order.fields[i].AS_ ==""{
							//t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = values[i].(float32)
						} else if order.fields[i].AS_ !="" {
							//t[order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].AS_] = values[i].(float32)
						}else {
							t[order.fields[i].Name] = values[i].(float32)
						}
					}
					break
				case "string":
					{
						if t[order.fields[i].Name] != nil && order.fields[i].AS_ ==""{
							//t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2String(values[i].([]byte))
						} else if order.fields[i].AS_ !="" {
							//t[order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].AS_] = byte2String(values[i].([]byte))
						}else {
							t[order.fields[i].Name] = byte2String(values[i].([]byte))
						}
					}
					break
				default:
					{
						if t[order.fields[i].Name] != nil && order.fields[i].AS_ ==""{
							//t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].Target.GetName() + "." + order.fields[i].Name] = getInitValue(values[i].([]byte))
						} else if order.fields[i].AS_ !="" {
							//t[order.fields[i].Name] = byte2Int(values[i].([]byte))
							t[order.fields[i].AS_] = getInitValue(values[i].([]byte))
						}else {
							t[order.fields[i].Name] = getInitValue(values[i].([]byte))
						}
						//t[order.fields[i].Name] = getInitValue(values[i].([]byte))
					}
					break
				}
			}

		}
		results = append(results, t)

	}
	return results
}