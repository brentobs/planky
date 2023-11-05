package main

import (
	"fmt"
	"os"
	// "strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

const (
	padding  = 2
	maxWidth = 80
)
var (
	term          = termenv.EnvColorProfile()
	keyword       = makeFgStyle("211")
	subtle        = makeFgStyle("241")
	dot           = colorFg(" • ", "236")

)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var languages = []string{"Python", "NodeJS", "Java", "Go"}
var pythonTypes = []string{"FastAPI", "Django", "CLI"}
var nodeTypes = []string{"Express", "NestJS", "CLI"}
var javaTypes = []string{"Spring", "Bla", "CLI"}
var goTypes = []string{"a", "b", "CLI"}
var languageTypes = map[string][]string{
    "Python": {"FastAPI", "Django", "CLI"},
    "NodeJS": {"Express", "NestJS", "CLI"},
    "Java":   {"Spring", "Bla", "CLI"},
    "Go":     {"a", "b", "CLI"},
}

type tickMsg time.Time

// type model struct {
// 	LangChoice   int
// 	LangChosen   bool
// 	TypeChoice   int
// 	TypeChosen   bool
// 	Quitting bool
// }

func main() {
	initialModel := model{
    LangChoice: 0,
		LangChosen: false,
		TypeChoice: 0,
		TypeChosen: false,
    Quitting: false,
	}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}


// func (m model) Init() tea.Cmd {
// 	return nil
// }

// // Main update function.
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	// Make sure these keys always quit
// 	if msg, ok := msg.(tea.KeyMsg); ok {
// 		k := msg.String()
// 		if k == "q" || k == "esc" || k == "ctrl+c" {
// 			m.Quitting = true
// 			return m, tea.Quit
// 		}
// 	}
//   if m.LangChosen && m.TypeChosen {
// 		return updateCreation(msg, m)
// 	} else if m.LangChosen {
// 		return updateLanguageType(msg, m)
// 	} else {
// 		return updateLanguage(msg, m)
// 	}
// }

// // The main view, which just calls the appropriate sub-view
// func (m model) View() string {
// 	var s string
// 	if m.Quitting {
// 		return "\n  See you later!\n\n"
// 	}
// if m.LangChosen && m.TypeChosen {
// 		s = creationView(m)
// 	} else if m.LangChosen {
// 		s = languageTypeView(m)
// 	} else {
// 		s = languageView(m)
// 	}
// 	return indent.String("\n"+s+"\n\n", 2)
// }

// Sub-update functions

// Update loop for the first view where you're choosing a task.
// func updateLanguage(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "j", "down":
// 			m.LangChoice++
// 			if m.LangChoice > 3 {
// 				m.LangChoice = 3
// 			}
// 		case "k", "up":
// 			m.LangChoice--
// 			if m.LangChoice < 0 {
// 				m.LangChoice = 0
// 			}
// 		case "enter":
// 			m.LangChosen = true
// 			return m, tickCmd()
// 		}
// 	}
// 	return m, nil
// }

// // Update loop for the second view after a choice has been made
// func updateLanguageType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "j", "down":
// 			m.TypeChoice++
// 			if m.TypeChoice > 3 {
// 				m.TypeChoice = 3
// 			}
// 		case "k", "up":
// 			m.TypeChoice--
// 			if m.TypeChoice < 0 {
// 				m.TypeChoice = 0
// 			}
// 		case "enter":
// 			m.TypeChosen = true
// 			return m, tickCmd()
// 		}
// 	}
// 	return m, nil
// }
// // Update loop for the second view after a choice has been made
// func updateCreation(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "j", "down":
// 			m.TypeChoice++
// 			if m.TypeChoice > 3 {
// 				m.TypeChoice = 3
// 			}
// 		case "k", "up":
// 			m.TypeChoice--
// 			if m.TypeChoice < 0 {
// 				m.TypeChoice = 0
// 			}
// 		case "enter":
// 			m.TypeChosen = true
// 			return m, tickCmd()
// 		}
// 	}
// 	return m, nil
// }

// Sub-views

// The first view, where you're choosing a task
// func languageView(m model) string {
// 	c := m.LangChoice

// 	tpl := "Which Language?\n\n"
// 	tpl += "%s\n\n"
// 	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

//   var choices string
//   for i, lang := range languages{
//     choices += checkbox(lang, c == i) + "\n"
//   }

// 	return fmt.Sprintf(tpl, choices,  "")
// }

// The second view, after a task has been chosen
// func languageTypeView(m model) string {

//   lang := m.LangChoice
// 	c := m.TypeChoice

// 	tpl := "Which Language?\n\n"
// 	tpl += "%s\n\n"
// 	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

//   selectedLanguage := languages[lang] // Get the selected language
// 	selectedTypes := languageTypes[selectedLanguage] // Get the corresponding types

//   var choices string
//   for i, t := range selectedTypes{
//     choices += checkbox(t, c == i) + "\n"
//   }

// 	return fmt.Sprintf(tpl, choices,  "")
// }
// func creationView(m model) string {

//   lang := m.LangChoice
// 	c := m.TypeChoice

// 	tpl := "Which Language?\n\n"
// 	tpl += "%s\n\n"
// 	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

//   selectedLanguage := languages[lang] // Get the selected language
// 	selectedTypes := languageTypes[selectedLanguage] // Get the corresponding types

//   var choices string
//   for i, t := range selectedTypes{
//     choices += checkbox(t, c == i) + "\n"
//   }

// 	return fmt.Sprintf(tpl, choices,  "")
// }

func checkbox(label string, checked bool) string {
	if checked {
		return colorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

// Utils

// Color a string's foreground with the given value.
func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
