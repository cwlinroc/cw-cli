package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateClaude(targetDirect string) error {
	err := os.WriteFile(filepath.Join(targetDirect, "CLAUDE.md"), []byte(claude_Template), 0o644)
	if err != nil {
		return fmt.Errorf("write CLAUDE.md: %w", err)
	}

	return nil
}

var claude_Template string = `# CLAUDE.md

@AGENTS.md

> Prefer editing ` + "`" + `AGENTS.md` + "`" + ` for root-level workspace guidance. Keep detailed component instructions in the nearest subproject instruction file.
`
