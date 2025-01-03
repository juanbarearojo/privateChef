# Sistemas de integración continua

## Criterios de elección

- La herramienta elegida deberá permitir al desarollador realizar todas la pruebas necesarias sin coste alguno hasta el día 15 de Febrero de 2025. Bajo este criterio sólo se permitirán aquellas herramientas que el plazo ded prueba dure hasta ese día o que sean gratuitas.
- No sea necesaria para ejecutar la pipeline una infraestructura virtual adicional que no haya sido creada en los objetivos anterirores. En caso de que sea necesario establecer uan estructura adicional como puede ser kubernetes o una máquina virtual será descartado ya que no es lo buscado en este objetivo.
- Necesidad de generar un archivo para que quede reflejado en el reposiotrio el trabajo realizadoo se puedan crear flujos más complejos más adelantes.

## Posibles Opciones

1. **Travis (https://www.travis-ci.com/)**: Es descartado porque no tiene prueba gratuita actualmente. El plan mas barato que tienen es de 15 euros al mes. 

2. **Circle ci(https://circleci.com/developer/)**: Tiene un plan gratuito que permite (Up to 6,000 build minutes, Up to 5 active users/month, Docker, Windows, Linux, Arm, macOS, self hosted runners, 30x concurrency). 

3. **Semaphore ci (https://semaphoreci.com/)** : Tiene una tier gratuita que es suficente para lo que vamos a hacer (Maximum 5 users,Concurrency, 40 parallel jobs, Free 7,000 cloud minutes/month, Access to 5 self-hosted agents)

4. ** 

5. 
