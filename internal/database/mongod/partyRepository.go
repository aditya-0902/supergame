package mongod

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aditya-0902/supergaming/internal/models"
	"github.com/aditya-0902/supergaming/internal/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) CreateParty(ctx context.Context, leaderID string) (string, error) {
	partyID := uuid.New().String()
	party := bson.M{
		"party_id":     partyID,
		"party_leader": leaderID,
		"members":      []string{leaderID},
		"invitations":  []string{},
		"created_at":   time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, party)
	if err != nil {
		return "", err
	}

	return partyID, nil
}

func (r *MongoRepository) InviteToParty(ctx context.Context, partyID, inviterID, inviteeID string) error {
	filter := bson.M{"party_id": partyID, "party_leader": inviterID}
	update := bson.M{
		"$addToSet": bson.M{
			"invitations": bson.M{
				"invitee_id": inviteeID,
				"status":     "pending",
				"invited_at": time.Now(),
			},
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("party not found or you are not the party leader")
	}
	return nil
}

func (r *MongoRepository) JoinParty(ctx context.Context, leader_id, partyID, userID string) error {
	filter := bson.M{
		"party_id": partyID,
		"invitations": bson.M{
			"$elemMatch": bson.M{"invitee_id": userID},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"invitations.$.status": "accepted",
		},
		"$addToSet": bson.M{
			"members": userID,
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("party not found or no invitation found for the user")
	}

	utils.AddParticipantToParty(leader_id, partyID, userID)

	return nil
}

func (r *MongoRepository) LeaveParty(ctx context.Context, leaderID, partyID, userID string) error {
	filter := bson.M{"party_id": partyID}
	update := bson.M{
		"$pull": bson.M{"members": userID},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("party not found")
	}
	utils.RemoveParticipantFromParty(leaderID, userID, partyID)
	return nil
}

func (r *MongoRepository) AcceptInvite(ctx context.Context, leaderID, partyID, userID string) error {
	return r.JoinParty(ctx, leaderID, partyID, userID)
}

func (r *MongoRepository) RejectInvite(ctx context.Context, partyID, userID string) error {
	filter := bson.M{"party_id": partyID}
	update := bson.M{
		"$pull": bson.M{
			"invitations": bson.M{"invitee_id": userID},
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("party not found")
	}
	return nil
}

func (r *MongoRepository) RemoveUser(ctx context.Context, partyID, userID string) error {
	fmt.Println(partyID, userID)
	filter := bson.M{"party_id": partyID}
	update := bson.M{"$pull": bson.M{"members": userID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	filter = bson.M{"party_id": partyID}
	var result models.Party

	err = r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}
	utils.RemoveParticipantFromParty(result.PartyLeader, userID, partyID)
	return nil
}

func (r *MongoRepository) GetPartyByID(ctx context.Context, partyID string) (*models.Party, error) {
	var party *models.Party
	filter := bson.M{"party_id": partyID}

	err := r.collection.FindOne(ctx, filter).Decode(&party)
	if err != nil {
		return &models.Party{}, err
	}
	fmt.Printf("%+v \n", party)
	return party, nil
}
