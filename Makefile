.POSIX:
.PHONY: *
.EXPORT_ALL_VARIABLES:

KUBECONFIG = $(shell pwd)/metal/kubeconfig.yaml
KUBE_CONFIG_PATH = $(KUBECONFIG)
DOTENV_FILE = $(shell pwd)/.env

ifeq (env,$(firstword $(MAKECMDGOALS)))
    COMMAND_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
    COMMAND_ARGS := $(subst :,\:,$(COMMAND_ARGS))
    $(eval $(COMMAND_ARGS):;@:)
else ifneq (,$(wildcard $(DOTENV_FILE)))
    include $(DOTENV_FILE)
    export 
endif


default: guard-env metal bootstrap external smoke-test post-install clean

configure: guard-env
	git config --global --add safe.directory $(shell pwd)
	./scripts/configure
	git status

metal: guard-env
	make -C metal env=${env_stage}

bootstrap: guard-env
	make -C bootstrap env=${env_stage}

external: guard-env
	[ "${env_stage}" = "devel" ] || make -C external

smoke-test: guard-env
	make -C test filter=Smoke

post-install: guard-env
	@[ "${env_stage}" = "devel" ] || ./scripts/hacks

tools:
	@[ -f $(shell pwd)/.docker.json ] || \
		echo "{}" >$(shell pwd)/.docker.json
	@docker run \
		--rm \
		--interactive \
		--tty \
		--network host \
		--env "NIX_USER=${USER}" \
		--env "NIX_USER_ID=$(shell id -u)" \
		--env "NIX_USER_GID=$(shell id -g)" \
		--env "KUBECONFIG=${KUBECONFIG}" \
		--volume "/var/run/docker.sock:/var/run/docker.sock" \
		--volume "$(shell pwd)/.docker.json:/root/.docker/config.json" \
		--volume $(shell pwd):$(shell pwd) \
		--volume ${HOME}/.ssh:/root/.ssh \
		--volume ${HOME}/.terraform.d:/root/.terraform.d \
		--volume homelab-tools-cache:/root/.cache \
		--volume homelab-tools-nix:/nix \
		--workdir $(shell pwd) \
		nixos/nix nix-shell

test: guard-env
	make -C test

clean: guard-env
	[ "${env_stage}" = "devel" ] || docker compose --project-directory ./metal/roles/pxe_server/files down
	make -C metal clean

docs:
	docker run \
		--rm \
		--interactive \
		--tty \
		--publish 8000:8000 \
		--volume $(shell pwd):/docs \
		squidfunk/mkdocs-material

git-hooks:
	pre-commit install

guard-env: 
	@[ "${env_stage}" ] || ( echo ">> Stage is not set in .env file use command 'make env stage:<target>"; exit 1 )
	@echo "Selected stage: ${env_stage}"
	#@[ "${env_stage}" != "devel" ] || ( ./scripts/update-dev-inventory )

env:
	@for cfg in $(COMMAND_ARGS); do \
		echo env_$$cfg | tr \: \= > "${DOTENV_FILE}"; \
	done
	@chown ${NIX_USER_ID}:${NIX_USER_GID} ${DOTENV_FILE}
