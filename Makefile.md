# Commandes Makefile

## Commandes disponibles

### **1. `make build`**
Construit l'image Docker.

### **2. `make run`**
Construit l'image Docker (si nécessaire) et lance le conteneur.

### **3. `make stop`**
Arrête et supprime le conteneur Docker en cours d'exécution.

### **4. `make clean`**
Nettoie les conteneurs et supprime l'image Docker.

### **5. `make logs`**
Affiche les logs du conteneur Docker en temps réel.

### **6. `make compose`**
Construit et lance les services définis dans le fichier `docker-compose.yml`.

### **7. `make test`**
Exécute les tests unitaires Go.

### **8. `make push`**
Construit l'image Docker, la tague, et la pousse sur Docker Hub.