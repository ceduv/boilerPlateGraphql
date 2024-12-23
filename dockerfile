# Étape 1 : Construire l'application Go
FROM golang:1.23.4 AS build

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers de configuration du module Go
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste des fichiers du projet
COPY . .

# Construire l'application Go
RUN go build -o server ./cmd/main.go

# Étape 2 : Exécuter l'application dans une image compatible avec glibc
FROM debian:bookworm-slim

# Installer les certificats SSL (nécessaires pour HTTPS si utilisé)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail
WORKDIR /root/

# Copier le binaire compilé depuis l'étape de construction
COPY --from=build /app/server .

# Exposer le port 8080
EXPOSE 8080

# Commande par défaut pour exécuter l'application
CMD ["./server"]