# Atelier 5
### Créer une API REST

**Branche Atelier 5**

#### Objectifs

Savoir créer des API Rest
Savoir appeler des API Rest

#### Etapes

- Pour recevoir les demandes globales de traitements et le suivi des tâches
- Pour recevoir une tâche et la réaliser

```mermaid
    flowchart TD;
    
    subgraph coordinator
    TaskManager
    end
    
    user--Add task-->TaskManager
    TaskManager[/POST /task/]-->TaskDispatcher
    
    subgraph worker
        TaskDispatcher-->ResizeTask
        TaskDispatcher-->PrintTask
    end
```

#### Aide
