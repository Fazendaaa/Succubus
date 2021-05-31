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

Apenas para deixar mais claro -- *Eu sei, é um desenho bem feio*:

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

Which can be accessed by just running the following command:

```shell
succ run
```

In your terminal in a Python + Django project to open a website running and start using containers:

- No Dockerfile required
- No Docker-compose required
- Not even Succubus installed if you want to

The whole idea is a effortless software development experience while providing safety and speed of development.

## Possibilities

The main idea behind this project is to open possibilities to encapsulate that projects that you are working on just like we are doing at the company that I work on -- you can take a look at this right [here](https://github.com/Fazendaaa/CFD) --; but, besides that, you can ask yourself:

- Would you like to run your project in 20 years like you run it today?
- Would you like to avoid that one project "intoxicate" another?
- Would you like to run the project in Windows, Mac and Linux without any issue?
- Would you like to make your work more accessible?

If you said *"aye"* to any of the previous of the questions you want Succubus. But, besides this, the idea is create a platform to easily, securely and fast patterns development and testing so you can use it yourself in a project or share with others -- maybe even landing a hand and help others understand better how to make a awesome project.

And you can use it in any language, cross-compile to any architecture afterwards because you will be running running containers under the hood; so if you are rocking that last M1 Mac you can easily share your work with your coworkers or even deploy it to the cloud.

## Project

The concept behind the whole "manifest" idea is to make it easier to deploy a new project trough a familiar interface so the programmer can work on it.

### Overview

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

The idea is to develop a project based in a set of rules:

1. A Project is formed by Objectives
2. Those Objectives are formed by Tasks
3. Those Tasks are a set of Rules
4. And each Rule is a sequence of Commands

You can look at this like a [docker-compose](https://docs.docker.com/compose/) joining forces with [npm](https://www.npmjs.com/) to help maintain healthy a project.

## Modelo de Corretor

Sotaques de código

As you saw you can think in Succubus as a "broker" to your project:

- It handles all the project details while you only worry to issue requests to it
- It manages to follow appropriate defaults
- Reduces variabilities in software development because the container is a industry standard technology

## TODO

If you read this far, please take a look at the [repo](https://github.com/Fazendaaa/Succubus) and maybe to drop a star in the project, give me some feedback or even dropping some of your code in it :)
