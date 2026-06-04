# Setup Alpine no Raspberry Pi 3B

O ambiente alvo é um Alpine instalado diretamente no Raspberry Pi 3B. O Docker é usado apenas na máquina de desenvolvimento para montar e validar o pacote APK em um ambiente Alpine compatível.

O pacote APK do `sensors-publisher-go` já instala:

1. `/etc/modules-load.d/w1.conf` com os módulos `w1-gpio` e `w1-therm`.
2. As dependências runtime declaradas no pacote.
3. O serviço OpenRC no runlevel `default`.
4. O serviço `modules` no runlevel `boot`.
5. A linha `dtoverlay=w1-gpio` em `/boot/usercfg.txt`, sem sobrescrever o arquivo.

Após a instalação, reinicie o Raspberry Pi para ativar o overlay 1-Wire quando a linha `dtoverlay=w1-gpio` tiver sido adicionada.

Também mantenha `/etc/apk/repositories` compatível com `dev/setup/rpi-3b-alpine/repositories` antes de instalar o pacote, caso o Alpine ainda não esteja apontando para os repositórios da versão usada.

O arquivo `packages.alpine` lista ferramentas úteis para o sistema de desenvolvimento/operação do Raspberry Pi, mas elas não são dependências runtime obrigatórias do `sensors-publisher-go`.
