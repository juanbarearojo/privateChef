# Decisión de la herramienta de gestión de dependencias

**Criterios de elección:**

- **Compatibilidad con el Ecosistema Moderno de Go:** La herramienta de gestión de dependencias debe alinearse con las versiones actuales de Go y el flujo de trabajo recomendado en 2024
- **Minimización de Dependencias Externas:** El proyecto debe poder ejecutarse sin necesidad de acceder constantemente a fuentes externas. La herramienta debe facilitar la administración de dependencias de forma local cuando sea necesario.

**Opciones:**

- **Go Modules (https://go.dev/doc/modules/gomod-ref):** Es el sistema oficial de gestión de dependencias desde Go 1.11 y, en 2024, se ha consolidado como el estándar recomendado para proyectos en Go. go.mod y go.sum permiten una administración eficiente de dependencias, asegurando versiones consistentes y controladas de cada módulo.


**Elección: go.mod**

El archivo go.mod se introdujo en Go 1.11, lanzado en agosto de 2018, como parte del sistema de módulos para gestionar dependencias. Este sistema se convirtió en el estándar oficial a partir de Go 1.13, lanzado en septiembre de 2019, cuando se activó por defecto el modo de módulos. Desde entonces, el uso de go.mod es la forma recomendada para manejar dependencias en proyectos Go.


