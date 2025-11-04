# Projet : Mini CRM

## Collaborateurs
TRAN Huu-Nghia

---

## 🧠 Description

Mini-CRM est une application minimale de gestion de contacts développée en **Go**.  
Elle permet d’ajouter, afficher, mettre à jour et supprimer des utilisateurs via un **menu interactif**, ou directement en ligne de commande à l’aide de **flags**.  
Les données sont stockées dans une **map en mémoire** et sont perdues à chaque fermeture du programme.

---

## ⚙️ Prérequis

- Go 1.18+
- Git (optionnel)

Vérifie ton installation de Go :

```bash
go version
```

# Structure du projet
```bash
Mini-CRM/
│
├── go.mod                # Fichier de configuration du module Go
├── main.go               # Point d’entrée de l’application
├── main_test.go          # Tests unitaires pour main.go
│
├── internal/
│   ├── app/
│   │   └── app.go        # Logique principale de l’application
│   │
│   └── storage/
│       ├── memory.go     # Stockage en mémoire (implémentation)
│       └── storage.go    # Interface de stockage et logique associée
│
└── README.md             # Documentation du projet
```
# Exécution normale
go run .

## Cela démarre
=== Mini-CRM Menu ===
1) Ajouter un contact
2) Lister les contacts
3) Lister les informaton d'un seul contact
4) Mettre à jour un contact
5) Supprimer un contact
6) Quitter

# Fonctionnalités 

## Ajout utilisateur 
### Ajout normal
```bash
1️⃣ Ajouter un contact
→ Entrer le nom :
→ Entrer l’email :
✅ Contact ajouté !
```

### Ajout depuis flag
```bash
go run . -name "test" -email "test@mail.com"
```

## Liste des utilisateurs
### Liste normale
```bash
2️⃣ Lister les contacts
📋 Liste des utilisateurs :
ID: 1 | Nom: Alice | Email: alice@mail.com
ID: 2 | Nom: Bob   | Email: bob@mail.com
```

## Update utilisateur 
```bash
3️⃣ Mettre à jour un contact
→ Entrer l’ID du contact à modifier :
→ Entrer le nouveau nom :
→ Entrer le nouvel email :
✅ Utilisateur avec l’ID 1 mis à jour avec succès
```

## Delete utilisateur 
```bash
4️⃣ Supprimer un contact
→ Entrer l’ID du contact à supprimer :
✅ Utilisateur avec l’ID 2 supprimé avec succès
```
