package database

import (
	"context"
	"fmt"

	"github.com/aditya-0902/supergaming/internal/database/mongod"
	"github.com/aditya-0902/supergaming/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	FriendRepository
	PartyRepository
}

type PartyRepository interface {
	CreateParty(ctx context.Context, leaderID string) (string, error)
	InviteToParty(ctx context.Context, partyID, inviterID, inviteeID string) error
	JoinParty(ctx context.Context, leader, partyID, userID string) error
	LeaveParty(ctx context.Context, leaderID, partyID, userID string) error
	AcceptInvite(ctx context.Context, leaderID, partyID, userID string) error
	RejectInvite(ctx context.Context, partyID, userID string) error
	RemoveUser(ctx context.Context, partyID, userID string) error
	GetPartyByID(ctx context.Context, partyID string) (*models.Party, error)
}

type FriendRepository interface {
	AddFriendRequest(ctx context.Context, fromUser, toUser string) error
	AcceptFriendRequest(ctx context.Context, toUser, requestID string) error
	RemoveFriend(ctx context.Context, userID, friendID string) error
	GetFriendsList(ctx context.Context, userID string) ([]models.Friend, error)
	RejectFriendRequest(ctx context.Context, toUser, requestID string) error
}

func NewMongoDatabase(uri string, opts ...mongod.Option) (Repository, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("couldn't create mongo client: %v", err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}

	return mongod.NewMongoRepository(client, opts...), nil
}
