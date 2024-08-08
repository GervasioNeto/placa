package main

import (
	"bytes"
	"fmt"
	"strings"
)

// signValue := Placa do carro

func main() {
	signValue := "HZB"
	signByteValue := []byte(signValue[:3])
	tesAlgo := map[string][]string{
		"Paraná":              []string{"AAA-BEZ", "RHA-RHZ"},
		"Sao Paulo":           []string{"BFA-GKI", "QSN-QSZ"},
		"Minas Gerais":        []string{"GKJ-HOK", "NXX-NYG", "OLO-OMH", "OOV-ORC", "OWH-OXK", "PUA-PZZ", "QMQ-QQZ", "QUA-QUZ", "QWR-QXZ", "RFA-RGD", "RMD-RNZ", "RTA-RVZ"},
		"Maranhão":            []string{"HOL-HQE", "NHA-NHT", "NMP-NNI", "NWS-NXQ", "OIR-OJQ", "OXQ-OXZ", "PSA-PTZ", "ROA-ROZ"},
		"Mato Grosso do Sul":  []string{"HQF-HTW", "NRF-NSD", "OOG-OOU", "QAA-QAZ", "REW-REZ", "RWA-RWJ"},
		"Ceará":               []string{"HTX-HZA", "NQL-NRE", "NUM-NVF", "OCB-OCU", "OHX-OIQ", "ORN-OSV", "OZA-OZA", "PMA-POZ", "RIA-RIN", "SAN-SBV"},
		"Sergipe":             []string{"HZB-IAP", "NVG-NVN", "OEJ-OES", "OZB-OZB", "QKN-QKZ", "QMA-QMP", "RQW-RRH"},
		"Rio Grande do Sul":   []string{"IAQ-JDO"},
		"Distrito Federal":    []string{"JDP-JKR", "OVM-OVV", "OZW-PBZ", "REC-REV"},
		"Bahia":               []string{"JKS-JSZ", "NTD-NTW", "NYH-NZZ", "OKI-OLG", "OUF-OVD", "OZC-OZV", "PJA-PLZ", "QTU-QTZ", "RCO-RDR", "RPA-RQL"},
		"Pará":                []string{"JTA-JWE", "NSE-NTC", "OBT-OCA", "OFI-OFW", "OSW-OTZ", "QDA-QEZ", "QVA-QVZ", "RWK-RXD"},
		"Amazonas":            []string{"JWF-JXY", "NOI-NPB", "OAA-OAO", "OXM-OXM", "PHA-PHZ", "QZA-QZZ"},
		"Goiás":               []string{"KAV-KFC", "NFC-NGZ", "NJX-NLU", "NVO-NWR", "OGH-OHA", "OMI-OOF", "PQA-PRZ", "QTN-QTS", "RBK-RCN", "SBW-SDS"},
		"Pernambuco":          []string{"KFD-KME", "NXU-NXW", "OYL-OYZ", "PCA-PED", "PEE-PFQ", "PFR-PGK", "PGL-PGU", "PGV-PGZ", "QYA-QYZ", "RZE-RZZ"},
		"Rio de Janeiro":      []string{"KMF-LVE", "RIO-RIO", "RIP-RKV"},
		"Piauí":               []string{"LVF-LWQ", "NHU-NIX", "ODU-OEI", "OUA-OUE", "OVW-OVY", "PIA-PIZ", "QRN-QRZ", "RSG-RST"},
		"Santa Catarina":      []string{"LWR-MMM", "OKD-OKH", "QHA-QJZ", "QTK-QTM", "RAA-RAJ", "RDS-REB", "RKW-RLP", "RXK-RYI"},
		"Paraíba":             []string{"MMN-MOW", "NPR-NQK", "OET-OFH", "OFX-OGG", "OXO-OXO", "QFA-QFZ", "QSA-QSM", "RLQ-RMC"},
		"Espírito Santo":      []string{"MOX-MTZ", "OCV-ODT", "OVE-OVF", "OVH-OVL", "OYD-OYK", "PPA-PPZ", "QRB-QRM", "RBA-RBJ", "RQM-RQV"},
		"Alagoas":             []string{"MUA-MVK", "NLV-NMO", "OHB-OHK", "ORD-ORM", "OXN-OXN", "QLA-QLM", "QTT-QTT", "QWG-QWL", "RGO-RGU", "SAA-SAJ"},
		"Tocantins":           []string{"MVL-MXG", "OLH-OLN", "OYA-OYC", "QKA-QKM", "QWA-QWF", "RSA-RSF"},
		"Rio Grande do Norte": []string{"MXH-MZM", "NNJ-NOH", "OJR-OKC", "OVZ-OWG", "QGA-QGZ", "RGE-RGM", "RGN-RGN"},
		"Acre":                []string{"MZN-NAG", "NXR-NXT", "OVG-OVG", "OXP-OXP", "QLU-QLZ", "QWM-QWQ"},
		"Roraima":             []string{"NAH-NBA", "NUH-NUL", "RZA-RZD"},
		"Rondônia":            []string{"NBB-NEH", "OHL-OHW", "OXL-OXL", "QRA-QRA", "QTA-QTJ", "RSU-RSZ"},
		"Amapá":               []string{"NEI-NFB", "QLN-QLT", "SAK-SAM"},
		"Mato Grosso":         []string{"JXZ-KAU", "NIY-NJW", "NPC-NPQ", "NTX-NUG", "OAP-OBS", "QBA-QCZ", "RAK-RAZ", "RRI-RRZ"},
		"Ainda não Definida":  []string{"RGV-RGZ", "RXE-RXJ", "RYJ-RYZ", "SDT-ZZZ"},
	}

	var state string
	for key, value := range tesAlgo {
		for _, intervals := range value {
			stateConstrains := strings.Split(intervals, "-")
			minValue := []byte(stateConstrains[0])
			maxValue := []byte(stateConstrains[1])

			if bytes.Compare(signByteValue, minValue) >= 0 && bytes.Compare(signByteValue, maxValue) <= 0 {
				state = key
				break
			}
		}
	}

	fmt.Printf("SIGN STATE: %s \n", state)
}
