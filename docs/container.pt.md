# Container

When developing a project using Succubus, a little disclaimer before pointing out other things:

> Succubus is not Docker-Compose

Even tho they share some resemblance, they are different tools made to account and handle different purpose:

- Succubus use network as a shared resource with the host machine
- Succubus only sees the current directory and its contents
- Succubus doesn't handle volumes

If you want to do those kinda of things, please take a look at this kinda of resources:

- [Docker-Compose docs](https://docs.docker.com/compose/)
- [Educative's Docker Compose Tutorial](https://www.educative.io/blog/docker-compose-tutorial)
- [Microsoft's Use Docker Compose](https://docs.microsoft.com/en-us/visualstudio/docker/tutorials/use-docker-compose)

## Dockerfile

<!-- Explain:
1. Layers
2. Caching
3. Multi-stage
-->

### Default

### Multi-staged

<!-- Explain:
1. System tag
2. Runner tag
3. Everything in between
-->

#### System

#### Runner
