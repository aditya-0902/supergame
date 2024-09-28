# supergame

Design and implement backend services for a lite social-presence system for online-games using Go(lang)
[ Level 1 ] REST APIs
Friends
- Add Friends: Users can send friend requests to other users.
- Accept/Reject Friend Requests: Users can accept or reject friend requests they receive.
- Remove Friends: Users can remove other users from their friend list.
- View Friend List: Users can view their current list of friends.
Party
- Create Party: Users can create a game party (short-sessions).
- Invite to Party: Users can invite their friends to join their game party.
- Join Party: Users can join a game party they have been invited to.
- Leave Party: Users can leave a game party they are currently in.
- Accept/Reject Party Invitations: Users can accept or reject game party invitations they receive.
- Remove/Kick Users from Party: The party leader can remove users from the party.
[ Level 2 ] Real-Time Communication ( websockets/grpc )
Notifications
- Online Status: Users can see the online status of their friends in real-time.
- Party Status: Users can see in real time which all friends in their party.
[ Level 3 ] Orchestration ( docker, kubernetes )
● Containerize your services using Docker
○ Make all the REST API available locally via docker-compose.
○ HTTP Clients like curl/POSTMAN will be used to test the APIs.




To run the container 

first use this command := "docker build . -t supergaming"

then run the command := "docker compose up"

The API formats are given in api.txt file where you can check the API calls

The request_id for friends  and party_id for party routes can be fetched from db




