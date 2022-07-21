# Atelier 6
### Gestion de la concurence

**Branche Atelier 6**

#### Objectifs

* _Distribuer_ nos tâches sur plusieurs instances et les appeler intelligemment.
* _Sécuriser_ l’usage en gérant un timeout

#### Etapes

Actuellement, le coordinateur n'est pas très intelligent : quand il reçoit une tâche, il la dispatche immédiatement à un worker.
Ce mécanisme, le round robin, a plusieurs inconvénients : 
* Il ne prend pas en compte la capacité de traitement du worker
* Il n'est pas possible de limiter le nombre de requêtes sur un worker

Nous allons changer le mécanisme en utilisant des channels

Voici les étapes : 
* Générer un identifiant unique lors de l'ajout d'un worker dans le pool
* Dans le status d'une tâche, permettre le stockage de l'identifiant du noeud qui le traite
* 

#### Aide