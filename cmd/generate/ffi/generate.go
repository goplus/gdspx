// Package gdextensionwrapper generates C code to wrap all of the gdextension
// methods to call functions on the gdextension_api_structs to work
// around the cgo C function pointer limitation.
package ffi

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"godot-ext/gdspx/cmd/gdextensionparser/clang"

	"github.com/iancoleman/strcase"
)

var (
	//go:embed ffi_wrapper.h.tmpl
	ffiWrapperHeaderFileText string

	//go:embed ffi_wrapper.go.tmpl
	ffiWrapperGoFileText string

	//go:embed ffi.go.tmpl
	ffiFileText string

	//go:embed manager_wrapper.go.tmpl
	wrapManagerGoFileText string
)
var (
	relDir = "../internal/ffi"
)

func Generate(projectPath string, ast clang.CHeaderFileAST) {
	err := GenerateGDExtensionWrapperHeaderFile(projectPath, ast)
	if err != nil {
		panic(err)
	}
	err = GenerateGDExtensionWrapperGoFile(projectPath, ast)
	if err != nil {
		panic(err)
	}
	err = GenerateGDExtensionInterfaceGoFile(projectPath, ast)
	if err != nil {
		panic(err)
	}
	err = GenerateManagerWrapperGoFile(projectPath, ast)
	if err != nil {
		panic(err)
	}

}

func GenerateGDExtensionWrapperHeaderFile(projectPath string, ast clang.CHeaderFileAST) error {
	tmpl, err := template.New("ffi_wrapper.gen.h").
		Funcs(template.FuncMap{
			"snakeCase": strcase.ToSnake,
		}).
		Parse(ffiWrapperHeaderFileText)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, ast)
	if err != nil {
		return err
	}

	filename := filepath.Join(projectPath, relDir, "ffi_wrapper.gen.h")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func GenerateGDExtensionWrapperGoFile(projectPath string, ast clang.CHeaderFileAST) error {
	funcs := template.FuncMap{
		"gdiVariableName":    gdiVariableName,
		"snakeCase":          strcase.ToSnake,
		"camelCase":          strcase.ToCamel,
		"goReturnType":       goReturnType,
		"goArgumentType":     goArgumentType,
		"goEnumValue":        goEnumValue,
		"add":                add,
		"cgoCastArgument":    cgoCastArgument,
		"cgoCastReturnType":  cgoCastReturnType,
		"cgoCleanUpArgument": cgoCleanUpArgument,
		"trimPrefix":         trimPrefix,
	}

	tmpl, err := template.New("ffi_wrapper.gen.go").
		Funcs(funcs).
		Parse(ffiWrapperGoFileText)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, ast)
	if err != nil {
		return err
	}

	headerFileName := filepath.Join(projectPath, relDir, "ffi_wrapper.gen.go")
	f, err := os.Create(headerFileName)
	f.Write(b.Bytes())
	f.Close()
	return err
}

func GenerateGDExtensionInterfaceGoFile(projectPath string, ast clang.CHeaderFileAST) error {
	funcs := template.FuncMap{
		"gdiVariableName":     gdiVariableName,
		"snakeCase":           strcase.ToSnake,
		"camelCase":           strcase.ToCamel,
		"goReturnType":        goReturnType,
		"goArgumentType":      goArgumentType,
		"goEnumValue":         goEnumValue,
		"add":                 add,
		"cgoCastArgument":     cgoCastArgument,
		"cgoCastReturnType":   cgoCastReturnType,
		"cgoCleanUpArgument":  cgoCleanUpArgument,
		"trimPrefix":          trimPrefix,
		"loadProcAddressName": loadProcAddressName,
	}

	tmpl, err := template.New("ffi.gen.go").
		Funcs(funcs).
		Parse(ffiFileText)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, ast)
	if err != nil {
		return err
	}

	headerFileName := filepath.Join(projectPath, relDir, "ffi.gen.go")
	f, err := os.Create(headerFileName)
	f.Write(b.Bytes())
	f.Close()
	return err
}

func GenerateManagerWrapperGoFile(projectPath string, ast clang.CHeaderFileAST) error {
	funcs := template.FuncMap{
		"gdiVariableName":    gdiVariableName,
		"snakeCase":          strcase.ToSnake,
		"camelCase":          strcase.ToCamel,
		"goReturnType":       goReturnType,
		"goArgumentType":     goArgumentType,
		"goEnumValue":        goEnumValue,
		"add":                add,
		"cgoCastArgument":    cgoCastArgument,
		"cgoCastReturnType":  cgoCastReturnType,
		"cgoCleanUpArgument": cgoCleanUpArgument,
		"trimPrefix":         trimPrefix,
		"isManagerMethod":    isManagerMethod,
		"getManagerFuncName": getManagerFuncName,
		"getManagerFuncBody": getManagerFuncBody,
	}

	tmpl, err := template.New("manager_wrapper.gen.go").
		Funcs(funcs).
		Parse(wrapManagerGoFileText)
	if err != nil {
		return err
	}

	items := []string{}
	for _, item := range ast.CollectGDExtensionInterfaceFunctions() {
		items = append(items, item.Name)
	}

	managerSet = make(map[string]bool)
	managers := []string{}
	for _, str := range items {
		managerSet[getManagerName(str)] = true
	}
	delete(managerSet, "")
	delete(managerSet, "string")
	delete(managerSet, "variant")
	delete(managerSet, "global")
	for item := range managerSet {
		managers = append(managers, item)
	}
	sort.Strings(managers)
	cppType2Go = map[string]string{
		"GdInt":    "int64",
		"GdFloat":  "float64",
		"GdObj":    "Object",
		"GdVec2":   "Vec2",
		"GdVec3":   "Vec3",
		"GdVec4":   "Vec4",
		"GdRect2":  "Rect2",
		"GdString": "string",
		"GdBool":   "bool",
		"GdColor":  "Color",
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, ManagerData{Ast: ast, Mangers: managers})
	if err != nil {
		return err
	}

	headerFileName := filepath.Join(projectPath, relDir, "../wrap/manager_wrapper.gen.go")
	f, err := os.Create(headerFileName)
	f.Write(b.Bytes())
	f.Close()
	return err
}

var (
	managerSet = map[string]bool{}
	cppType2Go = map[string]string{}
)

type ManagerData struct {
	Ast     clang.CHeaderFileAST
	Mangers []string
}

func getManagerName(str string) string {
	prefix := "GDExtensionSpx"
	str = str[len(prefix):]
	chs := []rune{}
	chs = append(chs, rune(str[0]), rune(str[1]))
	for _, ch := range str[2:] {
		if unicode.IsUpper(rune(ch)) {
			break
		}
		chs = append(chs, rune(ch))
	}
	result := strings.ToLower(string(chs))
	return result
}

func isManagerMethod(function *clang.TypedefFunction) bool {
	return managerSet[getManagerName(function.Name)]
}
func getFuncParamTypeString(typeName string) string {
	return cppType2Go[typeName]
}
func getManagerFuncName(function *clang.TypedefFunction) string {
	prefix := "GDExtensionSpx"
	sb := strings.Builder{}
	mgrName := getManagerName(function.Name)
	funcName := function.Name[len(prefix)+len(mgrName):]
	sb.WriteString("(")
	sb.WriteString("pself *" + mgrName)
	sb.WriteString("Mgr) ")
	sb.WriteString(funcName)
	sb.WriteString("(")
	count := len(function.Arguments)
	for i, arg := range function.Arguments {
		sb.WriteString(arg.Name)
		sb.WriteString(" ")
		typeName := getFuncParamTypeString(arg.Type.Primative.Name)
		sb.WriteString(typeName)
		if i != count-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")

	if function.ReturnType.Name != "void" {
		typeName := getFuncParamTypeString(function.ReturnType.Name)
		sb.WriteString(" " + typeName + " ")
	}
	return sb.String()
}
func getManagerFuncBody(function *clang.TypedefFunction) string {
	sb := strings.Builder{}
	prefixTab := "\t"
	params := []string{}
	// convert arguments
	for i, arg := range function.Arguments {
		sb.WriteString(prefixTab)
		typeName := arg.Type.Primative.Name
		argName := "arg" + strconv.Itoa(i)
		switch typeName {
		case "GdString":
			sb.WriteString(argName + "Str := ")
			sb.WriteString("NewCString(")
			sb.WriteString(arg.Name)
			sb.WriteString(")")
			sb.WriteString("\n" + prefixTab)
			sb.WriteString(argName + " := " + argName + "Str.ToGdString() ")

		default:
			sb.WriteString(argName + " := ")
			sb.WriteString("To" + typeName)
			sb.WriteString("(")
			sb.WriteString(arg.Name)
			sb.WriteString(")")
		}
		sb.WriteString("\n")
		params = append(params, argName)
	}

	// call the function
	sb.WriteString(prefixTab)
	if function.ReturnType.Name != "void" {
		sb.WriteString("retValue := ")
	}

	funcName := "Call" + trimPrefix(function.Name, "GDExtensionSpx")
	sb.WriteString(funcName)
	sb.WriteString("(")
	for i, param := range params {
		sb.WriteString(param)
		if i != len(params)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	if function.ReturnType.Name != "void" {
		sb.WriteString("\n" + prefixTab)
		sb.WriteString("return ")
		typeName := getFuncParamTypeString(function.ReturnType.Name)
		sb.WriteString("To" + strcase.ToCamel(typeName) + "(retValue)")
	}
	return sb.String()
}
