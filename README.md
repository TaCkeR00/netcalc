# netcalc 🌐

A fast, efficient IPv4 subnet calculator and CIDR tool built entirely in Go. 

Instead of relying on heavy string manipulation, netcalc uses underlying bitwise operations on uint32 types to calculate subnet masks, network IDs, broadcast addresses, and host ranges. It is designed to be a lightweight, modular CLI tool for network engineers, SREs, and developers.

## ✨ Features

* Multiple Input Formats: Calculate subnets using CIDR notation, IP + Mask, or by specifying the required number of hosts/subnets.
* Bitwise Efficiency: Core logic relies on raw bitwise operators (<<, >>, &, |, ^) for optimal performance.
* Subnet Plage Generator: Automatically generates a visual table of all available subnet ranges, usable hosts, and broadcast addresses.
* Class Identification: Automatically identifies traditional IP classes (A, B, C, D, E).

## 🚀 Installation

Ensure you have Go installed on your machine. You can build the executable directly from the source:

```bash
git clone https://github.com/TaCkeR00/netcalc.git
cd netcalc
go build -o netcalc ./cmd/netcalc
