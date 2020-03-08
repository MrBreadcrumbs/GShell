package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("  ________   _________.__             .__   .__   ")
	fmt.Println(" /  _____/  /   _____/|  |__    ____  |  |  |  |  ")
	fmt.Println("/   \\  ___  \\_____  \\ |  |  \\ _/ __ \\ |  |  |  |  ")
	fmt.Println("\\    \\_\\  \\ /        \\|   Y  \\\\  ___/ |  |__|  |__")
	fmt.Println(" \\______  //_______  /|___|  / \\___  >|____/|____/")
	fmt.Println("        \\/         \\/      \\/      \\/  by @TerminalJockey   ")
	fmt.Println("Gr33tz V1s3r1on CptSticky Sazoukis & all the mystiko crew")
	fmt.Println("~~~~~~~~~~~Hack the planet~~~~~~~~~~		")

	buffer := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>")
		input, err := buffer.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		parse(input)
	}
}

func parse(input string) {
	//format input for windows
	input = strings.Replace(input, "\r\n", "\n", -1)
	input = strings.TrimSuffix(input, "\n")

	//implements ls, add arguement functionality
	if strings.Compare(input, "ls") == 0 {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Println(err)
		}
		for _, file := range files {
			fmt.Println("-", file.Name(), file.Size())
		}

	} else if strings.Compare(input, "?") == 0 {
		fmt.Println("GShell is modeled after the unix command line")
		fmt.Println("ls -- lists files in current directory")
		fmt.Println("cd -- changes to target directory")
		fmt.Println("pwd -- prints current working directory")
		fmt.Println("cat -- prints target file contents to stdout")
		fmt.Println("GetInfo -- runs local enumeration script")
		fmt.Println("exit -- exits clean")

		//implements pwd
	} else if strings.Compare(input, "pwd") == 0 {
		currdir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(currdir)

		//exit clean
	} else if strings.Compare(input, "exit") == 0 {
		os.Exit(0)

		//implements cat
	} else if strings.HasPrefix(input, "cat ") == true {
		catArgs := strings.Split(input, " ")
		content, err := ioutil.ReadFile(catArgs[1])
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(content))

		//implements cd, add functionality for spaces in path
	} else if strings.HasPrefix(input, "cd") == true {

		cdArgs := strings.Split(input, " ")
		os.Chdir(cdArgs[1])

		//enumeration script implementation
	} else if strings.Compare(input, "GetInfo") == 0 {

		fmt.Println("Whoami results:")
		whoami := exec.Command("whoami", "/all")
		whoami.Stderr = os.Stderr
		whoami.Stdout = os.Stdout
		whoami.Run()

		fmt.Println("Net user results:")
		netUser := exec.Command("net", "user")
		netUser.Stderr = os.Stderr
		netUser.Stdout = os.Stdout
		netUser.Run()

		fmt.Println("Net share results: \n")
		netView := exec.Command("net", "share")
		netView.Stderr = os.Stderr
		netView.Stdout = os.Stdout
		netView.Run()

		fmt.Println("Systeminfo results:")
		sysinfo := exec.Command("systeminfo")
		sysinfo.Stdout = os.Stdout
		sysinfo.Stderr = os.Stderr
		sysinfo.Run()

		fmt.Println("IPConfig results:")
		ipconfig := exec.Command("ipconfig", "/all")
		ipconfig.Stderr = os.Stderr
		ipconfig.Stdout = os.Stdout
		ipconfig.Run()

		fmt.Println("Netstat results:")
		netstat := exec.Command("netstat", "-ano")
		netstat.Stdout = os.Stdout
		netstat.Stderr = os.Stderr
		netstat.Run()

		fmt.Println("Tasklist results:")
		tasklist := exec.Command("tasklist", "/v")
		tasklist.Stderr = os.Stderr
		tasklist.Stdout = os.Stdout
		tasklist.Run()

		fmt.Println("Searching for passwords in txt and ini files...")
		searcher := exec.Command("findstr", "/D:/Users/", "/si", "/m", "password", "*.ini", "*.txt")
		searcher.Stderr = os.Stderr
		searcher.Stdout = os.Stdout
		searcher.Run()
		fmt.Printf("\n")

		fmt.Println("Checking for Unquoted Service Paths...")
		servquot := exec.Command("powershell", "-Command", "gwmi -class Win32_Service -Property Name, DisplayName, PathName, StartMode | Where {$_.StartMode -eq \"Auto\" -and $_.PathName -notlike \"C:\\Windows*\" -and $_.PathName -notlike '\"*'} | select PathName,DisplayName,Name")
		servquot.Stderr = os.Stderr
		servquot.Stdout = os.Stdout
		servquot.Run()
		fmt.Printf("\n")

		fmt.Println("Getting named pipes...")
		namePipe := exec.Command("powershell", "-Command", "[System.IO.Directory]::GetFiles(\"\\\\.\\pipe\\\\\\\")")
		namePipe.Stderr = os.Stderr
		namePipe.Stdout = os.Stdout
		namePipe.Run()
		fmt.Printf("\n")

		fmt.Println("Checking for stored creds...")
		saveCreds := exec.Command("cmdkey", "/list")
		saveCreds.Stderr = os.Stderr
		saveCreds.Stdout = os.Stdout
		saveCreds.Run()

	} else {

		return

	}

}
