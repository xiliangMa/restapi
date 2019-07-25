## run api

    set env:
        MARIADB_DATABASE=uranus_local;MARIADB_USER=root;MARIADB_PASSWORD=abc123123;MARIADB_HOST=127.0.0.1
    
    go build:
        go build -o restapi
        
    run：
        bee run -gendoc=true -downdoc=true
        
    run by restapi：
        ./restapi

## build image

    docker build -t restapi-1.0.0 .


## run in kubernetes cluster

    kubectl apply -f deploy-k8s.yml


## How to implement application deployment in kubenets？

    https://blog.csdn.net/weixin_41806245/article/details/93745532
