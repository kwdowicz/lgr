package lgr

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func GenerateLogging(filename string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Skip functions without a receiver (not methods)
		if fn.Recv == nil {
			continue
		}

		fnName := fn.Name.Name

		// Generate the logging statements
		fn.Body.List = append([]ast.Stmt{
			&ast.ExprStmt{X: &ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: ast.NewIdent("log"), Sel: ast.NewIdent("Println")},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"Entering method ` + fnName + `"`}},
			}},
			&ast.DeferStmt{Call: &ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: ast.NewIdent("log"), Sel: ast.NewIdent("Println")},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"Exiting method ` + fnName + `"`}},
			}},
		}, fn.Body.List...)
	}

	// Format the modified AST
	outputFile, err := os.Create("main.go") // Create the output file
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	err = format.Node(outputFile, fset, node) // Write to the output file
	if err != nil {
		log.Fatal(err)
	}

	// No need to use ioutil.WriteFile now
}