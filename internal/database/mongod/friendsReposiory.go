package mongod

import (
	"context"
	"errors"
	"time"

	"github.com/aditya-0902/supergaming/internal/models"
	"github.com/aditya-0902/supergaming/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoRepository) AddFriendRequest(ctx context.Context, fromUser, toUser string) error {
	request := models.FriendRequest{
		RequestID: primitive.NewObjectID().Hex(),
		FromUser:  fromUser,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	filter := bson.M{"_id": toUser}
	update := bson.M{"$push": bson.M{"friend_requests": request}}

	_, err := m.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func (m *MongoRepository) AcceptFriendRequest(ctx context.Context, toUser, requestID string) error {

	var result struct {
		FriendRequests []struct {
			RequestID string `bson:"request_id"`
			FromUser  string `bson:"from_user"`
			Status    string `bson:"status"`
		} `bson:"friend_requests"`
	}

	filter := bson.M{"_id": toUser, "friend_requests.request_id": requestID}
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}

	var fromUser string
	for _, req := range result.FriendRequests {
		if req.RequestID == requestID && req.Status == "pending" {
			fromUser = req.FromUser
			break
		}
	}
	if fromUser == "" {
		return errors.New("friend request not found")
	}

	update := bson.M{
		"$set": bson.M{
			"friend_requests.$.status": "accepted",
		},
		"$push": bson.M{
			"friends": bson.M{
				"friend_id": fromUser,
				"added_at":  time.Now(),
			},
		},
	}

	_, err = m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	filter = bson.M{"_id": fromUser}

	update = bson.M{
		"$push": bson.M{
			"friend_requests": bson.M{
				"request_id": requestID,
				"from_user":  toUser,
				"status":     "accepted",
				"created_at": time.Now(),
			},
			"friends": bson.M{
				"friend_id": toUser,
				"added_at":  time.Now(),
			},
		},
	}

	updateOpts := options.Update().SetUpsert(true)

	_, err = m.collection.UpdateOne(ctx, filter, update, updateOpts)
	if err != nil {
		return err
	}

	utils.AddFriendToFriendOnlineStatusMap(fromUser, toUser)

	return nil
}

func (m *MongoRepository) RemoveFriend(ctx context.Context, userID, friendID string) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$pull": bson.M{"friends": bson.M{"friend_id": friendID}}}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	filter = bson.M{"_id": friendID}
	update = bson.M{"$pull": bson.M{"friends": bson.M{"friend_id": userID}}}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	utils.RemoveFriendFromFriendOnlineStatusMap(userID, friendID)
	return err
}

func (m *MongoRepository) GetFriendsList(ctx context.Context, userID string) ([]models.Friend, error) {
	var user models.User
	filter := bson.M{"_id": userID}
	projection := bson.M{"friends": 1, "_id": 0}

	err := m.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.Friends, nil
}

func (m *MongoRepository) RejectFriendRequest(ctx context.Context, toUser, requestID string) error {
	filter := bson.M{"_id": toUser, "friend_requests.request_id": requestID}
	update := bson.M{
		"$set": bson.M{"friend_requests.$.status": "rejected"},
	}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}
