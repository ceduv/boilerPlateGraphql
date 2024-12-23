# Variables
APP_NAME=my-go-app
PORT=8080

# Construire l'image Docker
build:
	docker build -t $(APP_NAME) .

# Stopper le conteneur (si nécessaire)
stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

# Nettoyer les images Docker
clean: stop
	docker rmi $(APP_NAME) || true

# Afficher les logs du conteneur
logs:
	docker logs -f $(APP_NAME)

################################################################
################## Lancer le conteneur Docker ##################
################################################################

# Lancer le conteneur Docker
run:
	air -c air.toml

################################################################
############################ Tools #############################
################################################################

# tools
gen:
	gqlgen generate

sql:
	sqlc generate