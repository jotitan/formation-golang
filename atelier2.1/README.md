# Atelier 2
### Structure de données et interface

#### Objectifs

* Initialiser un projet avec des modules
* Modéliser les structure de données pour permettre le redimensionnement d'image
* Utiliser des interfaces pour représenter une tâche pouvant s'exécuter

#### Etape 2 => Branche atelier2.1

Le code que nous avons écrit est bien mais il expose trop l'implémentation de nos tâches alors que nous souhaiterions dissimuler cela.  
Nous allons également ajouter la notion de suppression de tâche au taskmanager.

* Expliquer l'intérêt de la méthode NextId du manager et l'implémenter.
* Implémenter la méthode Remove du TaskManager et lancer les tests de taskmanager_test.go
* Modifier les tâches Print et Resize pour ne pas exposer l'implémentation


#### Aide
* Initialiser les modules dans un projet
```bash
go mod init project_name 
```
* [reflect.TypeOf()](https://pkg.go.dev/reflect#TypeOf) permet de connaître le type d'une structure
* Pour ne pas exposer le détail d'une tache, on peut créer une fonction qui créée la tâche, NewPrint par exemple, et on change la visibilité de Print
* Pour rappel, quand le nom d'une structure commence par une majuscule, elle est publique, sinon privée