package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"

	"goast/types"
)

func main() {
	var (
		path string
		err  error
	)
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		path, err = filepath.Abs(gopath)
		if err != nil {
			panic(err)
		}
		path = filepath.Join(path, "src")
	} else if srcpath := os.Getenv("SRCPATH"); srcpath != "" {
		path, err = filepath.Abs(srcpath)
		if err != nil {
			panic(err)
		}
	} else {
		path = "../"
	}

	a := types.AstTest{}
	typ := reflect.TypeOf(a)

	fset := token.NewFileSet()
	pkgPath := filepath.Join(path, typ.PkgPath())
	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, pkgAst := range pkgs {
		pkgDoc := doc.New(pkgAst, pkgPath, doc.AllDecls)
		matchTypeDoc(typ, pkgDoc)
	}
}

func matchTypeDoc(typ reflect.Type, pkgDoc *doc.Package) {
	for _, typeDoc := range pkgDoc.Types {
		if typ.Name() == typeDoc.Name {
			fmt.Println("========", typ.Name(), typeDoc.Methods)
			for _, m := range typeDoc.Methods {
				fmt.Println("===", m.Name, m.Doc)
			}
		}
	}
}
