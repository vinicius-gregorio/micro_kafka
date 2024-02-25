# Challenge Full Cycle - Event Driven Architecture



## How to run this program

In terminal, with docker installed run:
`docker-compose up -d`

Check if all services are up
`docker-compose ps` (-a flag if wanted)

You should see something like this:


```
control-center                          /etc/confluent/docker/run        Up      0.0.0.0:9021->9021/tcp
kafka                                   /etc/confluent/docker/run        Up      0.0.0.0:9092->9092/tcp
microservices_challenge_balance-app_1   /usr/local/bin/entrypoint. ...   Up      0.0.0.0:3003->3003/tcp
microservices_challenge_goapp_1         /usr/local/bin/entrypoint. ...   Up      0.0.0.0:8080->8080/tcp
microservices_challenge_mysql_1         docker-entrypoint.sh mysqld      Up      0.0.0.0:3306->3306/tcp, 33060/tcp
zookeeper                               /etc/confluent/docker/run        Up      0.0.0.0:2181->2181/tcp, 2888/tcp, 3888/tcp
```

## Checking

You can see the file `api.http` in the root folder, it contains all request this system has.
