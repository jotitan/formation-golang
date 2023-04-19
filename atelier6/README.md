# Atelier 6
### Gestion de la concurrence

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
* Générer un identifiant unique (uuid) lors de la création d'un worker
* Lors de l'enregistrement du worker auprès du coordinateur, transmettre l'uuid
* Dans le status d'une tâche, permettre le stockage de l'identifiant du noeud qui le traite
* Lors qu'un worker indique qu'une tâche est terminée (ou en erreur), transmettre l'uuid
* Le status running n'est pas utilisé actuellement (car la tâche est envoyée immédiatement au worker) : voir pour intégrer ce statut
* Il va falloir faire en sorte que le taskManager et le PoolWorker communique entre eux. Ajouter des tests.

#### Aide

* Pour générer un UUID, la librairie de [google](github.com/google/uuid) est facile à utiliser
* Rappel sur les channels : 
  * Pour limiter les appels en parallèle, on peut utiliser un channel avec une taille correspondant au nombre de tâches en parallèle
  * Pour échanger des données de manière threadsafe, le channel est idéal