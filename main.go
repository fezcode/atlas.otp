package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"atlas.otp/internal/storage"
	"atlas.otp/internal/ui"
)

var Version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("atlas.otp v%s\n", Version)
		return
	}
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		fmt.Println("Atlas OTP - Minimalist, secure terminal TOTP (2FA) manager.")
		fmt.Println("\nUsage:")
		fmt.Println("  atlas.otp        Start the TOTP manager TUI")
		fmt.Println("  atlas.otp -v     Show version")
		fmt.Println("  atlas.otp -h     Show this help")
		return
	}

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
