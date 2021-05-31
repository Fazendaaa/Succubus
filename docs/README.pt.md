# Succubus

> A maneira cloud-native de codar

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.pt.md)

## Como funciona?

- Succubus usa uma abordagem de desenvolvimento de software baseada em projetos através de um "Modelo de Corretor", o tornando claro e enforçando as regras para cada projeto que gerencia
- Succubu é expansível, para que você possa adicionar suas próprias premissas nele
- Funciona através de vários SOs e suas respectivas versões
- Reduz o tempo de desenvolvimento de projetos e entregas por rodar todo o código dentro de um ambiente controlado durante desenvolvimento

### Fundamentos

![basics](https://raw.githubusercontent.com/Fazendaaa/Succubus/master/assets/img/basics.jpg)

> Eu sei, é um desenho bem feio

Apenas para deixar mais claro:

1. Escreva um programa -- em qualquer linguagem, framework, versão
2. Você chama Succubus para manusear qualquer requisição que fizer
3. Succubus procura tal requisição no manifesto proveniente no projeto e lê o código dele
4. Executa tal requisição dentro do container

Você pode desenvolver seus próprios "sotaques" dentro do manifesto e padronizar eles para poder compartilhar-los entre projetos; ou até mesmo utilizar um já feito por outros.

Um exemplo disso seria o seguinte manifesto:

```yaml
image: python

objectives:
  base:
    run: python3 manage.py runserver
    test: python3 -m pytest .
    add: pip3 install $ARGV
    rm: pip3 uninstall $ARGV

  dev:
    anal: |
      python3 -m mccabe --min 5 ./src
      python3 -m bandit -r ./src
    linter: python3 -m mypy ./src
    doc: python3 -m sphinx-apidoc -o source ../
```

O qual os comandos podem ser acessados rodando o seguinte comando:

```shell
succ run
```

No seu terminal verá a mensagem do projeto de Python + Django rodando e acessível pelo seu navegador; rodando containers desta maneira:

- Sem precisar instalar o Python
- Sem precisar instalar o Django
- Sem precisar de Dockerfile
- Sem precisar de Docker-Compose
- Sem até mesmo precisar do Succubus instalado se quiser

A ideia toda por trás de desenvolvimento de software sem dores de cabeça é melhorar a experiência do progamador e prover segurança com velocidade.

## Possibilidades

A principal ideia por trás deste projeto é abrir possibilidades para encapsulamento de projetos que você já está trabalhando assim como foi feito na empresa que trabalho -- você pode saeber mais [aqui](https://github.com/Fazendaaa/CFD) --; mas, além disso, você pode se perguntar algumas coisas:

- *Você gostaria de rodar em 20 anos os projetos como roda agora?*
- *Você gostaria de evitar que um projeto "intoxicasse" outro?*
- *Você gostaria de rodar seu proeto no Windows, Mac e Linux sem problemas?*
- *Você gostaria de deixar seu trabalho mais acessível?*

Se você diss *"sim"* para alguma das perguntas anteriores, você quer o Succubus. E, além disso, a ideia de criar de é de criar uma plataforma de fácil, seguro e rápido acesso para desenvolvimento desses "sotaques" de projeto para você poder evoluir de maneria rápida através de cada tentativa no projeto e/ou compartilhar com outros o que fez -- *ou até mesmo dando uma "mão amiga" e ajudar outros a entender melhor como fazer um projeto de estourar a boca do balão*

E o melhor é que você pode usar em qualquer linguagem, compilar para toda arquitetura depois porque você vai estar rodando containers debaixo do capô; se você estiver o M1 Mac você pode facilmente compartilhar o seu trabalho com outros colaboradores e até mesmo publicar ele na nuvem.

## Projeto

O conceito por trás de toda a ideia de "manifesto" é de facilitar a publicação de um novo projeto através de uma interface familiar para qual o programador consegue trabalhar com.

### Overview

```shell
          Projeto
|--------------------------|
|        Objetivos         |
| |----------------------| |
| |       Tarefas        | |
| | |------------------| | |
| | |     Regras       | | |
| | | |--------------| | | |
| | | |   Comandos   | | | |
| | | |--------------| | | |
| | |------------------| | |
| |----------------------| |
|--------------------------|
```

O core da idea é desenvolver um projeto baseado no conceito de:

1. Um Projeto é formado por Objetivos
2. Esses Objetivos são formados por Tarefas
3. Essas Tarefas são um conjunto de Regras
4. E cada uma dessas Regras são uma sequência de Comandos

Você pode ver isso como se olha para um [docker-compose](https://docs.docker.com/compose/) juntando forças com o [npm](https://www.npmjs.com/) para manter um projeto saúdavel.

## Modelo de Corretor

Como você viu, o Succubus pode ser visto como um "Corretor" para o seu projeto:

- Dá conta dos detalhes enquanto você só tem que se preocupar com fazer pedidos para ele
- Supervisiona os padrões definidos
- Reduz as variabilidades no desenvolvimento de software porque a indústria segue containers como um padrão

## A Fazer

Se você chegou até aqui, por favor dê uma olhada no [repositório](https://github.com/Fazendaaa/Succubus) do projeto e talvez dê um like, feedback ou até mesmo melhoria de código nele :)
