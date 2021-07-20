# Project

## Specs

```yaml
[optional] name: ...
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

- `name`: nome do Projeto
- `image`: imagem Docker a ser utilizada
- `dockerfile`: caminho para o Dockerfile a ser utilizado no projeto ou `false` para não utilizar o Dockerfile presente no mesmo projeto
- `interact`: chamada de como invocar o interpretador a ser utilizado dentro do Projeto
- `context`: qual diretório a ser utilizado como base do Projeto
- `env_file`: arquivo para carregar todas as definições de variáveis de ambiente
- [objectives](#objectives): ações definidas pelo usuário a serem tomadas no presente Projeto
- [system](#system): definições de Docker-Compose a serem utilizados para expandir o horizonte de ambientes nos quais o Projeto podem se encontrar presentes

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
[optional] system:
              <system_name>:
                path: /path/to/docker-compose.yml
                objectives:
                - objetive_name_one
                - objetive_name_two
...
```

You can read more how can you take more advantage of it reading how the [container](container.md) containers works.

## Tasks

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

### Development

```yaml
...
[required]    dev:
[optional]        env_file: ...
[required]        anal: ...
[required]        linter: ...
[required]        doc: ...
```

### Extended

```yaml
...
[optional]    <extended>:
[optional]        env_file: ...
[optional]        <anyname>: ...
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
