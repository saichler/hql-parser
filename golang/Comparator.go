package hqlparser

import (
	"bytes"
	"errors"
	"strings"
)

type ComparatorOperation string

const (
	Equal              ComparatorOperation = "="
	NotEqual           ComparatorOperation = "!="
	GreaterThan        ComparatorOperation = ">"
	LessThan           ComparatorOperation = "<"
	GreaterThanOrEqual ComparatorOperation = ">="
	LessThanOrEqual    ComparatorOperation = "<="
	IN                 ComparatorOperation = " in "
)

var comparators = make([]ComparatorOperation, 0)

type Comparator struct {
	left  string
	op    ComparatorOperation
	right string
}

func (comparator *Comparator) Left() string {
	return comparator.left
}

func (comparator *Comparator) Right() string {
	return comparator.right
}

func (comparator *Comparator) Operation() ComparatorOperation {
	return comparator.op
}

func initComparators() {
	if len(comparators) == 0 {
		comparators = append(comparators, GreaterThanOrEqual)
		comparators = append(comparators, LessThanOrEqual)
		comparators = append(comparators, NotEqual)
		comparators = append(comparators, Equal)
		comparators = append(comparators, GreaterThan)
		comparators = append(comparators, LessThan)
		comparators = append(comparators, IN)
	}
}

func (comparator *Comparator) String() string {
	buff := bytes.Buffer{}
	buff.WriteString(comparator.left)
	buff.WriteString(string(comparator.op))
	buff.WriteString(comparator.right)
	return buff.String()
}

func (comparator *Comparator) Visualize(lvl int) string {
	buff := bytes.Buffer{}
	buff.WriteString(space(lvl))
	buff.WriteString("Comparator (")
	buff.WriteString(comparator.left)
	buff.WriteString(string(comparator.op))
	buff.WriteString(comparator.right)
	buff.WriteString(")\n")
	return buff.String()
}

func NewCompare(ws string) (*Comparator, error) {
	for _, op := range comparators {
		loc := strings.Index(ws, string(op))
		if loc != -1 {
			cmp := &Comparator{}
			cmp.left = strings.TrimSpace(strings.ToLower(ws[0:loc]))
			cmp.right = strings.TrimSpace(strings.ToLower(ws[loc+len(op):]))
			cmp.op = op
			if validateValue(cmp.left) != "" {
				return nil, errors.New(validateValue(cmp.left))
			}
			if validateValue(cmp.right) != "" {
				return nil, errors.New(validateValue(cmp.right))
			}
			return cmp, nil
		}
	}
	return nil, errors.New("Cannot find comparator operation in: " + ws)
}

func validateValue(ws string) string {
	bo := strings.Index(ws, "(")
	be := strings.Index(ws, ")")
	if bo != -1 || be != -1 {
		return "Value " + ws + " contain illegale brackets."
	}
	return ""
}
