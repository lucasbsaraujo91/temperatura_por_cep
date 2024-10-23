Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

Requisitos:

    O sistema deve receber um CEP válido de 8 digitos
    O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
    O sistema deve responder adequadamente nos seguintes cenários:
        Em caso de sucesso:
            Código HTTP: 200
            Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
        Em caso de falha, caso o CEP não seja válido (com formato correto):
            Código HTTP: 422
            Mensagem: invalid zipcode
        ​​​Em caso de falha, caso o CEP não seja encontrado:
            Código HTTP: 404
            Mensagem: can not find zipcode
    Deverá ser realizado o deploy no Google Cloud Run.

Dicas:

    Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
    Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
    Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
    Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
        Sendo F = Fahrenheit
        Sendo C = Celsius
        Sendo K = Kelvin

Entrega:

    O código-fonte completo da implementação.
    Testes automatizados demonstrando o funcionamento.
    Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
    Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

------------------------------------------------------------------------------------------------------------------------

Descrição:

    Para esse projeto utilizei o exercício anterior sobre a arquitetura clean architecture e o exercicio onde utilizamos go routines e duas apis para pesquisar CEP e a resposta mais rápido e a que utilizamos para responder a chamada da requisição.
    Por opção criei uma chamada apenas para trazer os dados do CEP na rota /address/06765000

Realizar os testes via brownser local: 

    Para realizar os testes locais http://localhost:8081/weather/zipcode/{cep}

Realizar os testes via aplicação em prod:

    Para realizar os testes da aplicação no ar: https://temperatura-por-cep-773278080263.us-central1.run.app/weather/zipcode/{cep}

Para rodas os testes locais:

    Na raiz os projeto via terminal rodar o seguinte comando go test -cover ./...
    
Link Cloud run

https://temperatura-por-cep-773278080263.us-central1.run.app/weather/zipcode/06765000


    

