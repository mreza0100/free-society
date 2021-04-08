# MicroService with golang


## A social media with graphql written in golang with microservices architecture

### Used technologies:
  Nats<br />
  GRPC<br />
  Docker<br />
  Gqlgen<br />
  Postgresql<br />
  Redis<br />
  Gin<br />

### Features:
  Post<br />
  Like<br />
  Advanced security<br />
  Session management<br />
  Follow<br />
  Timeline<br />

## In order to use this project first clone it
```
git clone https://github.com/mreza0100/golang-microService
cd ./golang-microService
```

#### Requirements
<pre> Unix like OS</pre>
<pre> Golang +1.16</pre>
<pre> Docker</pre>
<pre> Docker-compose</pre>

## Run development
### <pre>Running databases and other requirements like nats server</pre>
<pre>  bash ./scripts/develop/fire-requirements.sh </pre>
### <pre>Running services</pre>
<pre>  scripts provided for running each service is in /scripts/develop/services directory
  Each service has it's own process and terminal
  For running Hellgate service you should start all other services and then starting Hellgate service
  My recommendation is <a href="https://github.com/tmux/tmux">tmux</a>
  Now if Hellgate is up without any error you can go to <a href="http://localhost:10000">localhost:10000</a>
</pre>
### You can find more information about instructions in [Hellgate models](https://github.com/mreza0100/golang-microService-boilerplate/tree/master/services/hellgate/graph/schema)
  

