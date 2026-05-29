# Sobre o professor
Sou Mikael Menezes, aluno da universidade federal da paraíba e sou apaixonado por programação de sistemas
. Eu recomendo fortemente que você faça perguntas se tiver, tire dúvidas ou até entre em contato caso queira apenas discutir sobre assuntos diversos
Por isso eu divulgo meu contato para vocês. Email: mixaelmenezes@gmail.com, se quiser meu número do whatsapp, é só mandar um email pedindo.
E não tenha medo de perguntar, eu sempre digo pra todo mundo : "não existe pergunta besta, besta é não perguntar".

# Por que Go?
abaixo podemos ver uma linha que robustamente representa linguagens e suas complexidades/nível de acesso ao hardware, quanto mais à esquerda, mais acesso ao hardware a linguagem têm e à medida que a linha segue à direita, mais abstração com relação ao hardware temos. 
[Hardware] ----------------------------------------------------------------------------------> [Abstração]

 C / Zig / Rust          Go              JVM (Java/Kotlin)                             JavaScript / Lua / Python / Bash
  (Hard Systems)     (Cloud Sys)           (Applications)                                       (Pure Script)

o que chamamos de programação de sistemas normalmente se refere ao lado mais à esquerda da faixa, já que temos acesso e controle direto sobre o hardware,enquanto linguagens de script, normalmente costumam abstrair esse acesso para nós, sendo representadas à direita. Porém à medida que abstraimos o hardware, menos controle temos, mas ao mesmo tempo ganhamos mais facilidade ao programar.

Mas por que Go? 

Go ainda se encontra no escopo de linguagens de baixo nível, mas por ele possuir um coletor de lixo, ele facilita bastante a vida do programador, abdicando só um pouco da performance


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
