## run api

    bee run -gendoc=true -downdoc=true


## build image

    docker build -t restapi-1.0.0 .


## run in kubernetes cluster

    kubectl apply -f deploy-k8s.yml


## How to implement application deployment in kubenets？

    https://blog.csdn.net/weixin_41806245/article/details/93745532