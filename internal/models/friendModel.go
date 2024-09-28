package models

import "time"

type Friend struct {
	FriendID string    `bson:"friend_id" json:"friend_id"`
	AddedAt  time.Time `bson:"added_at" json:"added_at"`
}

type FriendRequest struct {
	RequestID string    `bson:"request_id" json:"request_id"`
	FromUser  string    `bson:"from_user" json:"from_user"`
	Status    string    `bson:"status" json:"status"` // pending, accepted, rejected
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type User struct {
	ID             string          `bson:"_id" json:"id"`
	Username       string          `bson:"username" json:"username"`
	Friends        []Friend        `bson:"friends" json:"friends"`
	FriendRequests []FriendRequest `bson:"friend_requests" json:"friend_requests"`
}

type FriendRequestModel struct {
	FromUser string `json:"from_user"`
	ToUser   string `json:"to_user"`
}
type AcceptRequsetModel struct {
	ToUser    string `json:"to_user"`
	RequestID string `json:"request_id"`
}

type RejectRequestModel struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
}
