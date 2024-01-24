package token

import "time"

// Maker is interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specified username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// Verify Token
	VerifyToken(token string) (*Payload, error)
}
