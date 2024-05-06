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
	srcPath := "scramble/annotation.go"
	src, err := os.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over all declarations to find structs with @scramble
	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok || genDecl.Doc == nil {
				continue
			}

			// Check if @scramble is in the comments
			for _, comment := range genDecl.Doc.List {
				if strings.Contains(comment.Text, "@scramble") {
					// Apply scrambling here, for example, by renaming fields
					for _, field := range structType.Fields.List {
						for _, name := range field.Names {
							name.Name = "Scrambled" + name.Name // Simple example of scrambling
						}
					}
				}
			}
		}
		return true
	})

	// Generate the new Go code from the modified AST
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		log.Fatal(err)
	}
	os.WriteFile("scrambled_output.go", buf.Bytes(), 0644)
}
