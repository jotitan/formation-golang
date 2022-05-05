# Atelier 4
### Gestion des tests

**Branche Atelier 4**

#### Objectifs

* simplifier l'écriture de test unitaire avec testify

#### Etapes

* Ajouter la dépendance _github.com/stretchr/testify_
* Modifier la fonction de test TestManager pour utiliser la librairie d'assertion testify
* Faire un test unitaire pour tester automatiquement que de la réduction de la photo **/resources/photo_test.jpg** fonctionne

#### Aide

* [Créer un répertoire temporaire](https://pkg.go.dev/io/ioutil#TempDir)
* [Lire le contenu d'un fichier](https://pkg.go.dev/io/ioutil#ReadFile)
* Pour vérifier qu'une photo a été réduite, faite simple, comparer le poids de l'image avant / après