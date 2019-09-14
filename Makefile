PROJECT_NAME = docnoc
DOCKER_USERNAME = mtusman
DOCKER_TAG_NAME = $(DOCKER_USERNAME)/$(PROJECT_NAME)
ABSOLUTE_CONFIG_FILE_PATH = $(PWD)/examples/docnoc_config.yaml

.PHONY: build run ssh creds push

build:
	docker build . -t $(PROJECT_NAME)

run:
	docker run -v /var/run/docker.sock:/var/run/docker.sock -v $(ABSOLUTE_CONFIG_FILE_PATH):/tmp/docnoc_config.yaml $(PROJECT_NAME)

up:
	docker-compose up

creds:
	echo $(DOCKER_PASSWORD) | docker login -u $(DOCKER_USERNAME) --password-stdin

push:
	docker tag $(PROJECT_NAME):latest $(DOCKER_TAG_NAME)
	docker push $(DOCKER_TAG_NAME)
