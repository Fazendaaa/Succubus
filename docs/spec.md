# Project

Before continuing please take a look at [What is Succubus?](what.md)

- [Project](#project)
  - [Specs](#specs)
  - [Objectives](#objectives)
  - [System](#system)
  - [Tasks](#tasks)
    - [Base](#base)
    - [Development](#development)
    - [Extended](#extended)
  - [Commands](#commands)

You can read more how can you take more advantage of it reading how the [container](container.md) containers works.

## Specs

```yaml
[optional] name: ...
[optional] tag: ...
[optional] image: ...
[optional] dockerfile: ...
[optional] interact: ...
[optional] context: ...
[optional] env_file: ...

[required] objectives:
...
[optional] system:
...
```

- `name`: Project name:
  - cannot contain upper case letters
  - cannot contain empty spaces
- `tag`: Project tag to be used as a versioning system:
  - only accepts [calendar versioning](https://en.wikipedia.org/wiki/Software_versioning#Date_of_release) or [semantic versioning](https://semver.org/)
- `image`: Docker image to be used:
  - it will check if available locally first, then will try to download if not available
  - string in the format `<registry-owner>/<image-name>`
- `dockerfile`: Dockerfile path or `false` value to not use the current defined Dockerfile in the Project folder
- `interact`: call to invoke the interpreter to run it inside the Project
- `context`: which directory to use as Project root
- `env_file`: file to load all environment definitions
- [objectives](#objectives): user defined actions to be taken in the current Project
- [system](#system): Docker-Compose definitions to expand and run the Project on it

## Objectives

```yaml
[required] objectives:
[required]    base:
...
[required]    dev:
...
[optional]    <extended>:
...
```

As pointed out before, `base` and `dev` definitions are required to properly defined a project, but the user has all the flexibility in the world to define them as long they are present.

## System

```yaml
...
[optional] system:
              <system_name>:
                path: /path/to/docker-compose.yml
                objectives:
                - objetive_name_one
                - objetive_name_two
...
```

Docker-Compose are a great and readable way to define a system, but they are limited to how to perform and debug it. Systems references allows us to properly couple Project definitions and run multiple scenarios without having to create and maintain many `docker-compose.yml`.

## Tasks

Tasks are the are the set of [Commands](#commands) in which every [Objective](#objectives) is defined upon.

### Base

```yaml
...
[required]    base:
[optional]        env_file: ...
[required]        run: ...
[required]        add: ...
[required]        rm: ...
[required]        test: ...
```

`Base` let us properly define the Tasks required to build a Project without having to figure out how to add a dependency, versioned it or even figure out if this same dependency has the minimum requirements like a compliance license needed in this Project.

### Development

```yaml
...
[required]    dev:
[optional]        env_file:
[required]        anal: ...
[required]        linter: ...
[required]        doc: ...
```

### Extended

```yaml
...
[optional]    <extended>:
[optional]        env_file:
[optional]        <anyname>: ...
...
```

## Commands

```yaml
...
base:
  run:
    env:
      TEST="foo"
    commands: echo "${TEST} bar"
...
```

```yaml
...
base:
  run:
    env:
      TEST="foo"
    commands: |
      echo "Hello, World"
      echo "${TEST} bar"
...
```
