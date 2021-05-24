# Succubus

## Configuration Description

```shell
          Project
|--------------------------|
|        Objectives        |
| |----------------------| |
| |        Tasks         | |
| | |------------------| | |
| | |      Rules       | | |
| | | |--------------| | | |
| | | |   Commands   | | | |
| | | |--------------| | | |
| | |------------------| | |
| |----------------------| |
|--------------------------|
```

### Obejctives

```yaml
[optional] image:
[optional] dockerfile:

[required] base:
...

[required] dev:
...

[optional] extended:
...
```

### Tasks

#### Base

```yaml
...
[required] base:
[required]     run: ...
[required]     add: ...
[required]     rm: ...
[required]     test: ...
```

#### Development

```yaml
...
[required] dev:
[required]     anal: ...
[required]     linter: ...
[required]     doc: ...
```

#### Extended

```yaml
...
[optional] extended:
[optional]     <anyname>: ...
...
```

### Rules

```yaml
...
base:
  run: echo "Hello, World"
  add: echo "Hello, World"
  rm: echo "Hello, World"
  test: echo "Hello, World"
...
```

### Commands

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
    rules: |
      echo "Hello, World"
      echo "${TEST} bar"
...
```

```yaml
...
base:
  run:
    env:
      TEST="foo"
    rules:
      hello:
        env: TEST="Hello"
        command: echo "${TEST}, World"
      echo "${TEST} bar"
...
```

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
