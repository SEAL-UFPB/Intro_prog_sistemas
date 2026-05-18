#!/bin/bash

# 1. Identifica a versão e baixa
GO_VERSION=$(wget -qO- https://go.dev/VERSION?m=text | head -n 1)
FILE_NAME="${GO_VERSION}.linux-amd64.tar.gz"

echo "Baixando $GO_VERSION..."
wget -q --show-progress "https://go.dev/dl/${FILE_NAME}"

# 2. Limpeza e Extração
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "$FILE_NAME"
rm "$FILE_NAME"

# 3. Identifica o shell real do usuário
if [ -f "$HOME/.zshrc" ]; then
    CONF_FILE="$HOME/.zshrc"
else
    CONF_FILE="$HOME/.bashrc"
fi

# 4. Adiciona ao PATH se ainda não estiver lá
if ! grep -q '/usr/local/go/bin' "$CONF_FILE"; then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> "$CONF_FILE"
    echo "PATH adicionado em $CONF_FILE"
else
    echo "PATH já configurado em $CONF_FILE"
fi
