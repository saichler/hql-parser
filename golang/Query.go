package hqlparser

import "strings"

type Query struct {
	query  string
	tables []string
	column []string
	where  *Expression
}

const (
	Select = "select"
	From   = "from"
	Where  = "where"
)

func (query *Query) Where() *Expression {
	return query.where
}

func (query *Query) Tables() []string {
	return query.tables
}

func (query *Query) Columns() []string {
	return query.column
}

func NewQuery(query string) (*Query, error) {
	cwql := &Query{}
	cwql.query = query
	e := cwql.init()
	return cwql, e
}

func (query *Query) split() (string, string, string) {
	sql := strings.TrimSpace(strings.ToLower(query.query))
	a := strings.Index(sql, Select)
	if a == -1 {
		return "", "", ""
	}
	b := strings.Index(sql, From)
	if b == -1 {
		return sql, "", ""
	}
	s := strings.TrimSpace(sql[a+len(Select) : b])
	f := strings.TrimSpace(sql[b+len(From):])
	c := strings.Index(f, Where)
	if c == -1 {
		return s, f, ""
	}
	w := strings.TrimSpace(f[c+len(Where):])
	f = strings.TrimSpace(f[0:c])
	return s, f, w
}

func (query *Query) init() error {
	s, f, w := query.split()
	if s != "" {
		columns := strings.Split(s, ",")
		query.column = make([]string, 0)
		for _, col := range columns {
			query.column = append(query.column, strings.TrimSpace(col))
		}
	}
	if f != "" {
		tables := strings.Split(f, ",")
		query.tables = make([]string, 0)
		for _, tbl := range tables {
			query.tables = append(query.tables, strings.TrimSpace(tbl))
		}
	}
	if w != "" {
		where, e := parseExpression(w)
		if e != nil {
			return e
		}
		query.where = where
	}
	return nil
}
