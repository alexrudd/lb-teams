# lb-teams

This is a playground project for exploring the use of [Liftbridge](https://liftbridge.io/) as an [event store](https://en.wikipedia.org/wiki/Event_store).

The goal of the project is to implement a simple **Member** aggregate that captures the relationship between a **User** and a **Team**. Membership is represented in the domain as a mapping of **User** to **Team** but will feed into a read model projection of team members.