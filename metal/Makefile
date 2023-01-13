.POSIX:


ifeq (dev, ${env_target})

default: cluster

cluster:
	k3d cluster start homelab-dev || k3d cluster create --config k3d-dev.yaml
	k3d kubeconfig get homelab-dev > kubeconfig.yaml

clean:
	k3d cluster delete homelab-dev

else

default: boot cluster

~/.ssh/id_ed25519:
	ssh-keygen -t ed25519 -P '' -f "$@"

boot: ~/.ssh/id_ed25519
	ansible-playbook \
		--inventory inventories/${env_target}.yml \
		boot.yml

cluster:
	ansible-playbook \
		--inventory inventories/${env_target}.yml \
		cluster.yml

clean:

endif

console:
	ansible-console \
		--inventory inventories/${env_target}.yml
