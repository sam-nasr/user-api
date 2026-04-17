package services

import "user-api/internal/models"
import "fmt"

// Interface defines the contract for user operations
type UserService interface {
	AddUser(user models.User)       // Add a new user
	GetUsers() []models.User        // Get all users
	DeleteUser(id int) error          // Delete user by ID
	UpdateUser(id int, updated models.User) bool
}

// InMemoryUserService stores users in memory (slice)
type InMemoryUserService struct {
	users []models.User // slice acting as in-memory database
}

// Constructor to initialize service with empty slice
func NewUserService() *InMemoryUserService {
	return &InMemoryUserService{
		users: []models.User{}, // initialize empty (not nil)
	}
}

// AddUser adds a user to the slice
func (s *InMemoryUserService) AddUser(user models.User) {
	s.users = append(s.users, user)
}

// GetUsers returns all users
func (s *InMemoryUserService) GetUsers() []models.User {
	return s.users
}

func (s *InMemoryUserService) UpdateUser(id int, updated models.User) bool {
	for i, u := range s.users {
		if u.ID == id {
			s.users[i] = updated
			return true
		}
	}
	return false
}

// DeleteUser removes user by ID
func (s *InMemoryUserService) DeleteUser(id int) error {
	var updated []models.User // new slice to store filtered users
	// make it empty (not nil) to avoid nil slice issues
	updated = []models.User{}
	found := false

	for _, u := range s.users {
		if u.ID != id {
			updated = append(updated, u)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("user with ID %d not found", id)
	}

	s.users = updated // replace old slice
	return nil
}