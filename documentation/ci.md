# Sistemas de Integración Continua

## Criterios de Elección

- La herramienta seleccionada debe permitir al desarrollador ejecutar todas las pruebas necesarias sin coste alguno hasta el 15 de febrero de 2025. Bajo este criterio, únicamente se considerarán herramientas cuyo periodo de prueba gratuita abarque esa fecha o que sean completamente gratuitas.
- No debe ser necesaria la configuración de una infraestructura virtual adicional que no haya sido establecida previamente en los objetivos anteriores. En caso de que se requiera infraestructura adicional, como Kubernetes o una máquina virtual, la herramienta será descartada, ya que no se ajusta a los requerimientos de este objetivo.
- Es imprescindible que la herramienta genere un archivo que permita documentar el trabajo realizado y facilite la creación de flujos más complejos en el futuro.

## Posibles Opciones

1. **Travis (https://www.travis-ci.com/)**: Descartado porque no ofrece un periodo de prueba gratuita actualmente en su página web. Específicamente cuando uno se va a registrar pone *Trials are currently unavailable. The issue has been identified, and we are preparing a fix. In the meantime, please reach out to support@travis-ci.com if you have any questions.* El plan más económico cuesta 15 euros al mes. Debido a esto se descarta como opción.

2. **CircleCI (https://circleci.com/developer/)**: Ofrece un plan gratuito que incluye hasta 6,000 minutos de construcción, un máximo de 5 usuarios activos por mes, y soporte para Docker, Windows, Linux, ARM, macOS, y runners autohospedados con concurrencia hasta 30x (https://circleci.com/pricing/). Genera un archivo `.yml`.

3. **SemaphoreCI (https://semaphoreci.com/)**: Dispone de un plan gratuito adecuado para nuestros objetivos, con un límite de 5 usuarios, hasta 40 trabajos paralelos, 7,000 minutos de uso en la nube por mes, y acceso a 5 agentes autohospedados (https://semaphoreci.com/pricing). Genera un archivo `.yml`.

4. **Harness (https://www.harness.io/)**: Cuenta con un plan gratuito muy completo (https://www.harness.io/pricing). Sin embargo, al intentar configurarlo, se verificó que requiere definir infraestructura adicional, como Kubernetes, una máquina virtual, o un agente hospedado. Por este motivo, se descarta como opción. Genera un archivo `.yml`.

5. **Jenkins (https://www.jenkins.io/)**: Es software de código abierto. Durante las pruebas de instalación, presentó problemas de compatibilidad con el JDK, ya que solo admite las versiones 17 y 21. Aunque se intentó con ambas versiones, hubo errores relacionados con la conexión a GitHub.

6. **Buddy Works (https://buddy.works/)**: Ofrece un plan gratuito que permite usuarios ilimitados así como concurrencia paralela (https://buddy.works/pricing). No requiere configuración de infraestructura adicional, pero no genera un archivo `.yml`, lo que lo descarta como opción.

7. **AppVeyor (https://www.appveyor.com/)**: Tiene un plan gratuito algo limitado en comparación al resto ya que incluye Unlimited public projects, 1 concurrent job, 5 self-hosted jobs, pero que es más que suficiente para el proyecto (https://www.appveyor.com/pricing/). Genera un archivo `appveyor.yml` y no es necesario crear ninguna infraestructura virtual adicional.

8. **GitHub Actions (https://github.com/features/actions)**: Es gratuita, genera un archivo.yml, no es necesario infraestructra adicional.

## Elección

- Se elige **CircleCI** para que ejecute los tests a través de la imagen Docker.
- Se elige **AppVeyor** para probar las distintas versiones de Go. Como se establece en este documento (https://go.dev/doc/devel/release#policy), las versiones mantenidas son las dos últimas versiones que son 1.23 y 1.22. Debido a esto, son probadas. También se prueba la versión 1.21 para comprobar si funcionaría una versión que ya no está mantenida.
- Se mantienen el resto de las pipeline del resto de Ci como muestra del trabajo realizado.


