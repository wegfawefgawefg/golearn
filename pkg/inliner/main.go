package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	srcPath := "inliner/in.go" // Path to the source file to process
	dstPath := "output.go"     // Path where the transformed file will be saved

	src, err := os.ReadFile(srcPath)
	if err != nil {
		log.Fatalf("Error reading source file: %s", err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error parsing source file: %s", err)
	}

	inlineFunctions(file) // Apply the inlining transformation

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		log.Fatalf("Error formatting transformed AST: %s", err)
	}

	if err := os.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Error writing transformed file: %s", err)
	}
	log.Println("Inlining complete. Output saved to", dstPath)
}

// inlineFunctions scans the AST for functions with //@inline and replaces calls to them.
func inlineFunctions(file *ast.File) {
	inlineCandidates := make(map[string]*ast.FuncDecl)

	// First pass: Identify functions with //@inline comments.
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			for _, comment := range fn.Doc.List {
				if strings.Contains(comment.Text, "//@inline") {
					inlineCandidates[fn.Name.Name] = fn
					break
				}
			}
		}
		return true
	})

	// Second pass: Replace calls to the identified functions with their bodies.
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			if ident, ok := node.Fun.(*ast.Ident); ok {
				if fn, found := inlineCandidates[ident.Name]; found {
					// Implement the specific logic to replace the call with the function body
					// This example simplifies the process and does not handle parameters and return values
					if len(fn.Body.List) == 1 {
						stmt, ok := fn.Body.List[0].(*ast.ReturnStmt)
						if !ok {
							break
						}
						expr := stmt.Results[0]
						clonedExpr := cloneAstNode(expr).(ast.Expr)
						node.Fun = clonedExpr
					}
				}
			}
		}
		return true
	})
}

// cloneAstNode is a helper to clone AST nodes, needs to handle all necessary node types
func cloneAstNode(n ast.Node) ast.Node {
	if n == nil {
		return nil
	}
	switch src := n.(type) {
	case *ast.BasicLit:
		cloned := *src
		return &cloned
	case *ast.Ident:
		cloned := *src
		return &cloned
		// Add more cases as needed for other types of nodes
	}
	panic("cloneAstNode: unexpected node type")
}
