# Instalação do go no linux 
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

para adicionar o go ao seu path utilize o comando abaixo:


```bash
export PATH=$PATH:/usr/local/go/bin
```

agora para testar se deu certo a instalação, saia e entre novamente no terminal, ou rode algum dos comandos abaixo: 


```bash 
source ~/.bashrc
```
para quem usa o shell padrão da maioria das distribuições linux , o bash .


ou :


```bash 
source ~/.zshrc
```
para quem usa o zsh como shell



e depois disso é só rodar o comando 

```bash 
go version

```
para ver a versão do go e garantir que está instalado 


para sua informação, os tutoriais a seguir levam como consideração a versão ****go version go1.26.3 linux/amd64**** do go 
