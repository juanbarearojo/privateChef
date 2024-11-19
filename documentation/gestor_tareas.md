# Decisión de la herramienta de gestión de tareas

Criterios de elección:

- **Comunidad Activa para Reducir Deuda Técnica a Largo Plazo:** Es fundamental que la herramienta cuente con una comunidad activa y respaldo continuo, ya que esto garantiza actualizaciones regulares y una mayor longevidad en el ecosistema de desarrollo. Esta base de soporte y la evolución continua ayudan a evitar la acumulación de deuda técnica, manteniendo la herramienta actualizada y reduciendo el riesgo de tener que migrar a una alternativa en el futuro debido a la obsolescencia.

- **Facilidad de Incorporación para Nuevos Desarrolladores:** La herramienta debe ser intuitiva y accesible, permitiendo que los nuevos miembros del equipo comprendan y utilicen rápidamente sus funcionalidades. Esta accesibilidad es clave en equipos en crecimiento, ya que facilita que los desarrolladores recién incorporados puedan contribuir al proyecto sin una curva de aprendizaje extensa, mejorando la productividad y promoviendo una colaboración ágil.

- **Madurez en el Mercado:** Es importante que la herramienta tenga suficiente tiempo en el mercado para que los desarrolladores hayan tenido oportunidad de identificar y corregir errores, consolidando así una base estable. Las herramientas nuevas o menos probadas pueden no ofrecer la misma confiabilidad.

- **Actualización Creciente del Proyecto:** La herramienta debe mostrar evidencia de actualizaciones continuas y soporte activo en su repositorio o comunidad. Esto asegura que el proyecto no quedará obsoleto a corto plazo.

Posibles opciones:

- **Just (https://github.com/casey/just):** Recomendado como gestor de tareas por su simplicidad y enfoque directo, Just permite definir y ejecutar comandos comunes en un archivo justfile. Esta herramienta proporciona una sintaxis moderna y amigable que facilita la ejecución de tareas sin la complejidad de herramientas tradicionales de construcción.

- **Make (https://www.gnu.org/software/make/):** Una herramienta de construcción clásica y robusta, conocida por su soporte en múltiples entornos y su capacidad para manejar dependencias de construcción. Aunque eficaz en proyectos grandes, Make puede resultar excesivamente complejo para la gestión de tareas simples, y su sintaxis es menos intuitiva para configuraciones básicas.

- **Task (https://taskfile.dev/):** Task utiliza un archivo de configuración en formato YAML, lo que hace que su sintaxis sea más legible y moderna en comparación con Make. Es una opción interesante para quienes buscan una gestión de tareas estructurada y clara. 

- **Mage (https://github.com/magefile/mage):** Ideal para quienes prefieren escribir tareas en código Go, Mage permite definir scripts de automatización sin un archivo de configuración específico, empleando funciones de Go directamente. 

- **Goyek (https://github.com/goyek/goyek):** Una biblioteca escrita en Go que permite definir tareas como código Go en lugar de archivos de configuración. Está diseñada para integrarse directamente en proyectos Go, ofreciendo ventajas como depuración fácil y portabilidad. 


**Elección: Task**

Mage ha sido descartado porque su repositorio no ha registrado actividad en más de un año. Aunque esto no implica necesariamente que el proyecto haya sido abandonado, se ha decidido no seleccionarlo como medida cautelar para reducir la deuda técnica. Si el proyecto llegara a quedar obsoleto (deprecated), podría representar un problema futuro. Make, por su parte, también ha sido descartado debido a que introduce una dificultad artificial al proyecto. Goyek ha sido descartado porque, aunque ofrece características interesantes para proyectos en Go, es una herramienta relativamente nueva en el mercado. Esto significa que no ha tenido suficiente tiempo para ser ampliamente testeada por la comunidad y corregir posibles errores significativos.

En cambio, tanto Just como Task se consideran opciones válidas. Ambos cuentan con una sintaxis sencilla, han sido actualizados recientemente, y sus repositorios muestran actividad reciente y estadísticas similares. Sin embargo, se ha optado por Task debido a que su licencia, la MIT License, fomenta más el desarrollo de código abierto y se alinea mejor con la visión del proyecto, en comparación con la CC0-1.0 License de Just pero son intercambiables. 
Ha sido una elección basada en preferencias personales.