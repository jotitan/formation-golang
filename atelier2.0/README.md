# Atelier 2.0
### Structure de données et interface

#### Objectifs

* Initialiser un projet avec des modules
* Modéliser les structure de données pour permettre le redimensionnement d'image
* Utiliser des interfaces pour représenter une tâche pouvant s'exécuter

Cet atelier se fait en deux étapes : atelier2.0 et atelier2.1

#### Etape 1 => Répertoire atelier2.0

##### Gestion des tâches

* Initialiser les modules avec comme nom de projet "formation-go"
* Modifier l'exemple de l'atelier 1 pour utiliser le package **logger** et lancer le test newgopher_test.go
* Dans le package model qui contient les structures utiles
  * L'interface Task représente une tâche à exécuter : lui ajouter une méthode **Do()**
  * Ajouter à la structure Print un champ message
  * Implémenter les méthodes nécessaires pour que **Print** soit une **Task**
  * Ajouter à la structure resize les champs nécessaires : 
    * largeur
    * longueur
    * chemin de l'image source
    * chemin de l'image destination
  * Implémenter dans Resize la méthode **Do()** afin de respecter l'interface Task.
  L'implémentation sera basique et affichera juste un message avec le logger :  "Run resize from, to, height, width". 
  _Exemple "Run resize img.jpg img2.jpg 400px 200px"_
  * Lancer les tests tasks_test.go

##### Manipulation de tableau
* Implémenter les méthodes de la structure Taskmanager
* Lancer le test taskmanager_test.go

#### Aide
* Initialiser les modules dans un projet (dans le répertoire src)
```bash
go mod init project_name 
```
* Lancer un test : se mettre dans le répertoire et 
 ```bash
 go test
 ```

* [reflect.TypeOf()](https://pkg.go.dev/reflect#TypeOf) permet de connaître le type d'une structure. (méthode Name pour le nom)
* Pour ne pas exposer le détail d'une tache, on peut créer une fonction qui créée la tâche, NewPrint par exemple, et on change la visibilité de Print
* Pour rappel, quand le nom d'une structure commence par une majuscule, elle est publique, sinon privée