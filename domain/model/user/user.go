package user

import (
	"context"
	"time"
)

type UserRepository interface {
	FindOne(ctx context.Context, id string) (*User, error)
}

type User struct {
	ID        string    `firestore:"id,omitempty" json:"id"`
	Name      string    `firestore:"name,omitempty" json:"name"`
	Email     string    `firestore:"email,omitempty" json:"email"`
	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at" json:"updated_at"`
}
