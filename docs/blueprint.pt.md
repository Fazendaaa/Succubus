{{ template:title }}

{{ template:logo }}

{{ template:badges }}

{{ template:description }}

<div align = "center">

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](.README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](.README.pt.md)

</div>

# What if NPM had a baby with Docker-Compose?

> you can also see it as a "Makefile with steroids" or just as an "augmented layer" to your language/framework... Or simply: "the effortless way to do things right"

Welcome to Fazendaaa's {{ pkg.name }}. This is version {{ pkg.version }}!


## Idea

Currently, in the company that I work for we have a CLI (Command Line Interface) made in [Python](https://www.python.org/) called `estat` and you can read more about it [here](https://github.com/Fazendaaa/CFD).

## What is Succubus?

- Succubus is a language/framework agnostic project manager manifest
- It works by declaring objectives inside the given project
- Each objective is a sequence of tasks
- And a task is formed by determined rules

Succubus is used to maintain Succubus, as you can see it in the [succ.yaml](./succ.yaml) presented here. And the best part to start using it is that you don't even need to install it to run and maintain all of your projects made using a plethora of languages/frameworks and theirs respective versions.

## Components

An example to [Python](https://www.python.org/) + [Django](https://www.djangoproject.com/) project -- just take a look at the [examples](./examples) folder -- would look something like it:

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

Then just open your terminal and type:

```shell
succ run
```

To see your website running in your `http://localhost`.

**And you can do all all of this without needing to install:**

- Python
- Django
- Or even Succubus for that matter

You can read more about the Project segmentation idea right [here](https://fazendaaa.github.io/Succubus/).

## Advantages

- Using many languages and frameworks without needing to install and maintain any of it.
- Running any project without need to configure your OS to properly run it.
- Maintainability:
  - Low entrance bar helps a lot; helping a "one project" person to avoid conflicts during releases and deploys and in many people projects can reduce the time to a person to familiarize with the new language/framework because the Project usage interface is the same
  - You can run an old project with some old language/framework version without needing to install it in your machine, encapsulating old code and making it easy to run in new scenarios
  - Even though you might not be running the latest Node LTS version, in a few years that might not be the case anymore; so help your future self or even a coworker understand what you are doing it today in that awesome project that you are currently working on
- Time to Market:
  - Cloud-native approach helps not only to deploy a project down the line to a cloud provider but also helps a client to receive the technology transfer later on

## Use Cases

You can always check the more how the concept behind breaking a Project and defining a project works in the [description](./docs/README.md) of it; but just to name a few Objectives that you can define in your project besides the `base` and `dev` ones:

- security
- homologation
- benchmark
- [chaos gorilla](https://www.gremlin.com/chaos-monkey/the-simian-army/)
- etc.

Once you get comfortable with the idea of breaking things apart and making your thought process easier to reproduce, it also makes it easier for others to understand and help you out with it; making the engine run without any problems.

## Why not Docker-Compose?

- Compose is more tailored to provide a "System Defined" experience trough a manifest which is more suited to accommodate other projects besides the one you are currently working on, as a Data Base, Reverse-Proxy, Memory Cache and etc.
- Compose is not suited to perform many tasks, just "run" one which is defined to allow a whole system to work, not a project.

## Why "NPM like"?

- A few years ago every project that I did was in [Node](https://nodejs.org/en/) + [TS](https://www.typescriptlang.org/). When I moved to another languages like Python, [Julia](https://julialang.org/) and [R](https://www.r-project.org/), missing some CLI "sugar" features like those provided by [NPM](https://www.npmjs.com/) and [Yarn](https://yarnpkg.com/) was a big wake up call that I was in a *no man's land*.
- Those missing features help a project development and maintainability to become almost effortless when compared to the current available way; when you start a Node/NPM project you are just one *yay/apt/apk/brew/or any other system package manager* away from starting to code and just like Go or even Haskell, the code you write is loosely coupled to the OS/Distro that you wrote in it, this helps to avoid any project breaking after systems updates. That's a far fetch dream to other languages which are heavily coupled to the ambient that they are running on; this lack of cohesion and conciseness was the seed to start Succubus.

## Goals

- The idea of this project is to create an OS/language/framework independent flow of project development.
- At the same time the tool allows the users to implement their own patterns inside the standard flow.
- The mantra behind all of this is: *"help yourself by helping other to achieve greatness"*.
- As previously pointed out, ambients end up screwing development flow and as someone bearing a unfinished Computer Science (CS) degree usually I can handle myself during this kind of situation but my coworkers running their code in a Mac, Windows or some other Linux environment usually have lost multiple hours during a project trying to make their code run in each others machines to debug it, slowing the pace and projects deadlines.

## Do more doing less

- The usa of container under the hood speeds up every development cycle that I've been part of; specially if you are looking to deploy it later to [AWS](https://aws.amazon.com/), [Heroku](https://www.heroku.com/), [Azure](https://azure.microsoft.com/en-us/), [Linode](https://www.linode.com/), [GCP](https://cloud.google.com/) and many others; briding the gap between serverless and non-serverless environments.
- Don't fight against making your project accessible to others therefore also making it more reproducible also to others; your cloud provider doesn't know that you are rocking the latest M1 Mac, but they will understand the same "packaging" technology that hey are used to use it internally as many others industry whales work on.
- This might sound like a shout to the void, but hear me for a second:
  - *The cloud provider doesn't know how to configure your latest version of the bleeding edge compiler that you are using it*
  - *Your client doesn't care about your well cared 200 steps build process*
  - *And your internet doesn't have the same years in your back as you do and saw all that you did*
- If you are just a rookie in the whole **"Programming World"** and this might see a little bit *too much* for you or even scary because your [NumPy](https://numpy.org/) + [Pandas](https://pandas.pydata.org/) never talked about all of this, my message to you is *"give it a try"* -- please take a look at [CFD](https://github.com/Fazendaaa/CFD) and take a look at many other tools that might be of your interesting, specially [Jinn](https://github.com/Fazendaaa/Jinn) if you don't know where or how to start it

More than a *Jake of all Trades* platform this project has as its core the purpose of sharing knowledge/learning uses case as its main virtue -- learn from your mistakes then share it as this might help others rather then just yourself

## Usage

### init

Shows already made made manifestos for you to choose from and use in your project allowing you to maintain an already project using something that would better fit your need:

```shell
succ init
```

### repl

```shell
succ repl
```

### add

```shell
succ add [ pkg01 pkg02 pkg03=1.3 ... ]
```

### rm

```shell
succ rm [ pkg01 pkg02 pkg03 ... ]
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

## Engines

I know that Docker Inc is kinda of a "hot topic" in some circles, that's why this tool tries to maintain itself behind an *Engine* abstraction layer, allowing a east swap in replacement if need to do so

## Installing

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
- [I'll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/)

### Norms

- [Shell Grammar Rules](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_10_02)

{{ template:license }}
