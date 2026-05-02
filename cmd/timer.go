package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(timerCmd)
}

var timerCmd = &cobra.Command{
	Use:   "timer <duration>",
	Short: "Countdown timer in the terminal",
	Long: `Display a live countdown timer. Press f to abort.
When time is up, press any key to dismiss.

Duration examples:
  cw timer 30        (bare number = seconds)
  cw timer 23s       cw timer 23sec
  cw timer 13m       cw timer 13min
  cw timer 5h        cw timer 5hr
  cw timer 1h30m     cw timer 1hr30min`,
	Args: cobra.ExactArgs(1),
	RunE: timerCmdRun,
}

//------------------------------------------------------------

var durTokenRe = regexp.MustCompile(`(?i)(\d+)(hr|min|sec|h|m|s)`)

func parseDuration(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty duration")
	}

	// bare integer → seconds
	if n, err := strconv.Atoi(s); err == nil {
		if n <= 0 {
			return 0, fmt.Errorf("duration must be greater than zero")
		}
		return time.Duration(n) * time.Second, nil
	}

	// validate: remove all valid tokens; any remaining non-whitespace is invalid
	remainder := durTokenRe.ReplaceAllString(s, "")
	if strings.TrimSpace(remainder) != "" {
		return 0, fmt.Errorf("invalid duration %q", s)
	}

	matches := durTokenRe.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("invalid duration %q", s)
	}

	var total time.Duration
	for _, m := range matches {
		n, _ := strconv.Atoi(m[1])
		switch strings.ToLower(m[2]) {
		case "h", "hr":
			total += time.Duration(n) * time.Hour
		case "m", "min":
			total += time.Duration(n) * time.Minute
		case "s", "sec":
			total += time.Duration(n) * time.Second
		}
	}
	if total <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}
	return total, nil
}

//------------------------------------------------------------

var doneStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("9")).
	Border(lipgloss.RoundedBorder()).
	Padding(0, 2)

type timerModel struct {
	t timer.Model
}

func (m timerModel) Init() tea.Cmd {
	return m.t.Init()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.t, cmd = m.t.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		m.t, _ = m.t.Update(msg)
		return m, tea.Printf("\a")
	case tea.KeyMsg:
		if m.t.Timedout() {
			return m, tea.Quit
		}
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m timerModel) View() string {
	if m.t.Timedout() {
		return doneStyle.Render("⏰  TIME UP  —  press any key") + "\n"
	}
	return "⏳ " + m.t.View() + "  (press q to quit)\n"
}

//------------------------------------------------------------

func timerCmdRun(cmd *cobra.Command, args []string) error {
	d, err := parseDuration(args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	m := timerModel{t: timer.New(d)}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running timer: %w", err)
	}
	return nil
}
