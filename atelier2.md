# Atelier 2
### Structure de données et interface

**Branche Atelier 2**

#### Objectifs

* Modéliser les structure de données pour permettre le redimensionnement d'image
* Utiliser des interfaces pour représenter une tâche pouvant s'exécuter

#### Etapes

* Modifier l'exemple précédent pour utiliser le package **logger** et lancer le test newgopher_test.go
* Créer une structure présentant une tâche de traitement d'une image : largeur, longueur, chemin image source et chemin image destination
* Utiliser une interface représentant une tâche avec une méthode **Do()**
* Mocker la méthode Do pour le traitement d'image en utilisant le logger pour afficher "Run resize from, to, width, height". Exemple "Run resize img.jpg img2.jpg 400px 200px"

Manipulation de tableau :
* Implémenter les méthodes de la structure task.Manager
  * Expliquer les méthodes manager.NextId() et task.GetId()
  * Certaines fonctions n'ont aucun intérêt sinon de manipuler un tableau 
* Lancer le test task_manager_test.go

#### Aide

* [reflect.TypeOf()](https://pkg.go.dev/reflect#TypeOf) permet de connaître le type d'une structure