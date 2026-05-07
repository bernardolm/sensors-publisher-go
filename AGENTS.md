# AGENTS

Este arquivo define regras obrigatórias para qualquer agente de IA que atue neste repositório.

## Regras, premissas e restrições mandatórias

- Código, GoDoc, logs, mensagens de erros, comentários, nomes de arquivos/pastas e mensagens de git commit escrever em en-US.
- Documentação sempre em pt-BR, com acentuação correta e vocabulário do **Brasil**.
- Escrita sempre objetiva, simples e pragmática; sem repetições de informação e sem prolixidade sem ganho real.
- Documentar brevemente usando o padrão GoDoc, todas as funções e métodos.
- Propagar `context.Context` entre camadas quando houver possibilidade de cancelamento/timeout/deadline.
- Não criar função dentro de função (incluindo closures para lógica de negócio ou utilitários).
- É proibido criar commits sem solicitação **explícita** do usuário. O agente deve apenas preparar as alterações no working tree e explicar o que mudou. Quando o git commit for solicitado, sempre crie micro commits separando as mudanças em micro contextos.
- É proibido alterar o estado do index/staging sem solicitação **explícita** do usuário. Não executar comandos como `git add`, `git restore --staged`, `git reset`, `git rm --cached` ou equivalentes para adicionar, remover ou reorganizar arquivos em staging por iniciativa própria.
- `git restore --staged` e `git reset` são comandos terminantemente proibidos sem permissão expressa e explícita do usuário, mesmo quando o agente acredita estar corrigindo um erro próprio ou deixando o repositório em estado melhor.
- A regra sobre index/staging é absoluta: mesmo que o staging pareça errado, sujo, incompleto, contraditório ou causado por erro anterior do próprio agente, é proibido tentar corrigir, limpar, reorganizar, desfazer ou refazer o staging sem autorização explícita do usuário. O agente deve apenas reportar o estado encontrado e continuar trabalhando no working tree quando possível.
