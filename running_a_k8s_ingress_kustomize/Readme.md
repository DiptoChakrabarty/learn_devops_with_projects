# Running a Kubernetes ingress (Kustomize edition)

## by [@selamanse](https://github.com/selamanse)

## DESCRIPTION OF PROJECT

Create three types of kubernetes resources

Two deployments of nginx which have different messages when we hit their main endpoint
Create two services for each of the deployments
Create a single ingress which redirects to first service on endpoint /first and second endpoint on endpoint /second

## Prerequisites

- nginx ingress controller (if not present [described here](https://docs.rancherdesktop.io/how-to-guides/setup-NGINX-Ingress-Controller/))
- port forward to ingress controller (if you already have a different one on port 80) `kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 8080:80`

## Steps to run

- `kubectl create ns kube-ingress`
- `kubectl kustomize kubernetes | kubectl apply -f -`
- optionally wait for service1: `kubectl wait --for=jsonpath='{.status.readyReplicas}'=2 deployment service1-deployment -n kube-ingress`
- optionally wait for service2: `kubectl wait --for=jsonpath='{.status.readyReplicas}'=2 deployment service2-deployment -n kube-ingress`
- `curl http://localhost/first`
- `curl http://localhost/second`

## ScreenShots

[![asciicast](https://asciinema.org/a/612618.svg)](https://asciinema.org/a/612618)
