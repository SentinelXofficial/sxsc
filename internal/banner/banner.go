package banner

import (
	"fmt"

	"github.com/SentinelXofficial/sxsc/internal/color"
	"github.com/SentinelXofficial/sxsc/internal/updater"
	"github.com/SentinelXofficial/sxsc/internal/version"
)

func Print() {
	fmt.Println()
	fmt.Print(color.CYN + `   _____            __  _            __  _  __` + color.RST + "\n")
	fmt.Print(color.CYN + `  / ___/___  ____  / /_(_)___  ___  / / | |/ /` + color.RST + "\n")
	fmt.Print(color.CYN + `  \__ \/ _ \/ __ \/ __/ / __ \/ _ \/ /  |   /` + color.RST + "\n")
	fmt.Print(color.CYN + ` ___/ /  __/ / / / /_/ / / / /  __/ /___/   |` + color.RST + "\n")
	fmt.Print(color.CYN + `/____/\___/_/ /_/\__/_/_/ /_/\___/_____/_/|_|` + color.RST + "\n")
	fmt.Println()
	fmt.Println(color.GRY + "  SentinelX SC" + color.RST + color.GRY + color.DIM + " — Web Vulnerability Scanner" + color.RST)
	fmt.Println(color.GRY + color.DIM + "  Author : WildanDev" + color.RST)
	fmt.Println()
	printVersionInfo()
}

func printVersionInfo() {
	latest, err := updater.FetchLatest()
	if err != nil {
		// silently skip if can't reach github
		return
	}
	if latest != version.Current {
		fmt.Printf(
			color.GRY+"  [INF] Current sxsc version: "+color.BOLD+"%s"+color.RST+color.YEL+" (outdated, latest: %s)"+color.RST+"\n\n",
			version.Current, latest,
		)
	} else {
		fmt.Printf(
			color.GRY+"  [INF] Current sxsc version: "+color.BOLD+"%s"+color.RST+color.GRN+" (latest)"+color.RST+"\n\n",
			version.Current,
		)
	}
}
