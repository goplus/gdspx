// Package gdextensionwrapper generates C code to wrap all of the gdextension
// methods to call functions on the gdextension_api_structs to work
// around the Cgo C function pointer limitation.
package cpp

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"text/template"

	"godot-ext/gdspx/cmd/gdextensionparser/clang"
	 . "godot-ext/gdspx/cmd/generate/common"

	"github.com/iancoleman/strcase"
)

var (
	//go:embed gdextension_spx_ext.cpp.tmpl
	gdSpxExtCpp string

	//go:embed gdextension_spx_ext.h.tmpl
	gdSpxExtH string
)
func Generate(projectPath string, ast clang.CHeaderFileAST) {
	err := generateGdCppFile(projectPath, ast,"gdextension_spx_ext.cpp")
	if err != nil {
		panic(err)
	}
	err = generateGdCppFile(projectPath, ast,"gdextension_spx_ext.h.h")
	if err != nil {
		panic(err)
	}
	
}


func generateGdCppFile(projectPath string, ast clang.CHeaderFileAST, fileName string) error {
	funcs := template.FuncMap{
		"gdiVariableName":    GdiVariableName,
		"snakeCase":          strcase.ToSnake,
		"camelCase":          strcase.ToCamel,
		"goReturnType":       GoReturnType,
		"goArgumentType":     GoArgumentType,
		"goEnumValue":        GoEnumValue,
		"add":                Add,
		"cgoCastArgument":    CgoCastArgument,
		"cgoCastReturnType":  CgoCastReturnType,
		"cgoCleanUpArgument": CgoCleanUpArgument,
		"trimPrefix":         TrimPrefix,
	}

	tmpl, err := template.New(fileName).
		Funcs(funcs).
		Parse(gdSpxExtH)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, ast)
	if err != nil {
		return err
	}

	headerFileName := filepath.Join(projectPath, RelDir, fileName)
	f, err := os.Create(headerFileName)
	f.Write(b.Bytes())
	f.Close()
	return err
}
