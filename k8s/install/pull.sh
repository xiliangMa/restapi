#!/bin/bash
images=(kube-proxy:v1.14.0 kube-scheduler:v1.14.0 kube-controller-manager:v1.14.0 kube-apiserver:v1.14.0 etcd:3.3.10 coredns:1.3.1 pause:3.1 )
for imageName in ${images[@]} ; do
docker pull xiliangma/$imageName
docker tag  xiliangma/$imageName k8s.gcr.io/$imageName
docker rmi  xiliangma/$imageName
done
