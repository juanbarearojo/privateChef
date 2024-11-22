# Decisión Test Runner y Biblioteca de Aserciones 

---

## Selección de Biblioteca de Aserciones

**Descripción:**

Es necesario seleccionar una biblioteca adecuada para escribir y validar pruebas en el proyecto. La decisión debe basarse en las necesidades del equipo, priorizando simplicidad, funcionalidad y estándares de la comunidad para garantizar un flujo de trabajo eficiente y mantenible.

**Criterios de aceptación:**

- Si existe una biblioteca ampliamente adoptada por la comunidad de Go, esta será priorizada para garantizar compatibilidad y reducir la curva de aprendizaje.
- Debe ser fácil de configurar e integrar en el flujo de trabajo existente.
- La herramienta debe contar con soporte activo, actualizaciones frecuentes y documentación adecuada.

**Opciones de Bibliotecas de Aserciones:**

1. **`testing` (https://pkg.go.dev/testing):**  
   Biblioteca nativa de Go para escribir y ejecutar pruebas. Es simple y se integra perfectamente con `go test`, ideal para proyectos que priorizan la compatibilidad con las herramientas estándar.

2. **Gomega (https://onsi.github.io/gomega/):**  
   Una biblioteca de aserciones expresiva que complementa Ginkgo, aunque puede utilizarse de forma independiente. Proporciona una amplia gama de matchers y una sintaxis legible que mejora la claridad y el mantenimiento de las pruebas.

3. **Testify (https://github.com/stretchr/testify):**  
   Una biblioteca ampliamente adoptada en la comunidad de Go, con métodos avanzados para validar resultados. Ofrece soporte para mocks y suites de pruebas, siendo ideal para proyectos con pruebas complejas.

---

## Selección de Test Runner

**Descripción:**

Es necesario elegir un test runner adecuado para el proyecto, considerando que está desarrollado en Go. El objetivo es contar con una herramienta eficiente que permita ejecutar pruebas unitarias y de integración de manera automatizada, optimizando el flujo de trabajo del equipo.

**Criterios de aceptación:**

- El test runner debe tener soporte activo para garantizar actualizaciones y reducir la deuda técnica del proyecto.
- Debe soportar la ejecución paralela de pruebas para maximizar la eficiencia.
- Debe permitir la generación de reportes sencillos que faciliten la revisión de resultados.

**Opciones de Test Runner:**

1. **`go test` (https://pkg.go.dev/cmd/go#hdr-Test_packages):**  
   Herramienta oficial de Go para ejecutar pruebas unitarias e integración, con soporte para pruebas paralelas y configuración mínima.

2. **Ginkgo (https://onsi.github.io/ginkgo/):**  
   Un framework de pruebas maduro para Go que facilita la escritura de especificaciones expresivas. Ginkgo permite organizar y ejecutar pruebas de manera estructurada y expresiva, con soporte avanzado para suites de pruebas y perfiles de ejecución.

---

## Elección:

- **Biblioteca de Aserciones:**  
  Se usará `testing` como base por ser nativa y estándar. Gomega y Testify se usarán de manera complementaria para enriquecer las pruebas con una sintaxis más expresiva y funcionalidad adicional cuando sea necesario.

- **Test Runner:**  
  `go test` será el test runner principal por su simplicidad, soporte oficial y capacidad para ejecución paralela. Ginkgo se considerará como opción para casos donde se necesite una estructura más robusta y descriptiva para las pruebas.
