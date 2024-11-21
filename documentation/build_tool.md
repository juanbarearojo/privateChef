# Decisión de la herramienta de construcción (Build Tool)

Es necesario decidir qué herramienta de construcción se va a utilizar para el proyecto. Los criterios de elección para el build tool son los siguientes:

- **Simplicidad de Uso:** La herramienta debe ser fácil de utilizar y permitir al programador buildear el código con una barrera de entrada mínima. Esto con un objetivo muy claro, facilitar y agilizar la carga de trabajo del desarrollador. 

- **Configuración Mínima:** La herramienta no debe requerir instalaciones adicionales, salvo las necesarias para Go, con el fin de mantener el entorno de desarrollo ligero y sin dependencias externas innecesarias.


 A continuación, se presentan las opciones consideradas:

- **Bazel (https://bazel.build/?hl=es-419):** Herramienta de construcción eficiente para proyectos complejos y de gran escala, especialmente útil cuando se manejan múltiples lenguajes o dependencias.

- **Go tooling (Go build) (https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies):** Herramientas nativas de Go para la compilación que ofrecen una solución simple y efectiva para proyectos sin necesidades avanzadas.


**Elección: Go tooling.**

Sabiendo que Go tiene una herramienta nativa y considerada estandar en muchos casos se plantea como la opción mas lógica. Cualquier persona que le interese 
el proyecto con tener Go sera suficiente para buildearlo. **Bazel** se descartó porque no se va a hacer un despliegue multiplataforma y por ende se iba a 
introducir complejidad innecesaria. 