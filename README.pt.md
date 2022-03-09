<!-- ⚠️ This README has been generated from the file(s) "docs/blueprint.pt.md" ⚠️--><h1 align="center">Succubus</h1>

<p align="center">
  <img src="https://raw.githubusercontent.com/Fazendaaa/Succubus/master/assets/img/logo.svg" alt="Logo" width="150" height="150" />
</p>

<p align="center">
		<a href="https://golang.org/"><img alt="Made With Go" src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg?style=flat-square" height="20"/></a>
<a href="https://www.docker.com"><img alt="Made With Docker" src="https://img.shields.io/badge/Made%20with-Docker-important?style=flat-square" height="20"/></a>
<a href="https://code.visualstudio.com/"><img alt="Made With VSCode" src="https://img.shields.io/badge/made%20with-vscode-blueviolet?style=flat-square" height="20"/></a>
<a href="https://codeclimate.com/github/Fazendaaa/Succubus/maintainability"><img alt="Maintainability" src="https://api.codeclimate.com/v1/badges/f4f7ed3bf34aa8c570b7/maintainability" height="20"/></a>
<a href="https://codeclimate.com/github/Fazendaaa/Succubus/test_coverage"><img alt="Test Coverage" src="https://api.codeclimate.com/v1/badges/f4f7ed3bf34aa8c570b7/test_coverage" height="20"/></a>
<a href="https://www.codacy.com/gh/Fazendaaa/Succubus/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Fazendaaa/Succubus&amp;utm_campaign=Badge_Grade"><img alt="Codacy Badge" src="https://app.codacy.com/project/badge/Grade/0b27f064a67d411a872719802ab140dc" height="20"/></a>
	</p>


<p align="center">
  <b>A universal project manager based on cloud-native</b></br>
  <sub><sub>
</p>

<br />


<div align = "center">

[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg?longCache=true&style=for-the-badge)](https://saythanks.io/to/lucas.carotta%40outlook.com)

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](./README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](./README.pt.md)

</div>


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#como-seria-se-o-npm-tivesse-um-beb-com-o-docker-compose)

# ➤ Como seria se o NPM tivesse um bebê com o Docker-Compose?

> você pode também ver como um "Makefile com esteróides" ou apenas uma "camada de expansão" para as suas linguagens/frameworks... Ou simplesmente: "a maneira sem esforço de fazer as coisas certas"

Bem vindo ao pacote Succubus de Fazendaaa. Esta é a versão 0.0.0!


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#ideia)

## ➤ Ideia

Atualmente na empresa que eu trabalho nós possuímos uma CLI (Command Line Interface) feita em [Python](https://www.python.org/) chamada `estat` e você pode ler mais sobre ela [aqui](https://github.com/Fazendaaa/CFD).


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#o-que--succubus)

## ➤ O que é Succubus?

- Succubus é um manifesto agnostica de linguagem/frameworks para gerencia e desenvolvimento de projetos
- Ela funciona por declarar objetivos dentro de um projeto
- Cada objetivo é uma sequência de tarefas
- E cada tarefa é formada por determinadas regras
- Suporte para sua [empresa](docs/company.pt.md)/seus projetos poderem ter a sua versao dela que atenda melhor as suas necessidades

Succubus é utilizado para manter Succubus, você pode ver no arquivo [succ.yaml](./succ.yaml) aqui declarado. E a melhor parte é que você pode utilizar ele sem precisar instalá-lo e as respectivas linguagens, frameworks e pacotes que irá utilizar.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#componentes)

## ➤ Componentes

Um exemplo para um projeto [Python](https://www.python.org/) + [Django](https://www.djangoproject.com/) -- dê uma olhada na pasta [examples](./examples) -- seria algo mais ou menos assim:

```yaml
image: python
interact: python

objectives:
  base:
    run: python manage.py runserver
    test: python -m pytest .
    add: pip3 install $ARGV
    rm: pip3 uninstall $ARGV

  dev:
    anal: |
      python -m mccabe --min 5 ./src
      python -m bandit -r ./src
    linter: python -m mypy ./src
    doc: python -m sphinx-apidoc -o source ../
```

Apenas abra o seu terminal e digite:

```shell
succ run
```

E veja seu site rodando em `http://localhost`.

**O melhor de tudo é que você pode fazer isso sem precisar instalar:**

- Python
- Django
- Ou até mesmo Succubus

Você pode ler mais sobre toda essa ideia de segmentação [aqui](https://fazendaaa.github.io/Succubus/).


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#vantagens)

## ➤ Vantagens

- Usar várias linguagens e frameworks sem precisar instalar ou manter elas.
- Rodar qualquer projeto sem ter que configurar seu Sistema Operacional (SO) para poder rodar ela.
- Manutenabilidade:
  - Baixa barreira de entrada ajuda bastante; ajudar em "projetos de uma pessoa" evitar conflitos durante lançamentos e entregas, além de ajudar projetos de várias pessoas ao reduzir o tempo que uma nova pessoa a se familiarizar com a nova linguagem/framework porque o Projeto usa uma interface comum de desenvolvimento
  - Você pode rodar um projeto antigo em alguma versão antiga da linguagem/framework sem precisar instalar na sua máquina, encapsulando código antigo e facilitando o uso dele em novos cenários
  - Mesmo se você não tiver utilizando a versão LTS do Node, em alguns anos este não pode ser o mesmo cenário; ajudar o futuro você ou até mesmo um colega de trabalho a entender o que está fazendo hoje neste projeto incrível que está trabalhando
- Time to Market:
  - Abordagem cloud-native ajuda não apenas a entregar um projeto para rodar em produção no seu provedor de nuvem mas também auxilia um cliente que vá rodar ele on-premise


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#casos-de-uso)

## ➤ Casos de uso

Você pode sempre verificar como o conceito por trás de segmentar um Projeto e definir suas características na [descrição](./docs/spec.pt.md); porém para nomear alguns Objetivos que você precisa declarar além dos obrigatórios -- `base` e `dev` -- você poderia fazer alguns para:

- segurança
- homologação
- benchmark
- [chaos gorilla](https://www.gremlin.com/chaos-monkey/the-simian-army/)
- etc.

Uma vez que você ficar confortável com a idea de quebrar as coisas em segmentos, seu processo de pensamento fica claro para outros reproduzirem e te ajudarem com isso; como fazer um motor rodar sem problemas.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#por-que-no-o-docker-compose)

## ➤ Por que não o Docker-Compose?

- Compose é mais complexo e feito para ser de visão de "Definido por Sistema" experiência com o manifesto que é mais adequado para acomodar outros projetos além do qual está atualmente trabalhando, como Banco de Dados, Proxy Reverso, Cache em Memória e etc.
- Compose é não foi feito para comportar várias tarefas, apenas fazer uma coisa que bem feita que seria rodar um sistema feito de vários projetos e não rodar várias coisas em um projeto.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#por-que-como-npm)

## ➤ Por que "como NPM"?

- Alguns anos atrás cada projeto que trabalhei era feito em [Node](https://nodejs.org/en/) + [TS](https://www.typescriptlang.org/). Quando comecei a mexer com outras linguagens como Python, [Julia](https://julialang.org/) e [R](https://www.r-project.org/), a falta de algumas facilidades de CLI que são providas pelo [NPM](https://www.npmjs.com/) e [Yarn](https://yarnpkg.com/) que foi uma grande "chamada" de que eu me encontrava em uma *terra de ninguém*.
- Essas facilidades que ajudam o desenvolvimento de um projeto e sua manutenabilidade ao ponto de se tornar quase que sem eforço quando comparado com agumas maneiras atuais de se fazer isso; quando você inicia um projeto Node/NPM você está apenas um passo de um *yay/apt/apk/brew/ou qualquer outro gerenciador de pacotes de sistemas* de começar seu projeto como até mesmo Go ou Haskkel, o código que escreve é pouco acoplado no SO/Distro que você escreve, isso ajuda evitar que o projeto pare de funcionar depois de algum update de SO. Este sonho longínquo para outras linguagens que são altamente acopladas no ambiente que elas são executadas; esta falta de coesão e concisão são a semente por trás do ínicio do Succubus.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#metas)

## ➤ Metas

- A idéia dete projeto é cirar um fluxo de desenvolvimento independente do SO/linguagem/framework.
- Ao mesmo tempo que a ferramenta permite aos usuários implementarem os próprios "sotaques" dentro do padrão de desenvolvimento.
- O mantar por trás de tudo isso é: *"ajude você por ajudar os outros a alcançar o melhor"*.
- Como pontuado anteriormente, o ambiente acabam atrapalhando o fluxo de desenvolvimento e alguém não terminou a graduação dem Ciências da Computação (CC) geralmente eu consigo lidar bem com este tipo de situação mas muitas vezes pessoas que trabalham comigo rodam seu código no Mac, Windows ou qualquer outro ambiente Linux que geralmente levam várias horas durante o desenvolvimento de um projeto tentando fazer com que o código rode na máquina dos outros para debugar, reduzindo o passo de desenvolvimento e apertando os prazos de entregas. Então a ideia é compartilhar maneiras de evitar isso.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#faa-mais-fazendo-menos)

## ➤ Faça mais fazendo menos

- O uso de containers por debaixo do capô ajuda acelerar o ciclo de desenvolvimento do qual participei; epecialmente se você está pensando em publicar depois na [AWS](https://aws.amazon.com/), [Heroku](https://www.heroku.com/), [Azure](https://azure.microsoft.com/en-us/), [Linode](https://www.linode.com/), [GCP](https://cloud.google.com/) e vários outros; reduzindo a lacuna entre serverless e não-serverles ambiente.
- Não lute contra fazer seu projeto mais acessível para outros e por consequência fazer mais reproduzível para outros; seu provedor de nuvem não sabe que você está utilizando o Mac M1 mais recente, mas eles entendem a mesma tecnologia de "enpacotamento" que eles usam internamente para outros projetos que as "baleias" da indústrias trabalham.
- Isso pode soar como um grito ao vento, mas por favor me ouça:
  - *Seu provedor de nuvem não sabe como configurar o compilador bleeding edge que você usa*
  - *Seu cliente não liga sobre a granularidade de 200 passos no processo de build*
  - *E seu estagiário não possuí nas costas os mesmos anos de experiência que você e as marcas de batalha que você passou*
- Se você é apenas um novato em todo o **Mundo de Programação** e pode ver tudo isso e pensar que é *um pouco demais* para você e até mesmo estar assustado porque seu curso de [NumPy](https://numpy.org/) + [Pandas](https://pandas.pydata.org/) nunca falou sobre isto, minha mensagem para ti é *"dê uma chance"* -- por favor dê uma olhada em [CFD](https://github.com/Fazendaaa/CFD) e veja  vários outras ferramentas que possam ser interessantes para ti, especialmente [Jinn](https://github.com/Fazendaaa/Jinn) se você não souber onde começar.

Mais do que uma plataforma *"Faz tudo"*, este projeto tem ao seu cerne o propósito de compartilhar conhecimento/aprendizado e casos de uso como sua virtude principal -- aprender com as suas falhas e compartilhar elas pode ajudar os outros além de só você mesmo.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#uso)

## ➤ Uso

### init

Lista manifestos já feitos para você escolher e usar no seu projeto, permitindo que você mantenha um projeto no qual já trabalha usando algo que atenda as suas necessidades:

```shell
succ init
```

### interactive

```shell
succ interactive
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


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#executando)

## ➤ Executando

Você não precisa instalar Go para rodar esta ferramenta, apenas o Docker. E para fazer isso e dar um *test-drive*, você pode fazer isso apenas colando a seguinte linha no seu terminal:

```shell
shell'docker run -it --volume $(pwd):/project --workdir /project fazenda/succubus
```

Esta abordagem possuí algumas limitações porem é uma grande maneira de molhar seus dedos e iniciar logo a usar a ferramenta; se suas necessidades não forem suportadas por esta maneira você pode sempre procurar [instalar](#instalar) ela.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#motores)

## ➤ Motores

Eu sei que a empresa Docker é meio que um "B.O." em alguns circulos, por isso que a ideia desta ferramenta é procurar se manter atrás de uma camada de *Motor* de abstração, permitindo uma troca se necessária.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#instalar)

## ➤ Instalar

### Go

```shell
go install github.com/Fazendaaa/Succubus@latest
```

Caso decida em utilizar o binario distribuido atravez do Go, só se lembre de usar `Succubus` ao invés de `succubus` para executar os comandos.

#### Sem carregar

Provavelmente está faltando as seguintes configurações:

```shell
export GOPATH="$HOME/go/"
export PATH="$PATH:$GOPATH/bin/"
```

### Binário

Primeiro de uma olhada no projeto [zyedidia/eget](https://github.com/zyedidia/eget)

```shell
curl https://zyedidia.github.io/eget.sh | sh
./eget Fazendaaa/Succubus
mv Succubus $HOME/.local/bin/succubus
```

### Docker

Você não precisa instalar o Go para usar esta ferramenta, apenas Docker. E para testar você só precisa rodar a seguinte linha no seu terminal:

```shell
alias succubus='docker run -it --volume ${PWD}:${PWD} --workdir ${PWD} fazenda/succubus'
```

E então executar o seguinte comando para ter certeza que tudo está funcionando como o esperado:

```shell
succubus --help
```

Após você confirmar que tudo está funcionando direitinho, pode fazer o seguinte para iniciar um projeto:

```shell
succubus init .
```


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#uninstalling)

## ➤ Uninstalling

### Go

```shell
rm $GOPATH/bin/Succubus
```

### Binary

```shell
rm $HOME/.local/bin/succubus
```

### Docker

```shell
docker rmi --force fazenda/succubus
```


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#autor)

## ➤ Autor

Apenas [eu](https://github.com/Fazendaaa) porque o projeto já mencionado foi implementado por mim mesmo. Por conhecer cada linha daquele código, fazer o port seria mais fácil dessa maneira.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#contribuindo)

## ➤ Contribuindo

Veja mais sobre como contribuir [aqui](CONTRIBUTING.md). Aqui temos uma lista dos contribuidores:


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#contributors)

## ➤ Contributors
	

| [<img alt="Fazendaaa" src="https://avatars2.githubusercontent.com/u/12137236?s=460&u=75ec76d6f0c577de2ebfa4eae77cc4c4ad17ec06&v=4" width="100">](https://twitter.com/the_fznd) | [<img alt="You?" src="https://joeschmoe.io/api/v1/random" width="100">](https://github.com/andreasbm/web-config/blob/master/CONTRIBUTING.md) |
|:--------------------------------------------------:|:--------------------------------------------------:|
| [Fazendaaa](https://twitter.com/the_fznd)        | [You?](https://github.com/andreasbm/web-config/blob/master/CONTRIBUTING.md) |



[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#a-fazer)

## ➤ A fazer


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#referncias)

## ➤ Referências

### Repositórios

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [devspace/awesome-github-templates](https://github.com/devspace/awesome-github-templates)

### Blog-posts

- [Static code analysis for golang](https://levelup.gitconnected.com/static-code-analysis-for-golang-5f24b555d227)
- [Compilers 101 - Overview and Lexer](https://dev.to/lefebvre/compilers-101---overview-and-lexer-3i0m)
- [THE BASICS: 7 Alternatives to Docker: All-in-One Solutions and Standalone Container Tools](https://jfrog.com/knowledge-base/the-basics-7-alternatives-to-docker-all-in-one-solutions-and-standalone-container-tools/)
- [I'll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/)
- [Getting Started with Code Coverage for Golang](https://about.codecov.io/blog/getting-started-with-code-coverage-for-golang/)
- [Tutorial: How to create a CLI tool in Golang](https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f)
- [How to Create a CLI in Go in Minutes](https://dzone.com/articles/how-to-create-a-cli-in-go-in-few-minutes)

### Normas

- [Shell Grammar Rules](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_10_02)

### Forums

- [Why do I get a “cannot assign” error when setting value to a struct as a value in a map? [duplicate]](https://stackoverflow.com/a/32751792/7092954)


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/water.png)](#license)

## ➤ License
	
Licensed under [AGPL-3.0](https://opensource.org/licenses/AGPL-3.0).
