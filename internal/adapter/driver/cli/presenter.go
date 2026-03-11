package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
	"github.com/fatih/color"
	terminal "golang.org/x/term"
)

// ColorPresenter é o adaptador driver que implementa port.Presenter
// com saída colorida no terminal.
type ColorPresenter struct {
	red   *color.Color
	green *color.Color
	warn  *color.Color
	cyan  *color.Color
	blue  *color.Color
	white *color.Color
}

// NewColorPresenter cria o presenter com paleta de cores.
func NewColorPresenter() *ColorPresenter {
	return &ColorPresenter{
		red:   color.New(color.FgRed, color.Bold),
		green: color.New(color.FgGreen),
		warn:  color.New(color.FgYellow, color.Bold),
		cyan:  color.New(color.FgCyan, color.Bold),
		blue:  color.New(color.FgBlue, color.Bold),
		white: color.New(color.FgWhite, color.Bold),
	}
}

func (p *ColorPresenter) RenderMatch(secretName string) {
	p.green.Printf("  ✓ ")
	fmt.Printf("Found in secret: ")
	p.white.Printf("%s\n", secretName)
}

func (p *ColorPresenter) RenderTermStart(term string) {
	fmt.Println()
	p.cyan.Printf("🔍 Searching for term: ")
	p.blue.Printf("%s\n", term)
	fmt.Println()
}

func (p *ColorPresenter) RenderTermSummary(result *domain.SearchResult) {
	fmt.Println()
	if !result.HasMatches() {
		p.warn.Println("  ⚠ No secrets found")
	} else {
		p.cyan.Printf("  📊 Total: ")
		fmt.Printf("%d secret(s) found\n", result.Count())
	}
}

func (p *ColorPresenter) RenderSeparator() {
	fmt.Println()
	color.New(color.FgCyan).Println(strings.Repeat("─", p.terminalWidth()))
}

func (p *ColorPresenter) RenderFinalSummary(totalMatches int) {
	fmt.Println()
	fmt.Println()
	width := min(p.terminalWidth(), 63)
	sep := strings.Repeat("═", width)
	color.New(color.FgCyan, color.Bold).Println(sep)
	p.cyan.Printf("  🎯 Search completed! Total matches: ")
	color.New(color.FgGreen, color.Bold).Printf("%d\n", totalMatches)
	color.New(color.FgCyan, color.Bold).Println(sep)
	fmt.Println()
}

func (p *ColorPresenter) RenderWarning(message string) {
	fmt.Println()
	p.warn.Printf("⚠️  Warning: %s\n", message)
	p.cyan.Println("   → Using system environment variables instead")
	fmt.Println()
}

func (p *ColorPresenter) RenderError(message string, hints []string) {
	p.red.Printf("\n❌ Error: %s\n", message)
	if len(hints) > 0 {
		fmt.Println()
		for _, hint := range hints {
			switch {
			case strings.HasPrefix(hint, "•"):
				p.cyan.Printf("   %s\n", hint)
			case strings.HasPrefix(hint, "💡"):
				fmt.Printf("\n   %s\n", hint)
			default:
				fmt.Printf("   %s\n", hint)
			}
		}
	}
	fmt.Println()
}

func (p *ColorPresenter) terminalWidth() int {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return 80
	}
	return width
}
