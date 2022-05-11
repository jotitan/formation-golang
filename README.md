# Formation Golang

7 ateliers à réaliser

* [Atelier 1](atelier1/README.md)
* [Atelier 2](atelier2.0/README.md)
* [Atelier 3](atelier3/README.md)
* [Atelier 4](atelier4/README.md)
* [Atelier 5](atelier5/README.md)
* [Atelier 6](atelier6/README.md)
* [Atelier 7](atelier7/README.md)

### Projet

L'objectif de projet est d'exécuter des tâches diverses par le biais d'un serveur qui va distribuer ces tâches.
La distribution des tâches sera assez basique (round robin)

Les tâches : 
* Redimensionner une image
* Afficher un hello world avec le prénom de l'utilisateur
* Tâche fictive qui met un temps aléatoire à s'exécuter

```mermaid

flowchart LR;



subgraph atelier
    Coordinateur
    Coordinateur--"Execute tache"-->Executor1
    Coordinateur--"Execute tache"-->Executor2
end

User--"Envoie tache"-->Coordinateur

```