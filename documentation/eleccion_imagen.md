# Decisión de la imagen elegida para el contenedor

Antes de establecer ningún criterio es necesario tener en cuenta que elementos son necesarios para el proyecto. El primero es que se está trabajando con Go en linux, ningún biblioteca que no es está en la standard library es usada por lo que no es necesario instalación adicional en ese aspecto. Luego una instalación necesaria es la del gestor de tareas que en este caso es TASK(https://github.com/go-task/task).

## Criterios de elección

En este caso vamos a tomar de referencia las recomendaciones en las buenas prácticas de docker (https://docs.docker.com/build/building/best-practices/) como criterios para la selección de la imagen para el contenedor. En base a esto se destacan dos frases que se establecen en esta documentación "When building your own image from a Dockerfile, ensure you choose a minimal base image that matches your requirements". Bajo esta idea se favorecerán aquellas imágenes que tengan lo mínimo para funcionar y por ende tengan un tamaño menor. El tamaño en cuenta será tomado una vez sean instalados los elementos necesarios establecidos antes. Otro detalle a tener en cuenta establecido en las buenas prácticas de docker es que recomienda instalar imágenes de Docker Official Images (https://hub.docker.com/search?badges=official), de Verified Publisher (https://hub.docker.com/search?badges=verified_publisher), o de Docker-Sponsored Open Source (https://hub.docker.com/search?badges=open_source), por lo cual solo serán consideradas aquellas imagenes que al buscar en docker hub estén en algunas de estas categorías.





