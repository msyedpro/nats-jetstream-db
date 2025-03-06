# nats-jetstream-db

This is a Go Publisher and Consumer Project which connects to NATS JetStream Cluster Servers via Docker Compose. 
The data consumed by the consumer are appended to PostgreSQL database which is also served up Docker Compose.

## Starting

To start the NATS Streaming cluster and PostgresSQL database server, run this command:

```
$ docker-compose -f docker-compose.yaml up

 ✔ Network nats-jetstream-db_postgres-network  Created           0.1s 
 ✔ Network nats-jetstream-db_default           Created           0.1s 
 ✔ Container nats-jetstream-db-nats2-1         Created           0.3s 
 ✔ Container nats-jetstream-db-nats3-1         Created           0.3s 
 ✔ Container nats-jetstream-db-nats1-1         Created           0.3s 
 ✔ Container nats-jetstream-db-database-1      Created           0.4s 
Attaching to database-1, nats1-1, nats2-1, nats3-1
```

## Connection to the running NATS Streamng Cluster and PostgresSQL servers

Kick off the Go programs the publisher and then the consumer, which will connect to the NATS Streaming cluster and database:

```
$  go run .\publisher\main.go
2025/03/06 11:21:10 Publish message 1
2025/03/06 11:21:10 Publish message 2
2025/03/06 11:21:10 Publish message 3
```

```
$  go run .\consumer\main.go
2025/03/06 10:31:04 Received message: items.1
2025/03/06 10:31:04 Received message: items.2
2025/03/06 10:31:04 Received message: items.3
```

## Data storage of consumed data from NATS server

Check that the data consumed by the Consumer has been added to the Postgres SQL database by runing Docker command and SELECT statement below:

```
$  docker exec -ti nats-jetstream-db-database-1  psql -U dbuser -d dbuser
psql (17.4 (Debian 17.4-1.pgdg120+2))
Type "help" for help.

dbuser=# SELECT * from item;

id |  name   |          created
----+---------+----------------------------
  1 | items.1 | 2025-03-06 10:09:29.23009
  2 | items.2 | 2025-03-06 10:09:29.24906
  3 | items.3 | 2025-03-06 10:09:29.267487

```

## NATS Streaming Cluster Resilience

Test the resilience of the NATS Streaming Cluster by stopping one the NATS servers running container via Docker Desktop

![image](https://github.com/user-attachments/assets/ae9a60ae-7529-4482-afd7-15684ca65ba1)

Re-run the the Publisher and Consumer Go programs:

```
$  go run .\publisher\main.go
2025/03/06 11:21:10 Publish message 1
2025/03/06 11:21:10 Publish message 2
2025/03/06 11:21:10 Publish message 3
```

```
$  go run .\consumer\main.go
2025/03/06 10:31:04 Received message: items.1
2025/03/06 10:31:04 Received message: items.2
2025/03/06 10:31:04 Received message: items.3
```

## Clean up

To remove the NATS Streaming cluster and Postgres SQL, run:

```
$ docker-compose -f docker-compose.yaml down
[+] Running 6/6
 ✔ Container nats-jetstream-db-nats3-1         Removed                                                                                             2.2s 
 ✔ Container nats-jetstream-db-nats1-1         Removed                                                                                             2.1s 
 ✔ Container nats-jetstream-db-nats2-1         Removed                                                                                             0.1s 
 ✔ Container nats-jetstream-db-database-1      Removed                                                                                             1.4s 
 ✔ Network nats-jetstream-db_postgres-network  Removed                                                                                             0.8s 
 ✔ Network nats-jetstream-db_default           Removed                                                                                             1.5s 
```
