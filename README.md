# mongo-db-cluster-sample

### How to run sample

* setup environment with docker-compose
    
    ```bash
    docker-compose up -d
    ```
    
* setup replica set in mongo's containers
    
    ```bash
    bash set-primary.sh
    ```
    
* run application
    
    ```bash
    go run cmd/mongo-db-cluster-sample/main.go
    ```
