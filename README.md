# Private Chef

## Descripción del Problema

Como estudiante con poco tiempo durante la semana, es esencial para mí optimizar el tiempo que dedico a cocinar. Para gestionar mi tiempo de manera eficiente, suelo practicar el **meal prep**. El **meal prep** consiste en preparar todas las comidas de la semana en uno o dos días, organizando y cocinando grandes cantidades de comida que luego se almacenan para su consumo durante los días laborales. Esto me permite evitar cocinar a diario, lo que ahorra tiempo y facilita seguir una dieta planificada.

Sin embargo, uno de los mayores retos de hacer meal prep es planificar qué cocinar con los ingredientes que ya tengo disponibles en casa. A menudo, al final de la semana me doy cuenta de que no he utilizado ciertos ingredientes y terminan por estropearse o sobran, lo que genera desperdicio de alimentos. Encontrar las recetas adecuadas para aprovechar al máximo los ingredientes disponibles puede ser difícil, lo que a veces me lleva a comprar más de lo necesario o no utilizar correctamente lo que ya tengo en casa.

### Objetivos

- **Facilitar la planificación de comidas semanales**, optimizando el uso de los ingredientes disponibles en casa para evitar el desperdicio de alimentos al final de la semana.

- **Abordar la dificultad de encontrar recetas** que se ajusten a los ingredientes ya disponibles, haciendo más sencillo el proceso de organización del meal prep.

- **Reducir el desperdicio de alimentos**, asegurando que se utilicen la mayor cantidad posible de ingredientes ya comprados, sin necesidad de adquirir más productos innecesarios.

## Lenguaje

- **Lenguaje de Programación Principal**: Go

## Build Tool

- **Compilador**: gc (Go Compiler), ejecutado a través de `go build`

## Gestor de dependencias

-  **Gestor**: Go Modules (go.mod). El proyecto utiliza Go Modules como sistema de gestión de dependencias, permitiendo una administración eficiente y controlada de las versiones de cada módulo, asegurando consistencia en los entornos de desarrollo y producción.

## Tareas Definidas en el Taskfile

- **Instalación de dependencias:** task install
- **Compilación del proyecto:** task build
- **Ejecución de pruebas:** task test
- **Verificación de sintaxis:** task check
- **Tarea por defecto:** task default

## orden check

Ejecutar task check

## Documentación adicional:

- [Licencia](./LICENSE)
- [Configuración del proyecto](./documentation/configuracion_repositorio.png)
- [Historias de usuario](./documentation/user_stories.md)
- [Milestones](./documentation/milestones.md)

