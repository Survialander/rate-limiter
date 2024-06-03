# Rate limiter


## Descrição
Um Rate Limiter é um mecanismo utilizado em sistemas de computação e aplicações web para controlar a quantidade de requisições ou operações que um usuário pode realizar em um determinado intervalo de tempo. Esse controle é importante por várias razões:

- Prevenção: Protege o sistema contra ataques DoS, onde um grande número de requisições pode sobrecarregar e desestabilizar o sistema.
- Gerenciamento de Recursos: Auxilia na gestão eficiente dos recursos do servidor, assegurando que eles não sejam monopolizados por um pequeno número de usuários intensivos.
- Equidade: Proporciona um melhor uso dos recursos do sistema, fazendo a distribuição de maneira balanceada entre os usuários da aplicação.

Um Rate Limiter pode funcionar várias maneiras, limitando requisições por IP, por token de API, entre outros.

Este limiter foi desenvolvido para funcionar como um middleware, capturando as requisições feitas a API e fazendo o controle de acessos. Para configurar o limiter basta alterar as variáveis de ambiente no arquivo .env do seguinte modo:
```
REQUEST_LIMIT=3 // Limite de requisições por segundo
TIMEOUT_TIME_IN_SECONDS=30 // Tempo em segundos que o IP vai ser bloqueado caso exceda o limite de requisições
CUSTOM_TOKENS = [] // Tokens com valores de limite e tempo de bloqueio customizados
```

OBS: Ao configurar tokens customizados atente-se ao formato JSON correto:
```
{
  "token": "TOKEN"
  "requestLimit": 10
  "timeoutTimeInSeconds": 20
}
```

## Dependências
- Docker

## Executando o projeto
Após clonar o repositório basta rodar o seguinte comando na pasta raiz do projeto:
```
docker-compose up
```

A api estará disponível na porta `:8080`
