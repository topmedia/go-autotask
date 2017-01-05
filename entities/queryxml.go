package entities

import (
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/beevik/etree"
)

type QueryExpression struct {
	Field string
	Op    string
	Value string
}

func (qe *QueryExpression) ToQueryXML() *QueryXML {
	q := NewQueryXML()
	q.FieldExpression(qe.Field, qe.Op, qe.Value)
	return q
}

type QueryXML struct {
	Doc  *etree.Document
	Qxml *etree.Element
}

func (q *QueryXML) ToReader() io.Reader {
	return strings.NewReader(q.String())
}

func (q *QueryXML) String() string {
	tmpl, err := ioutil.ReadFile(path.Join("templates", "query.xml"))

	if err != nil {
		return ""
	}

	out, err := q.Doc.WriteToString()

	if err != nil {
		return ""
	}

	return strings.Replace(string(tmpl), "{queryxml}", out, 1)
}

func (q *QueryXML) Entity(name string) {
	e := q.Qxml.CreateElement("entity")
	e.SetText(name)
}
func (q *QueryXML) FieldExpression(name, op, value string) {
	qry := q.Qxml.CreateElement("query")
	f := qry.CreateElement("field")
	f.SetText(name)
	e := f.CreateElement("expression")
	e.SetText(value)
	e.CreateAttr("op", op)
}

func NewQueryXML() *QueryXML {
	q := &QueryXML{}
	q.Doc = etree.NewDocument()
	q.Qxml = q.Doc.CreateElement("queryxml")

	return q
}
