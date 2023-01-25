#!/bin/sh

helm repo add argo https://argoproj.github.io/argo-helm

VALUES="values.yaml"
kubectl get ingress argocd-server --namespace argocd \
    || VALUES="values-seed.yaml"

helm template \
    --dependency-update \
    --include-crds \
    --namespace argocd \
    --values "${VALUES}" \
    --set gitops.stage=${env_stage} \
    argocd . \
    | kubectl apply -n argocd -f -

kubectl -n argocd wait --timeout=60s --for condition=Established \
       crd/applications.argoproj.io \
       crd/applicationsets.argoproj.io
