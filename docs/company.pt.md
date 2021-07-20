# Empresarial

Este documento demonstra alguns dos beneficios ao se utilizar Succubus como uma ferramenta de gerenciamento de Projetos na sua empresa; como uma ferramenta de uso diário.

Alguns dos benefícios estão relacionados com:

- Zero necessidade de instalar compiladores/interpretadores nas máquinas dos desenvolvedores do seu time
- Você pode remover o `Dockerfile` do repositório -- evitando que alguém mexa com ele apenas para poder passar em algum processo de CI/CD
- Evita que algum update de copilador quebre o sistema operacional e vice-versa
- Seu time pode usar uma infinidade de sistemas operacionais nas suas mais diferentes versõe e arquiteturas -- isso não quebrará o ciclo de desenvolvimento do Projeto
- Você pode atualizar todo a versão seu framework/compilador apenas mudando uma linha
- etc.

Além de tudo isso, você não precisa saber Go ou até mesmo nada relacionado ao Dokcer ou até mesmo os Projetos para ajudar Succubus para atender suas necessidades. A ideia no geral é permitir que todo mundo se beneficie disso sem precisar tocar uma única linha de código.

- [Empresarial](#empresarial)
  - [Cenários](#cenários)
  - [Diligência](#diligência)
  - [Zero Confiança](#zero-confiança)
  - [Construa a sua versão](#construa-a-sua-versão)

## Cenários

- Você atualmente está tocando um Projeto no qual vários desenvolvedores estão com problemas no processo de CI/CD porque alguém esqueceu de adicionar uma dependência que tinha na máquina dele
- Você não consegue testar facilmente várias versões de pacotes/linguagens sem ter dores de cabeça quando tudo falha
- Você quer aproveitar os benefícios do "jeito container de viver" sem ter que treinar seu time todo nisso
- Rodar o seu Projeto nas mais diversas arquiteturas de CPU? Sem problemas, *"o rei está morto, longa vida ao rei"*
- Quer testar o [GraalVM](https://www.graalvm.org/) sem ter que instalar ele? Ou até mesmo remover ele após decidir que quer permanecer do jeito que as coisas estão? Sem problemas

## Diligência

> Este código é Free and Open Source Software (FOSS) e sempre o será!

Tudo isso é para permitir que este projeto permita que várias pessoas entendam melhor a tecnologia de container e a adapte para poder atender melhor a necessidade deles, em todos os projetos imagináveis que existem aí a fora. Por isso uma implementação sólida é necessária e isso é uma das premissas das quais o próprio Succubus é utilizado para manter o Succubus.

## Zero Confiança

- Você pode implementar sua própria versão
- Você pode mudar esta versão
- Você pode fazer o que quiser -- com tanto que siga a [LICENÇA](../LICENSE)

## Construa a sua versão

Você pode sempre construir todo o conceito de Projeto apresentado aqui em outra linguagem como Python, Java, Rusta e etc.; porém você pode também adaptar esta implementação através de uma interface genérica na qual se ajusta ao que precisa:
