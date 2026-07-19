# Especificações Técnicas

## Chaves Primárias UUID

O Postgres é a única fonte responsável por gerar chaves primárias UUID das tabelas persistidas pelo serviço.

O código da aplicação não deve gerar, reparar, substituir ou enviar valores de PK `id` em `INSERT` ou `UPSERT`. O `base.Model.ID` é mapeado como campo somente-leitura para escrita, e repositories devem hidratar o valor gerado pelo banco via `RETURNING` ou leitura posterior.

Registros já persistidos com UUID zero são problema de dados e devem ser corrigidos operacionalmente fora da aplicação, por exclusão física ou migration controlada. A aplicação não deve conter lógica de reparo automático para PK zerada.

Restrições adicionais em tabelas existentes devem ser criadas por migrations com `ALTER TABLE`, nunca por alteração de `CREATE TABLE` esperando recriação da base.

## Query-First Para Repositories e Aggregators

A regra de execução é tentar primeiro GORM query builder.

Quando chaves compostas, `VALUES`, `CTE` ou `unnest` deixarem o GORM pouco claro, usar raw SQL parametrizado.

Manter Go para mutação de objetos em memória, cache, publish, HTTP e propagação para ponteiros.

O escopo auditado inclui `pkg/adapter/repository`, `pkg/adapter/aggregator`, `pkg/usecase/entity`, `pkg/usecase/hubspot` e `pkg/usecase/doare`.

### Regras

- Consolidar a lógica `entity_fonte` primeiro, com fallback para `entity.source/source_code`, em helpers internos de repository.
- Suportar múltiplos pares `(source, source_code)` em uma única query para donor, donation e subscription.
- Usar consulta única por pares nos processors de donor, donation e subscription em `pkg/usecase/entity/handlers/*/processor/batch.go`.
- Em `pkg/usecase/entity/handlers/donation/persister/batch.go`, consultar os pares persistidos após o batch.
- Em `fillPersistedFieldsBySourceAndSourceCode` de donor e subscription, buscar todos os pares em uma query só.
- Em `subscription.ListSourceCodesByDonations`, relacionar doações recebidas com `assinatura_fonte` primeiro e fallback em `subscription.source/source_code`.
- Em `donor.ListByEmailSince`, usar filtros SQL com arrays.
- Padronizar `upsertSources` de donor, donation e subscription em helper comum de repository.
- Não mover para SQL deduplicação de batches, propagação de duplicatas, builders de payload HubSpot, chunks HTTP ou filtros de API externa Doare.

### Compatibilidade

- A API `ListBySourceCodePairs` existe em `pkg/contract/` para donor, donation e subscription.
- Não há mudança planejada de schema ou migration.
- Raw SQL deve ser sempre parametrizado; `source_code` nunca deve ser interpolado.
- Consultas sem source explícito, como `List`, `ListByDonors` e `ListSince`, ficam fora desta execução até definição da relação de fonte vencedora quando uma entidade tiver múltiplas fontes.

### Testes e Aceite

- Cobrir SQL dry-run para garantir que relação vence legado, fallback só entra quando não há relação equivalente, pares não viram produto cartesiano e códigos vazios são ignorados.
- Cobrir processors e persisters com batches misturando sources diferentes e source codes repetidos.
- Rodar `go test -count=1 ./...`.
- Validação manual: `make doare-subscriptions` não deve cair em `unresolved donor IDs after fill+insert`.

## Persistência e entrega de medições meteorológicas

### Fluxos e workers

- O worker de coleta executa `sensor → modelo de domínio → storage SQLite` e não depende de rede. Cada medição contém sensor, tipo, unidade, valor `float64` e instante da coleta com fuso horário.
- Na inicialização, os sensores configurados são inseridos ou atualizados no SQLite pelo seu `id` estável. O cadastro preserva os metadados públicos necessários aos formatadores, mesmo que o hardware deixe de estar disponível depois.
- Os use cases de MQTT e InfluxDB executam `storage SQLite → sensor SQLite → formatação → cliente do destino`. O use case PostgreSQL executa `storage SQLite → repository PostgreSQL`. Cada use case controla o próprio fluxo e entrega uma tarefa agnóstica ao seu worker.
- Ao iniciar, todos os workers são criados. Falhas de conexão com destinos externos são tratadas no ciclo daquele destino e nunca interrompem a coleta local.
- Sem variável de ambiente, a coleta ocorre a cada 10 segundos, MQTT a cada 15 segundos e PostgreSQL a cada 1 minuto. InfluxDB usa 1 minuto por padrão.

### Storage local e entregas

- `measurement` é a fonte canônica local das coletas. Contém `id`, `id_sensor`, `collected_at`, `value`, `class` e `unit_of_measurement`.
- `sensor` é a fonte canônica local dos metadados dos sensores, independentemente do tipo de medição. A coluna `class` recebe diretamente `Sensor.Class()` e identifica a informação entregue pela instância, como `temperature`, `atmospheric_pressure` ou `humidity`. A tabela contém `id`, `class`, `icon`, `manufacturer`, `model`, `name`, `picture`, `time` e `unit_of_measurement`.
- `measurement.class` recebe diretamente a classe do sensor, como `temperature`, `atmospheric_pressure` ou `humidity`. `unit_of_measurement` registra a unidade efetivamente coletada.
- A aplicação normaliza os instantes gravados para `America/Sao_Paulo`, inclusive os timestamps automáticos do GORM e a sessão PostgreSQL.
- `publication` registra entregas bem-sucedidas. Contém `id`, `id_measurement`, `sent_at` e `destination`, onde destino é `mqtt`, `postgres` ou `influxdb`.
- Uma medição está pendente para um destino enquanto não existir uma linha em `publication` para o par `(id_measurement, destination)`.
- Após cada entrega bem-sucedida, somente o use case daquele destino grava a respectiva `publication`; os demais destinos continuam pendentes até sua própria confirmação.
- As consultas de pendência são ordenadas por `measurement.collected_at` e `measurement.id`. Uma retentativa usa o instante original da coleta no payload, nunca o instante da retentativa.
- O DDL comum está em [`dev/database/ddl/ddl.sql`](dev/database/ddl/ddl.sql). `id_sensor` preserva a origem necessária para reconstruir cada envio.
- A infraestrutura SQLite cria e migra o schema antes de disponibilizar a conexão aos repositories.
- `Sensor`, `Measurement` e `Publication` possuem repositories próprios; nenhum repository persiste outra entidade.
- PostgreSQL é um destino de publicação. O use case lê exclusivamente o SQLite, persiste os metadados do sensor e a medição no PostgreSQL e registra o sucesso em `publication` no SQLite.
