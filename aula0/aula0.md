# Instalação do go no linux 
## Instalação automática 
Caso você esteja utilizando um computador com o sistema operacional linux e com a arquitetura x64, criei um comando específico que já automatiza a instalação para você

```bash
wget -qO- https://raw.githubusercontent.com/BASE-UFPB/Intro_prog_sistemas/main/aula0/goinstall.sh | bash
```
```bash
[ -f "$HOME/.$(basename $SHELL)rc" ] && . "$HOME/.$(basename $SHELL)rc" && go version
```


## Instalação manual
vá para o link :https://go.dev/dl/ e baixe o arquivo *.tar.gz equivalente à arquitetura do seu computador, normalmente é amd64 
rode o comando abaixo na pasta onde você baixou o arquivo 
```bash
sudo rm -rf /usr/local/go
```
para remover versões prévias do go no seu sistema

e o comando :

```bash 
sudo tar -C /usr/local -xzf *.tar.gz
```
para extrair os arquivos no seu sistema

para guardar o go ao seu path utilize o comando abaixo:


para quem usa o shell padrão da maioria das distribuições linux , o bash .

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc
```


ou :


```bash 
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc && source ~/.zshrc
```
para quem usa o zsh como shell



para verificar se deu certo, podemos só rodar o comando :

```bash 
go version

```
e ele vai mostrar a versão do go e garantir que está instalado 


para sua informação, os tutoriais a seguir levam como consideração a versão ****go version go1.26.3 linux/amd64**** do go 
