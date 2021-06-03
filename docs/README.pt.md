# Succubus

> A maneira cloud-native de codar

[![English README](https://img.shields.io/badge/Language-EN-blue.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.md)
[![Portuguese README](https://img.shields.io/badge/Linguagem-PT-green.svg?longCache=true&style=for-the-badge)](https://github.com/Fazendaaa/Succubus/blob/master/docs/README.pt.md)

## O que é?

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

E se você precisar iteragir com alguma coisa, apenas para ver se vai funcionar ou não, basta:

```shell
succ interact
```

Uma tela similar a seguinte irá aparecerce:

```shell
Python 3.9.5 (default, May  9 2021, 14:00:28)
[GCC 10.2.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

E aí você terá uma experiencia de REPL (Read-Evaluate-Print-Loop) sem precisar instalar o Python, e, a melhor parte disso, é que depois este será o mesmo ambiente entregue para o sistema de produção.

# Índice

- [Succubus](#succubus)
  - [O que é?](#o-que-é)
    - [Fundamentos](#fundamentos)
- [Índice](#índice)
  - [Security and Reliability](#security-and-reliability)
  - [Universe](#universe)
  - [Possibilidades](#possibilidades)
  - [Manifesto de Projeto](#manifesto-de-projeto)
    - [Overview](#overview)
  - [Modelo de Corretor](#modelo-de-corretor)
    - [Por que Corretor?](#por-que-corretor)
  - [Pipe Production Environment](#pipe-production-environment)
  - [A Fazer](#a-fazer)
  - [Apêndice](#apêndice)
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

## Possibilidades

A principal ideia por trás deste projeto é abrir possibilidades para encapsulamento de projetos que você já está trabalhando assim como foi feito na empresa que trabalho -- você pode saeber mais [aqui](https://github.com/Fazendaaa/CFD) --; mas, além disso, você pode se perguntar algumas coisas:

- *Você gostaria de rodar em 20 anos os projetos como roda agora?*
- *Você gostaria de evitar que um projeto "intoxicasse" outro?*
- *Você gostaria de rodar seu proeto no Windows, Mac e Linux sem problemas?*
- *Você gostaria de deixar seu trabalho mais acessível?*

Se você diss *"sim"* para alguma das perguntas anteriores, você quer o Succubus. E, além disso, a ideia de criar de é de criar uma plataforma de fácil, seguro e rápido acesso para desenvolvimento desses "sotaques" de projeto para você poder evoluir de maneria rápida através de cada tentativa no projeto e/ou compartilhar com outros o que fez -- *ou até mesmo dando uma "mão amiga" e ajudar outros a entender melhor como fazer um projeto de estourar a boca do balão*

E o melhor é que você pode usar em qualquer linguagem, compilar para toda arquitetura depois porque você vai estar rodando containers debaixo do capô; se você estiver o M1 Mac você pode facilmente compartilhar o seu trabalho com outros colaboradores e até mesmo publicar ele na nuvem.

## Manifesto de Projeto

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

### Por que Corretor?

Um [corretor](https://dictionary.cambridge.org/dictionary/english/broker) segundo o Dicionário de Cambridge:

> uma pessoa que compra e vende moeda estrangeira, ações em companias, etc. [...]:
> "Eu liguei pro meu corretor buscando conselhos para investir no mercado de ações"

Dito isso, Succubus age como tal [corretor](https://www.betterteam.com/broker-job-description) para um programador:

- Fica responsável por detalhes que um programador comum normalmente não está ciente mas isso não significa que não deveriam ser seguidos
- Processa todos as transações requisitadas seguindo a regra
- Process all the required transactions by the rule
- Shows great knowledge about the market and how it behaves
- Helps you grow and understand better the discussed subject
- Allows some tweaks to proper fit the matter at hand to your best interest

That definition put together with the whole [Project Manifest](#project-manifest) concept, allow the broker -- Succubus -- to proper operate by the given guidelines.

## Pipe Production Environment

Trazendo todo o conceito da filosofia de produção [Just In Time (JIT)](http://people.brunel.ac.uk/~mastjjb/jeb/or/jit.html) -- por favor não confunda com o [JIT](https://www.freecodecamp.org/news/just-in-time-compilation-explained/) utilizado em compiladores --; fazendo com que o processo de produção se torne mais eficiente por padrão por evitar processos de validação ou de configurações tardios. Isso também permite fácil reprodução de toda a cadeia de processo na sua máquina apenas olhando para a parte do serviço que te importa, o Projeto; isso também nos permite melhor enteder e fazer translações de alguns tipos dos conceitos de microsserviços mas em um olhar mais profundo na sua máquina e apenas olhando o que te importa, quase como uma abordagem baseada em *comportamento do serviço*.

Essa abordagem é uma ferramente para enforçar alguns conceitos bem aclamados de desenvolvimento de software como:

- [Domain Driven Development](https://en.wikipedia.org/wiki/Domain-driven_design)
- [Test Drive Development](https://en.wikipedia.org/wiki/Test-driven_development)
- [Service Oriented Delivery](https://www.enterprise-architecture.com/docs/deliverSOARealiseEA.pdf)
- E muitos outros

E você pode ver isso sem medo de ter uma granularidade muito fina que acabe sobrecarregando seu processo todo de desenvolvimento; uma pessoa não pode dar conta de uma complexidade de uma empresa de desenvolvimento de software de 500 funcionários durante o processo de desenvolvimento de software; porém em ambos os cenários eles podem aumentar ou reduzir a mensurabilidade de granularidade durante o processo de desenvolvimento para ver se vai te trazer ou não um **retorno de investimento** para você, você também poderá reduzir seu **custo de propriedade**, além de maximizar modificação de workflows e a utilização de **testes de hipóteses**.

Este processo todo também ajuda automatizar algumas tarefas posteriores como:

- Automação de builds
- Automação de distribuição de software
- Análise de código
- Validações de segurança
- Coleta de métricas
- e muitas outras

Isso tudo porque você controla as regras do jogo, isso gera uma crescente reproducibilidade por uma margem significativa como consequência. Você pode olhar tal processo como o de uma FAB (Planta de Fabricar Semicondutores) na qual, por funcionar 24h durante sete dias por semana, muitos times são necessários para fazer com que tal processo funcione sem nenhum contra-tempos; FABs normalmente tem [95% do seu processo automatizado](https://www.youtube.com/watch?v=FP_g-as29x0&feature=youtu.be) o que permite que muitos times diferentes trabalhem no mesmo produto durante diferentes etapas da cadeia. Em um ambiente de desenvolvimento de software isso significaria que alguém pode tomar conta do processo onde o time/pessoa anterior parou sem se preocupar em ter que configurar nada, apenas continuaria trabalhando logo onde o processo anterior a chegada dele parou.

![factory](https://raw.githubusercontent.com/Fazendaaa/Succubus/master/assets/img/factory.jpg)

E isso em termos de desenvolvimento de software significa que a pessoa atual não terá com lidar com nenhuma "batata quente" nas mãos dela devido há alguma coisa feita pelo procedimento anterior da cadeia, assim como tal segurança também protege o processo que vem logo em seguida de ter que lidar com alguma coisa do tipo, seja ele um processo automatico, uma pessoa ou um time; não ter que ficar modificando as coisas e torcer que elas deem certo é uma ótima segurança de desenvolvimento por reduzir a margem de erro.

E a cereja no topo do bolo é que tudo isso é feito com uma mistura de várias tecnologias atuais que já foram testadas em batalhas e são abertas para todos poderem melhorar o processo como um todo, melhorando a experiência de desenvolvimento fazendo tornando uma tecnologia "testada em guerra"; isso significa que um passo na cadeia pode fazer uma modificação que, na próxima iteração, acabe facilitando a cadeia como um todo.

## A Fazer

Se você chegou até aqui, por favor dê uma olhada no [repositório](https://github.com/Fazendaaa/Succubus) do projeto e talvez dê um like, feedback ou até mesmo melhoria de código nele :)

## Apêndice


### DevOps

- [Addressing The Detrimental Effects of Context Switching With DevOps](https://insights.sei.cmu.edu/blog/addressing-the-detrimental-effects-of-context-switching-with-devops/)

### Cloud-Native

- [O Inferno de Dante do Cloud Native](https://fazenda.hashnode.dev/o-inferno-de-dante-do-cloud-native)

### JIT

- [Sennheiser Factory Tour - Hanover, Germany](https://youtu.be/5es8zggYM7A)
