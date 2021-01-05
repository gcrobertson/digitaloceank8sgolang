# digitaloceank8sgolang

## Purpose

The purpose of this project is to learn Kubernetes [K8s] in greater detail by producing a usable application in Golang with K8s.

## Architecture

The project architecture is guided by learning and trying out new things, and not necessarily what is most efficient.

Certain patterns are unsuitable for a real production environment but completely fine for a single node setup or development environment. 

Examples:
- `secret` and `configMap` are not managed in the most efficient way. It is more about exploring the different k8s patterns.
- Using a volume of type `hostPath` is not OK in a multi-node production setting

There is no practical use for MySQL at the moment. I remember replication being complex and want to eventually try it in a K8s environment.

## Development vs Production

Production uses a service of type `LoadBalancer` while development uses a service of type `NodePort`.

Development is run locally via minikube and will pull insecure Docker images that are built locally.

Both development and production are managed through the `kubectl` CLI making use of: 
>
> kubectl config get-contexts
>
> kubectl config use-context < context > 
>

## Creating the project locally

See: bash/dev-env.sh

## Accessing the project URL locally:

See: bash/url.sh

## Creating the project on DigitalOcean K8s cluster

It is necessary to add the `context` so that `kubectl` can be the mechanism to access and modify the K8s cluster.

See: bash/prod-env.sh

## Why doesn't this project actually work?

The Dockerfile pulls from my private Dockerhub account and this is not something I am comfortable sharing right now.

To build your own Dockerfile locally and to commit and push it to your own Dockerhub,

See: _build/docker-build.md

Put in your own credentials and it should work after you connect to your own Dockerhub account.  If not, feel free to email me for help:

gregcrobertson@gmail.com