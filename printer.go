package gengo

import (
	"bytes"
	"fmt"
)

type Printer struct {
	pkgPath GoImport
	pkg     string
	buffer  *bytes.Buffer

	importsByPath  map[GoImport]string
	importsByAlias map[string]GoImport
}

func New(pkgPath GoImport) *Printer {
	return &Printer{
		pkgPath: pkgPath,
		pkg:     pkgName(pkgPath),
		buffer:  bytes.NewBuffer(nil),

		importsByPath:  map[GoImport]string{},
		importsByAlias: map[string]GoImport{},
	}
}

func (p *Printer) Write(b []byte) (n int, err error) {
	return p.buffer.Write(b)
}

func (p *Printer) P(f string, args ...any) {
	a := make([]any, len(args))
	for i := range args {
		a[i] = args[i]
		switch x := args[i].(type) {
		case GoIdent:
			a[i] = p.Ident(x)
		}
	}

	p.buffer.WriteString(fmt.Sprintf(f+"\n", a...))
}

func (p *Printer) Ident(name GoIdent) string {
	if name.Import == "" || name.Import == p.pkgPath {
		return name.Name
	}
	return fmt.Sprintf("%s.%s", p.getImportAlias(name.Import), name.Name)
}

func (p *Printer) getImportAlias(path GoImport) string {
	if alias, ok := p.importsByPath[path]; ok {
		return alias
	}
	alias := pkgName(path)
	index := 0
	for {
		if _, ok := p.importsByAlias[alias]; ok { // exists
			alias = fmt.Sprintf("%s%d", pkgName(path), index)
			index++
		} else {
			break
		}
	}
	p.importsByPath[path] = alias
	p.importsByAlias[alias] = path

	return alias
}

func (p *Printer) Bytes() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("// Code generated. DO NOT EDIT.\n")
	buf.WriteString("\n")
	buf.WriteString("package " + p.pkg + "\n")
	buf.WriteString("\n")
	buf.WriteString("import (\n")
	for alias, path := range p.importsByAlias {
		fmt.Fprintf(buf, "\t%s %q\n", alias, path)
	}
	buf.WriteString(")\n")
	buf.Write(p.buffer.Bytes())

	output, err := FormatSource(buf.String())
	if err != nil {
		return nil, err
	}
	return []byte(output), nil
}
