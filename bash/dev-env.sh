#!/bin/bash

#   Restart minikube if Running ...
#
#
#
#
#
status=$(minikube status | grep -i "host: ")
if [ "${status}" != 'host: Stopped' ]

    then 

    minikube stop

fi

#   Manually verify K8s is prepared to pull insecure docker images [local docker images] ...
#
#
#
#
#
echo 📌 ' Note: Make sure the Deployment can pull local Docker Image ...'
echo 🔎 ' Check Deployment YML: spec.template.spec.containers[0].imagePullPolicy: IfNotPresent'
# read -n 1 -s -r -p "Press any key to continue"

#   Start minikube
#
#
#
#
#
minikube start --insecure-registry=true
eval $(minikube docker-env)

#   Use the correct, local kubectl context
#
#
#
#
#
kubectl config use-context minikube
echo 🏄 ' Done! The local environment is ready!'

# LINK[s]: 
# - https://stackoverflow.com/questions/40144138/pull-a-local-image-to-run-a-pod-in-kubernetes
# - https://minikube.sigs.k8s.io/docs/handbook/pushing/



#   minikube -p minikube docker-env
#   eval $(minikube -p minikube docker-env)