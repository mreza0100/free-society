# MicroService with golang


## A social media with graphql written in golang with microservices architecture

### Used technologies:
	Nats
	GRPC
	Docker
	GraphQL&Gqlgen
	Postgresql
	Redis
	Gin

### Design patterns:
	DDD
	Repository based
	Saga
  
### Principles:
	12 Factor


### Features:
	Post
	Upload 4 pictures for every post
	Like
	Notification for like
	Advanced security
	Session management
	Follow
	Timeline
	Reshare post with your followers
	Avatar for user
	Database rollback on error
	Server side cookies
  
### Task list:
      Add more e2e tests
      TDD
      Notification for follow
      Notification for posts
      Kubernetes for deployment
      Using node.js for hellgate service
      Graphdb for follower suggestion
      Block feature  

## In order to use this project, clone it first
	git clone https://github.com/mreza0100/free-society
	cd ./free-society


#### Requirements
	Unix like OS
	Golang +1.16
	Docker
	Docker-compose

## Development
### Running databases and other requirements, like nats server
	bash ./scripts/develop/fire-requirements.sh 

### Running services 
<pre>scripts provided for running each service is in /scripts/develop/services directory.
Each service has it's own process and terminal and it's restarting the process on file saving.
For running Hellgate service you should start all other services and then starting Hellgate service.
My recommendation is <a href="https://github.com/tmux/tmux">tmux.</a>
Now if Hellgate is up without any error you can go to <a href="http://localhost:10000">localhost:10000.</a>
</pre>
### You can find more information about instructions in [Hellgate models](https://github.com/mreza0100/golang-microService-boilerplate/tree/master/services/hellgate/graph/schema)
