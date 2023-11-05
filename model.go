package main

import (
  tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	LangChoice int
	LangChosen bool
	TypeChoice int
	TypeChosen bool
	Quitting   bool
}

func (m model) Init() tea.Cmd {
	return nil
}

