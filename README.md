# Projet : Mini CRM

## Collaborateurs
TRAN Huu-Nghia

---

## 🧠 Description

Mini-CRM est une application minimale de gestion de contacts développée en **Go**.  
Elle permet d'ajouter, afficher, mettre à jour et supprimer des utilisateurs via un **menu interactif**, ou directement en ligne de commande à l'aide de **flags**.  

**Persistance des données :**
- Les données sont maintenant **sauvegardées automatiquement** dans un fichier `contacts.json` à la racine du projet
- L'application charge automatiquement les contacts existants au démarrage
- Toutes les modifications (ajout, mise à jour, suppression) sont immédiatement persistées dans le fichier JSON
- L'ancienne implémentation en mémoire (`MemoryStore`) est conservée mais n'est plus utilisée par défaut

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
├── go.sum                # Fichier de dépendances
├── contacts.json         # 💾 Fichier de persistance des contacts (généré automatiquement)
├── main.go               # Point d'entrée de l'application
├── main_test.go          # Tests unitaires pour main.go
│
├── cmd/                  # Commandes Cobra CLI
│   ├── root.go           # Commande racine (initialise JSONStore)
│   ├── add.go            # Commande pour ajouter un contact
│   ├── update.go         # Commande pour mettre à jour un contact
│   ├── delete.go         # Commande pour supprimer un contact
│   ├── get.go            # Commande pour obtenir un contact par ID
│   └── getAll.go         # Commande pour lister tous les contacts
│
├── internal/
│   ├── app/
│   │   └── app.go        # Logique métier et handlers
│   │
│   ├── config/
│   │   └── config.go     # Gestion de la sérialisation/désérialisation JSON
│   │
│   └── storage/
│       ├── storage.go    # Interface Storer et définition Contact
│       ├── memory.go     # ⚠️ Implémentation en mémoire (conservée mais non utilisée)
│       └── json.go       # ✅ Implémentation avec persistance JSON (utilisée par défaut)
│
└── README.md             # Documentation du projet
```

**Note sur l'architecture :**
- L'interface `Storer` permet de basculer facilement entre différentes implémentations de stockage
- `JSONStore` est actuellement utilisé par défaut (voir `cmd/root.go` ligne 38)
- `MemoryStore` est conservé pour référence ou tests mais n'est plus le store par défaut
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
=====================
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
→ Entrer l'ID du contact à supprimer :
✅ Utilisateur avec l'ID 2 supprimé avec succès
```

---

## 🚀 CLI avec Cobra - Guide des Commandes

L'application Mini-CRM est maintenant disponible en tant qu'outil CLI utilisant **Cobra**. Vous pouvez l'utiliser de deux manières :
- **Mode interactif** : L'application vous guide avec des prompts
- **Mode CLI** : Utilisation directe avec des sous-commandes et flags

### 💾 Persistance des données

**Toutes les opérations sont automatiquement sauvegardées** :
- Le fichier `contacts.json` est créé automatiquement à la racine du projet au premier ajout
- Les contacts sont chargés automatiquement au démarrage de l'application
- Chaque modification (ajout, mise à jour, suppression) est immédiatement persistée
- Les données survivent à la fermeture de l'application

**Emplacement du fichier :**
```bash
# Le fichier est créé dans le répertoire de travail actuel
./contacts.json

# Exemple de contenu :
[
  {
    "id": 1,
    "name": "Alice Martin",
    "email": "alice@mail.com"
  },
  {
    "id": 2,
    "name": "Bob Smith",
    "email": "bob@company.com"
  }
]
```

### Compilation de l'exécutable

```bash
# Compiler l'application
go build -o gomincrm

# Rendre l'exécutable (Unix/Linux/macOS)
chmod +x gomincrm

# Exécuter
./gomincrm
```

### 📋 Aide et Documentation

```bash
# Aide générale - Liste toutes les commandes disponibles
./gomincrm --help
./gomincrm -h

# Aide sur une commande spécifique
./gomincrm [commande] --help
./gomincrm [commande] -h

# Exemples :
./gomincrm add --help
./gomincrm update --help
./gomincrm delete --help
```

---

## Sous-commandes disponibles

### 1️⃣ **add** - Ajouter un contact

Ajoute un nouveau contact au système CRM.

**2 modes d'utilisation :**

#### Mode interactif (sans flags)
```bash
./gomincrm add
```
→ L'application vous demandera le nom et l'email

#### Mode avec flags
```bash
./gomincrm add -n "Nom" -e "email@example.com"
./gomincrm add --name "Nom" --email "email@example.com"
```

**Flags disponibles :**
| Flag | Raccourci | Description | Obligatoire |
|------|-----------|-------------|-------------|
| `--name` | `-n` | Nom du contact | Oui (en mode flags) |
| `--email` | `-e` | Email du contact | Oui (en mode flags) |

**Exemples d'utilisation :**
```bash
./gomincrm add                                      # Mode interactif
./gomincrm add -n "Alice" -e "alice@mail.com"       # Mode flags
./gomincrm add --name "Bob" --email "bob@test.com"  # Mode flags (format long)
```

---

### 2️⃣ **list** - Lister tous les contacts

Affiche la liste complète de tous les contacts enregistrés.

**Utilisation :**
```bash
./gomincrm list
```

**Flags disponibles :**
Aucun flag pour cette commande.

---

### 3️⃣ **get** - Obtenir un contact par ID

Affiche les informations détaillées d'un contact spécifique.

**2 modes d'utilisation :**

#### Mode avec argument
```bash
./gomincrm get [ID]
```

#### Mode interactif (sans argument)
```bash
./gomincrm get
```
→ L'application vous demandera l'ID du contact

**Exemples d'utilisation :**
```bash
./gomincrm get 1      # Affiche le contact avec l'ID 1
./gomincrm get 5      # Affiche le contact avec l'ID 5
./gomincrm get        # Mode interactif
```

---

### 4️⃣ **update** - Mettre à jour un contact

Met à jour le nom et/ou l'email d'un contact existant.

**2 modes d'utilisation :**

#### Mode interactif (sans flags)
```bash
./gomincrm update
```
→ L'application vous guidera pour entrer l'ID et les nouvelles informations

#### Mode avec flags
```bash
./gomincrm update -i [ID] -n "Nouveau nom" -e "nouvel@email.com"
./gomincrm update --id [ID] --name "Nouveau nom" --email "nouvel@email.com"
```

**Flags disponibles :**
| Flag | Raccourci | Description | Obligatoire |
|------|-----------|-------------|-------------|
| `--id` | `-i` | ID du contact à mettre à jour | Oui (en mode flags) |
| `--name` | `-n` | Nouveau nom du contact | Non* |
| `--email` | `-e` | Nouvel email du contact | Non* |

*Au moins un des deux champs (name ou email) doit être fourni

**Exemples d'utilisation :**
```bash
./gomincrm update                                    # Mode interactif
./gomincrm update -i 1 -n "Jane"                     # Mettre à jour uniquement le nom
./gomincrm update -i 1 -e "jane@newmail.com"         # Mettre à jour uniquement l'email
./gomincrm update -i 1 -n "Jane" -e "jane@mail.com"  # Mettre à jour les deux
```

---

### 5️⃣ **delete** - Supprimer un contact

Supprime un contact du système.

**2 modes d'utilisation :**

#### Mode avec argument
```bash
./gomincrm delete [ID]
```

#### Mode interactif (sans argument)
```bash
./gomincrm delete
```
→ L'application vous demandera l'ID du contact à supprimer

**Flags disponibles :**
| Flag | Raccourci | Description | Obligatoire |
|------|-----------|-------------|-------------|
| `--id` | `-i` | ID du contact à supprimer | Non (peut être passé comme argument) |

**Exemples d'utilisation :**
```bash
./gomincrm delete 1       # Supprime le contact avec l'ID 1
./gomincrm delete 5       # Supprime le contact avec l'ID 5
./gomincrm delete         # Mode interactif
./gomincrm delete -i 3    # Avec flag (alternative)
```

---

## 📊 Exemple de workflow complet

```bash
# 1. Ajouter plusieurs contacts
./gomincrm add -n "Alice Martin" -e "alice@mail.com"
./gomincrm add -n "Bob Smith" -e "bob@company.com"
./gomincrm add -n "Charlie Brown" -e "charlie@test.com"

# 2. Lister tous les contacts
./gomincrm list

# 3. Voir les détails d'un contact spécifique
./gomincrm get 2

# 4. Mettre à jour un contact
./gomincrm update -i 2 -e "bob.smith@newcompany.com"

# 5. Mettre à jour nom et email
./gomincrm update -i 1 -n "Alice Johnson" -e "alice.j@newmail.com"

# 6. Supprimer un contact
./gomincrm delete 3

# 7. Vérifier la liste finale
./gomincrm list
```

---

## � Récapitulatif des commandes

| Commande | Description | Mode interactif | Flags principaux |
|----------|-------------|-----------------|------------------|
| `add` | Ajouter un contact | ✅ | `-n`, `-e` |
| `list` | Lister tous les contacts | ❌ | Aucun |
| `get [ID]` | Obtenir un contact | ✅ | Argument ID |
| `update` | Mettre à jour un contact | ✅ | `-i`, `-n`, `-e` |
| `delete [ID]` | Supprimer un contact | ✅ | Argument ID ou `-i` |

---

## 🔑 Points importants

- **Mode interactif** : Lancez la commande sans flags, l'application vous guidera
- **Mode CLI** : Utilisez les flags pour des opérations rapides ou de l'automatisation
- **Aide contextuelle** : Utilisez `--help` ou `-h` après n'importe quelle commande pour voir sa documentation
- Les **IDs** sont générés automatiquement et commencent à 1
- Les **données** sont stockées dans `contacts.json` et **persistantes entre les sessions**
- Le fichier JSON est créé automatiquement dans le répertoire de travail actuel
- Tous les **messages** sont en anglais

### 🔄 Changement de mode de stockage

Si vous souhaitez revenir au mode en mémoire (non persistant) :
1. Ouvrez `cmd/root.go`
2. Ligne 38, remplacez `storage.NewJsonStore()` par `storage.NewMemoryStore()`
3. Recompilez avec `go build -o gomincrm`

**Comparaison des modes :**
| Mode | Fichier | Persistance | Utilisation |
|------|---------|-------------|-------------|
| `JSONStore` | `json.go` | ✅ Oui (contacts.json) | **Par défaut** |
| `MemoryStore` | `memory.go` | ❌ Non (perdu à la fermeture) | Tests/Développement |
