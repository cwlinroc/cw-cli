package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateRazor(targetDirect string, fileName string) error {

	nameSpace, err := cs_namespace(targetDirect)

	if err != nil {
		fmt.Println("Error getting namespace: ", err)
		return err
	}

	//create cshtml file
	{
		file, err := os.Create(filepath.Join(targetDirect, fileName+".cshtml"))

		if err != nil {
			fmt.Println("Error creating file: ", err)
			return err
		}

		defer file.Close()

		fileText := fmt.Sprintf(cshtml_Template, nameSpace, fileName)

		file.WriteString(fileText)
	}

	//create cshtml.cs file
	{
		file, err := os.Create(filepath.Join(targetDirect, fileName+".cshtml.cs"))

		if err != nil {
			fmt.Println("Error creating file: ", err)
			return err
		}

		defer file.Close()

		fileText := fmt.Sprintf(cshtml_cs_Template, nameSpace, fileName)

		file.WriteString(fileText)
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
