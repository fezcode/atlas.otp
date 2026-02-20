package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var s strings.Builder

	// Header
	s.WriteString(titleStyle.Render(" ATLAS.OTP ") + " " + headerStyle.Render("Secure One-Time Passwords") + "\n\n")

	switch m.state {
	case stateNormal, stateConfirmDelete:
		if len(m.store.Accounts) == 0 {
			s.WriteString(itemStyle.Render("No accounts yet. Press 'a' to add one."))
		} else {
			for i, acc := range m.store.Accounts {
				cursor := " "
				style := itemStyle
				if m.cursor == i {
					cursor = ">"
					style = selectedItemStyle
				}
				s.WriteString(style.Render(fmt.Sprintf("%s %s", cursor, acc.Name)) + "\n")
			}
		}

		if len(m.store.Accounts) > 0 {
			acc := m.store.Accounts[m.cursor]
			code, expiresIn := getOTP(acc.Secret)
			
			s.WriteString("\n")
			s.WriteString(headerStyle.Render(fmt.Sprintf("  %s Code:", acc.Name)) + "\n")
			
			codeStr := otpCodeStyle.Render(fmt.Sprintf(" %s %s ", code[:3], code[3:]))
			if m.copied {
				codeStr += " " + successStyle.Render("Copied!")
			}
			s.WriteString("  " + codeStr + "\n")

			// Progress bar
			barWidth := 20
			filled := (expiresIn * barWidth) / 30
			bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
			s.WriteString(fmt.Sprintf("  %s %ds\n", progressBarStyle.Render(bar), expiresIn))
		}

		if m.state == stateConfirmDelete {
			s.WriteString("\n" + deleteWarnStyle.Render("  Really delete this account? (y/n)"))
		}

	case stateAdding:
		s.WriteString(headerStyle.Render("  Add New Account") + "\n\n")
		
		nameStyle := lipgloss.NewStyle()
		secStyle := lipgloss.NewStyle()
		if m.inputField == 0 {
			nameStyle = nameStyle.Foreground(gold)
		} else {
			secStyle = secStyle.Foreground(gold)
		}

		s.WriteString("  " + nameStyle.Render("Name:  ") + m.inputName.View() + "\n")
		s.WriteString("  " + secStyle.Render("Secret: ") + m.inputSec.View() + "\n")
		s.WriteString("\n  " + dimStyle.Render("Press TAB to switch, ENTER to save, ESC to cancel."))
	}

	// Help
	s.WriteString("\n\n" + helpStyle.Render(m.help.View(m.keys)))

	return appStyle.Render(s.String())
}

var deleteWarnStyle = lipgloss.NewStyle().
	Foreground(red).
	Bold(true)
