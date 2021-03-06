build-docs:
	@docker buildx build \
		--file docs/Dockerfile \
		--platform linux/amd64 \
		--push --tag fazenda/appnest \
		docs/

doc:
	@docker run --rm \
		--volume ${PWD}:${PWD} \
		--workdir ${PWD} fazenda/appnest \
		--package docs/package.json \
		--config docs/blueprint.json \
		--input docs/blueprint.pt.md \
		--output README.pt.md
	@docker run --rm \
		--volume ${PWD}:${PWD} \
		--workdir ${PWD} fazenda/appnest \
		--package docs/package.json \
		--config docs/blueprint.json \
		--input docs/blueprint.md \
		--output README.md

run:
	@go run ./main.go

docker:
	@docker buildx build \
		--platform linux/amd64 \
		--load --tag succubus \
		.
	@docker run -it \
		--volume ${PWD}:${PWD} \
		--workdir ${PWD} \
		succubus
