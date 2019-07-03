#!/usr/bin/env bash

# docker-compose exec mongodb-artitrator mongo --eval \
    # 'config={"_id":"TestMongoReplicaSet","members":[{"_id":0,"host":"mongodb-artitrator:27017", arbiterOnly : true},{"_id":1,"host":"mongodb-replica-1:27017"},{"_id":2,"host":"mongodb-replica-2:27017"}]};rs.initiate(config)'

docker-compose exec mongodb-replica-1 mongo --eval \
'
rs.initiate({"_id" : "TestMongoReplicaSet", members : [
    {"_id" : 0, priority : 3, host : "mongodb-replica-1:27017"},
    {"_id" : 1, host : "mongodb-replica-2:27017"},
    {"_id" : 2, host : "mongodb-arbitrator:27017", arbiterOnly : true}
]});
'

#docker-compose exec mongo0 mongo --eval \
#    'config={"_id":"rs0","members":[{"_id":0,"host":"mongo0:27017"},{"_id":1,"host":"mongo1:27017"},{"_id":2,"host":"mongo2:27017"}]};rs.initiate(config)'


# rs.remove("mongodb-artitrator:27017")
# rs.addArb("mongodb-arbitrator:27017") 
