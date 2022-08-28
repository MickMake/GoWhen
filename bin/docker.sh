#!/bin/bash

TEMPLATE_GIT_REPO="$(jq -r '."git-repo"' $HOME/.GoWhen/config.json)"
TEMPLATE_GIT_DIR="/GoWhen"

docker build -t gowhen \
	-f Dockerfile \
	--build-arg 'GO_REPO_TOKEN=glpat-42424242424242424242' \
	--build-arg "TEMPLATE_GIT_REPO=${TEMPLATE_GIT_REPO}" \
	--build-arg "TEMPLATE_GIT_DIR=${TEMPLATE_GIT_DIR}" \
	.

docker run --rm -it gowhen:latest /usr/local/bin/GoWhen help
docker run --rm -it gowhen:latest /bin/sh

