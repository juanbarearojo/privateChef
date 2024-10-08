# Private Chef - Milestones

## [M0] Milestone 0: **Definición del dominio del problema**

 **Objetivo**: Realizar un análisis detallado de las Historias de Usuario (HU) para identificar y definir los elementos clave del dominio del problema, estableciendo la metodología para el modelado.

+ **Entregable**: 
    - Código que refleje los elementos del dominio, incluyendo clases para Ingrediente, Receta, Inventario, Semana, aquellas más consideradas por el programador, y las características específicas de cada uno de ellos.
    - El código debe reflejar las relaciones entre estos elementos, especialmente incorporando el inventario y cómo interactúa con los ingredientes.

+ **Viabilidad**: 
    - El modelo es válido si incluye todas las entidades clave del dominio y sus relaciones correctamente implementadas en el código.
    - Se considera viable si el código refleja la naturaleza de los ingredientes, en términos de cómo se degradan con el tiempo y las cantidades de estos, de manera eficiente y especialmente sencilla.
    - Se considera viable cuando otro desarrollador puede entender el modelo y continuar el desarrollo siguiendo las directrices establecidas, sin ambigüedades.
    - Se considera viable cuando el diseño permite la implementación eficiente de algoritmos que toman decisiones óptimas en cada paso basándose en la información disponible, facilitando la selección adecuada de recetas y el uso eficiente de ingredientes.
    - Se considera viable cuando las recetas tienen una forma unificada y usan la misma medida usada en el ingrediente.
    - La representación de la semana no incluye información no necesaria para la solución del problema.
    - La estructura de la receta permitirá representar el estilo de cocina al que pertenece. 

### Historias de Usuario Asociadas: [HU1], [HU2], [HU3]

---

## [M1] Milestone 1: **Implementación de Lista de Recetas para Minimizar el Desperdicio de Ingredientes**

+ **Objetivo**: Implementar una solución que genere una lista óptima de recetas que maximice el uso de los ingredientes disponibles, especialmente los perecederos, para minimizar el desperdicio. El resultado debe ser una lista de recetas, no de ingredientes.

+ **Entregable**:
    - Código con las clases y módulos que calculen la lista de recetas óptimas basadas en los ingredientes disponibles, priorizando el uso de ingredientes perecederos. 
    - Implementación de algoritmos que optimicen la selección de recetas para minimizar el desperdicio de ingredientes.
    - Tests unitarios que verifiquen el correcto funcionamiento de los módulos y la optimización del uso de ingredientes.

+ **Viabilidad**: 
    - Los tests deben demostrar que, dada una lista de ingredientes disponibles, el programa genera una lista de recetas que maximiza el uso de ingredientes perecederos y minimiza el desperdicio. No será necesario maximizar el uso de ingredientes no perecederos. El uso ingredientes de uso común no deberá ser tomado en cuenta a la hora de toma decisiones para la elección de recetas. 
    - Se considera viable cuando los resultados cumplen con los criterios de optimización establecidos, los tests pasan correctamente y el código se ajusta al modelo establecido.
    - Para que sea condiderado viable en este milestone no se necesario considerar la preferencias del usuario respecto al estilo de cocina.

### Historias de Usuario Asociadas: [HU1], [HU2], [HU3]


---

## [M2] Milestone 2: **Sugerencia de Recetas Basadas en Preferencias del Usuario**

+ **Objetivo**: Similar al Milestone 1, pero incluyendo las preferencias del usuario respecto al tipo de cocina (por ejemplo, italiana, japonesa, vegetariana) en la sugerencia de recetas.

+ **Entregable**: 
    - Código con las clases y módulos que realicen la búsqueda de recetas en base a ingredientes disponibles y preferencias del usuario.
    - Tests unitarios que verifiquen el correcto funcionamiento de los módulos y la optimización del uso de ingredientes basada en las preferencias del usuario a la hora de cocinar.


+ **Viabilidad**: 
    - Los tests deben verificar que las recetas sugeridas coinciden con las preferencias del usuario en cuanto al tipo de cocina y utilizan los ingredientes disponibles.
    - Siga respetando las viabilidad del objetivo 1 manteniendo la busqueda del desperdicio mínimo de ingredientes usados

### Historias de Usuario Asociadas: [HU3]


---

