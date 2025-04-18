package todos

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
