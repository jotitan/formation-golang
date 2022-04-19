# Atelier 3
### Utilisation d'une librairie externe

**Branche atelier3**

#### Objectifs

Utilisation de librairies externes :
* pour redimensionner nos images
* simplifier l'écriture de test unitaire avec testify

#### Etapes

* Utiliser go mod pour ajouter les dépendances suivantes :
  * github.com/nfnt/resize
  * github.com/stretchr/testify
* Modifier la fonction de test TestManager pour utiliser la librairie d'assertion testify
* Appeler [la librairie de redimensionnement](https://github.com/nfnt/resize) d'image dans resize_image.go
* Tester la fonction Resize à partir du main

#### Aide
* [Go module](https://go.dev/ref/mod#go-get)