package gengo

import "strings"

type GoImport string

func (i GoImport) String() string {
	return string(i)
}

func (i GoImport) Ident(name string) GoIdent {
	return GoIdent{
		Import: i,
		Name:   name,
	}
}

type GoIdent struct {
	Import GoImport
	Name   string
}

func pkgName(importPath GoImport) string {
	parts := strings.Split(string(importPath), "/")
	if len(parts) == 1 {
		return parts[0]
	}
	return parts[len(parts)-1]
}
