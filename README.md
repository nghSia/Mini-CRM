# Projet : Mini CRM

## Collaborateurs

FAZER Nino - TRAN Huu-Nghia

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
Mini-CRM/
│
├── main.go               # Point d’entrée de l’application
├── menu/
│   └── menu.go           # Menu interactif (Add, Get, Update, Delete)
├── crudContact/
│   └── crudContact.go    # Logique CRUD et gestion de la map users
└── user/
    └── user.go           # Définition de la struct User

# Exécution normale
go run .

## Cela démarre
=== Mini-CRM Menu ===
1) Ajouter un contact
2) Lister les contacts
3) Supprimer un contact
4) Mettre à jour un contact
5) Quitter

# Fonctionnalités 

## Ajout utilisateur 
### Ajout normal
1️⃣ Ajouter un contact
→ Entrer le nom :
→ Entrer l’email :
✅ Contact ajouté !

### Ajout depuis flag
go run . -name "test" -email "test@mail.com"

## Liste des utilisateurs
### Liste normale
2️⃣ Lister les contacts
📋 Liste des utilisateurs :
ID: 1 | Nom: Alice | Email: alice@mail.com
ID: 2 | Nom: Bob   | Email: bob@mail.com

### Liste depuis flag
go run . -userList

## Update utilisateur 
3️⃣ Mettre à jour un contact
→ Entrer l’ID du contact à modifier :
→ Entrer le nouveau nom :
→ Entrer le nouvel email :
✅ Utilisateur avec l’ID 1 mis à jour avec succès

## Delete utilisateur 
4️⃣ Supprimer un contact
→ Entrer l’ID du contact à supprimer :
✅ Utilisateur avec l’ID 2 supprimé avec succès

