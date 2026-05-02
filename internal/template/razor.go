package template

import (
	"fmt"
	"os"
	"path/filepath"

	"cw/internal/file"
)

func GenerateRazor(targetDirect string, fileName string) error {
	nameSpace, err := file.ExtractCSNamespace(targetDirect)
	if err != nil {
		return fmt.Errorf("get namespace: %w", err)
	}

	fileText := fmt.Sprintf(cshtml_Template, nameSpace, fileName)
	err = os.WriteFile(filepath.Join(targetDirect, fileName+".cshtml"), []byte(fileText), 0o644)
	if err != nil {
		return fmt.Errorf("write razor view: %w", err)
	}

	fileText = fmt.Sprintf(cshtml_cs_Template, nameSpace, fileName)
	err = os.WriteFile(filepath.Join(targetDirect, fileName+".cshtml.cs"), []byte(fileText), 0o644)
	if err != nil {
		return fmt.Errorf("write razor code-behind: %w", err)
	}

	return nil
}

var cshtml_cs_Template string = `using Microsoft.AspNetCore.Mvc.RazorPages;
namespace %s
{
    public class %sModel : PageModel
    {
        public void OnGet()
        {
        }
    }
}`

var cshtml_Template string = `@page
@model %s.%sModel
@{
}`
