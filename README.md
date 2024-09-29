¡Me alegra que te haya gustado cómo quedó! Vamos a actualizar el archivo `README.md` en tu proyecto de GitHub para reflejar todos los cambios recientes y agregar los créditos correctamente.

### Aquí te dejo una versión actualizada del archivo `README.md`:

```markdown
# HackintoshChecker

HackintoshChecker es una herramienta escrita en **Go** que verifica la compatibilidad de tu hardware para la instalación de **Hackintosh**. El programa recopila información clave del sistema como CPU, RAM, disco, UEFI y Secure Boot, y muestra un veredicto con colores (verde para OK, amarillo para advertencias y rojo para problemas).

## Características

- **Verificación del procesador (CPU)**: Detecta si es Intel o AMD y evalúa su compatibilidad con Hackintosh.
- **Verificación de la memoria RAM**: Evalúa si tienes suficiente RAM para instalar macOS.
- **Verificación del disco**: Verifica si hay suficiente espacio en el disco para la instalación.
- **Modo UEFI**: Detecta si el sistema está en modo UEFI o Legacy (BIOS).
- **Secure Boot**: Verifica si Secure Boot está habilitado o deshabilitado (solo en Windows).
- **Salida con colores**: Verde para OK, Amarillo para advertencias y Rojo para problemas críticos.

## Requisitos

- **Go**: Necesitarás tener Go instalado para ejecutar este proyecto. Puedes descargarlo desde [golang.org](https://golang.org/dl/).
- **Windows o Linux**: El programa está diseñado para ambos sistemas operativos.

## Instalación

1. **Clona este repositorio**:
   ```bash
   git clone https://github.com/tuusuario/HackintoshChecker.git
   ```

2. **Instala las dependencias** del proyecto:
   ```bash
   go get github.com/shirou/gopsutil/cpu
   go get github.com/shirou/gopsutil/disk
   go get github.com/shirou/gopsutil/host
   go get github.com/shirou/gopsutil/mem
   go get github.com/shirou/gopsutil/net
   ```

## Compilación

Para **PowerShell**:
```bash
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o HackintoshChecker.exe
```

Para **CMD** (símbolo del sistema de Windows):
```bash
set GOOS=windows && set GOARCH=amd64 && go build -o HackintoshChecker.exe
```

## Uso

1. **Ejecuta el programa**:
   ```bash
   ./HackintoshChecker
   ```

2. El programa te mostrará los resultados con colores para que puedas determinar si tu hardware es compatible con Hackintosh.

## Ejemplo de salida

```plaintext
HackintoshChecker - Versión: 1.0
Creadores: Martin Oviedo & Daedalus

----------------------------------------

### Informacion de la PC ###
----------------------------------------
CPU: Intel Core i7-7700HQ
Núcleos: 4, Frecuencia: 2.80 GHz
----------------------------------------
Tienes suficiente memoria RAM para instalar macOS (mínimo 8GB).
Tienes suficiente espacio libre en el disco (mínimo 50 GB).
El sistema está en modo UEFI, listo para Hackintosh.
Secure Boot está deshabilitado, listo para Hackintosh.

### Veredicto Final ###
Tu CPU es Intel, compatible con Hackintosh.
Tienes suficiente memoria RAM para instalar macOS.
Tienes suficiente espacio en disco.
El sistema está listo para Hackintosh.
```

## Contribuir

¡Las contribuciones son bienvenidas! Si tienes ideas para mejorar el proyecto o encuentras algún problema, abre un **pull request** o un **issue** en el repositorio.

## Créditos

- **Martin Oviedo**: Desarrollador principal.
- **Daedalus**: Asistente extraordinario.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

## Arte ASCII

El programa ahora incluye un genial arte ASCII al inicio:

```plaintext
    __  __           __   _       __             __  
   / / / /___ ______/ /__(_)___  / /_____  _____/ /_ 
  / /_/ / __ `/ ___/ //_/ / __ \/ __/ __ \/ ___/ __ \
 / __  / /_/ / /__/ ,< / / / / / /_/ /_/ (__  ) / / /
/_/ /_/\__,_/\___/_/|_/_/_/ /_/\__/\____/____/_/ /_/ 
   ________              __                          
  / ____/ /_  ___  _____/ /_____  _____              
 / /   / __ \/ _ \/ ___/ //_/ _ \/ ___/              
/ /___/ / / /  __/ /__/ ,< /  __/ /                  
\____/_/ /_/\___/\___/_/|_|\___/_/                   
```

¡Gracias por usar HackintoshChecker! 😊
```

### **Cambios incluidos en el README:**
1. Actualicé la descripción del proyecto para incluir las nuevas características como los colores y el arte ASCII.
2. Añadí las instrucciones de compilación para PowerShell y CMD.
3. Incluí un ejemplo de salida para mostrar cómo funciona el programa.
4. Añadí los créditos a **Martin Oviedo** y **Daedalus** (¡claro que sí! 😄).
5. Agregué el arte ASCII en el archivo README para que también luzca en el repositorio.

¡Con este `README.md`, tu proyecto se verá mucho más profesional y estará listo para recibir visitas en GitHub!
