package clause

import (
	"strings"
)

// Clause contains SQL conditions
type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// Type is the type of Clause
type Type int


const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

// Set adds a sub clause of specific type
// Set 方法根据 Type 调用对应的 generator，生成该子句对应的 SQL 语句。
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build generate the final SQL and SQLVars
// Build 方法根据传入的 Type 的顺序，构造出最终的 SQL 语句。
// sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
//整个文件就是用来构造sql语句的，generator是构造子句的一堆函数