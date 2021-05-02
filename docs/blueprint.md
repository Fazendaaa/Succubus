{{ template:title }}

{{ template:logo }}

{{ template:badges }}

{{ template:description }}

Welcome to Fazendaaa's {{ pkg.name }}. This is version {{ pkg.version }}!

Made with:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

## Ideia

Currently, in the company that I work for we have a CLI (Command Line Interface) made in [Python](https://www.python.org/) called `estat` that works like the following:

```shell
estat add ggplot2 Yang-Tang/shinyjqui VGAM jbryer/DTedit
```

This works in a [R](https://www.r-project.org/); if you know [Node.js](https://nodejs.org/en/) you may find it similar as [npm](https://www.npmjs.com/) which this tool is heavily based on. The CLI will fetch the given packages and add them to the current DESCRIPTION file -- which is the equivalent `package.json`, listing the packages added to the project.

So far, so good. But the turning point is the following: **I DON'T HAVE R INSTALLED IN MY MACHINE**

All the company projects have a `Dockefile` in it; so the CLI does the following:

1. Build the current Dockerfile as a helper container
2. Bind the current project to the helper container, installs the mentioned packages in it
3. Adds the packages to the DESCRIPTION file truncating its version

The main advantage of this approach is to help avoid errors during development, especially because in any project you might find yourself using:

1. A different language version
2. Different Operational Systems (OSes)
3. Forgetting to list all packages used in the project
4. Also forgetting to list OS-based dependencies
5. And also forgetting to list environment variables
6. etc.

The idea is for the development process to be *frictionless*; as all the developers can switch branches and projects and start to work on them without having to worry to install any dependency or figuring out what the other developer did to that project to work.

Besides this, for that matter, this will work as well for other projects based in:

- Python
- Julia
- npm/yarn -- Node.js, Vue, React, etc.
- etc.

As the main ideia is to work as a abstraction layer to handle projects the language, its version, and OS related packages should not matter only the what is described in the Dockerfile.

And as npm allows us to run the project and test it so `estat` does it as well. Each project that contains a `docker-compose.yml` file is executed when running a simple `estat run`, avoiding that developers no familizired with [docker-compose](https://docs.docker.com/compose/) forget to add the `--build` flag to it and, besides that, the tool also goes trough the process of checking each listed image in it for updates as also checking the base image for updates before running it -- sometimes developers forget to check them for updates that might fix their problems.

As `estat` have grown so much and making it available as FOSS (Free and open-source software) was always the idea but the project still in development and not having a properly defined scope, I decided to break its main features in other projects:

- [Succubus](https://github.com/Fazendaaa/Succubus): universal package manager based on cloud-native
- [Jinn](https://github.com/Fazendaaa/Jinn): universal project manager built to expand Succubus capabilities
- [Baba Yaga](https://github.com/Fazendaaa/BabaYaga): universal cloud-native manager built to expand Jinn and Succubus capabilities
- [Wendigo](https://github.com/Fazendaaa/Wendigo): universal project translator from cloud-native projects to other infra technologies
- [Shōjō](https://github.com/Fazendaaa/Shojo): LaTex package manager

## Components

Succubus goes trough the `succ.yaml` file looking for the commands to run when doing one of the following tasks, working as a "manifesto". If they are not described or the file it's not presented the program will follow default steps to each language based on Jinn workflow.

### init

```shell
succ init
```

### add

```shell
succ add [ pkg01 pkg02 pkg03=1.3 githubUser/pkg ... ]
```

### run

```shell
succ run
```

### test

```shell
succ test
```

## Installing

You don't need to install Go to run this tool, just Docker. And to do so to give it a try, you can do it just by running the following line in your terminal:

```shell
alias succ='docker run -it --volume $(pwd):/project --workdir /project fazenda/succubus'
```

And then running the following to check whether or not is working properly:

```shell
succ --help
```

## Author

Only [me](https://github.com/Fazendaaa) because the aforementioned project was implemented by yours only. By knowing each line of that code wrote doing the port would be more easily done this way.

## Contributing

Check more about this in [CONTRIBUTING.md](CONTRIBUTING.md). Here we have a list of some of our contributors:

{{ template:contributors }}

## TODO

## References

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

{{ template:license }}
