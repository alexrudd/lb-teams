# lb-teams

This is a playground project for exploring the use of [Liftbridge](https://liftbridge.io/) as an [event store](https://en.wikipedia.org/wiki/Event_store).

The goal of the project is to implement a simple **User** aggregate that captures the relationship between a **User** and a **Team**. Events from the **User** aggregate will then be built into a projection that lists all teams and their members.

The commands supported by the User aggregate are:

* CreateTeam - `curl localhost:8080/CreateTeam -d '{"userId":"bob","teamId":"party"}'`
* JoinTeam - `curl localhost:8080/JoinTeam -d '{"userId":"alice","teamId":"party"}'`
* LeaveTeam  - `curl localhost:8080/LeaveTeam -d '{"userId":"alice","teamId":"party"}'`
* ChangeOwner - `curl localhost:8080/ChangeOwner -d '{"userId":"bob","newOwnerUserId":"alice"}'`
* DisbandTeam - `curl localhost:8080/DisbandTeam -d '{"userId":"alice"}'`