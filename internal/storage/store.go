package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"atlas.otp/internal/model"
)

type Store struct {
	mu       sync.Mutex
	filePath string
	Accounts []model.Account
}

func NewStore() (*Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(home, ".atlas")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	return &Store{
		filePath: filepath.Join(configDir, "otp.json"),
		Accounts: []model.Account{},
	}, nil
}

func (s *Store) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.Accounts)
}

func (s *Store) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s.Accounts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *Store) Add(acc model.Account) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if acc.ID == "" {
		acc.ID = time.Now().Format("20060102150405")
	}
	s.Accounts = append(s.Accounts, acc)
}

func (s *Store) Delete(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= 0 && index < len(s.Accounts) {
		s.Accounts = append(s.Accounts[:index], s.Accounts[index+1:]...)
	}
}
