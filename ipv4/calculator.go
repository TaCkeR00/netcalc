package ipv4

import (
	"fmt"
	"math"
)

func GetClass(ip Address) Class {
	filter := uint8(((ip >> 24) & MaxUint8))
	switch {
	case filter >= 1 && filter <= 127:
		return A
	case filter >= 128 && filter <= 191:
		return B
	case filter >= 192 && filter <= 223:
		return C
	case filter >= 224 && filter <= 239:
		return D
	case filter >= 240 && filter <= 255:
		return E
	default:
		return INVALID
	}
}

func GetSubnetMaxNumber(ip, mask Address) int {
	cls := GetClass(ip)
	resBit := GetReservedBitNumber(mask)
	var bitNumber uint8

	switch cls {
	case A:
		bitNumber = 24 - (32 - resBit)
	case B:
		bitNumber = 16 - (32 - resBit)
	case C:
		bitNumber = 8 - (32 - resBit)
	case D:
		bitNumber = 0
	case E:
		bitNumber = 0
	default:
		bitNumber = 0
	}
	if bitNumber > 0 {
		return int(math.Pow(2, float64(bitNumber)))
	} else {
		return 0
	}
}

func GetIpMaskFromCidr(cidr string) (Address, Address) {
	var a, b, c, d, Reserved uint8
	fmt.Sscanf(cidr, "%d.%d.%d.%d/%d", &a, &b, &c, &d, &Reserved)
	ip := Init(a, b, c, d)
	mask := GetMaskFromResersedBitsNumber(Reserved)
	return ip, mask
}

func GetMaskFromResersedBitsNumber(ReservedBitsNumber uint8) Address {
	return Address(^uint32(0) << (32 - ReservedBitsNumber))

}

func GetHostMaxNumber(mask Address) int {
	return int(math.Pow(2, float64(32-GetReservedBitNumber(mask)))) - 2
}

func GetSubNetId(ip, mask Address) Address {
	return ip & mask
}

func GetHostId(ip, mask Address) Address {
	return (mask ^ MaxUint32) & ip
}

func GetBroadCast(ip, mask Address) Address {
	return GetSubNetId(ip, mask) | (mask ^ MaxUint32)
}

func GetFirstAddressInPlage(ip Address) Address {
	var Addr Address
	cls := GetClass(ip)

	switch cls {
	case A:
		Addr = ip & GetMaskFromResersedBitsNumber(8)
	case B:
		Addr = ip & GetMaskFromResersedBitsNumber(16)
	case C:
		Addr = ip & GetMaskFromResersedBitsNumber(24)
	default:
		Addr = MaxUint32
	}

	return Addr
}

func CheckAddress(ip, mask Address) bool {
	cls := GetClass(ip)
	switch cls {
	case A:
		return ReservedBitsA <= GetReservedBitNumber(mask)
	case B:
		return ReservedBitsB <= GetReservedBitNumber(mask)
	case C:
		return ReservedBitsC <= GetReservedBitNumber(mask)
	default:
		return true
	}
}

func CheckHostsAndSubnets(ip Address, hosts, subnest uint8) bool {
	cls := GetClass(ip)

	switch cls {
	case A:
		return hosts+subnest+ReservedBitsA <= 32
	case B:
		return hosts+subnest+ReservedBitsB <= 32
	case C:
		return hosts+subnest+ReservedBitsC <= 32
	default:
		return false
	}
}

func GetMaskFromIpSubnet(ip Address, subnet uint8) Address {
	cls := GetClass(ip)
	var Addr Address

	switch cls {
	case A:
		Addr = GetMaskFromResersedBitsNumber(ReservedBitsA + subnet)
	case B:
		Addr = GetMaskFromResersedBitsNumber(ReservedBitsB + subnet)
	case C:
		Addr = GetMaskFromResersedBitsNumber(ReservedBitsC + subnet)
	}
	return Addr
}

func GetSubnetNumFromIpHost(ip Address, host uint8) uint8 {
	cls := GetClass(ip)
	var subnet uint8

	switch cls {
	case A:
		subnet = 32 - uint8(ReservedBitsA+host)
	case B:
		subnet = 32 - uint8(ReservedBitsB+host)
	case C:
		subnet = 32 - uint8(ReservedBitsC+host)
	}
	return subnet
}

func GetReservedBitNumber(mask Address) uint8 {
	var counter uint8 = 0
	for (mask & (1 << counter)) == 0 {
		counter++
	}
	return 32 - counter
}

func CidrNotation(ip, mask Address) string {
	counter := 0
	for (mask & (1 << counter)) == 0 {
		counter++
	}
	a, b, c, d := ip.GetAddressRepresentation()
	return fmt.Sprintf("%d.%d.%d.%d/%d", a, b, c, d, GetReservedBitNumber(mask))
}
