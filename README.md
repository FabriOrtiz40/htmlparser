# HTML Link Parser (Gophercise #4)

Este proyecto implementa un **parser de links HTML en Go**, capaz de extraer todas las etiquetas `<a>` de un documento HTML y devolver, para cada una, su `href` y el texto visible asociado.

El ejercicio forma parte de **Gophercises** y está enfocado en aprender a trabajar con:
- parsing de HTML real
- estructuras en forma de árbol
- recorridos DFS
- separación de responsabilidades
- testing orientado a comportamiento

---

## 📌 Objetivo

Dado un HTML, obtener una estructura como:

```go
type Link struct {
    Href string
    Text string
}
```

Ejemplo de entrada:

```html
<a href="/dog">
  <span>Dog</span>
  and
  <b>Cat</b>
</a>
```

Salida esperada:

```go
Link{
    Href: "/dog",
    Text: "Dog and Cat",
}
```

---

## 🧠 Enfoque y diseño

### 1. HTML como árbol
Se utiliza el paquete:

```go
golang.org/x/net/html
```

para parsear el HTML en un **árbol de nodos**, en lugar de trabajar con texto plano o expresiones regulares.

---

### 2. Recorrido DFS
El árbol HTML se recorre usando **Depth-First Search (DFS)** para:
- visitar todos los nodos
- detectar etiquetas `<a>`
- evitar links anidados

---

### 3. Separación de responsabilidades

| Función | Responsabilidad |
|--------|-----------------|
| Parse | Punto de entrada y orquestación |
| extractLinks | Recorrer el árbol y detectar `<a>` |
| extractText | Extraer texto visible de un nodo |
| cleanText | Normalizar whitespace |

---

## 📂 Estructura del proyecto

```
htmlparser/
├── go.mod
└── link/
    ├── link.go
    └── link_test.go
```

---

## 🚀 Uso

```go
html := `<a href="/dog">Dog</a>`

links, err := link.Parse(strings.NewReader(html))
if err != nil {
    log.Fatal(err)
}

fmt.Println(links)
```

---

## ⚠️ Reglas implementadas

- Se extrae el atributo `href`
- Se concatena todo el texto visible
- Se eliminan saltos de línea, tabs y espacios duplicados
- Se ignoran comentarios HTML
- Se ignoran links anidados

---

## 🧪 Testing

```bash
go test ./link -v
```

o

```bash
go test ./...
```

---

## 📚 Dependencias

```bash
go get golang.org/x/net/html
```

---

## ✅ Estado

- Implementación completa
- Tests pasando
- Bonus de nested links resuelto
