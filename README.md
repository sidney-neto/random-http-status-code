<h1 align="center">
    <img alt="Semana_OmniStack" src="https://res.cloudinary.com/practicaldev/image/fetch/s--J3E8KS70--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://cdn-images-1.medium.com/max/1024/1%2AE33brkN6zivLSb-D9i-CdQ.png" width="200px" height="200px" />
</h1>

# Random HTTP Status Code
Servidor HTTP na porta 8080 que retorna um código de status HTTP aleatório
## Exemplo em Go
```console
$ git clone https://github.com/sidney-neto/random-http-status-code
$ go run webserver.go
# Abrir no navegador - http://localhost:8080
```
## Exemplo em Docker 
https://hub.docker.com/r/sidneyramosneto/random-http-status-code
```console
$ git clone https://github.com/sidney-neto/random-http-status-code
$ docker build -t random-http-status-code .
$ docker run -it -p 8080:8080 random-http-status-code
# ou
$ docker run -d --rm -p 8080:8080 --name random-http-status-code random-http-status-code
# Abrir no navegador - http://localhost:8080
```
