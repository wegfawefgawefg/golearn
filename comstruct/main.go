package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	. "trying-mo/eu"
)

func main() {
	src_path := "comstruct/in.go"
	src, err := os.ReadFile(src_path)
	Elf(err)

	// output_path := "comstruct/out.go"

	// Create a new scanner for the provided src
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Walk through the AST and print struct fields and comments
	ast.Inspect(file, func(n ast.Node) bool {
		// Check if the node is a struct type
		t, ok := n.(*ast.TypeSpec)
		if ok {
			s, ok := t.Type.(*ast.StructType)
			if ok {
				println("Struct:", t.Name.Name)
				for _, field := range s.Fields.List {
					if len(field.Names) > 0 {
						for _, name := range field.Names {
							print("Field:", name.Name, " ")
						}
					}
					if field.Comment != nil {
						println("Comment:", field.Comment.Text())
					} else if field.Doc != nil {
						println("Comment:", field.Doc.Text())
					} else {
						println("No comment")
					}
				}
			}
		}
		return true
	})
}
