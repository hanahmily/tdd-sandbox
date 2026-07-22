package tally

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"sort"
	"strings"
	"testing"
)

// TestPublicAPISurfaceIsExactlyTotal enforces R1's exact-surface promise: the
// tally package must export the single identifier Total and nothing else.
//
// Caller compilation alone cannot prove exclusivity — a coder could add another
// exported symbol (e.g. func Sum) to the unprotected tally_impl.go and every
// behavioural test plus the build would stay green. This guard closes that gap
// by parsing the package's non-test source and asserting the exported top-level
// declarations are precisely {Total}.
//
// Unlike the behavioural tests, this is a structural invariant guard: it is
// green from contract time and must remain green through implementation.
func TestPublicAPISurfaceIsExactlyTotal(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", func(fi fs.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatalf("parsing tally package sources: %v", err)
	}

	pkg, ok := pkgs["tally"]
	if !ok {
		t.Fatalf("package %q not found among parsed packages", "tally")
	}

	var exported []string
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if !d.Name.IsExported() {
					continue
				}
				if d.Recv != nil {
					exported = append(exported, "method "+d.Name.Name)
					continue
				}
				exported = append(exported, d.Name.Name)
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						if s.Name.IsExported() {
							exported = append(exported, s.Name.Name)
						}
					case *ast.ValueSpec:
						for _, name := range s.Names {
							if name.IsExported() {
								exported = append(exported, name.Name)
							}
						}
					}
				}
			}
		}
	}

	sort.Strings(exported)
	if len(exported) != 1 || exported[0] != "Total" {
		t.Errorf("exported API surface = %v, want [Total]", exported)
	}
}
