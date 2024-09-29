package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Códigos de color ANSI
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func main() {
	// Mostrar el arte ASCII en el inicio
	fmt.Println(`    __  __           __   _       __             __  
   / / / /___ ______/ /__(_)___  / /_____  _____/ /_ 
  / /_/ / __ '/ ___/ //_/ / __ \/ __/ __ \/ ___/ __ \
 / __  / /_/ / /__/ ,< / / / / / /_/ /_/ (__  ) / / /
/_/ /_/\__,_/\___/_/|_/_/_/ /_/\__/\____/____/_/ /_/ 
   ________              __                          
  / ____/ /_  ___  _____/ /_____  _____              
 / /   / __ \/ _ \/ ___/ //_/ _ \/ ___/              
/ /___/ / / /  __/ /__/ ,< /  __/ /                  
\____/_/ /_/\___/\___/_/|_|\___/_/                   
`)

	// Mostrar el nombre del programa, la versión y los creadores con colores
	fmt.Println("Creadores: Martin Oviedo & Daedalus")

	// Separar los créditos con una línea
	fmt.Println("----------------------------------------\n")
	fmt.Println("\n### Informacion de la PC ###")
	// Veredicto y consejos
	var veredictoCPU, veredictoRAM, veredictoDisk, veredictoUEFI, veredictoSecureBoot string

	// Definir la variable secureBoot antes de usarla
	var secureBoot bool

	// Información del procesador
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cpuInfo {
		fmt.Println("----------------------------------------")
		fmt.Printf("CPU: %s\nNúcleos: %d, Frecuencia: %.2f GHz\n", c.ModelName, c.Cores, c.Mhz/1000)
		fmt.Println("----------------------------------------")

		// Veredicto CPU
		if strings.Contains(c.ModelName, "Intel") {
			veredictoCPU = fmt.Sprintf("%sTu CPU es Intel, compatible con Hackintosh.%s", Green, Reset)
		} else if strings.Contains(c.ModelName, "AMD") {
			veredictoCPU = fmt.Sprintf("%sTu CPU es AMD, es posible instalar Hackintosh, pero requerirá parches adicionales.%s", Yellow, Reset)
		} else {
			veredictoCPU = fmt.Sprintf("%sNo se pudo identificar el tipo de CPU. Verifica si tu procesador es Intel o AMD.%s", Red, Reset)
		}
	}

	// Información de la memoria
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memoria total: %v MB\n", memInfo.Total/1024/1024)
	// Veredicto RAM
	if memInfo.Total >= 8*1024*1024*1024 {
		veredictoRAM = fmt.Sprintf("%sTienes suficiente memoria RAM para instalar macOS (mínimo 8GB).%s", Green, Reset)
	} else if memInfo.Total >= 4*1024*1024*1024 {
		veredictoRAM = fmt.Sprintf("%sTienes suficiente memoria RAM para instalar macOS, pero se recomienda tener al menos 8GB.%s", Yellow, Reset)
	} else {
		veredictoRAM = fmt.Sprintf("%sNo tienes suficiente memoria RAM para instalar macOS (mínimo 4GB).%s", Red, Reset)
	}

	// Información de los discos
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Disco total: %v GB, Usado: %.2f%%\n", diskInfo.Total/1024/1024/1024, diskInfo.UsedPercent)
	// Veredicto Disco
	if diskInfo.Free > 50*1024*1024*1024 {
		veredictoDisk = fmt.Sprintf("%sTienes suficiente espacio libre en el disco (mínimo 50 GB).%s", Green, Reset)
	} else {
		veredictoDisk = fmt.Sprintf("%sNo tienes suficiente espacio libre en el disco. Se recomienda al menos 50 GB libres.%s", Red, Reset)
	}

	// Información del host (BIOS, UEFI, sistema)
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sistema operativo: %s %s (%s)\n", hostInfo.OS, hostInfo.Platform, hostInfo.PlatformVersion)
	fmt.Printf("Tiempo de arranque: %v segundos\n", hostInfo.Uptime)

	// Información de las interfaces de red
	netInfo, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range netInfo {
		fmt.Printf("Interfaz de red: %s, Dirección MAC: %s\n", iface.Name, iface.HardwareAddr)
	}

	// Verificar modo UEFI
	isUEFI, err := checkUEFIMode()
	if err != nil {
		veredictoUEFI = fmt.Sprintf("No se pudo determinar el modo UEFI: %v", err)
	} else {
		if isUEFI {
			veredictoUEFI = fmt.Sprintf("%sEl sistema está en modo UEFI, listo para Hackintosh.%s", Green, Reset)
		} else {
			veredictoUEFI = fmt.Sprintf("%sEl sistema está en modo Legacy (BIOS). Necesitarás cambiar a UEFI para instalar Hackintosh.%s", Yellow, Reset)
		}
	}

	// Verificar Secure Boot (solo en Windows)
	if runtime.GOOS == "windows" {
		secureBoot, err = checkSecureBootWindows()
		if err != nil {
			veredictoSecureBoot = fmt.Sprintf("No se pudo determinar el estado de Secure Boot: %v", err)
		} else {
			if secureBoot {
				veredictoSecureBoot = fmt.Sprintf("%sSecure Boot está habilitado. Deberás deshabilitarlo en la BIOS para instalar Hackintosh.%s", Yellow, Reset)
			} else {
				veredictoSecureBoot = fmt.Sprintf("%sSecure Boot está deshabilitado, listo para Hackintosh.%s", Green, Reset)
			}
		}
	} else {
		veredictoSecureBoot = "La verificación de Secure Boot no está implementada para este sistema operativo."
	}

	// Mostrar el veredicto final
	fmt.Println("\n### Veredicto Final ###")
	fmt.Println(veredictoCPU)
	fmt.Println(veredictoRAM)
	fmt.Println(veredictoDisk)
	fmt.Println(veredictoUEFI)
	fmt.Println(veredictoSecureBoot)

	fmt.Println("\n### Consejos para Hackintosh ###")
	if isUEFI && (runtime.GOOS == "windows" && !secureBoot) {
		fmt.Println("Tu sistema está casi listo para Hackintosh. Revisa los componentes específicos como la tarjeta gráfica y otros periféricos para asegurar compatibilidad.")
	} else if !isUEFI {
		fmt.Println("Debes cambiar el sistema al modo UEFI desde la configuración de BIOS antes de proceder con la instalación de Hackintosh.")
	}
	fmt.Println("\nGracias por usar HackintoshChecker")
	// Usar bufio.NewReader para esperar a que el usuario presione una tecla
	fmt.Println("\nPresiona 'Enter' para salir...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n') // Espera hasta que el usuario presione Enter
}

// checkUEFIMode verifica si el sistema está en modo UEFI.
func checkUEFIMode() (bool, error) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("mountvol")
		output, err := cmd.Output()
		if err != nil {
			return false, err
		}
		return bytes.Contains(output, []byte("EFI")), nil
	case "linux":
		if _, err := os.Stat("/sys/firmware/efi"); err == nil {
			return true, nil
		}
		return false, nil
	default:
		return false, fmt.Errorf("verificación de modo UEFI no soportada para %s", runtime.GOOS)
	}
}

// checkSecureBootWindows verifica si Secure Boot está habilitado en Windows.
func checkSecureBootWindows() (bool, error) {
	cmd := exec.Command("powershell", "-Command", "Confirm-SecureBootUEFI")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("error ejecutando PowerShell: %v, %s", err, stderr.String())
	}

	result := strings.TrimSpace(out.String())
	if result == "True" {
		return true, nil
	} else if result == "False" {
		return false, nil
	} else {
		return false, fmt.Errorf("resultado inesperado: %s", result)
	}
}

/*Para compilar con Powershell use

Instrucción para PowerShell:
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o HackintoshChecker.exe

Instrucción para CMD:
set GOOS=windows && set GOARCH=amd64 && go build -o HackintoshChecker.exe

recuerde instalar MinDW
*/
