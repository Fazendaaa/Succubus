# Succubus

> The cloud-native way to code

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.pt.md)

## How does it work?

- Succubus uses a project driven approach to handle software development trough a "Broker Model", making it clear and enforcing the rules to each and every project that it handles it.
- Succubus is expansible, so that you can add your own "taste" to it
- Works across many different OSes and their respective versions
- Reduces the time to develop projects and deploy them by running it the code inside a control environment during development

### Basics

![basics](https://raw.githubusercontent.com/Fazendaaa/Succubus/master/assets/img/basics.jpg)

> I know, it's a pretty ugly drawing

Just to make it clear:

1. You write a program -- any language, any framework, any version
2. You call Succubus to handle any request that you issued
3. Succubus looks up this request in the provided manifest then reads up the code
4. Executes the invoked command in a container

You can develop your own patterns inside the manifest standard and share them across projects; or just use one provided by a third party.

A example of this manifest is the following:

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

- No Python required
- No Django required
- No Dockerfile required
- No Docker-compose required
- Not even Succubus installed if you want to

The whole idea is a effortless software development experience while providing safety and speed of development.

## Possibilities

The main idea behind this project is to open possibilities to encapsulate that projects that you are working on just like we are doing at the company that I work on -- you can take a look at this right [here](https://github.com/Fazendaaa/CFD) --; but, besides that, you can ask yourself:

- *Would you like to run your project in 20 years like you run it today?*
- *Would you like to avoid that one project "intoxicate" another?*
- *Would you like to run the project in Windows, Mac and Linux without any issue?*
- *Would you like to make your work more accessible?*

If you said *"aye"* to any of the previous of the questions you want Succubus. But, besides this, the idea is create a platform to easily, securely and fast patterns development and testing so you can use it yourself in a project or share with others -- *maybe even landing a hand and help others understand better how to make a awesome project*.

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

The core idea is to develop a project based in a set of concepts:

1. A Project is formed by Objectives
2. Those Objectives are formed by Tasks
3. Those Tasks are a set of Rules
4. And each Rule is a sequence of Commands

You can look at this like a [docker-compose](https://docs.docker.com/compose/) joining forces with [npm](https://www.npmjs.com/) to help maintain healthy a project.

## Broker Model

As you saw you can think in Succubus as a "broker" to your project:

- It handles all the project details while you only worry to issue requests to it
- It manages to follow appropriate defaults
- Reduces variabilities in software development because the container is a industry standard technology

## TODO

If you read this far, please take a look at the [repo](https://github.com/Fazendaaa/Succubus) and maybe to drop a star in the project, give me some feedback or even dropping some of your code in it :)
