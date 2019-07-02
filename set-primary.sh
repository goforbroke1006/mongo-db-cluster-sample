#!/usr/bin/env bash

docker-compose exec mongodb-replica-1 mongo --eval \
    'config={"_id":"TestMongoReplicaSet","members":[{"_id":0,"host":"mongodb-replica-1:27017"},{"_id":1,"host":"mongodb-replica-2:27017"},{"_id":2,"host":"mongodb-replica-3:27017"}]};rs.initiate(config)'


#docker-compose exec mongo0 mongo --eval \
#    'config={"_id":"rs0","members":[{"_id":0,"host":"mongo0:27017"},{"_id":1,"host":"mongo1:27017"},{"_id":2,"host":"mongo2:27017"}]};rs.initiate(config)'

