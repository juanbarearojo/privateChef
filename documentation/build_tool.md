Como se establece en el juanbarearojo#17. Es necesario decidir que build tool que va a ser utilizada. 

- **Bazel:** Herramienta de construcción eficiente para proyectos complejos y de gran escala, especialmente útil cuando se manejan múltiples lenguajes o dependencias.

- **Goreleaser:** Enfocada en la distribución y lanzamiento de aplicaciones en Go, ideal para la generación automática de binarios y despliegues multiplataforma.

- **Go tooling (Go build):** Herramientas nativas de Go para la compilación que ofrecen una solución simple y efectiva para proyectos sin necesidades avanzadas.



Los criterios de elección para el build tool son los siguientes:

- **Simplicidad de Uso:** La herramienta debe ser fácil de utilizar y permitir al programador buildear el código con una barrera de entrada mínima. Esto con un objetivo muy claro, facilitar y agilizar la carga de trabajo del desarrollador. 

- **Configuración Mínima:** La herramienta no debe requerir instalaciones adicionales, salvo las necesarias para Go, con el fin de mantener el entorno de desarrollo ligero y sin dependencias externas innecesarias.


**Elección: Go tooling.**

Sabiendo que Go tiene una herramienta nativa y considerada estandar en muchos casos se plantea como la opción mas lógica. Cualquier persona que le interese 
el proyecto con tener Go sera suficiente para buildearlo. **Goreleaser** se descartó porque no se va a hacer un despliegue multiplataforma y por ende se iba a 
introducir complejidad innecesaria. **Bazel** fue descartadó por un motivo muy parecido