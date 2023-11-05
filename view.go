package main

import (
	"fmt"
	"github.com/muesli/reflow/indent"
)

// var term = termenv.EnvColorProfile()

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if m.LangChosen && m.TypeChosen {
		s = creationView(m)
	} else if m.LangChosen {
		s = languageTypeView(m)
	} else {
		s = languageView(m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

// Sub-view functions
func languageView(m model) string {
	c := m.LangChoice

	tpl := "Which Language?\n\n"
	tpl += "%s\n\n"
	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

  var choices string
  for i, lang := range languages{
    choices += checkbox(lang, c == i) + "\n"
  }

	return fmt.Sprintf(tpl, choices,  "")
}
func languageTypeView(m model) string {

  lang := m.LangChoice
	c := m.TypeChoice

	tpl := "Which Language?\n\n"
	tpl += "%s\n\n"
	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

  selectedLanguage := languages[lang] // Get the selected language
	selectedTypes := languageTypes[selectedLanguage] // Get the corresponding types

  var choices string
  for i, t := range selectedTypes{
    choices += checkbox(t, c == i) + "\n"
  }

	return fmt.Sprintf(tpl, choices,  "")
}
func creationView(m model) string {

  lang := m.LangChoice
	c := m.TypeChoice

	tpl := "Which Language?\n\n"
	tpl += "%s\n\n"
	tpl += "j/k, up/down: select" + " - " + "enter: choose" + " - " + "q, esc: quit"

  selectedLanguage := languages[lang] // Get the selected language
	selectedTypes := languageTypes[selectedLanguage] // Get the corresponding types

  var choices string
  for i, t := range selectedTypes{
    choices += checkbox(t, c == i) + "\n"
  }

	return fmt.Sprintf(tpl, choices,  "")
}
