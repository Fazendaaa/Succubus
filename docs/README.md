# Succubus

> The cloud-native way to code

## How does it work?

- Succubus uses a project driven approach to handle software development, making it clear and enforcing the rules to each and every project that it handles it.
- Succubus is expansible, so that you can add your own "taste" to it
- Works across many different OSes and their respective versions
- Reduces the time to develop projects and deploy them by running it the code inside a control environment during development

### Basics

![basics](https://raw.githubusercontent.com/Fazendaaa/Succubus/master/assets/img/basics.jpg)

Just to better explain it to you:

1. You write a program -- any language, any framework, any version
2. You call Succubus to handle any request that you issued
3. Succubus looks up this request in the provided manifest then reads up the code
4. Executes the invoked command in a container

As you saw you can think in Succubus as a "broker" to your project:

- It handles all the project details while you only worry to issue requests to it
- It manages to follow appropriate defaults
- Reduces variabilities in software development because the container is a industry standard technology

You can develop your own patterns inside the manifest standard and share them across projects; or just use one provided by a third party.

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
