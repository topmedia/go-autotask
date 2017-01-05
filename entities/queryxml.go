package entities

import (
	"io"
	"strings"

	"github.com/beevik/etree"
)

const XMLtemplate = `<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:tns="http://autotask.net/ATWS/v1_5/" xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
  <env:Body>
    <tns:query xmlns="http://autotask.net/ATWS/v1_5/">
      <sXML><![CDATA[
        {queryxml}
]]></sXML>
    </tns:query>
  </env:Body>
</env:Envelope>
`

type QueryCondition interface {
	ToQueryXML() *QueryXML
}

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

func (q *QueryXML) ToQueryXML() *QueryXML {
	return q
}

func (q *QueryXML) String() string {
	out, err := q.Doc.WriteToString()

	if err != nil {
		return ""
	}

	return strings.Replace(XMLtemplate, "{queryxml}", out, 1)
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
