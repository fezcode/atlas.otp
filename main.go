package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"atlas.otp/internal/storage"
	"atlas.otp/internal/ui"
)

func main() {
	store, err := storage.NewStore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing store: %v\n", err)
		os.Exit(1)
	}

	if err := store.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading accounts: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(ui.NewModel(store), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
