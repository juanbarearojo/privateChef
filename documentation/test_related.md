# Decisión Test Runner y Biblioteca de Aserciones 


**Criterios de elección comunes para ambas herramientas:**

- En este caso como se establece en los Go Proverbs (https://go-proverbs.github.io/), que se han considerado como unas buenas prácticas para el lenguaje, "A little copying is better than a little dependency". Por lo general el lenguaje de Go incentiva a la no dependencia de librerías externas. Por esto se dará prioridad a herramientas que vengan instaladas por defecto con Go. Es decir si una herramienta está en Standard library(https://pkg.go.dev/std) le será dado preferencia.
- En caso de tener una puntuación en https://snyk.io/advisor/golang será utilizada como un reflejo de la salud del proyecto. Esto ya que en teoría, snyk advisor da la puntuación para una herramienta esta basada en 4 conceptos (security, popularity, maintenance, community).
- Debe de tener una versión estable lanzada desde hace más de un año para poder ser utilizada. Esto con el objetivo de aunque haya herramientas prometodoras, se busca no caer en el FOMO (Fear Of Missing Out). Que una herramienta sea la más nueva no significa que sea la mejor. 

---


## Selección de Biblioteca de Aserciones

Antes de nada creo que es necesario explicar un concepto sobre como Go trabaja el tema de las aserciones. Es importante destacar que Go maneja errores de manera explícita ya que las funciones devuelven tanto un objeto como un valor de error. Por lo que en Go es posible trabajar usando los errores que esté da como opción base o con  biblitecas externas.

**Opciones de Bibliotecas de Aserciones:**

1. **Testing (https://pkg.go.dev/testing):**  
   Biblioteca nativa de Go para escribir y ejecutar pruebas. No hay métodos de aserción especializados sino que son usadas construcciones simples del 
   lenguaje. Esta es la forma más básica de trabajar los errores en Go. Está dentro del Standard library. 

2. **Gomega (https://onsi.github.io/gomega/):**  
   Una biblioteca de aserciones diseñada para trabajar junto con Ginkgo están basadas en el BDD, pero que también puede usarse con 'go test'. Gomega proporciona una forma de trabajar con aserciones. En snyk advisor cuenta con 96/100 (https://snyk.io/advisor/golang/github.com/onsi/gomega) lo que indica un buen estado de salud.
   La versión más actual disponible en su repositorio es la v1.36.0(https://github.com/onsi/gomega/releases/tag/v1.36.0) que pero cuenta con varias versiones estables lanzadas este año.

3. **Testify (https://github.com/stretchr/testify):**  
   Una biblioteca ampliamente adoptada en la comunidad de Go que complementa `testing`. Incluye métodos especificos para tipos de aserciones comunes. Algunos de estos ejemplos es .Equal o .Nil. Pensado para trabajar con la metodología TDD. Es importante resaltar que en snyk advisor cuando se busca testify (https://snyk.io/advisor/search?source=golang&q=testify) aparecen solo reseñas que apuntan a forks del repo original, no hay ninguna que apunta al repositorio de stretch que son los autores de la biblioteca. En este caso se ha escogido esta reseña debido a ser la más reciente (https://snyk.io/advisor/golang/github.com/01ne/testify) con una salud de 78/100. Es necesario instalación al no ser parte de la Standard library. La última version es la v1.10.0(https://github.com/stretchr/testify/releases/tag/v1.10.0) lanzada este año.

4. **Ghost (https://github.com/rliebz/ghost):**
   Ghost es una biblioteca de aserciones bastante nueva que permiten un enfoque algo distinto para trabajar con aserciones. Ghost proporciona cuatro métodos principales para realizar comprobaciones: Should, ShouldNot, Must y MustNot. Los métodos Should y ShouldNot verifican si una aserción ha tenido éxito, permitiendo que la prueba continúe incluso si la aserción falla, similar a t.Error. Por otro lado, Must y MustNot detienen la ejecución de la prueba si la aserción no pasa, de manera análoga a t.Fatal. La filosofía de esta herramienta  no obligarte, como hacen otras muchas biblitecas, a que en el momento que falle un test pare la ejecución. Aunque sea prometodora su versión mas reciente es la v0.2.0(https://github.com/rliebz/ghost/releases/tag/v0.2.0) lo que nos indica que todavía le falta madurez al proyecto. En snyk advisor todavía no hay ninguna reseña a esta herramienta.

5. **GoCheck (https://labix.org/gocheck):**
   GoCheck es una extensión de la biblioteca testing de Go que proporciona funcionalidades avanzadas para pruebas unitarias y de integración. Su uso ha bajado en comparación con bibliotecas como Testify. En snyk advisor cuenta con una salud de 62/100 (https://snyk.io/advisor/golang/gopkg.in/check.v1) pero se establece que su maintenance es "INACTIVE" cosa que podemos comprobar si nos metemos en el repositorio que lleva más de 4 años sin tener un commit(https://github.com/go-check/check). Es necesario instalación adicional al no ser parte de la Standard library.


---

## Selección de Test Runner

**Opciones de Test Runner:**

1. **`go test` (https://pkg.go.dev/cmd/go#hdr-Test_packages):**  
   Establecer que `go test` es el comando a través del cual el test runner es ejecutado ya que forma parte de Testing que forma parte de la Standard library. Supone la herramienta oficial de Go para ejecutar pruebas unitarias e integración. En el ecosistema de Go es el test runner más utilizado de todos. Al ser parte del propio lenguaje va a contar con soporte para el mismo. 

2. **Ginkgo (https://onsi.github.io/ginkgo/):**  
   Un framework de pruebas maduro para Go que incluye un **test runner** facilita la escritura de especificaciones expresivas. Ginkgo utiliza su propio ejecutable de línea de comandos, denominado ginkgo, como su test runner principal. En el propio repositorio se menciona que esta pensado para ser usado en conjunto con Gomega.  También es necesario decir que favorece el desarollo basado en la metodología BDD. Podemos ver que su última versión es la v2.22.0(https://github.com/onsi/ginkgo/releases/tag/v2.22.0) además de contar con varias versiones estables lanzadas este año.

---

## Elección:

**Selección de Biblioteca de Aserciones:**

Se selecciona `testing` como biblioteca de aserciones debido a que viene instalada por defecto con Go, lo que elimina la necesidad de instalaciones adicionales y simplifica la configuración. La gran mayoría de bibliotecas de aserciones para Go sirven como un complemento para la base del paquete estandar. Es verdad que hay bibliotecas que toman otro enfoque a la hora de hacer aserciones. Todo esto con el objetivo de respetar el primer criterio establecedio. También cumple los otros dos criterios al ser parte del propio lenguaje por lo que aseguramos salud y versión estable.

**Selección de Test Runner:**

El test runner seleccionado es `go test`, ya que es la herramienta estándar incluida por defecto con Go, lo que prioriza un criterio clave: evitar herramientas que requieran instalación adicional. El hecho de que tenga soporte oficial lo hacen la mejor opción para el proyecto. Hay muchos frameworks  para test pero la gran mayoría usan go test por debajo lo que refuerza la elección. 
