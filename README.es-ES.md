

# Modificaciones
### Se agregó el tipo "Original" a las tareas HTTP, el cual admite los siguientes dos formatos:

1. Copiar directamente desde el navegador
```
:authority: cm.bilibili.com
:method: POST
:path: /cm/api/fees/pc
:scheme: https
accept: application/json
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6
cache-control: no-cache
content-length: 318
content-type: application/json
cookie: i-wanna-go-back=-1; 
origin: https://www.bilibili.com
pragma: no-cache
referer: https://www.bilibili.com/
sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Microsoft Edge";v="100"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Windows"
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 Edg/100.0.1185.50
```
2. Contenido original, ten en cuenta que este método requiere completar manualmente la URL, como se muestra a continuación
```
GET https://www.baidu.com/?tn=62095104_43_oem_dg HTTP/1.1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9
Cache-Control: max-age=0
Connection: keep-alive
Cookie: BD_HOME=1;
Host: www.baidu.com
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36
sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Windows"
```
### Al usar el tipo "Original", se pueden utilizar las siguientes variables
```
{now}      将被替换成当前时间戳（秒）
{now_mill} 将被替换成当前时间戳（毫秒）
{now_nano} 将被替换成当前时间戳（纳秒）
```

# A continuación, el contenido original del README

# gocron - Sistema de gestión de tareas programadas
[![Downloads](https://img.shields.io/github/downloads/ouqiang/gocron/total.svg)](https://github.com/ouqiang/gocron/releases)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/ouqiang/gocron/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/ouqiang/gocron.svg?label=Release)](https://github.com/ouqiang/gocron/releases)

# Descripción del proyecto
Sistema ligero de programación y gestión centralizada de tareas, desarrollado en Go, diseñado para reemplazar el crontab de Linux. [Ver documentación](https://github.com/ouqiang/gocron/wiki)

Las tareas de retraso originales se han separado en un proyecto independiente: [Cola de retraso](https://github.com/ouqiang/delay-queue)

## Características
* Interfaz web para la gestión de tareas programadas
* Expresiones de tiempo tipo crontab, precisas hasta el segundo
* Reintento en caso de fallo en la ejecución de la tarea
* Finalización forzada en caso de tiempo de ejecución agotado
* Configuración de dependencias de tareas (la tarea B se ejecuta tras completar la tarea A)
* Control de permisos de cuenta
* Tipos de tarea
    * Tarea Shell
    > Ejecuta comandos shell en los nodos de tarea; admite la ejecución simultánea en múltiples nodos.
    * Tarea HTTP
    > Accede a una URL especificada y la ejecuta directamente el programador, sin depender de los nodos de tarea.
* Visualización del registro de resultados de ejecución de tareas
* Notificaciones de resultados de tareas; admite correo electrónico, Slack y Webhook

### Capturas de pantalla
![流程图](https://raw.githubusercontent.com/ouqiang/gocron/master/assets/screenshot/scheduler.png)
![任务](https://raw.githubusercontent.com/ouqiang/gocron/master/assets/screenshot/task.png)
![Slack](https://raw.githubusercontent.com/ouqiang/gocron/master/assets/screenshot/notification.png)
    
### Plataformas compatibles
> Windows, Linux, Mac OS

### Requisitos del entorno
> MySQL


## Descarga
[releases](https://github.com/ouqiang/gocron/releases)  

[Actualizaciones de versión](https://github.com/ouqiang/gocron/wiki/版本升级)

## Instalación

### Instalación binaria
1. Descomprimir el archivo.  
2. `cd directorio_descomprimido`   
3. Iniciar        
* Iniciar el programador        
  * Windows: `gocron.exe web`   
  * Linux, Mac OS:  `./gocron web`
* Iniciar el nodo de tarea (escucha por defecto en 0.0.0.0:5921)
  * Windows:  `gocron-node.exe`
  * Linux, Mac OS:  `./gocron-node`
4. Acceder desde el navegador a http://localhost:5920

### Instalación desde el código fuente

- Instalar Go 1.11+
- `go get -d github.com/ouqiang/gocron`
- `export GO111MODULE=on` 
- Compilar `make`
- Iniciar
    * gocron `./bin/gocron web`
    * gocron-node `./bin/gocron-node`


### Docker

```shell
docker run --name gocron --link mysql:db -p 5920:5920 -d ouqg/gocron
```

Configuración: /app/conf/app.ini

Registros: /app/log/cron.log

La imagen no incluye gocron-node; este debe compilarse junto con la aplicación específica.


### Desarrollo

1. Instalar Go 1.9+, Node.js y Yarn
2. Instalar dependencias frontend `make install-vue`
3. Iniciar gocron y gocron-node `make run`
4. Iniciar el servidor Node `make run-vue` y acceder a http://localhost:8080

Al acceder a http://localhost:8080, las solicitudes de API se reenviarán a gocron.

`make` para compilar

`make run` compila y ejecuta

`make package` para empaquetar 
> Genera un archivo comprimido para el sistema actual: gocron-v1.5-darwin-amd64.tar.gz gocron-node-v1.5-darwin-amd64.tar.gz

`make package-all` genera archivos comprimidos para Windows, Linux y Mac

### Comandos

* gocron
    * -v: mostrar la versión

* gocron web
    * --host: por defecto 0.0.0.0
    * -p: puerto, especifica el puerto (por defecto 5920)
    * -e: entorno de ejecución, dev|test|prod; el modo dev muestra más información de registros (por defecto prod)
    * -h: mostrar ayuda
* gocron-node
    * -allow-root: permite ejecutar como root en plataformas *nix
    * -s ip:port: dirección de escucha  
    * -enable-tls: habilitar TLS    
    * -ca-file: archivo de certificado CA  
    * -cert-file: archivo de certificado  
    * -key-file: archivo de clave privada
    * -h: mostrar ayuda
    * -v: mostrar la versión

## Lista de tareas pendientes
- [x] Actualizaciones de versión
- [x] Activar, desactivar y eliminar tareas en masa
- [x] Comunicación HTTPS entre programador y nodos de tarea
- [x] Agrupación de tareas
- [x] Soporte multiusuario
- [x] Control de permisos

## Componentes utilizados
* Framework web [Macaron](http://go-macaron.com/)
* Programador de tareas [Cron](https://github.com/robfig/cron)
* ORM [Xorm](https://github.com/go-xorm/xorm)
* Framework UI [Element UI](https://github.com/ElemeFE/element)
* Gestión de dependencias [Govendor](https://github.com/kardianos/govendor)
* Framework RPC [gRPC](https://github.com/grpc/grpc)

## Comentarios y soporte
Enviar un [issue](https://github.com/ouqiang/gocron/issues/new)

## Registro de cambios

v1.5
--------
* Reescritura del frontend con Vue + ElementUI
* Notificaciones de tareas
    * Se agregaron notificaciones WebHook
    * Plantillas de notificación personalizadas
    * Envío de notificaciones al coincidir palabras clave en el resultado de la tarea
* Muestra la próxima fecha de ejecución de la tarea en la lista de tareas

v1.4
--------
* Las tareas HTTP admiten solicitudes POST
* Detención manual en el panel de tareas Shell en ejecución
* Intervalo de reintento en caso de fallo configurable por el usuario
* Corrección de error 403 en llamadas a la API

v1.3
--------
* Soporte para inicio de sesión multiusuario
* Implementación de control de permisos de usuario


v1.2.2
--------
* Se añadió un captcha gráfico en la página de inicio de sesión
* Soporte para actualizaciones desde versiones anteriores
* Activación, desactivación y eliminación de tareas en masa
* Autenticación mutua HTTPS entre programador y nodos de tarea
* Corrección del recuento total de registros en la lista de tareas

v1.1
--------

* Las tareas pueden ejecutarse simultáneamente en múltiples nodos
* Por defecto, se prohíbe ejecutar nodos de tarea como root en plataformas *nix
* Se añadieron marcadores de posición predefinidos en los comandos de tareas secundarias, permitiendo que ejecuten acciones según el resultado de la tarea principal
* Eliminación del módulo de demonio (daemon)
* Los registros de acceso web se muestran en la terminal
