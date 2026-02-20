package ui

import (
	"time"

	"atlas.otp/internal/model"
	"atlas.otp/internal/storage"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xlzd/gotp"
)

type state int

const (
	stateNormal state = iota
	stateAdding
	stateConfirmDelete
)

type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type Model struct {
	store      *storage.Store
	cursor     int
	state      state
	inputName  textinput.Model
	inputSec   textinput.Model
	inputField int // 0: Name, 1: Secret
	help       help.Model
	keys       keyMap
	lastErr    string
	copied     bool
}

func NewModel(store *storage.Store) Model {
	tiName := textinput.New()
	tiName.Placeholder = "Account Name (e.g. GitHub)"
	tiName.Focus()

	tiSec := textinput.New()
	tiSec.Placeholder = "Secret Key (JBSWY3DPEHPK3PXP)"

	return Model{
		store:    store,
		state:    stateNormal,
		inputName: tiName,
		inputSec:  tiSec,
		help:     help.New(),
		keys:     keys,
	}
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		return m, doTick()
	case tea.KeyMsg:
		switch m.state {
		case stateNormal:
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.keys.Up):
				if m.cursor > 0 {
					m.cursor--
				}
				m.copied = false
			case key.Matches(msg, m.keys.Down):
				if m.cursor < len(m.store.Accounts)-1 {
					m.cursor++
				}
				m.copied = false
			case key.Matches(msg, m.keys.Add):
				m.state = stateAdding
				m.inputField = 0
				m.inputName.Focus()
				return m, nil
			case key.Matches(msg, m.keys.Delete):
				if len(m.store.Accounts) > 0 {
					m.state = stateConfirmDelete
				}
			case key.Matches(msg, m.keys.Copy):
				if len(m.store.Accounts) > 0 {
					acc := m.store.Accounts[m.cursor]
					code, _ := getOTP(acc.Secret)
					clipboard.WriteAll(code)
					m.copied = true
				}
			}
		case stateAdding:
			if msg.Type == tea.KeyEsc {
				m.state = stateNormal
				m.inputName.Reset()
				m.inputSec.Reset()
				return m, nil
			}
			if msg.Type == tea.KeyTab || msg.Type == tea.KeyDown || msg.Type == tea.KeyUp {
				if m.inputField == 0 {
					m.inputField = 1
					m.inputName.Blur()
					m.inputSec.Focus()
				} else {
					m.inputField = 0
					m.inputSec.Blur()
					m.inputName.Focus()
				}
				return m, nil
			}
			if msg.Type == tea.KeyEnter {
				if m.inputField == 0 {
					m.inputField = 1
					m.inputName.Blur()
					m.inputSec.Focus()
					return m, nil
				}
				// Save
				name := m.inputName.Value()
				secret := m.inputSec.Value()
				if name != "" && secret != "" {
					m.store.Add(model.Account{Name: name, Secret: secret})
					m.store.Save()
					m.state = stateNormal
					m.inputName.Reset()
					m.inputSec.Reset()
				}
				return m, nil
			}
			var cmd tea.Cmd
			if m.inputField == 0 {
				m.inputName, cmd = m.inputName.Update(msg)
			} else {
				m.inputSec, cmd = m.inputSec.Update(msg)
			}
			return m, cmd
		case stateConfirmDelete:
			if msg.String() == "y" {
				m.store.Delete(m.cursor)
				m.store.Save()
				if m.cursor >= len(m.store.Accounts) && m.cursor > 0 {
					m.cursor--
				}
				m.state = stateNormal
			} else if msg.String() == "n" || msg.Type == tea.KeyEsc {
				m.state = stateNormal
			}
			return m, nil
		}
	}
	return m, nil
}

func getOTP(secret string) (string, int) {
	totp := gotp.NewDefaultTOTP(secret)
	now := time.Now().Unix()
	code := totp.At(now)
	expiresIn := 30 - (now % 30)
	return code, int(expiresIn)
}
