package constants

import "os"

var MONGO_URI string

const FRIEND_DATABASE = "friends"
const FRIEND_COLLECTION = "friends"
const PARTY_DATABASE = "party"
const PARTY_COLLECTION = "party"

func init() {
	MONGO_URI = os.Getenv("MONGO_URI")
}
