# Atelier 2.1
### Structure de données et interface

#### Objectifs

* Refactorer du code

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
* Pour ne pas exposer le détail d'une tache, on peut créer une fonction qui créée la tâche, NewPrint par exemple, et on change la visibilité de Print
* Pour rappel, quand le nom d'une structure commence par une majuscule, elle est publique, sinon privée