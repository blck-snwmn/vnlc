package vnlc

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is analysis.Analyzer that count variable name length
var Analyzer = &analysis.Analyzer{
	Name:     "vnlc",
	Doc:      "count variable name length",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Ref golang.org/x/tools/go/analysis/passes/loopclosure/loopclosure.go
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.AssignStmt)(nil),
	}
	count := 0
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			if n.Tok != token.DEFINE {
				return
			}
			for _, e := range n.Lhs {
				ast.Inspect(e, func(nn ast.Node) bool {
					id, ok := nn.(*ast.Ident)
					if !ok || id.Obj == nil {
						return true
					}
					switch id.Obj.Kind {
					case ast.Var, ast.Con:
						pass.Reportf(id.Pos(), "len(%s) = %d (%+v) in ast.AssignStmt", id.Name, len(id.Name), n)
						count += len(id.Name)
					}
					return true
				})
			}
		case *ast.GenDecl:
			if n.Tok != token.VAR && n.Tok != token.CONST {
				return
			}
			ast.Inspect(n, func(n ast.Node) bool {
				v, ok := n.(*ast.ValueSpec)
				if !ok {
					return true
				}
				for _, n := range v.Names {
					ast.Inspect(n, func(n ast.Node) bool {
						id, ok := n.(*ast.Ident)
						if !ok || id.Obj == nil {
							return true
						}
						switch id.Obj.Kind {
						case ast.Var, ast.Con:
							pass.Reportf(id.Pos(), "len(%s) = %d (%v) in ast.GenDecl", id.Name, len(id.Name), id.Obj)
							count += len(id.Name)
						}
						return true
					})
				}
				// Namesの要素以外見る必要なし
				return false
			})
		}
	})
	pass.Reportf(token.NoPos, "total length is %d", count)
	return nil, nil
}
