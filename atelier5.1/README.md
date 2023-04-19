# Atelier 5.1
### Connecter le coordinateur au(x) worker(s)

#### Objectifs

* Savoir gérer un pool de server

#### Etapes

Le coordinateur va recevoir des requêtes et les déléguer à un pool de worker. 
Voici les différentes étapes

- Lorsqu'un worker se connecte, il connaît l'url du coordinateur et s'y connecte
- Lorsqu'une tâche est reçue, le coordinateur va la déléguer à un des workers connectés
- Lorsque le worker a fini la tâche, il va notifier le coordinateur

Le code à ajouter : 
- Sur le coordinateur, ajouter une route _POST_ **/register** où le worker indique son url
- Implémenter une version simple du pool worker pour exécuter une tâche (pas toujours le même worker)
  - L'interface **TaskSenderToWorker** permet de décorréler la gestion du pool de l'action d'envoi (notamment pour mocker) 
  - Lors du démarrage d'un worker, s'enregistrer auprès du coordinateur
  - Modifier le test TestCompleteChain afin de tester la chaine : lancement du coordinateur, du worker et ajout d'une tache print jusqu'à la notification

#### Aide

* Pour trier une liste : sort.Slice