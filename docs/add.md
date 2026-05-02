# add

Interactively scaffold a file. Prompts you to pick a target directory with an interactive directory picker, then asks for a filename (where applicable), and writes the generated file(s) into that directory.

Use [`touch`](touch.md) instead when you want to skip the directory prompt and create the file in the current working directory.

## Subcommands

### `add cs [name]`

Creates a C# class file (`<Name>.cs`).

The namespace is derived automatically by walking up the directory tree until a `.csproj` file is found, then joining the intermediate directory names.

```sh
cw add cs
# prompts: pick directory → pick name → writes MyClass.cs
```

**Output:** `<Name>.cs`

```csharp
namespace My.Project.SubDir;

public class Name
{

}
```

---

### `add razor [name]`

Creates a Razor Page pair.

```sh
cw add razor Index
# prompts: pick directory → writes Index.cshtml + Index.cshtml.cs
```

**Output:** `<Name>.cshtml` + `<Name>.cshtml.cs`

```html
@page @model My.Project.Pages.IndexModel @{ }
```

```csharp
using Microsoft.AspNetCore.Mvc.RazorPages;
namespace My.Project.Pages
{
    public class IndexModel : PageModel
    {
        public void OnGet()
        {
        }
    }
}
```

---

### `add code [name]`

Creates a VS Code workspace file.

```sh
cw add code MyWorkspace
# prompts: pick directory → writes MyWorkspace.code-workspace
```

**Output:** `<Name>.code-workspace`

```json
{
    "folders": [{ "path": "." }],
    "settings": {}
}
```

---

### `add page`

Creates a SvelteKit page file. The directory name is used as the page title.

```sh
cw add page
# prompts: pick directory → writes +page.svelte
```

**Output:** `+page.svelte`

```html
<svelte:head>
    <title>dirname</title>
    <meta name="description" content="dirname" />
</svelte:head>
```

---

### `add editorconfig`

Creates a `.editorconfig` with opinionated defaults for C#, Markdown, and general source files.

```sh
cw add editorconfig
# prompts: pick directory → writes .editorconfig
```

**Highlights:**

- 4-space indentation, UTF-8, trim trailing whitespace
- CRLF line endings globally; Markdown preserves trailing whitespace
- C# file-scoped namespaces, system directives sorted first, unused-using warnings

