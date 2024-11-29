# Decisión Test Runner y Biblioteca de Aserciones 


**Criterios de elección comunes para ambas herramientas:**

- Compatibilidad Nativa: Dar prioridad a herramientas que vengan instaladas por defecto con Go para evitar dependencias externas y configuraciones adicionales.
- El test runner debe tener soporte activo para garantizar actualizaciones y reducir la deuda técnica del proyecto.
- Haber tenido en el último año una actualización para asegurar la salud del proyecto.
- Tener suficente madurez el proyecto para asegurar una versión estable del mismo
- Salud del Proyecto: Preferir herramientas con una salud superior a 50 en Snyk Advisor para minimizar riesgos de mantenimiento o problemas de seguridad.

---



## Selección de Biblioteca de Aserciones

**Descripción:**

Es necesario seleccionar una biblioteca adecuada para escribir y validar pruebas en el proyecto. La decisión debe basarse en las necesidades del equipo, priorizando simplicidad, funcionalidad y estándares de la comunidad para garantizar un flujo de trabajo eficiente y mantenible.

**Opciones de Bibliotecas de Aserciones:**

1. **Testing (https://pkg.go.dev/testing):**  
   Biblioteca nativa de Go para escribir y ejecutar pruebas. Es simple, se integra perfectamente con `go test`, y es ideal para proyectos que priorizan la compatibilidad con las herramientas estándar.

2. **Gomega (https://onsi.github.io/gomega/):**  
   Una biblioteca de aserciones expresiva diseñada para trabajar junto con Ginkgo, pero que también puede usarse con el paquete estándar `testing`. Gomega proporciona una amplia variedad de matchers y una sintaxis expresiva, mejorando la claridad y el mantenimiento de las pruebas.

3. **Testify (https://github.com/stretchr/testify):**  
   Una biblioteca ampliamente adoptada en la comunidad de Go que complementa `testing`. Incluye aserciones, soporte para mocks, y suites de pruebas, siendo ideal para proyectos con necesidades avanzadas en pruebas unitarias.

4. **Ghost (https://github.com/rliebz/ghost):**
   Ghost es una biblioteca de aserciones. Ghost proporciona cuatro métodos principales para realizar comprobaciones: Should, ShouldNot, Must y MustNot. Los métodos Should y ShouldNot verifican si una aserción ha tenido éxito, permitiendo que la prueba continúe incluso si la aserción falla, similar a t.Error. Por otro lado, Must y MustNot detienen la ejecución de la prueba si la aserción no pasa, de manera análoga a t.Fatal


---

## Selección de Test Runner

**Descripción:**

Es necesario elegir un test runner adecuado para el proyecto, considerando que está desarrollado en Go. El objetivo es contar con una herramienta eficiente que permita ejecutar pruebas unitarias y de integración de manera automatizada, optimizando el flujo de trabajo del equipo.


**Opciones de Test Runner:**

1. **`go test` (https://pkg.go.dev/cmd/go#hdr-Test_packages):**  
   Herramienta oficial de Go para ejecutar pruebas unitarias e integración, con soporte para pruebas paralelas y configuración mínima.

2. **Ginkgo (https://onsi.github.io/ginkgo/):**  
   Un framework de pruebas maduro para Go que incluye un **test runner** facilita la escritura de especificaciones expresivas. Ginkgo utiliza su propio ejecutable de línea de comandos, denominado ginkgo, como su test runner principal. Este ejecutable ofrece funcionalidades avanzadas para generar, ejecutar, filtrar y perfilar suites de pruebas escritas con Ginkgo.

3. **gotestsum (https://github.com/gotestyourself/gotestsum):**  
   Gotestsum es una herramienta para ejecutar pruebas en Go que utiliza go test -json para proporcionar una salida formateada y un resumen de los resultados de las pruebas. Está diseñada tanto para el desarrollo local como para su integración en sistemas de integración continua

4. **gotestdox (https://github.com/bitfield/gotestdox):**
   Gotestdox es una herramienta de línea de comandos que ejecuta pruebas en Go y formatea los resultados en una documentación legible. Convierte los nombres de las pruebas escritas en notación CamelCase en oraciones descriptivas, facilitando la comprensión del comportamiento que se está probando

5. **GoConvey (https://github.com/smartystreets/goconvey):**  
   GoConvey es una herramienta que combina biblioteca de aserciones, framework de pruebas y test runner.Ofrece tanto una biblioteca de aserciones como un servidor web para visualizar resultados de pruebas en tiempo real. Se integra con `go test` y permite escribir pruebas de comportamiento. Es útil para proyectos que requieren una interfaz visual para revisar resultados.



---

## Elección:

**Selección de Biblioteca de Aserciones:**

Se selecciona `testing` como biblioteca de aserciones debido a que viene instalada por defecto con Go, lo que elimina la necesidad de instalaciones adicionales y simplifica la configuración. Se descarta la opción de `Ghost` debido a su baja salud en Snyk Advisor, lo que plantea dudas sobre su mantenimiento y estabilidad futura.

**Selección de Test Runner:**

El test runner seleccionado es `go test`, ya que es la herramienta estándar incluida por defecto con Go, lo que prioriza un criterio clave: evitar herramientas que requieran instalación adicional. Su simplicidad, soporte oficial y capacidad para ejecución paralela lo convierten en la mejor opción para el proyecto. Aunque herramientas como Ginkgo y Gotestdox son viables para proyectos que requieren estructuras más avanzadas, se elige `go test` por ser la solución oficial y por preferencias del programador, lo que asegura compatibilidad nativa y una experiencia más eficiente.
