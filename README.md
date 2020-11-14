# lb-teams

This is a playground project for exploring the use of [Liftbridge](https://liftbridge.io/) as an [event store](https://en.wikipedia.org/wiki/Event_store).

The goal of this project is to implement a simple domain of **Users** and **Teams**, where users can be invited to teams, accept invites to teams, leave teams, and remove other team members (assuming they're the team owner).

The commands supported by the User aggregate are:

* InviteUserToTeam - `curl localhost:8080/InviteUserToTeam -d '{"userId":"alice","teamId":"party"}'`
* AcceptInvite - `curl localhost:8080/AcceptInvite -d '{"userId":"alice","teamId":"party"}'`
* LeaveTeam  - `curl localhost:8080/LeaveTeam -d '{"userId":"alice"}'`
