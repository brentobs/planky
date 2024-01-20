package main

import (

	"github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	if m.LangChosen && m.TypeChosen {
		return updateCreation(msg, m)
	} else if m.LangChosen {
		return updateLanguageType(msg, m)
	} else {
		return updateLanguage(msg, m)
	}
}

// Update functions for different views
func updateLanguage(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.LangChoice++
			if m.LangChoice > len(languageTypes)-1 {
				m.LangChoice = len(languageTypes)-1
			}
		case "k", "up":
			m.LangChoice--
			if m.LangChoice < 0 {
				m.LangChoice = 0
			}
		case "enter":
			m.LangChosen = true
			return m, tickCmd()
		}
	}
	return m, nil
}

// Update loop for the second view after a choice has been made
func updateLanguageType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
  selectedLanguage := languages[m.LangChoice] // Get the selected language
	selectedTypes := languageTypes[selectedLanguage] // Get the corresponding types
		switch msg.String() {
		case "j", "down":
			m.TypeChoice++
			if m.TypeChoice > len(selectedTypes)-1 {
				m.TypeChoice = len(selectedTypes)-1
			}
		case "k", "up":
			m.TypeChoice--
			if m.TypeChoice < 0 {
				m.TypeChoice = 0
			}
		case "enter":
			m.TypeChosen = true
			return m, tickCmd()
		}
	}
	return m, nil
}
// Update loop for the second view after a choice has been made
func updateCreation(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.TypeChoice++
			if m.TypeChoice > 3 {
				m.TypeChoice = 3
			}
		case "k", "up":
			m.TypeChoice--
			if m.TypeChoice < 0 {
				m.TypeChoice = 0
			}
		case "enter":
			m.TypeChosen = true
			return m, tickCmd()
		}
	}
	return m, nil
}

