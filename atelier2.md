# Atelier 2
### Structure de données et interface

#### Objectifs

* Initialiser un projet avec des modules
* Modéliser les structure de données pour permettre le redimensionnement d'image
* Utiliser des interfaces pour représenter une tâche pouvant s'exécuter

#### Etapes

##### Etape 1 => Branche atelier2.0

* Initialiser les modules avec comme nom de projet "formation-go"
* Modifier l'exemple de l'atelier 1 pour utiliser le package **logger** et lancer le test newgopher_test.go
* Créer une structure présentant une tâche de traitement d'une image : largeur, longueur, chemin image source et chemin image destination
* Utiliser une interface représentant une tâche avec une méthode **Do()**
* Mocker la méthode Do pour le traitement d'image en utilisant le logger pour afficher "Run resize from, to, width, height". Exemple "Run resize img.jpg img2.jpg 400px 200px"

Manipulation de tableau :
* Implémenter les méthodes de la structure Taskmanager
* Lancer le test taskmanager_test.go

##### Etape 2 => Branche atelier2.1

Le code que nous avons écrit est bien mais il expose trop l'implémentation de nos tâches alors que nous souhaiterions dissimuler cela.
Nous allons également ajouter la notion de suppression de tâche au taskmanager.

* Expliquer l'utilisation des méthodes NextId du manager
* Implémenter la méthode Remove du TaskManager et lancer les tests de taskmanager_test.go


#### Aide
* Initialiser les modules dans un projet
```bash
go mod init project_name 
```
* [reflect.TypeOf()](https://pkg.go.dev/reflect#TypeOf) permet de connaître le type d'une structure