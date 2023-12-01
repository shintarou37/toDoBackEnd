package user

import (
	"context"
	"time"
	// "golang.org/x/crypto/bcrypt"
)

type Repository interface {
	FindOne(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	ID        string    `firestore:"id,omitempty" json:"id"`
	Name      string    `firestore:"name,omitempty" json:"name"`
	Email     string    `firestore:"email,omitempty" json:"email"`
	Password  []byte    `firestore:"password,omitempty" json:"password"`
	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at" json:"updated_at"`
}
