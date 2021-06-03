# Succubus

> The cloud-native way to code

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.pt.md)

## What is it?

- Succubus uses a project driven approach to handle software development trough a "Broker Model", making it clear and enforcing the rules to each and every project that it handles it.
- Succubus is expansible, so that you can add your own "flavor" to it
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
interact: python3

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

And if you need to interact with something just to see whether or not will work, just:

```shell
succ interact
```

A screen like the following should pop up:

```shell
Python 3.9.5 (default, May  9 2021, 14:00:28)
[GCC 10.2.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

And then you will have a REPL (Read-Evaluate-Print-Loop) experience without ever needing to install Python and the great thing about it is that later on this will be the same environment shipped to a production environment.

# Index

- [Succubus](#succubus)
  - [What is it?](#what-is-it)
    - [Basics](#basics)
- [Index](#index)
  - [Security and Reliability](#security-and-reliability)
  - [Universe](#universe)
  - [Possibilities](#possibilities)
  - [Project Manifest](#project-manifest)
    - [Overview](#overview)
  - [Broker Model](#broker-model)
    - [Why broker?](#why-broker)
  - [Pipe Production Environment](#pipe-production-environment)
  - [TODO](#todo)
  - [Appendix](#appendix)
    - [DevOps](#devops)
    - [Cloud-Native](#cloud-native)
    - [JIT](#jit)

## Security and Reliability

- Succubus avoids that any OS related variable make your project bound to it
- At the same time Succubus avoids that any Project related configuration ends up breaking your OS
- Easy to start any current or new Project without having to worry about prior knowledge about any system requirements -- **as long you can run containers in it**
- [Least-Privilege](https://docs.microsoft.com/en-us/windows-server/identity/ad-ds/plan/security-best-practices/implementing-least-privilege-administrative-models) by default in any working Project
- As David Patterson once said that he didn't believed in [security by obscurity](https://cacm.acm.org/magazines/2019/2/234352-a-new-golden-age-for-computer-architecture/fulltext) all the variables are made clear and the process easily reproducible
- No-locking and easy of usage to debug, change platforms, reproduce environments and replace vendors being compilers, packages or even Project technology itself when refactoring

## Universe

And the best part is that this works as well in projects based in:

- [Node](https://hub.docker.com/_/node)
- [Go](https://hub.docker.com/_/golang)
- [Ruby](https://hub.docker.com/_/ruby)
- [PHP](https://hub.docker.com/_/php)
- [Java](https://hub.docker.com/_/openjdk)
- [Perl](https://hub.docker.com/_/perl)
- [Erlang](https://hub.docker.com/_/erlang)
- [Rust](https://hub.docker.com/_/rust)
- [Clojure](https://hub.docker.com/_/clojure)
- [Swift](https://hub.docker.com/_/swift)
- [C](https://hub.docker.com/_/gcc)
- [Haskell](https://hub.docker.com/_/haskell)
- [R](https://hub.docker.com/_/r-base)
- [Dart](https://hub.docker.com/_/dart)
- And many, many, others

The possibilities are endless, you can take a look at [Docker Hub](https://hub.docker.com/search?type=image) to look at more community driven options or you can even roll up your sleeves and make your own language more accessible to other trough this -- *or even use a non public available technology just fine because Succubus works the best to help you out, even in a protective corporative environment*

And making the [Project Manifest](#project-manifest) a many-to-one way of development, you can rise your productiveness without having to worry about the context-switching overhead related to

## Possibilities

The main idea behind this project is to open possibilities to encapsulate that projects that you are working on just like we are doing at the company that I work on -- you can take a look at this right [here](https://github.com/Fazendaaa/CFD) --; but, besides that, you can ask yourself:

- *Would you like to run your project in 20 years like you run it today?*
- *Would you like to avoid that one project "intoxicate" another?*
- *Would you like to run the project in Windows, Mac and Linux without any issue?*
- *Would you like to make your work more accessible?*

If you said *"aye"* to any of the previous of the questions you want Succubus. But, besides this, the idea is create a platform to easily, securely and fast patterns development and testing so you can use it yourself in a project or share with others -- *maybe even landing a hand and help others understand better how to make a awesome project*.

And you can use it in any language, cross-compile to any architecture afterwards because you will be running running containers under the hood; so if you are rocking that last M1 Mac you can easily share your work with your coworkers or even deploy it to the cloud.

## Project Manifest

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

### Why broker?

A [broker](https://dictionary.cambridge.org/dictionary/english/broker) according to Cambridge Dictionary:

> a person who buys and sells foreign money, shares in companies, etc. [...]:
> "I called my broker for advice about investing in the stock market"

That being said, Succubus acts like a [broker](https://www.betterteam.com/broker-job-description) to a programmer:

- It handles the nitty-gritty that the common programmer usually are not aware of but that doesn't mean that it should not be followed
- Process all the required transactions by the rule
- Shows great knowledge about the market and how it behaves
- Helps you grow and understand better the discussed subject
- Allows some tweaks to proper fit the matter at hand to your best interest

That definition put together with the whole [Project Manifest](#project-manifest) concept, allow the broker -- Succubus -- to proper operate by the given guidelines.

## Pipe Production Environment

Bringing the whole [Just In Time (JIT)](http://people.brunel.ac.uk/~mastjjb/jeb/or/jit.html) production philosophy to the "code development" environment -- please don't mistake the whole [JIT](https://www.freecodecamp.org/news/just-in-time-compilation-explained/) used in compilers with this one --; making the whole pipe process more efficient by default by avoiding late binding or some late checks down the process. This also allows to easily reproduce the whole "production" environment in your machine but looking only at a part of a service, a project; this allows us to better understand and made a translation of some sorts of the microservices idea but in looking even deeper in in your machine, kinda like a *behavior-service* approach.

This approach is a tool to enforce some well-know software developments like:

- [Domain Driven Development](https://en.wikipedia.org/wiki/Domain-driven_design)
- [Test Drive Development](https://en.wikipedia.org/wiki/Test-driven_development)
- [Service Oriented Delivery](https://www.enterprise-architecture.com/docs/deliverSOARealiseEA.pdf)
- And many others

And you can see this without fear of having a too fine-grained control that overweights your whole process; a one person project cannot handle the complexity of a 500 employers company software development pipe process; but both of them can increase or decrease the granularity measuring during the process whether or not they will provide a measurable **return of investment** to you, reduce your **total cost of owning** this process or even trying many workflows without having to changing any code to check a **hypothesis test**.

## TODO

If you read this far, please take a look at the [repo](https://github.com/Fazendaaa/Succubus) and maybe to drop a star in the project, give me some feedback or even dropping some of your code in it :)

## Appendix

### DevOps

- [Addressing The Detrimental Effects of Context Switching With DevOps](https://insights.sei.cmu.edu/blog/addressing-the-detrimental-effects-of-context-switching-with-devops/)

### Cloud-Native

- [O Inferno de Dante do Cloud Native](https://fazenda.hashnode.dev/o-inferno-de-dante-do-cloud-native) -- my rant about all of this """new""" development wave, unfortunately I still haven't publish it in English, probably later on I will do so

### JIT

- [Sennheiser Factory Tour - Hanover, Germany](https://youtu.be/5es8zggYM7A)
