#!/bin/sh

helm repo add argo https://argoproj.github.io/argo-helm

helm template \
    --dependency-update \
    --include-crds \
    --namespace argocd \
    argocd . \
    | kubectl apply -n argocd -f -

kubectl -n argocd wait --timeout=60s --for condition=Established \
       crd/applications.argoproj.io \
       crd/applicationsets.argoproj.io
