# sensors-publisher-go

## Desenvolvimento

As tarefas Go são executadas em contêineres Docker. Consulte [dev/docker/README.md](dev/docker/README.md) para os pré-requisitos e comandos disponíveis.

O serviço persiste cada temperatura e os metadados do sensor no SQLite local antes de qualquer integração de rede. MQTT, PostgreSQL e InfluxDB leem exclusivamente essa base local; PostgreSQL recebe somente uma replicação de escrita do SQLite. O fluxo, as tabelas e as regras de reenvio estão definidos em [SPECS.md](SPECS.md#persistência-e-entrega-de-temperaturas). As variáveis disponíveis estão em [.env.sample](.env.sample).

## Raspberry Pi 3B

O artefato oficial é um pacote Alpine `.apk`, com binário, serviço OpenRC e configuração em `/etc/sensors-publisher-go/config.env`.

```bash
make package-pi
```

Para instalar um pacote local sem assinatura:

```bash
sudo apk add --allow-untrusted bin/sensors-publisher-go-*.apk
```

## ds18b20

Para usar o sensor 1-Wire, carregue os módulos antes de iniciar o serviço:

```bash
sudo modprobe w1-gpio
sudo modprobe w1-therm
```
