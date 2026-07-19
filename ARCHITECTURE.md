# Arquitetura

Este documento é a fonte canônica das regras de arquitetura do projeto. Regras de negócio e o fluxo textual do domínio são canônicos: código, testes, diagramas e documentação devem se adaptar a eles.

## Convenções de Go adotadas

Go define módulos e pacotes, mas não define uma arquitetura em camadas. Cada diretório com arquivos Go é um pacote próprio; diretórios aninhados não compartilham visibilidade, tipos ou dependências. Os nomes e limites dos pacotes devem comunicar uma responsabilidade coesa e estável.

As convenções de `cmd/` e `pkg/` seguem o uso disseminado na comunidade Go. `cmd/` contém os pontos de entrada dos executáveis e `pkg/` contém os pacotes da aplicação. As regras deste documento prevalecem sobre convenções genéricas.

## Princípios

- Não criar ou mover regras de negócio para fora de `pkg/domain/` e `pkg/usecase/`.
- Propagar `context.Context` quando a operação puder ser cancelada, expirar ou tiver *deadline*.
- Não criar funções dentro de funções, incluindo *closures* para regras de negócio ou utilitários.
- Interfaces de contrato ficam exclusivamente em `pkg/contract/` e só podem ser alteradas após consulta ao usuário. Interfaces só existem quando desacoplam uma variação real; não devem ser criadas para impor um padrão.
- Uma camada é agnóstica às camadas que a chamam, instanciam ou usam. A direção permitida abaixo descreve somente dependências de importação.
- Um pacote deve ter responsabilidade coesa. Não criar pacotes de conveniência, genéricos ou de uma única chamada sem fronteira técnica ou de domínio clara.

## Camadas e fronteiras

### `cmd/` — ponto de entrada e composição

`cmd/cli/main.go` recupera a configuração, instancia a infraestrutura comum e os repositórios, injeta as dependências nos use cases e inicia os seus fluxos. Não contém regra de negócio, consulta a banco, transformação de dados ou lógica de publicação. Nenhuma outra camada pode importar `cmd/`.

Pode importar `pkg/adapter/`, `pkg/contract/`, `pkg/usecase/`, `pkg/infrastructure/` e `pkg/infrastructure/config/`.

SQLite local é a única dependência obrigatória para iniciar a aplicação. MQTT, PostgreSQL e InfluxDB são destinos opcionais: seus use cases só são criados quando `MQTT_HOST`, `POSTGRES_DSN` ou `INFLUX_URL`, respectivamente, possuem valor. Falhas de inicialização desses destinos desativam somente o fluxo correspondente e não impedem coleta e armazenamento local.

### `pkg/domain/` — modelo e regras invariantes

Contém o vocabulário do negócio, entidades, tipos de valor e regras específicas de cada entidade. `pkg/domain/model/` não importa outra camada do projeto. Anotações e métodos de mapeamento GORM nas entidades são permitidos.

Pode depender somente de outros pacotes de `pkg/domain/`.

### `pkg/usecase/` — comportamento da aplicação

É o caminho de entrada para as funcionalidades da aplicação. Contém as regras de comportamento entre modelos: coleta, armazenamento local, formatação, publicação e retentativas. Somente `cmd/` compõe e inicia os use cases; adaptadores e infraestrutura não os conhecem.

O worker é um agendador agnóstico: recebe uma tarefa, inicia sua goroutine e a executa no intervalo configurado. Ele não pode conhecer sensores, leituras, repositórios, destinos, publicação ou qualquer outra regra de negócio. Cada use case implementa e entrega ao worker a própria tarefa, incluindo consulta, transformação, publicação e registro de sucesso.

Cada use case executável fica em uma pasta ou subpasta de `pkg/usecase/` e declara seu ponto de entrada em `usecase.go`. Os fluxos agendados são iniciados pelo `cmd/cli/main.go`; o use case inicia o seu worker internamente. O use case coordena diretamente as etapas de formatação e envio por clientes técnicos concretos, sem uma abstração genérica de publisher.

Pode depender de `pkg/domain/`, `pkg/contract/`, `pkg/adapter/`, de outros pacotes de `pkg/usecase/` e dos clientes técnicos necessários em `pkg/infrastructure/`. Não pode depender de `cmd/`.

### `pkg/adapter/` — fronteiras de entrada e saída

Traduz dados e protocolos entre a aplicação e o exterior. Contém repositórios, agregadores, montadores e DTOs, sem deslocar regra de negócio para fora de `domain` e `usecase`.

Em `pkg/adapter/repository/` e `pkg/adapter/aggregator/`, aplicar a estratégia *query-first*: usar primeiro o *query builder* do GORM; usar SQL puro parametrizado quando necessário. Go fica restrito à manipulação em memória e aos efeitos fora do banco.

Toda entidade de negócio possui repository próprio. Toda consulta, inclusive relações do GORM, permanece no repository da entidade correspondente à tabela principal consultada. Não criar *pools* de banco por método em `pkg/adapter/repository/`. Repositories não criam nem migram schemas; essa responsabilidade pertence à infraestrutura de banco.

Pode depender de `pkg/contract/`, `pkg/domain/`, outros pacotes de `pkg/adapter/`, `pkg/infrastructure/` e `pkg/infrastructure/config/`. Não pode depender de `cmd/` ou `usecase/`.

### `pkg/contract/` — contratos públicos

Contém os contratos de interface necessários à aplicação. Atualmente, o único contrato é o de sensor, pois existem implementações intercambiáveis. Todas as camadas, exceto `pkg/domain/`, podem importar `contract`.

Pode depender somente de tipos de `pkg/domain/`. Não contém implementação, regra de negócio ou dependência de infraestrutura.

### `pkg/infrastructure/` — recursos técnicos

Implementa preocupações técnicas agnósticas ao negócio e reutilizáveis: configuração, *logging*, conexões, criação e migração de schemas, clientes externos, métricas e utilitários. Migrações podem conhecer objetos do banco, mas não importam tipos de domínio. A camada não contém regras de negócio e não conhece quem a chama.

Pode depender de outros pacotes de `pkg/infrastructure/`, de `pkg/infrastructure/config/` e, para expor uma porta técnica, de `pkg/contract/`. Nunca pode importar `pkg/domain/`, `pkg/usecase/`, `pkg/adapter/` ou `cmd/`.

## Dependências permitidas

```text
cmd ────────────────> adapter, contract, usecase, config, infrastructure
contract ───────────> domain
usecase ────────────> contract, domain, adapter, infrastructure, usecase
adapter ────────────> contract, domain, adapter, config, infrastructure
domain ─────────────> domain
config ─────────────> infrastructure
infrastructure ────> contract, config, infrastructure
```

Dependências não representadas nesse fluxo exigem atualização prévia deste documento e da configuração `.go-arch-lint.yml` no mesmo conjunto de alterações.

## Mapa atual de pacotes

### `cmd/`

- `cmd/cli/` — ponto de entrada do executável e composição dos fluxos agendados.

### `pkg/domain/`

- `pkg/domain/model/sensor/` — entidade de metadados persistidos e classe de medição do sensor.
- `pkg/domain/model/measurement/` — entidade `Measurement` e tipos de medição meteorológica.
- `pkg/domain/model/publication/` — entidade de publicação bem-sucedida e os destinos suportados.

### `pkg/contract/`

- `sensor.go` — contrato público de sensor.

### `pkg/usecase/`

- `sensorcollector/temperature/` — coleta medições meteorológicas dos sensores e as persiste no SQLite local; o caminho atual é preservado.
- `homeassistant/` — recupera leituras pendentes, formata o payload do Home Assistant e publica no MQTT.
- `datareplicator/postgres/` — publica sensores e medições pendentes no PostgreSQL e registra o sucesso em `publication`; o caminho atual é preservado.
- `datareplicator/influxdb/` — publica leituras pendentes no InfluxDB.
- `lcdpublisher/` — ponto de entrada para publicação direta em LCD.
- `stdoutpublisher/` — ponto de entrada para publicação direta na saída padrão.
- `worker/` — agendador genérico, sem conhecimento de domínio.

Os arquivos `datareplicator/datareplicator.go` e `sensorcollector/sensorcollector.go` apenas reservam os respectivos pacotes-raiz; os use cases concretos ficam nas subpastas.

### `pkg/adapter/`

- `repository/measurement/` — persistência e consultas cuja tabela principal é `measurement`.
- `repository/measurement/sqlite/` — implementação SQLite do repository de measurement.
- `repository/measurement/postgres/` — persistência concreta usada pelo destino de publicação PostgreSQL.
- `repository/sensor/` — persistência e consulta da entidade `Sensor`, com implementações SQLite e PostgreSQL.
- `repository/publication/` — persistência da entidade `Publication`, com implementação SQLite.
- `aggregator/`, `assembler/` e `dto/` — fronteiras reservadas para suas responsabilidades específicas.

### `pkg/infrastructure/`

- `config/` — carregamento e interpretação de configuração.
- `database/` — utilitários técnicos compartilhados de banco, incluindo a normalização de instantes para `America/Sao_Paulo`.
- `database/sqlite/`, `database/postgres/` e `database/influxdb/` — clientes, conexões técnicas e migrations de banco.
- `messaging/mqtt/` — cliente técnico MQTT.
- `sensor/ds18b20/` e `sensor/mock/` — implementações técnicas do contrato de sensor.
- `logger/`, `context/`, `errors/`, `http/server/` e `metrics/` — recursos técnicos compartilhados.

## Validação automatizada

O arquivo [`.go-arch-lint.yml`](.go-arch-lint.yml) valida as dependências de importação entre as camadas. Execute `make format` para formatar o código e `make lint` para executar os linters, incluindo a validação arquitetural.

`deepScan` permanece ativado para validar chamadas e injeções de dependência, além das fronteiras de importação. A configuração separa MQTT, InfluxDB e logger como componentes técnicos usados diretamente pelos use cases. Nenhum pacote de infraestrutura pode importar `pkg/usecase/`.

O linter valida fronteiras de importação. As regras estruturais que não são inferíveis de imports — por exemplo, existência de `usecase.go`, inicialização de workers pelo use case e ausência de regras de negócio no `cmd` — são obrigatórias e devem ser revisadas junto com cada alteração.
