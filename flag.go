package main

import (
  "flag"
)

var (
	showCPUInfo  *bool
	showCPUUsage *bool
	showRAM      *bool
	showDisk     *bool
	showLocalIP  *bool
	showPublicIP *bool
	showAll      *bool
	show5TopRAM  *bool
)

func initFlags() {
  //init flags
  showCPUInfo = flag.Bool("cpu-info", false, "Show CPU information")
  showCPUUsage = flag.Bool("cpu-usage", false, "Show CPU usage")
  showRAM = flag.Bool("ram", false, "Show RAM usage")
  showDisk = flag.Bool("disk", false, "Show disk usage")
  showLocalIP = flag.Bool("local-ip", false, "Show local IP address")
  showPublicIP = flag.Bool("public-ip", false, "Show public IP address")
  show5TopRAM = flag.Bool("top5-ram", false, "Show top 5 process that consume the most memory")
  showAll = flag.Bool("all", false, "Show all stats")
  flag.Parse()

  if flag.NFlag() == 0 {
    *showAll = true
  }
}
