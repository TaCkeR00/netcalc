package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/TaCkeR00/netcalc/ipv4"
)

func main() {
	ipStr := flag.String("ip", "", "set ip")
	maskStr := flag.String("mask", "", "set mask")
	cidrNotation := flag.String("cidr", "", "set cidr")
	hosts := flag.Uint("host-num", 0, "set hosts number")
	subnets := flag.Uint("subnet-num", 0, "set subnets number")
	flag.Parse()

	// 1. Process the input into an IP and Mask
	ip, mask := processInput(*ipStr, *maskStr, *cidrNotation, *hosts, *subnets)

	if !ipv4.CheckAddress(ip, mask) {
		log.Fatal("invalid address")
	}

	// 2. Print all the calculated information
	printNetworkDetails(ip, mask)
}

func processInput(ipStr, maskStr, cidrStr string, hosts, subnets uint) (ipv4.Address, ipv4.Address) {
	var a, b, c, d uint8
	var ip, mask ipv4.Address

	if (hosts != 0 || subnets != 0) && ipStr != "" && maskStr == "" && cidrStr == "" {
		fmt.Sscanf(ipStr, "%d.%d.%d.%d", &a, &b, &c, &d)
		ip = ipv4.Init(a, b, c, d)

		if cls := ipv4.GetClass(ip); cls == ipv4.INVALID {
			log.Fatal("invalid ip")
		}

		var subnetsBitNum, hostsBitNum uint8
		hostsBitNum = uint8(math.Ceil(math.Log2(float64(hosts))))

		if subnets == 0 {
			subnetsBitNum = ipv4.GetSubnetNumFromIpHost(ip, hostsBitNum)
		} else {
			subnetsBitNum = uint8(math.Ceil(math.Log2(float64(subnets))))
		}

		if !ipv4.CheckHostsAndSubnets(ip, hostsBitNum, subnetsBitNum) {
			log.Fatal("invalid host and subnet num")
		}

		mask = ipv4.GetMaskFromIpSubnet(ip, subnetsBitNum)

	} else if ipStr != "" && maskStr != "" && cidrStr == "" && hosts == 0 && subnets == 0 {
		fmt.Sscanf(ipStr, "%d.%d.%d.%d", &a, &b, &c, &d)
		ip = ipv4.Init(a, b, c, d)

		if cls := ipv4.GetClass(ip); cls == ipv4.INVALID {
			log.Fatal("invalid ip")
		}

		fmt.Sscanf(maskStr, "%d.%d.%d.%d", &a, &b, &c, &d)
		mask = ipv4.Init(a, b, c, d)

	} else if ipStr == "" && maskStr == "" && cidrStr != "" && hosts == 0 && subnets == 0 {
		ip, mask = ipv4.GetIpMaskFromCidr(cidrStr)

		if cls := ipv4.GetClass(ip); cls == ipv4.INVALID {
			log.Fatal("invalid ip")
		}
	} else if ipStr != "" && maskStr == "" && cidrStr == "" && hosts == 0 && subnets == 0 {
		fmt.Sscanf(ipStr, "%d.%d.%d.%d", &a, &b, &c, &d)
		ip = ipv4.Init(a, b, c, d)

		if cls := ipv4.GetClass(ip); cls == ipv4.INVALID {
			log.Fatal("invalid ip")
		} else if cls != ipv4.D && cls != ipv4.E {
			log.Fatal("invalid arguments")
		}
	} else {
		log.Fatal("invalid arguments")
	}

	return ip, mask
}

func printNetworkDetails(ip, mask ipv4.Address) {
	cls := ipv4.GetClass(ip)
	if cls == ipv4.D || cls == ipv4.E {
		fmt.Println("ip:", ip)
		fmt.Println("class:", cls)
		return // Exit early for D and E
	}

	cidr := ipv4.CidrNotation(ip, mask)
	subnetMax := ipv4.GetSubnetMaxNumber(ip, mask)
	hostMax := ipv4.GetHostMaxNumber(mask)
	subnetId := ipv4.GetSubNetId(ip, mask)
	hostId := ipv4.GetHostId(ip, mask)
	broadcast := ipv4.GetBroadCast(ip, mask)

	fmt.Println("ip:", ip)
	fmt.Println("mask:", mask)
	fmt.Println("CIDR notation:", cidr)
	fmt.Println("class:", cls)
	fmt.Println("subnet Number:", subnetMax)
	fmt.Println("host number:", hostMax)
	fmt.Println("subnet id:", subnetId)
	fmt.Println("host id:", hostId)
	fmt.Println("broadcast address:", broadcast)

	if subnetMax > 1 {
		printSubnetRanges(ip, subnetMax, hostMax)
	}
}

func printSubnetRanges(ip ipv4.Address, subnetMax int, hostMax int) {
	hoststoAdd := hostMax - 1
	addr := ipv4.GetFirstAddressInPlage(ip)

	fmt.Println("Plage:")
	fmt.Println("subnet \t\t\t plage \t\t     broadcast")
	for range subnetMax {
		fmt.Print(addr, " | ")
		addr++
		fmt.Print(addr, " - ")
		addr += ipv4.Address(hoststoAdd)
		fmt.Print(addr, " | ")
		addr++
		fmt.Print(addr)
		addr++
		fmt.Println()
	}
}
