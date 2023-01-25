#!/bin/sh

echo "${env_stage}"
VALUES="values.yaml"

kubectl get ingress gitea --namespace gitea \
    || VALUES="values-seed.yaml"

helm template \
    --include-crds \
    --namespace argocd \
    --values "${VALUES}" \
    --set gitops.stage=${env_stage} \
    argocd . \
    | kubectl apply -n argocd -f -
