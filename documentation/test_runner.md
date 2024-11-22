# Decisión test Runner 

**Descripción:**

Es necesario elegir un test runner adecuado para el proyecto, considerando que está desarrollado en Go. El objetivo es contar con una herramienta eficiente que permita ejecutar pruebas unitarias y de integración de manera automatizada, optimizando el flujo de trabajo del equipo.

**Criterios de aceptación:**

- El test runner debe tener soporte activo para garantizar actualizaciones y reducir la deuda técnica del proyecto.

- Debe soportar la ejecución paralela de pruebas para maximizar la eficiencia.

- Debe permitir la generación de reportes sencillos que faciliten la revisión de resultados.


**Opciones de Test Runner:**

1. **`go test` (https://pkg.go.dev/cmd/go#hdr-Test_packages):**
 Herramienta oficial de Go para ejecutar pruebas unitarias e integración, con soporte para pruebas paralelas y configuración mínima.

2. **Ginkgo (https://onsi.github.io/ginkgo/):** Un framework de pruebas maduro para Go que facilita la escritura de especificaciones expresivas. Ginkgo se basa en la infraestructura de pruebas nativa de Go y se complementa con la biblioteca Gomega para aserciones. Juntos, permiten expresar de manera clara y precisa la intención detrás de las pruebas.


**Elección:**  
`go test` será el test runner principal, con la posibilidad de complementar con Testify para enriquecer las pruebas cuando sea necesario.
