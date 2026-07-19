# AGENTS

Este arquivo define regras obrigatórias para qualquer agente de IA que atue neste repositório.

## Princípios obrigatórios

- Regras de negócio e o fluxo textual são os escopos canônicos. Elas vêm do mundo real e das necessidades reais do domínio. Toda alteração deve respeitá-los estritamente como verdade canônica local.
- O código, testes, diagramas, documentação e demais artefatos devem sempre se adaptar ao escopo canônico; nunca o contrário.
- Evitar duplicidade documental: manter detalhe em uma fonte canônica e, nos demais arquivos, criar link se realmente necessário. Ao encontrar duplicidade, consolidar na fonte canônica no mesmo conjunto de alterações.

- Arquitetura e fronteiras de camada: [`ARCHITECTURE.md`](ARCHITECTURE.md), fonte canônica exclusiva das regras de arquitetura.

## Regras, premissas e restrições mandatórias

- Antes de qualquer ação, é obrigatório a leitura e validação do alinhamenta do que foi pedido com os documentos definidos acima neste arquivo.
- Sempre que a necessidade de alteração de interfaces em `pkg/contract/`, o usuário deverá ser consultado antes da alteração.
- Seguir estritamente as fronteiras de camada em [`ARCHITECTURE.md`](ARCHITECTURE.md).
- Em repositories e aggregators, seguir a regra query-first descrita em [`SPECS.md`](SPECS.md): GORM query builder primeiro, raw SQL parametrizado quando necessário e Go apenas para manipulação em memória ou efeitos fora do banco.
- Código, GoDoc, logs, mensagens de erros, comentários, nomes de arquivos/pastas e mensagens de git commit escrever em en-US.
- Novos artefatos de banco de dados criados em `dev/postgres/initdb.d`, incluindo nomes de tabelas, colunas, constraints, índices e comandos de migration, devem usar pt-BR.
- Documentação sempre em pt-BR, com acentuação correta e vocabulário do **Brasil**.
- Escrita sempre objetiva, simples e pragmática; sem repetições de informação e sem prolixidade sem ganho real.
- Documentar brevemente usando o padrão GoDoc, todas as funções e métodos.
- Nunca apagar comentários deixados pelo usuário no código. Ao alterar trechos próximos, preservar esses comentários e adaptar a implementação ao redor deles.
- Não introduzir novos pools de banco por método em `pkg/adapter/repository/`.
- Priorizar API builder do GORM em queries de Postgres.
- Propagar `context.Context` entre camadas quando houver possibilidade de cancelamento/timeout/deadline.
- Toda query de banco de dados ou mesmo relações com GORM, deve ficar na camada `pkg/adapter/repository/` tabela principal consultada; outras partes do código devem reutilizá-la.
- Não criar ou mover regras de negócio para fora de `pkg/domain/` e `pkg/usecase/`.
- `cmd/` só faz parsing/dispatch.
- `pkg/usecase/` contém regras de comportamento e regras de negócio entre modelos.
- `pkg/domain/model/` contém regras de negócio específicas por entidade.
- `pkg/infrastructure/` é sempre agnóstico a negócio: nunca importar `pkg/domain/`, `pkg/usecase/`, `pkg/adapter/` e fazer citações a domínios do negócio.
- `pkg/infrastructure/` deve conter qualquer helper **técnico** que **não** dependa de regras de negócio nem de tipos de domínio e **possa** servir outras partes do código.
- `cmd/cli/main.go` recupera configurações, instancia infraestrutura comum e injeta dependências nos _use cases_.
- Não criar função dentro de função (incluindo closures para lógica de negócio ou utilitários).
- É proibido criar commits sem solicitação **explícita** do usuário. O agente deve apenas preparar as alterações no working tree e explicar o que mudou. Quando o git commit for solicitado, sempre crie micro commits separando as mudanças em micro contextos.
- É proibido alterar o estado do index/staging sem solicitação **explícita** do usuário. Não executar comandos como `git add`, `git restore --staged`, `git reset`, `git rm --cached` ou equivalentes para adicionar, remover ou reorganizar arquivos em staging por iniciativa própria.
- `git restore --staged` e `git reset` são comandos terminantemente proibidos sem permissão expressa e explícita do usuário, mesmo quando o agente acredita estar corrigindo um erro próprio ou deixando o repositório em estado melhor.
- A regra sobre index/staging é absoluta: mesmo que o staging pareça errado, sujo, incompleto, contraditório ou causado por erro anterior do próprio agente, é proibido tentar corrigir, limpar, reorganizar, desfazer ou refazer o staging sem autorização explícita do usuário. O agente deve apenas reportar o estado encontrado e continuar trabalhando no working tree quando possível.
