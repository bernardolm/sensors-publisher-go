# Ambiente Docker

O Docker é o ambiente canônico para executar tarefas Go, compilar artefatos e montar o pacote APK.

## Pré-requisitos

- Docker com Compose e Buildx ativos.
- Arquivo `.env` na raiz do repositório.

Go não é necessário na máquina host.

## Comandos

```bash
make run
make watch
make format
make lint
make arch-check
make test
make build
make package-pi
make image-pi
```

`make build` sempre exporta em `bin/` os binários para macOS Silicon, Linux x86_64 e Raspberry Pi 3B. `make package-pi` usa o binário ARMv7 já exportado para gerar o APK. `make image-pi` carrega localmente a imagem ARMv7 e não publica em registry.
