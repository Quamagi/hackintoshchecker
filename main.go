package main

import (
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

func main() {
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
		fmt.Printf("CPU: %s, Núcleos: %d, Frecuencia: %.2f GHz\n", c.ModelName, c.Cores, c.Mhz/1000)
		// Veredicto CPU
		if strings.Contains(c.ModelName, "Intel") {
			veredictoCPU = "Tu CPU es Intel, compatible con Hackintosh."
		} else if strings.Contains(c.ModelName, "AMD") {
			veredictoCPU = "Tu CPU es AMD, es posible instalar Hackintosh, pero requerirá parches adicionales."
		} else {
			veredictoCPU = "No se pudo identificar el tipo de CPU. Verifica si tu procesador es Intel o AMD."
		}
	}

	// Información de la memoria
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memoria total: %v MB\n", memInfo.Total/1024/1024)
	// Veredicto RAM
	if memInfo.Total >= 4*1024*1024*1024 {
		veredictoRAM = "Tienes suficiente memoria RAM para instalar macOS (mínimo 4GB)."
		if memInfo.Total >= 8*1024*1024*1024 {
			veredictoRAM += " Se recomienda tener al menos 8 GB de RAM para un mejor rendimiento."
		}
	} else {
		veredictoRAM = "No tienes suficiente memoria RAM para instalar macOS (mínimo 4GB)."
	}

	// Información de los discos
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Disco total: %v GB, Usado: %.2f%%\n", diskInfo.Total/1024/1024/1024, diskInfo.UsedPercent)
	// Veredicto Disco
	if diskInfo.Free > 50*1024*1024*1024 {
		veredictoDisk = "Tienes suficiente espacio libre en el disco (mínimo 50 GB)."
	} else {
		veredictoDisk = "No tienes suficiente espacio libre en el disco. Se recomienda al menos 50 GB libres."
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
			veredictoUEFI = "El sistema está en modo UEFI, listo para Hackintosh."
		} else {
			veredictoUEFI = "El sistema está en modo Legacy (BIOS). Necesitarás cambiar a UEFI para instalar Hackintosh."
		}
	}

	// Verificar Secure Boot (solo en Windows)
	if runtime.GOOS == "windows" {
		secureBoot, err = checkSecureBootWindows()
		if err != nil {
			veredictoSecureBoot = fmt.Sprintf("No se pudo determinar el estado de Secure Boot: %v", err)
		} else {
			if secureBoot {
				veredictoSecureBoot = "Secure Boot está habilitado. Deberás deshabilitarlo en la BIOS para instalar Hackintosh."
			} else {
				veredictoSecureBoot = "Secure Boot está deshabilitado, listo para Hackintosh."
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
