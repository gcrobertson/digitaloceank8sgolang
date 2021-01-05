#!/bin/bash

#   Stop minikube if running ...
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

kubectl config use-context do-sfo2-millicent-cluster
echo 🏄 ' Done! The production environment is ready!'