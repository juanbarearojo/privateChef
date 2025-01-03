# Sistemas de Integración Continua

## Criterios de Elección

- La herramienta seleccionada debe permitir al desarrollador ejecutar todas las pruebas necesarias sin coste alguno hasta el 15 de febrero de 2025. Bajo este criterio, únicamente se considerarán herramientas cuyo periodo de prueba gratuita abarque esa fecha o que sean completamente gratuitas.
- No debe ser necesaria la configuración de una infraestructura virtual adicional que no haya sido establecida previamente en los objetivos anteriores. En caso de que se requiera infraestructura adicional, como Kubernetes o una máquina virtual, la herramienta será descartada, ya que no se ajusta a los requerimientos de este objetivo.
- Es imprescindible que la herramienta genere un archivo que permita documentar el trabajo realizado y facilite la creación de flujos más complejos en el futuro.

## Posibles Opciones

1. **Travis (https://www.travis-ci.com/)**: Descartado porque no ofrece un periodo de prueba gratuita actualmente en su página web. Especificamente cuando uno se va registar poner *Trials are currently unavailable. The issue has been identified, and we are preparing a fix. In the meantime, please reach out to support@travis-ci.com if you have any questions.* El plan más económico cuesta 15 euros al mes. Debido a esto se descarta como opción.

2. **CircleCI (https://circleci.com/developer/)**: Ofrece un plan gratuito que incluye hasta 6,000 minutos de construcción, un máximo de 5 usuarios activos por mes, y soporte para Docker, Windows, Linux, ARM, macOS, y runners autohospedados con concurrencia hasta 30x (https://circleci.com/pricing/). Genera un archivo `.yml`.

3. **SemaphoreCI (https://semaphoreci.com/)**: Dispone de un plan gratuito adecuado para nuestros objetivos, con un límite de 5 usuarios, hasta 40 trabajos paralelos, 7,000 minutos de uso en la nube por mes, y acceso a 5 agentes autohospedados (https://semaphoreci.com/pricing). Genera un archivo `.yml`.

4. **Harness (https://www.harness.io/)**: Cuenta con un plan gratuito muy completo (https://www.harness.io/pricing). Sin embargo, al intentar configurarlo, se verificó que requiere definir infraestructura adicional, como Kubernetes, una máquina virtual, o un agente hospedado. Por este motivo, se descarta como opción. Genera un archivo `.yml`.

5. **Jenkins (https://www.jenkins.io/)**: Es software de código abierto. Durante las pruebas de instalación, presentó problemas de compatibilidad con el JDK, ya que solo admite las versiones 17 y 21. Aunque se intentó con ambas versiones, persistieron errores relacionados con la conexión a GitHub.

6. **Buddy Works (https://buddy.works/)**: Ofrece un plan gratuito que permite usuarios ilimitado así como concurrencia paralela (https://buddy.works/pricing). No requiere configuración de infraestructura adicional, pero no genera un archivo `.yml`, lo que lo descarta como opción.

