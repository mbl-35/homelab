.POSIX:


ifeq (devel, ${env_stage})

default: cluster

cluster:
	k3d cluster start homelab-dev || k3d cluster create --config k3d-dev.yaml
	k3d kubeconfig get homelab-dev > ${KUBECONFIG}
	chown ${NIX_USER_ID}:${NIX_USER_GID} ${KUBECONFIG}

clean:
	k3d cluster delete homelab-dev && \
		[ ! -f ${KUBECONFIG} ] || rm ${KUBECONFIG}

else

default: boot cluster

~/.ssh/id_ed25519:
	ssh-keygen -t ed25519 -P '' -f "$@"

boot: ~/.ssh/id_ed25519
	ansible-playbook \
		--inventory inventories/${env_stage}.yml \
		boot.yml

cluster:
	ansible-playbook \
		--inventory inventories/${env_stage}.yml \
		cluster.yml

clean:

endif

console:
	ansible-console \
		--inventory inventories/${env_stage}.yml
