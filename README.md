# Membres du groupe : 
- Ulysse GUILLOT
- Philippe BONNAFOUS
- Nicolas THEAU

![gopher](./assets/gopher.gif)

### Commande générique :

```bash
go run main.go -action=<action> -name=<nom> -tel=<numéro>
```

### `add`
Ajoute un contact au carnet.

```bash
go run main.go -action=add -name=alice -tel=0123456789
```

---

### `search`
Recherche un contact par son nom.

```bash
go run main.go -action=search -name=alice
```

---

### `list`
Affiche tous les contacts enregistrés.

```bash
go run main.go -action=list
```

---

### `delete`
Supprime un contact du carnet.

```bash
go run main.go -action=delete -name=alice
```

---

### `exist`
Vérifie si un contact existe.

```bash
go run main.go -action=exist -name=alice
```

---

### `modify`
Modifie le numéro de téléphone d’un contact existant.

```bash
go run main.go -action=modify -name=alice -tel=0987654321
```

### Tests automatisés

```bash
go test
```

