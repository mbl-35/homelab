.POSIX:

default: namespace argocd root

namespace:
	kubectl create namespace argocd --dry-run=client --output=yaml \
		| kubectl apply -f -

.PHONY: argocd
argocd:
	cd argocd && ./apply.sh
	[ ! -f ./Chart.lock ] || chown ${NIX_USER_ID}:${NIX_USER_GID} ./Chart.lock

.PHONY: root
root:
	cd root && ./apply.sh
