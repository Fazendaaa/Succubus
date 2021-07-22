# Projeto

Antes de continuar, por favor dê uma olhada em [O que é Succubus?](what.pt.md)

- [Projeto](#projeto)
  - [Specs](#specs)
  - [Objetivos](#objetivos)
  - [Sistema](#sistema)
  - [Tarefas](#tarefas)
    - [Base](#base)
    - [Desenvolimento](#desenvolimento)
    - [Estendidas](#estendidas)
  - [Comandos](#comandos)

## Specs

```yaml
[optional] name: ...
[optional] tag: ...
[optional] image: ...
[optional] dockerfile: ...
[optional] interact: ...
[optional] context: ...
[optional] env_file: ...
[optional] init: ...

[required] objectives:
...
[optional] system:
...
```

- `name`: nome do Projeto:
  - não pode conter letras em maiúsculo
  - não pode conter espaços em branco
- `tag`: marcador do Projeto a ser utilizada como `dev`, `hom` ou mesmo `prod`:
  - não pode conter letras em maiúsculo
  - não pode conter espaços em branco
- `version`: vers'ao do Projeto a ser usada:
  - aceita apenas [versionamento de calendario](https://en.wikipedia.org/wiki/Software_versioning#Date_of_release) ou [versionamento semântico](https://semver.org/)
- `image`: imagem Docker a ser utilizada:
  - irá verificar primeiro se a imagem está disponível localmente, do contrário irá procurar baixar ela
  - string no formato `<registry-owner>/<image-name>`
- `dockerfile`: caminho para o Dockerfile a ser utilizado no projeto ou `false` para não utilizar o Dockerfile presente no mesmo projeto
- `interact`: chamada de como invocar o interpretador a ser utilizado dentro do Projeto
- `context`: qual diretório a ser utilizado como base do Projeto
- `env_file`: arquivo para carregar todas as definições de variáveis de ambiente
- `init`: comando a ser seguido quando o container iniciar pela primeira vez
- [objectives](#objectives): ações definidas pelo usuário a serem tomadas no presente Projeto
- [system](#system): definições de Docker-Compose a serem utilizados para expandir o horizonte de ambientes nos quais o Projeto podem se encontrar presentes

## Objetivos

```yaml
[required] objectives:
[required]    base:
...
[required]    dev:
...
[optional]    <extended>:
...
```

## Sistema

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

## Tarefas

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

### Desenvolimento

```yaml
...
[required]    dev:
[optional]        env_file: ...
[required]        anal: ...
[required]        linter: ...
[required]        doc: ...
```

### Estendidas

```yaml
...
[optional]    <extended>:
[optional]        env_file: ...
[optional]        <anyname>: ...
...
```

## Comandos

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
