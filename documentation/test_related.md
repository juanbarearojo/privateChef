# Decisión Test Runner y Biblioteca de Aserciones 


**Criterios de elección comunes para ambas herramientas:**

- En este caso como se establece en los Go Proverbs (https://go-proverbs.github.io/), que se han considerado como unas buenas prácticas para el lenguaje, "A little copying is better than a little dependency". Por lo general el lenguaje de Go incentiva a la no dependencia de librerías externas. Por esto se dará prioridad a herramientas que vengan instaladas por defecto con Go.
- En caso de tener una puntuación en https://snyk.io/advisor/golang será utilizada como un reflejo de la salud del proyecto. Esto ya que en teoría, snyk advisor da la puntuación para una herramienta esta basada en 4 conceptos (security, popularity, maintenance, community).
- Debe de tener una versión estable lanzada desde hace más de un año para poder ser utilizada. Esto con el objetivo de aunque haya herramientas prometodoras se busca no caer en el FOMO (Fear Of Missing Out). Que herramienta sea la más nueva no significa que sea la mejor. 

---


## Selección de Biblioteca de Aserciones

**Opciones de Bibliotecas de Aserciones:**

1. **Testing (https://pkg.go.dev/testing):**  
   Biblioteca nativa de Go para escribir y ejecutar pruebas. No hay métodos de aserción especializados sino que son usadas construcciones simples del 
   lenguaje. Es importante destacar que Go maneja errores de manera explícita ya que las funciones devuleven tanto un objeto como un error. 

2. **Gomega (https://onsi.github.io/gomega/):**  
   Una biblioteca de aserciones expresiva diseñada para trabajar junto con Ginkgo están basadas en el BDD, pero que también puede usarse con el paquete estándar. Gomega proporciona una amplia variedad de matchers y una sintaxis expresiva, mejorando la claridad y el mantenimiento de las pruebas.

3. **Testify (https://github.com/stretchr/testify):**  
   Una biblioteca ampliamente adoptada en la comunidad de Go que complementa `testing`. Incluye métodos especificos que incluyen contrucciones simples del lenguaje. Algunos de estos ejemplos es .Equal o .Nil. Basado en la metodología TDD. 

4. **Ghost (https://github.com/rliebz/ghost):**
   Ghost es una biblioteca de aserciones. Ghost proporciona cuatro métodos principales para realizar comprobaciones: Should, ShouldNot, Must y MustNot. Los métodos Should y ShouldNot verifican si una aserción ha tenido éxito, permitiendo que la prueba continúe incluso si la aserción falla, similar a t.Error. Por otro lado, Must y MustNot detienen la ejecución de la prueba si la aserción no pasa, de manera análoga a t.Fatal.La filosofía de esta herramienta  no obligarte, como hacen otras muchas biblitecas, a que en el momento que falle un test pare la ejecución. Te permite definir cuales son aquellas que verdaderamente son cruciales para tu programa.  

5. **GoCheck (https://labix.org/gocheck):**
   GoCheck es una extensión de la biblioteca testing de Go que proporciona funcionalidades avanzadas para pruebas unitarias y de integración. Aunque su adopción ha disminuido en comparación con bibliotecas como Testify, sigue siendo una opción robusta para quienes necesitan características adicionales sin dejar de utilizar el ecosistema estándar de Go. No busca sustuir sino complementar. 


---

## Selección de Test Runner

**Opciones de Test Runner:**

1. **`go test` (https://pkg.go.dev/cmd/go#hdr-Test_packages):**  
   Herramienta oficial de Go para ejecutar pruebas unitarias e integración, con soporte para pruebas paralelas y configuración mínima.

2. **Ginkgo (https://onsi.github.io/ginkgo/):**  
   Un framework de pruebas maduro para Go que incluye un **test runner** facilita la escritura de especificaciones expresivas. Ginkgo utiliza su propio ejecutable de línea de comandos, denominado ginkgo, como su test runner principal. Este ejecutable ofrece funcionalidades avanzadas para generar, ejecutar, filtrar y perfilar suites de pruebas escritas con Ginkgo. Basado en BDD.

---

## Elección:

**Selección de Biblioteca de Aserciones:**

Se selecciona `testing` como biblioteca de aserciones debido a que viene instalada por defecto con Go, lo que elimina la necesidad de instalaciones adicionales y simplifica la configuración. La gran mayoría de bibliotecas de aserciones para Go sirven como un complemento para la base del paquete estandar. Es verdad que hay bibliotecas que toman otro enfoque a la hora de hacer aserciones Bajo esta premisa se han puesto el resto de bibliotecas de aserciones para mostrar otras posibles opciones en el ecosistema de Go. 

**Selección de Test Runner:**

El test runner seleccionado es `go test`, ya que es la herramienta estándar incluida por defecto con Go, lo que prioriza un criterio clave: evitar herramientas que requieran instalación adicional. El hecho de que tenga soporte oficial lo hacen la mejor opción para el proyecto. Aunque herramientas como Ginkgo son viables para proyectos que se busque usar BDD. Al descartar Gomega que esta estrechamente relacionada con Ginkgo es otro motivo. Esto debido a estar pensadas para usar en conjunto. Hay muchos frameworks pero la gran mayoría usan go test por debajo lo que refuerza la elección. 
