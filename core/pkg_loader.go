package core

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"

	"github.com/samber/lo"
)

type File struct {
	Pkg  *Package  // Package to which this file belongs.
	File *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	// typeName string // Name of the constant type.
	Types []*ast.TypeSpec

	Constants   []*ast.GenDecl
	Funcs       []*ast.FuncDecl
	StructTypes []string // Names of the constant values.
	trimPrefix  string
	lineComment bool
}

func (f *File) GenDecl(node ast.Node) bool {
	if fnode, ok := node.(*ast.FuncDecl); ok {
		f.Funcs = append(f.Funcs, fnode)
		return true
	}

	decl, ok := node.(*ast.GenDecl)

	if !ok {
		return true
	}
	if decl.Tok == token.TYPE {
		T := decl.Specs[0].(*ast.TypeSpec)
		f.Types = append(f.Types, T)

		typeName := T.Name.Name

		f.StructTypes = append(f.StructTypes, typeName)
	} else if decl.Tok == token.CONST {

		f.Constants = append(f.Constants, decl)
	} else {
	}

	return true
}

func FilterT[T any](src []*ast.TypeSpec) []*ast.TypeSpec {
	return lo.Filter(src, func(item *ast.TypeSpec, _ int) bool {
		_, ok := item.Type.(T)
		return ok
	})
}

type Package struct {
	Name  string
	Defs  map[*ast.Ident]types.Object
	Files []*File
}

func (p *Package) GetTypes() []*ast.TypeSpec {
	return lo.Reduce(
		p.Files,
		func(
			acc []*ast.TypeSpec,
			file *File,
			_ int,
		) []*ast.TypeSpec {
			return append(acc, file.Types...)
		},
		[]*ast.TypeSpec{},
	)
}

func (p *Package) GetConstants() []*ast.GenDecl {
	return lo.Reduce(
		p.Files,
		func(
			acc []*ast.GenDecl,
			file *File,
			_ int,
		) []*ast.GenDecl {
			return append(acc, file.Constants...)
		},
		[]*ast.GenDecl{},
	)
}

func (p *Package) GetFuncs() []*ast.FuncDecl {
	return lo.Reduce(
		p.Files,
		func(
			acc []*ast.FuncDecl,
			file *File,
			_ int,
		) []*ast.FuncDecl {
			return append(acc, file.Funcs...)
		},
		[]*ast.FuncDecl{},
	)
}

//#endregion

// IsDirectory reports whether the named file is a directory.
func IsDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
