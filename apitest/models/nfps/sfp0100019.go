// Package models will define request and response message struct
// Version: v0.0.1
package models

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/sort"
)

type FP010019I struct {
	Data string `json:"data" validate:"required"`
}

type FP010019O struct {
	Data string `json:"data" validate:"required"`
}

var Spec0200RequestToPay = &iso8583.MessageSpec{
	Name: "Lookup and account verification request",
	Fields: map[int]field.Field{
		0: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message Type Indicator",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		1: field.NewBitmap(&field.Spec{
			Length:      16,
			Description: "Bitmap",
			Enc:         encoding.HexToASCII,
			Pref:        prefix.Hex.Fixed,
		}),
		2: field.NewString(&field.Spec{
			Length:      19,
			Description: "Primary Account Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		3: field.NewString(&field.Spec{
			Length:      6,
			Description: "Processing Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		4: field.NewString(&field.Spec{
			Length:      12,
			Description: "Transaction Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Settlement Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		6: field.NewString(&field.Spec{
			Length:      12,
			Description: "Billing Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		7: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transmission Date & Time",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		8: field.NewString(&field.Spec{
			Length:      8,
			Description: "Billing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		9: field.NewString(&field.Spec{
			Length:      8,
			Description: "Settlement Conversion Rate",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		10: field.NewString(&field.Spec{
			Length:      8,
			Description: "Cardholder Billing Conversion Rate",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		11: field.NewString(&field.Spec{
			Length:      6,
			Description: "Systems Trace Audit Number (STAN)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			// Pad:         padding.Left('0'),
		}),
		12: field.NewString(&field.Spec{
			Length:      6,
			Description: "Local Transaction Time",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		13: field.NewString(&field.Spec{
			Length:      4,
			Description: "Local Transaction Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		14: field.NewString(&field.Spec{
			Length:      4,
			Description: "Expiration Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		15: field.NewString(&field.Spec{
			Length:      4,
			Description: "Settlement Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.None,
		}),
		16: field.NewString(&field.Spec{
			Length:      4,
			Description: "Currency Conversion Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		17: field.NewString(&field.Spec{
			Length:      4,
			Description: "Capture Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		18: field.NewString(&field.Spec{
			Length:      4,
			Description: "Merchant Type",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		19: field.NewString(&field.Spec{
			Length:      3,
			Description: "Acquiring Institution Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		20: field.NewString(&field.Spec{
			Length:      3,
			Description: "PAN Extended Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		21: field.NewString(&field.Spec{
			Length:      3,
			Description: "Forwarding Institution Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		22: field.NewString(&field.Spec{
			Length:      3,
			Description: "Point of Sale (POS) Entry Mode",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		23: field.NewString(&field.Spec{
			Length:      3,
			Description: "Card Sequence Number (CSN)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         Right('0'),
		}),
		24: field.NewString(&field.Spec{
			Length:      3,
			Description: "Function Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		25: field.NewString(&field.Spec{
			Length:      2,
			Description: "Point of Service Condition Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		26: field.NewString(&field.Spec{
			Length:      2,
			Description: "Point of Service PIN Capture Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		27: field.NewString(&field.Spec{
			Length:      1,
			Description: "Authorizing Identification Response Length",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		28: field.NewString(&field.Spec{
			Length:      9,
			Description: "Transaction Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		29: field.NewString(&field.Spec{
			Length:      9,
			Description: "Settlement Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		30: field.NewString(&field.Spec{
			Length:      9,
			Description: "Transaction Processing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		31: field.NewString(&field.Spec{
			Length:      9,
			Description: "Settlement Processing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		32: field.NewString(&field.Spec{
			Length:      11,
			Description: "Acquiring Institution Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		33: field.NewString(&field.Spec{
			Length:      11,
			Description: "Forwarding Institution Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		34: field.NewString(&field.Spec{
			Length:      28,
			Description: "Extended Primary Account Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		35: field.NewString(&field.Spec{
			Length:      37,
			Description: "Track 2 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		36: field.NewString(&field.Spec{
			Length:      104,
			Description: "Track 3 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		37: field.NewString(&field.Spec{
			Length:      12,
			Description: "Retrieval Reference Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		38: field.NewString(&field.Spec{
			Length:      6,
			Description: "Authorization Identification Response",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		39: field.NewString(&field.Spec{
			Length:      2,
			Description: "Response Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		40: field.NewString(&field.Spec{
			Length:      3,
			Description: "Service Restriction Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		41: field.NewString(&field.Spec{
			Length:      8,
			Description: "Card Acceptor Terminal Identification",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		42: field.NewString(&field.Spec{
			Length:      15,
			Description: "Card Acceptor Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		43: field.NewComposite(&field.Spec{
			Length:      40,
			Description: "Card Acceptor Name/Location",
			Pref:        prefix.ASCII.Fixed,
			Tag: &field.TagSpec{
				Sort: sort.StringsByInt,
			},
			Subfields: map[string]field.Field{
				"1": field.NewString(&field.Spec{
					Length:      4,
					Description: "Terminal owner ID",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"2": field.NewString(&field.Spec{
					Length:      16,
					Description: "Terminal ID Card Acceptor Name",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"3": field.NewString(&field.Spec{
					Length:      2,
					Description: "Terminal Type",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"4": field.NewString(&field.Spec{
					Length:      4,
					Description: "Branch ID",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         nil,
				}),
				"5": field.NewString(&field.Spec{
					Length:      2,
					Description: "Terminal State",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"6": field.NewString(&field.Spec{
					Length:      4,
					Description: "Province Code (Region ID)",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"7": field.NewString(&field.Spec{
					Length:      2,
					Description: "PromptPay Flag",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"8": field.NewString(&field.Spec{
					Length:      4,
					Description: "Filler",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         padding.Left(' '),
				}),
				"9": field.NewString(&field.Spec{
					Length:      2,
					Description: "Country Code",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
			},
		}),
		44: field.NewString(&field.Spec{
			Length:      99,
			Description: "Additional Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		45: field.NewString(&field.Spec{
			Length:      76,
			Description: "Track 1 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		46: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data (ISO)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		47: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		48: field.NewComposite(&field.Spec{
			Length:      999,
			Description: "Additional data Private(ITMX segment)",
			Pref:        prefix.ASCII.LLL,
			Tag: &field.TagSpec{
				Sort: sort.StringsByInt,
			},
			Subfields: map[string]field.Field{
				"1": field.NewString(&field.Spec{
					Length:      12,
					Description: "Sender Fee",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         padding.Left('0'),
				}),
				"2": field.NewString(&field.Spec{
					Length:      12,
					Description: "Transferer Fee / From Account Fee",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         padding.Left('0'),
				}),
				"3": field.NewString(&field.Spec{
					Length:      12,
					Description: "Transferee Fee / To Account Fee",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         padding.Left('0'),
				}),
				"4": field.NewString(&field.Spec{
					Length:      3,
					Description: "Sending Bank",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"5": field.NewString(&field.Spec{
					Length:      40,
					Description: "From Account No",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"6": field.NewString(&field.Spec{
					Length:      50,
					Description: "From Account Name",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"7": field.NewString(&field.Spec{
					Length:      128,
					Description: "Receiving ID (Proxy ID) ",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"8": field.NewString(&field.Spec{
					Length:      12,
					Description: "Proxy Type",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"9": field.NewString(&field.Spec{
					Length:      3,
					Description: "Destination Bank",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"10": field.NewString(&field.Spec{
					Length:      40,
					Description: "To Account No / Destination E-wallet No",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"11": field.NewString(&field.Spec{
					Length:      50,
					Description: "To Account Display Name / Destination E-wallet Display Name",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"12": field.NewString(&field.Spec{
					Length:      50,
					Description: "To Account Name / Destination E-wallet Name",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"13": field.NewString(&field.Spec{
					Length:      1,
					Description: "PIN Block Type",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"14": field.NewString(&field.Spec{
					Length:      10,
					Description: "One time password",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"15": field.NewString(&field.Spec{
					Length:      40,
					Description: "Additional Note",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"16": field.NewString(&field.Spec{
					Length:      16,
					Description: "Transaction Reference",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"17": field.NewString(&field.Spec{
					Length:      128,
					Description: "Sender ID (Proxy ID)",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"18": field.NewString(&field.Spec{
					Length:      12,
					Description: "Sender Proxy Type",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"19": field.NewString(&field.Spec{
					Length:      16,
					Description: "Credit Notification Reference",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"20": field.NewString(&field.Spec{
					Length:      20,
					Description: "Bill Reference 1",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"21": field.NewString(&field.Spec{
					Length:      20,
					Description: "Bill Reference 2",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"22": field.NewString(&field.Spec{
					Length:      20,
					Description: "Bill Reference 3",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"23": field.NewString(&field.Spec{
					Length:      100,
					Description: "Bill Presentment URL",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
				"24": field.NewString(&field.Spec{
					Length:      16,
					Description: "Expiry Date",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         Right(' '),
				}),
			},
		}),
		49: field.NewString(&field.Spec{
			Length:      3,
			Description: "Transaction Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		50: field.NewString(&field.Spec{
			Length:      3,
			Description: "Settlement Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		51: field.NewString(&field.Spec{
			Length:      3,
			Description: "Cardholder Billing Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		52: field.NewString(&field.Spec{
			Length:      16,
			Description: "PIN Data",
			Enc:         encoding.HexToASCII,
			Pref:        prefix.Hex.Fixed,
		}),
		53: field.NewString(&field.Spec{
			Length:      16,
			Description: "Security Related Control Information",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		54: field.NewString(&field.Spec{
			Length:      120,
			Description: "Additional Amounts",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		55: field.NewString(&field.Spec{
			Length:      999,
			Description: "EMV DATA (ICC RELATED DATA)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		56: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (ISO)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		57: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		58: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		59: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		60: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		61: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (Private)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		62: field.NewComposite(&field.Spec{
			Length:      999,
			Description: "ITMX ADDITIONAL POS INFORMATION",
			Pref:        prefix.ASCII.LLL,
			Tag: &field.TagSpec{
				Sort: sort.StringsByInt,
			},
			Subfields: map[string]field.Field{
				"1": field.NewString(&field.Spec{
					Length:      1,
					Description: "Terminal Type",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"2": field.NewString(&field.Spec{
					Length:      1,
					Description: "Terminal Entry Capability",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"3": field.NewString(&field.Spec{
					Length:      1,
					Description: "Chip Condition Code",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"4": field.NewString(&field.Spec{
					Length:      1,
					Description: "Service Development",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
					Pad:         nil,
				}),
				"5": field.NewString(&field.Spec{
					Length:      2,
					Description: "Not Applicable",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"6": field.NewString(&field.Spec{
					Length:      1,
					Description: "Chip Transaction Indicator",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"7": field.NewString(&field.Spec{
					Length:      1,
					Description: "Chip card authentication reliability indicator",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"8": field.NewString(&field.Spec{
					Length:      2,
					Description: "Mail/Phone/Electronic commerce indicator",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"9": field.NewString(&field.Spec{
					Length:      1,
					Description: "Not applicable (SMS) only",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
				"10": field.NewString(&field.Spec{
					Length:      1,
					Description: "Partial Authorization Indicator",
					Enc:         encoding.ASCII,
					Pref:        prefix.ASCII.Fixed,
				}),
			}}),
		63: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (Private)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		64: field.NewString(&field.Spec{
			Length:      8,
			Description: "Message Authentication Code (MAC)",
			Enc:         encoding.HexToASCII,
			Pref:        prefix.Hex.Fixed,
		}),
		70: field.NewString(&field.Spec{
			Length:      3,
			Description: "Network management information code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		90: field.NewString(&field.Spec{
			Length:      42,
			Description: "Original Data Elements",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		102: field.NewString(&field.Spec{
			Length:      28,
			Description: "FROM ACCOUNT",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		103: field.NewString(&field.Spec{
			Length:      28,
			Description: "TO ACCOUNT",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
	},
}

type ISO85830200RequestToPayData struct {
	F0   *field.String
	F1   *field.Bitmap
	F2   *field.String
	F3   *field.String
	F4   *field.String
	F5   *field.String
	F6   *field.String
	F7   *field.String
	F8   *field.String
	F9   *field.String
	F10  *field.String
	F11  *field.String
	F12  *field.String
	F13  *field.String
	F14  *field.String
	F15  *field.String
	F16  *field.String
	F17  *field.String
	F18  *field.String
	F19  *field.String
	F20  *field.String
	F21  *field.String
	F22  *field.String
	F23  *field.String
	F24  *field.String
	F25  *field.String
	F26  *field.String
	F27  *field.String
	F28  *field.String
	F29  *field.String
	F30  *field.String
	F31  *field.String
	F32  *field.String
	F33  *field.String
	F34  *field.String
	F35  *field.String
	F36  *field.String
	F37  *field.String
	F38  *field.String
	F39  *field.String
	F40  *field.String
	F41  *field.String
	F42  *field.String
	F43  *ISO85830200F43Data
	F44  *field.String
	F45  *field.String
	F46  *field.String
	F47  *field.String
	F48  *ISO85830200RequestToPayF48Data
	F49  *field.String
	F50  *field.String
	F51  *field.String
	F52  *field.String
	F53  *field.String
	F54  *field.String
	F55  *field.String
	F56  *field.String
	F57  *field.String
	F58  *field.String
	F59  *field.String
	F60  *field.String
	F61  *field.String
	F62  *ISO85830200F62Data
	F63  *field.String
	F64  *field.String
	F70  *field.String
	F90  *field.String
	F102 *field.String
	F103 *field.String
}
type ISO85830200RequestToPayF48Data struct {
	F1  *field.String
	F2  *field.String
	F3  *field.String
	F4  *field.String
	F5  *field.String
	F6  *field.String
	F7  *field.String
	F8  *field.String
	F9  *field.String
	F10 *field.String
	F11 *field.String
	F12 *field.String
	F13 *field.String
	F14 *field.String
	F15 *field.String
	F16 *field.String
	F17 *field.String
	F18 *field.String
	F19 *field.String
	F20 *field.String
	F21 *field.String
	F22 *field.String
	F23 *field.String
	F24 *field.String
}

// type ISO85830210RequestToPayData ISO85830200RequestToPayData

//
// type ISO85830210RequestToPayF48Data struct {
//	F1  *field.String
//	F2  *field.String
//	F3  *field.String
//	F4  *field.String
//	F5  *field.String
//	F6  *field.String
//	F7  *field.String
//	F8  *field.String
//	F9  *field.String
//	F10 *field.String
//	F11 *field.String
//	F12 *field.String
//	F13 *field.String
//	F14 *field.String
//	F15 *field.String
//	F16 *field.String
//	F17 *field.String
//	F18 *field.String
//	F19 *field.String
//	F20 *field.String
//	F21 *field.String
//	F22 *field.String
//	F23 *field.String
//	F24 *field.String
//	F25 *field.String
//	F26 *field.String
//	F27 *field.String
//	F28 *field.String
//	F29 *field.String
//	F30 *field.String
//	F31 *field.String
//	F32 *field.String
//	F33 *field.String
// }
func (*FP010019I) GetServiceKey() string {
	return "FP010019"
}