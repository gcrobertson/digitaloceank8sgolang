#!/bin/bash
# chmod 700 <this-script>

# Colors because they're great!
# link: https://misc.flogisoft.com/bash/tip_colors_and_formatting
BGBlue="\e[44m"
BGReset="\e[49m"
FUnderline="\e[4m"
FGGreen="\e[32m"
FReset="\e[0m"


# K8s accessibility in development is made via: http://<public-cluster-ip>:<node-port>
#
# Note: The entire reason this script is necessary is exposing the deployment:
#       - type `NodePort`       for development
#       - type `LoadBalancer`   for production

# 1. cluster-ip:
#
#
# The <public-cluster-ip> can be found in minikube via: 
# > kubectl cluster-info
# 
#
# The following 1 liner will also return the public-cluster-ip:
# link: https://stackoverflow.com/questions/61861754/how-to-get-k8s-master-ip-address
clusterIP=$(kubectl get nodes --selector=node-role.kubernetes.io/master -o jsonpath='{$.items[*].status.addresses[?(@.type=="InternalIP")].address}')

# 2. <node-port> for the service that exposes the deployment: 
# 
# 
# The <node-port> can be found in minikube via:
# > kubectl describe service millicent-deployment-service | grep -i nodePort
#
#
# For now, it is just hard-coded into the service definition file itself as 32603.
# 
# @TODO: An improvement would be to make this node port dynamic.
nodePort="32603"


echo -e "\n\n\n\n\n\n\n\n${BGBlue}DEVELOPMENT${BGReset} application URL:\n\n${FUnderline}${FGGreen}http://${clusterIP}:${nodePort}${FReset}\n\n\n\n\n\n\n\n"