# Atelier 5.2
### Créer une API REST avec la librairie Gin

#### Objectifs

* Savoir utiliser la librairie Gin Gonic pour simplifier le développement d'API Rest

#### Etapes

Supprimer toutes dépendances à la librairie standard pour utiliser la librairie GIN

#### Aide

* [Librairie GIN](github.com/gin-gonic/gin) à ajouter
* Vous pouvez garder le wrapServer et changer uniquement le handler par celui de Gin
* Pour créer un endpoint avec un paramètre dans l'url, la syntaxte est : "/tasks/:id"
* Pour renvoyer une erreur avec un code http particulier, la fonction **AbortWithError** vous aidera
* Pour gérer correctement les erreurs 405 (méthode non autorisée), configurer la propriété **HandleMethodNotAllowed** de l'engine à true