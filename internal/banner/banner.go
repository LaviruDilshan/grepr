package banner

import (
	"fmt"

	"github.com/fatih/color"
)

const art = `
 ██████╗ ██████╗ ███████╗██████╗ ██████╗ 
██╔════╝ ██╔══██╗██╔════╝██╔══██╗██╔══██╗
██║  ███╗██████╔╝█████╗  ██████╔╝██████╔╝
██║   ██║██╔══██╗██╔══╝  ██╔═══╝ ██╔══██╗
╚██████╔╝██║  ██║███████╗██║     ██║  ██║
 ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝
`

const version = "v2.0.0"

func Print() {
	red := color.New(color.FgRed, color.Bold)
	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow)
	white := color.New(color.FgWhite)
	green := color.New(color.FgGreen)

	red.Println(art)

	cyan.Printf("  %s  ", "grepr")
	yellow.Printf("%-10s", version)
	white.Println("– Lightweight URL filter tool for bug bounty hunters and security researchers")

	fmt.Println()
	green.Println("  ╔══════════════════════════════════════════════════════════╗")
	green.Println("  ║              Developer Information                       ║")
	green.Println("  ╠══════════════════════════════════════════════════════════╣")

	green.Print("  ║  ")
	white.Printf("%-54s", "Author   : Laviru Dilshan From Ovate Security")
	green.Println("  ║")

	green.Print("  ║  ")
	white.Printf("%-54s", "Website  : https://lavirudilshan.com")
	green.Println("  ║")

	green.Print("  ║  ")
	white.Printf("%-54s", "GitHub   : https://github.com/LaviruDilshan")
	green.Println("  ║")

	green.Print("  ║  ")
	white.Printf("%-54s", "Twitter  : @LaviruDilshan")
	green.Println("  ║")

	green.Print("  ║  ")
	white.Printf("%-54s", "Purpose  : Authorized security testing only")
	green.Println("  ║")

	green.Println("  ╚══════════════════════════════════════════════════════════╝")

	fmt.Println()
	color.New(color.FgRed, color.Bold).Println("  [!] FOR AUTHORIZED SECURITY TESTING ONLY – MISUSE IS ILLEGAL")
	fmt.Println()
}
