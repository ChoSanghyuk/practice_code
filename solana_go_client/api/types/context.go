package types

type ContextKey string

const (
	RequestID ContextKey = "reqID"
	UserID    ContextKey = "userID"
)
