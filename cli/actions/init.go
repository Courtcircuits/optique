package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const URL = "https://github.com/Courtcircuits/optique/cli"

type Initialization struct {
	URL     string
	Name    string
	Version string
}

var defaultInitialization = &Initialization{
	Name:    "optique",
	URL:     "https://github.com/baptistebronsin/javoue",
	Version: "latest",
}

func NewInitialization(name string) Initialization {
	defaultInitialization.Name = name
	_, err := tea.NewProgram(initialModel(name)).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return *defaultInitialization
}

func Initialize(generation Initialization) {
	err := createProjectFolder(generation.Name)
	if err != nil {
		fmt.Println("Error creating project folder:", err)
		os.Exit(1)
	}
	err = cloneTemplate("https://github.com/Courtcircuits/optique", generation.Name)
	if err != nil {
		fmt.Println("Error cloning template:", err)
		os.Exit(1)
	}
	err = goBack()
	if err != nil {
		fmt.Println("Error going back:", err)
		os.Exit(1)
	}
}

func createProjectFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return err
	}
	return nil
}

func goBack() error {
	return os.Chdir("..")
}

func cloneTemplate(url string, name string) error {
	cmd := exec.Command("git", "clone", url, name)
	if err := cmd.Run(); err != nil {
		return err
	}

	// go to project folder
	err := os.Chdir(name)
	if err != nil {
		return err
	}

	current_dir, err := os.Getwd()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(current_dir)

	if err != nil {
		return err
	}

	folders_to_delete := []string{}
	files_to_delete := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() != "template" {
				folders_to_delete = append(folders_to_delete, entry.Name())
			}
		} else {
			files_to_delete = append(files_to_delete, entry.Name())
		}
	}

	for _, entry := range folders_to_delete {
		err = os.RemoveAll(entry)
		if err != nil {
			return err
		}
	}
	for _, entry := range files_to_delete {
		err = os.Remove(entry)
		if err != nil {
			return err
		}
	}

	// go to template folder
	err = os.Chdir("template")
	if err != nil {
		return err
	}

	entries, err = os.ReadDir(".")
	for _, entry := range entries {
		val, err := exec.Command("mv", entry.Name() , current_dir).CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(val))
	}
	// move all to parent folder
	err = goBack()
	if err != nil {
		return err
	}

	// remove template folder
	err = os.RemoveAll("template")
	return nil
}

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Generate ( •_•)>⌐■-■ ]")
	blurredButton = fmt.Sprintf("%s", blurredStyle.Render("[ Generate ( •_•)>⌐■-■ ]"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func initialModel(name string) model {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32
		switch i {
		case 0:
			t.Placeholder = "Module URL: example (github.com/you/" + name + ")"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.CharLimit = 255
		case 1:
			t.Placeholder = "Optique template version"
			t.PromptStyle = blurredStyle
			t.TextStyle = blurredStyle
			t.CharLimit = 10
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				defaultInitialization.URL = m.inputs[0].Value()
				defaultInitialization.Version = m.inputs[1].Value()

				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
