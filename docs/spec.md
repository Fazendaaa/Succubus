# Project

## Specs

```yaml
[optional] name:
[optional] image:
[optional] dockerfile:
[optional] context:
[optional] env_file:

[required] objectives:

[optional] system:
...
```

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

## System

```yaml
...
[optional] system: /path/to/docker-compose.yml
...
```

You can read more how can you take more advantage of it reading how the [container](container.md) containers works.

## Tasks

### Base

```yaml
...
[required]    base:
[optional]        env_file:
[required]        run: ...
[required]        add: ...
[required]        rm: ...
[required]        test: ...
```

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

## Rules

```yaml
...
base:
  run: echo "Hello, World"
  add: echo "Hello, World"
  rm: echo "Hello, World"
  test: echo "Hello, World"
...
```

## Commands

```yaml
...
base:
  run: |
    echo "Hello, World"
    echo "foo bar"
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
