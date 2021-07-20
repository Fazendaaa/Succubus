all: doc

doc:
	@docker buildx build \
		--file docs/Dockerfile \
		--platform linux/amd64 \
		--load --tag appnest \
		docs/
# Jerry Rig
	@docker run --rm \
		--volume $(shell pwd)/docs:/home/node/app \
		--workdir /home/node/app appnest \
		npx @appnest/readme generate -o render.md
	@docker run --rm \
		--volume $(shell pwd)/docs:/home/node/app \
		--workdir /home/node/app appnest \
		npx @appnest/readme generate -i blueprint.pt.md -o render.pt.md
	@cp docs/render.md README.md
	@cp docs/render.pt.md README.pt.md
	@rm docs/render.*
