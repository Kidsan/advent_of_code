package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_hexToBinary(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "example", args: args{input: "D2FE28"}, want: "110100101111111000101000"},
		{
			name: "example",
			args: args{
				input: "020D708041258C0B4C683E61F674A1401595CC3DE669AC4FB7BEFEE840182CDF033401296F44367F938371802D2CC9801A980021304609C431007239C2C860400F7C36B005E446A44662A2805925FF96CBCE0033C5736D13D9CFCDC001C89BF57505799C0D1802D2639801A900021105A3A43C1007A1EC368A72D86130057401782F25B9054B94B003013EDF34133218A00D4A6F1985624B331FE359C354F7EB64A8524027D4DEB785CA00D540010D8E9132270803F1CA1D416200FDAC01697DCEB43D9DC5F6B7239CCA7557200986C013912598FF0BE4DFCC012C0091E7EFFA6E44123CE74624FBA01001328C01C8FF06E0A9803D1FA3343E3007A1641684C600B47DE009024ED7DD9564ED7DD940C017A00AF26654F76B5C62C65295B1B4ED8C1804DD979E2B13A97029CFCB3F1F96F28CE43318560F8400E2CAA5D80270FA1C90099D3D41BE00DD00010B893132108002131662342D91AFCA6330001073EA2E0054BC098804B5C00CC667B79727FF646267FA9E3971C96E71E8C00D911A9C738EC401A6CBEA33BC09B8015697BB7CD746E4A9FD4BB5613004BC01598EEE96EF755149B9A049D80480230C0041E514A51467D226E692801F049F73287F7AC29CB453E4B1FDE1F624100203368B3670200C46E93D13CAD11A6673B63A42600C00021119E304271006A30C3B844200E45F8A306C8037C9CA6FF850B004A459672B5C4E66A80090CC4F31E1D80193E60068801EC056498012804C58011BEC0414A00EF46005880162006800A3460073007B620070801E801073002B2C0055CEE9BC801DC9F5B913587D2C90600E4D93CE1A4DB51007E7399B066802339EEC65F519CF7632FAB900A45398C4A45B401AB8803506A2E4300004262AC13866401434D984CA4490ACA81CC0FB008B93764F9A8AE4F7ABED6B293330D46B7969998021C9EEF67C97BAC122822017C1C9FA0745B930D9C480"},
			want: "000000100000110101110000100000000100000100100101100011000000101101001100011010000011111001100001111101100111010010100001010000000001010110010101110011000011110111100110011010011010110001001111101101111011111011111110111010000100000000011000001011001101111100000011001101000000000100101001011011110100010000110110011111111001001110000011011100011000000000101101001011001100100110000000000110101001100000000000001000010011000001000110000010011100010000110001000000000111001000111001110000101100100001100000010000000000111101111100001101101011000000000101111001000100011010100100010001100110001010100010100000000101100100100101111111111001011011001011110011100000000000110011110001010111001101101101000100111101100111001111110011011100000000000001110010001001101111110101011101010000010101111001100111000000110100011000000000101101001001100011100110000000000110101001000000000000001000010001000001011010001110100100001111000001000000000111101000011110110000110110100010100111001011011000011000010011000000000101011101000000000101111000001011110010010110111001000001010100101110010100101100000000001100000001001111101101111100110100000100110011001000011000101000000000110101001010011011110001100110000101011000100100101100110011000111111110001101011001110000110101010011110111111010110110010010101000010100100100000000100111110101001101111010110111100001011100101000000000110101010100000000000001000011011000111010010001001100100010011100001000000000111111000111001010000111010100000101100010000000001111110110101100000000010110100101111101110011101011010000111101100111011100010111110110101101110010001110011100110010100111010101010111001000000000100110000110110000000001001110010001001001011001100011111111000010111110010011011111110011000000000100101100000000001001000111100111111011111111101001101110010001000001001000111100111001110100011000100100111110111010000000010000000000010011001010001100000000011100100011111111000001101110000010101001100000000011110100011111101000110011010000111110001100000000011110100001011001000001011010000100110001100000000010110100011111011110000000001001000000100100111011010111110111011001010101100100111011010111110111011001010000001100000000010111101000000000101011110010011001100101010011110111011010110101110001100010110001100101001010010101101100011011010011101101100011000001100000000100110111011001011110011110001010110001001110101001011100000010100111001111110010110011111100011111100101101111001010001100111001000011001100011000010101100000111110000100000000001110001011001010101001011101100000000010011100001111101000011100100100000000100110011101001111010100000110111110000000001101110100000000000000010000101110001001001100010011001000010000100000000000001000010011000101100110001000110100001011011001000110101111110010100110001100110000000000000001000001110011111010100010111000000000010101001011110000001001100010000000010010110101110000000000110011000110011001111011011110010111001001111111111101100100011000100110011111111010100111100011100101110001110010010110111001110001111010001100000000001101100100010001101010011100011100111000111011000100000000011010011011001011111010100011001110111100000010011011100000000001010101101001011110111011011111001101011101000110111001001010100111111101010010111011010101100001001100000000010010111100000000010101100110001110111011101001011011101111011101010101000101001001101110011010000001001001110110000000010010000000001000110000110000000000010000011110010100010100101001010001010001100111110100100010011011100110100100101000000000011111000001001001111101110011001010000111111101111010110000101001110010110100010100111110010010110001111111011110000111110110001001000001000000000010000000110011011010001011001101100111000000100000000011000100011011101001001111010001001111001010110100010001101001100110011100111011011000111010010000100110000000001100000000000000001000010001000110011110001100000100001001110001000000000110101000110000110000111011100001000100001000000000111001000101111110001010001100000110110010000000001101111100100111001010011011111111100001010000101100000000010010100100010110010110011100101011010111000100111001100110101010000000000010010000110011000100111100110001111000011101100000000001100100111110011000000000011010001000000000011110110000000101011001001001100000000001001010000000010011000101100000000001000110111110110000000100000101001010000000001110111101000110000000000101100010000000000101100010000000000110100000000000101000110100011000000000011100110000000001111011011000100000000001110000100000000001111010000000000100000111001100000000001010110010110000000000010101011100111011101001101111001000000000011101110010011111010110111001000100110101100001111101001011001001000001100000000011100100110110010011110011100001101001001101101101010001000000000111111001110011100110011011000001100110100000000010001100111001111011101100011001011111010100011001110011110111011000110010111110101011100100000000101001000101001110011000110001001010010001011011010000000001101010111000100000000011010100000110101000101110010000110000000000000000010000100110001010101100000100111000011001100100000000010100001101001101100110000100110010100100010010010000101011001010100000011100110000001111101100000000100010111001001101110110010011111001101010001010111001001111011110101011111011010110101100101001001100110011000011010100011010110111100101101001100110011000000000100001110010011110111011110110011111001001011110111010110000010010001010000010001000000001011111000001110010011111101000000111010001011011100100110000110110011100010010000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexToBinary(tt.args.input); got != tt.want {
				t.Errorf("hexToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		binaryDefinition string
	}
	tests := []struct {
		name string
		args args
		want packet
	}{
		{
			name: "Literal Representation",
			args: args{
				binaryDefinition: "110100101111111000101000",
			},
			want: packet{
				version:      6,
				binary:       "110100101111111000101000",
				packetTypeId: 4,
				value:        2021,
			},
		},
		{
			name: "Operator Type 1",
			args: args{
				binaryDefinition: "00111000000000000110111101000101001010010001001000000000",
			},
			want: packet{
				version:      1,
				binary:       "00111000000000000110111101000101001010010001001000000000",
				packetTypeId: 6,
				subPackets: []packet{
					{
						binary:               "",
						packetTypeId:         0,
						version:              0,
						subPacketLength:      0,
						subPacketLengthValue: 0,
						numberOfSubPackets:   0,
						value:                10,
						subPackets:           []packet{},
					},
					{
						binary:               "",
						packetTypeId:         0,
						version:              0,
						subPacketLength:      0,
						subPacketLengthValue: 0,
						numberOfSubPackets:   0,
						value:                20,
						subPackets:           []packet{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := binaryString{}
			for _, value := range strings.Split(tt.args.binaryDefinition, "") {
				b = append(b, value)
			}
			got := b.Parse()
			if !reflect.DeepEqual(got.version, tt.want.version) {
				t.Errorf("Version = %v, want %v", got.version, tt.want.version)
			}

			if !reflect.DeepEqual(got.packetTypeId, tt.want.packetTypeId) {
				t.Errorf("packetTypeId = %v, want %v", got.packetTypeId, tt.want.packetTypeId)
			}
			if !reflect.DeepEqual(got.value, tt.want.value) {
				t.Errorf("Value = %v, want %v", got.value, tt.want.value)
			}

			for i, p := range tt.want.subPackets {
				if p.value != got.subPackets[i].value {
					t.Errorf("Value = %v, want %v", got.subPackets[i].value, p.value)
				}
			}
		})
	}
}

func TestSumVersions(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Operator Type 1",
			args: args{
				hex: "8A004A801A8002F478",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := hexToBinary(tt.args.hex)
			b := binaryString{}
			for _, value := range strings.Split(bs, "") {
				b = append(b, value)
			}
			parsed := b.Parse()
			actual := parsed.sumVersions()
			if actual != tt.want {
				t.Errorf("version sum = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				hex: "C200B40A82",
			},
			want: 3,
		},
		{
			name: "Example",
			args: args{
				hex: "04005AC33890",
			},
			want: 54,
		},
		{
			name: "Example",
			args: args{
				hex: "880086C3E88112",
			},
			want: 7,
		},
		{
			name: "Example",
			args: args{
				hex: "CE00C43D881120",
			},
			want: 9,
		},
		{
			name: "Example",
			args: args{
				hex: "D8005AC2A8F0",
			},
			want: 1,
		},
		{
			name: "Example",
			args: args{
				hex: "F600BC2D8F",
			},
			want: 0,
		},
		{
			name: "Example",
			args: args{
				hex: "9C005AC2F8F0",
			},
			want: 0,
		},
		{
			name: "Example",
			args: args{
				hex: "9C0141080250320F1802104A08",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := hexToBinary(tt.args.hex)
			b := binaryString{}
			for _, value := range strings.Split(bs, "") {
				b = append(b, value)
			}
			parsed := b.Parse()
			actual := parsed.part2()
			if actual != tt.want {
				t.Errorf("version sum = %v, want %v", actual, tt.want)
			}
		})
	}
}
