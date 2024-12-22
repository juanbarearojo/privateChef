# Decisión de la imagen elegida para el contenedor

Antes de establecer ningún criterio es necesario tener en cuenta que elementos son necesarios para el proyecto. El primero es que se está trabajando con Go en linux, ningún biblioteca que no es está en la standard library es usada por lo que no es necesario instalación adicional en ese aspecto además de la propia de Go. Luego una instalación necesaria es la del gestor de tareas que en este caso es task (https://taskfile.dev/installation/).

## Criterios de elección

En este caso vamos a tomar de referencia las recomendaciones en las buenas prácticas de docker (https://docs.docker.com/build/building/best-practices/) como criterios para la selección de la imagen para el contenedor. En base a esto se destacan se dice "When building your own image from a Dockerfile, ensure you choose a minimal base image that matches your requirements". Bajo esta idea se favorecerán aquellas imágenes que tengan lo mínimo para funcionar y por ende tengan un tamaño menor. El tamaño en cuenta será tomado una vez sean instalados los elementos necesarios establecidos antes. Otro detalle a tener en cuenta establecido en las buenas prácticas de docker es que recomienda instalar imágenes de Docker Official Images (https://hub.docker.com/search?badges=official), de Verified Publisher (https://hub.docker.com/search?badges=verified_publisher), o de Docker-Sponsored Open Source (https://hub.docker.com/search?badges=open_source), por lo cual solo serán consideradas aquellas imagenes que al buscar en docker hub estén en algunas de estas categorías.

## Posibles opciones

- **golang:Alpine (https://hub.docker.com/_/golang)**: No es necesario de instalación del lenguaje de go. Es una Docker Official Image. Esta imagen incluye alpine Linux, Go y ca-certificates. Prueba de funcionamiento en [Prueba golang:alpine](./imagenes_prueba/golang_alpine.png). El commit donde se ha probado el DockerFile ha sido [golang:alpine commit](https://github.com/juanbarearojo/privateChef/pull/34/commits/e567fd234c4a6e687ad669c90005435e45f5f9ae)


- **Okteto: (https://hub.docker.com/r/okteto/golang)**: No es necesario de instalación del lenguaje de go. Es una Verified Publisher. Esta imagen incluye las herramientas Gin, Delve, Gopls, Controller-gen y Air. Esto se puede comprobar viendo el repositorio donde esta el DockerFile(https://github.com/okteto/devenv/blob/main/golang/Dockerfile) Ninguna de estas herramientas son usadas en el proyecto. Prueba de funcionamiento en [Prueba okteto golang](./imagenes_prueba/okteto_golang.png). El commit donde se ha probado el DockerFile ha sido [okteto commit](https://github.com/juanbarearojo/privateChef/pull/34/commits/eb4c06d2aac1e147d5bd7f489529188ddc308c82)


- **Bitnami: (https://hub.docker.com/r/bitnami/golang)**: No es necesario de instalación del lenguaje de go. Es una Verified Publisher. Esta imagen incluye las herramientas build-essential, ca-certificates, curl, git, pkg-config, procps y unzip. Esto se puede comprobar viendo el repositorio donde esta el DockerFile(https://github.com/bitnami/containers/blob/main/bitnami/golang/1.23/debian-12/Dockerfile). De estas herramientas solo es usada git. Prueba de funcionamiento en [Prueba bitnami golang](./imagenes_prueba/bitnami.png). El commit donde se ha probado el DockerFile ha sido [bintami commit](https://github.com/juanbarearojo/privateChef/pull/34/commits/0ebf60d4f1af61c8032ceda4923438b0128efa0e)


- **Ubuntu (https://hub.docker.com/_/ubuntu)**: Es una imagen del sistema opertativo vacía por lo que es necesario instalar Go y se tiene que realizar apt update para poder hacerlo. Además, se tuvo que cambiar a useradd porque adduser no está disponible de forma nativa en la imagen. También es necesario instalar ca-certificates y actualizarlos. rueba de funcionamiento en [Prueba ubuntu:latest](./imagenes_prueba/ubuntu.png). El commit donde se ha probado el DockerFile ha sido [ubuntu commit](https://github.com/juanbarearojo/privateChef/pull/34/commits/7c7494dc256840675366f10697e6d01ef64b6c7d)


- **Debian (https://hub.docker.com/_/debian)**: Es una imagen del sistema opertativo vacía por lo que es necesario instalar Go. Necesidad de instalr wget y tar  ayq eu la versión de golang de la imagen es la 1.22 y por ende daba fallos ya que la versión del go.mod es la 1.23. Prueba de funcionamiento en [Prueba debian:latest](./imagenes_prueba/debian.png). El commit donde se ha probado el DockerFile ha sido [debian commit](https://github.com/juanbarearojo/privateChef/pull/34/commits/c2751aa1a2949fb8275b285d61ae805f37a3a597)
