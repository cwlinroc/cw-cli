package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateEditorConfig(targetDirect string) error {

	file, err := os.Create(filepath.Join(targetDirect, ".editorconfig"))

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}

	defer file.Close()

	file.WriteString(editorConfig_Template)

	return nil
}

var editorConfig_Template string = `root = true

[*]
indent_style = space
indent_size = 4
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true
end_of_line = crlf

[*.md]
trim_trailing_whitespace = false

# C# files
[*.cs]

# Organize usings
dotnet_sort_system_directives_first = true
dotnet_separate_import_directive_groups = false

# Remove unnecessary usings
dotnet_diagnostic.IDE0005.severity = warning

# File header (optional)
# file_header_template = unset

# Code style rules for usings
csharp_using_directive_placement = outside_namespace

# Additional formatting rules
csharp_new_line_before_open_brace = all
csharp_new_line_before_else = true
csharp_new_line_before_catch = true
csharp_new_line_before_finally = true

# Namespace preferences (C# 10+)
csharp_style_namespace_declarations = file_scoped:warning`
