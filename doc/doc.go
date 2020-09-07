package doc

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Function struct {
	f interface{}
}

// Info string for a function
func Info(f interface{}) string {
	return NewFunction(f).Doc()
}

// NewFunction creates the Function from a function arg
func NewFunction(f interface{}) *Function {
	return &Function{f: f}
}

// Get the name and path of a func
func (fn *Function) Path() string {
	return runtime.FuncForPC(reflect.ValueOf(fn.f).Pointer()).Name()
}

// Get the name of a func (with package path)
func (fn *Function) Name() string {
	splitFuncName := strings.Split(fn.Path(), ".")
	return splitFuncName[len(splitFuncName)-1]
}

// Get description of a func
func (fn *Function) Doc() string {
	fileName, _ := runtime.FuncForPC(reflect.ValueOf(fn.f).Pointer()).FileLine(0)
	funcName := fn.Name()
	fset := token.NewFileSet()

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")
	myDoc := doc.New(pkg, importPath, doc.AllDecls)
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return theFunc.Doc
		}
	}
	return ""
}
