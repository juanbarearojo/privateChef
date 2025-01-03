# Decisión Test Runner y Biblioteca de Aserciones 


**Criterios de elección comunes para ambas herramientas:**

- En este caso como se establece en los Go Proverbs (https://go-proverbs.github.io/), que se han considerado como unas buenas prácticas para el lenguaje, "A little copying is better than a little dependency". Por lo general el lenguaje de Go incentiva a la no dependencia de librerías externas. Por esto se dará prioridad a herramientas que vengan instaladas por defecto con Go. Es decir si una herramienta está en Standard library(https://pkg.go.dev/std) le será dado preferencia.
- En caso de tener una puntuación en https://snyk.io/advisor/golang será utilizada como un reflejo de la salud del proyecto. Esto ya que en teoría, snyk advisor da la puntuación para una herramienta esta basada en 4 conceptos (security, popularity, maintenance, community).
- Debe de tener una versión estable lanzada desde hace más de un año para poder ser utilizada. 

---

## Selección de Biblioteca de Aserciones

Antes de nada creo que es necesario explicar un concepto sobre como Go trabaja el tema de los errores. Go no tiene aserciones de forma nativa. Es importante destacar que Go maneja errores de manera explícita ya que las funciones devuelven tanto un objeto como un valor de error. Por lo que en Go es posible trabajar usando los errores que esté da como opción base o con biblitecas externas.

**Opciones de Bibliotecas de Aserciones:**

1. **Gomega (https://onsi.github.io/gomega/):**  
   Una biblioteca de aserciones diseñada para trabajar junto con Ginkgo, pero que también puede usarse con 'go test'. En snyk advisor cuenta con 96/100 (https://snyk.io/advisor/golang/github.com/onsi/gomega) lo que indica un buen estado de salud. La versión más actual disponible en su repositorio es la v1.36.0(https://github.com/onsi/gomega/releases/tag/v1.36.0) pero cuenta con varias versiones estables lanzadas este año. Al no ser parte de la Standard library es necesario instalación adicional.

2. **Testify (https://github.com/stretchr/testify):**  
   Una biblioteca ampliamente adoptada en la comunidad de Go. Incluye métodos especificos para tipos de aserciones comunes. Es importante resaltar que en snyk advisor cuando se busca testify (https://snyk.io/advisor/search?source=golang&q=testify) aparecen solo reseñas que apuntan a forks del repo original, no hay ninguna que apunta al repositorio de stretch que son los autores de la biblioteca. En este caso se ha escogido esta reseña debido a ser la más reciente (https://snyk.io/advisor/golang/github.com/01ne/testify) con una salud de 78/100. La última version es la v1.10.0(https://github.com/stretchr/testify/releases/tag/v1.10.0) lanzada este año. Es necesario instalación al no ser parte de la Standard library. 

3. **Be (https://pkg.go.dev/github.com/rliebz/ghost/be):**
   Be es una biblioteca de aserciones bastante nueva. Su creador dice en la documentación que está pensada para ser usada con ghost (https://github.com/rliebz/ghost). Proporciona un conjunto de aserciones básicas para trabajar. Aunque sea prometodora su versión mas reciente es la v0.2.0(https://pkg.go.dev/github.com/rliebz/ghost/be?tab=versions) lo que nos indica que todavía le falta madurez al proyecto. En snyk advisor todavía no hay ninguna reseña a esta herramienta. Es necesario instalación adicional al no ser parte de la Standard library.

4. **GoCheck (https://labix.org/gocheck):**
   GoCheck es una extensión de la biblioteca testing de Go que proporciona funcionalidades avanzadas para pruebas unitarias. En snyk advisor cuenta con una salud de 62/100 (https://snyk.io/advisor/golang/gopkg.in/check.v1) pero se establece que su maintenance es "INACTIVE" cosa que podemos comprobar si nos metemos en el repositorio que lleva más de 4 años sin tener un commit(https://github.com/go-check/check). Es necesario instalación adicional al no ser parte de la Standard library.


---

## Selección de Test Runner

**Opciones de Test Runner:**

1. **Testing (https://pkg.go.dev/testing):**  
   Test runner de Go. Es la herramienta oficial de Go para ejecutar test. Forma parte de la Standard library de Go. `go test` ((https://pkg.go.dev/cmd/go#hdr-Test_packages) es el comando a través del cual el test runner es ejecutado. 

2. **Ginkgo (https://onsi.github.io/ginkgo/):**  
   Un framework de pruebas para Go que incluye un **test runner**. ginkgo es también es el nombre del test runner y también el comando a través del cual el test runner es ejecutado. En el propio repositorio se menciona que esta pensado para ser usado en conjunto con Gomega. Podemos ver que su última versión es la v2.22.0(https://github.com/onsi/ginkgo/releases/tag/v2.22.0) además de contar con varias versiones estables lanzadas este año. No encontré una reseña en snyk advisor para esta herramienta.

---

## Elección:

**Selección de Biblioteca de Aserciones:**

En este caso se decide que no se va trabajar con ninguna biblioteca de aserciones. Se va a trabajar con el manejo de errores que Go da forma nativa. Todo esto con el objetivo de respetar el primer criterio establecedio. Al ser elementos del propio lenguaje se asegura también salud y una versión estable.

**Selección de Test Runner:**

El test runner seleccionado es testing, ya que es la herramienta estándar incluida por defecto con Go, lo que prioriza un criterio clave: evitar herramientas que requieran instalación adicional. Al formar parte de la Standard library cumple los otros dos criterios teniendo salud y una versión estable.
