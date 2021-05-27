{{ template:title }}

{{ template:logo }}

{{ template:badges }}

{{ template:description }}

# What if NPM had a baby with Docker-Compose?

> or just a Makefile with steroids

Welcome to Fazendaaa's {{ pkg.name }}. This is version {{ pkg.version }}!

## Ideia

Currently, in the company that I work for we have a CLI (Command Line Interface) made in [Python](https://www.python.org/) called `estat` and you can read more about it [here](https://github.com/Fazendaaa/CFD).

## What is Succubus?

- Succubus is a language/framework agnostic project manager manifest
- It works by declaring objectives inside the given project
- Each objective is a sequence of tasks
- And a task is formed by determined rules

Succubus is used to maintain Succubus, as you can see it in the [succ.yaml](./succ.yaml) presented here. And the best part to start using it is that you don't even need to install it to run and maintain all of your projects made using a plethora of languages/frameworks and theirs respective versions.

## Components

An example to [Python](https://www.python.org/) + [Django](https://www.djangoproject.com/) project would look something like it:

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

Then just open your terminal and type:

```shell
succ run
```

To see your website running in your `http://localhost`.

## Advantages

- Using many languages and frameworks without needing to install and maintain any of it
- Running any project without need to configure your OS to properly run it
- Maintainability:
  - Low entrance bar helps a lot; helping a "one project" person to avoid conflicts during releases and deploys and in many people projects can reduce the time to a person to familiarize with the new language/framework because the Project usage interface is the same
  - 
- Time to Market:
  - 

## Why not Docker-Compose?

- Compose is more tailored to provide a "System Defined" experience trough a manifest which is more suited to accommodate other projects besides the one you are currently working on, as a Data Base, Reverse-Proxy, Memory Cache and etc.
- Compose is not suited to perform many tasks, just "run" one which is defined to allow a whole system to work, not a project

## Usage

### init

Shows already made made manifestos for you to choose from and use in your project allowing you to maintain an already project using something that would better fit your need:

```shell
succ init
```

### add

```shell
succ add [ pkg01 pkg02 pkg03=1.3 ... ]
```

### rm

```shell
succ add [ pkg01 pkg02 pkg03=1.3  ... ]
```

### run

```shell
succ run
```

### test

```shell
succ test
```

## Running

You don't need to install Go to run this tool, just Docker. And to do so to give it a try, you can do it just by running the following line in your terminal:

```shell
alias succ='docker run -it --volume $(pwd):/project --workdir /project fazenda/succubus'
```

And then running the following to check whether or not is working properly:

```shell
succ --help
```

This approach has some limitations but is a great way to tip your toes and start right way using the tool; if your needs aren't meet by it, you can always [install](#installing) the tool.

## Installing

## How does it work?

### Basics

Let me paint a picture for you: *"When you access a remote desktop all the things you did stay there, right?"*... **YES**

You can see the

### Intermediate

<!--
1. Multiarch
2. Cloud native
3. Docker-compose
4. K8s
5. Making CI/CD accessible and the standard
-->

### Advanced

<!--
1. Multistage-build as cache:
   1. system
   2. everything in between
   3. runner
2. Clarity first:
   1. explicit usage of env vars -- isn't inheritance/overload if you don't know what are you doing
-->

## Author

Only [me](https://github.com/Fazendaaa) because the aforementioned project was implemented by yours only. By knowing each line of that code wrote doing the port would be more easily done this way.

## Contributing

Check more about this in [CONTRIBUTING.md](CONTRIBUTING.md). Here we have a list of some of our contributors:

{{ template:contributors }}

## TODO

## References

### Repositories

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [devspace/awesome-github-templates](https://github.com/devspace/awesome-github-templates)

### Blog-posts

- [Static code analysis for golang](https://levelup.gitconnected.com/static-code-analysis-for-golang-5f24b555d227)
- [Compilers 101 - Overview and Lexer](https://dev.to/lefebvre/compilers-101---overview-and-lexer-3i0m)
- [THE BASICS: 7 Alternatives to Docker: All-in-One Solutions and Standalone Container Tools](https://jfrog.com/knowledge-base/the-basics-7-alternatives-to-docker-all-in-one-solutions-and-standalone-container-tools/)

### Norms

- [Shell Grammar Rules](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_10_02)

{{ template:license }}
