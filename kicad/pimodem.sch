EESchema Schematic File Version 4
LIBS:pimodem-cache
EELAYER 26 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "PiModem"
Date "08/10/17"
Rev "1.0"
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L pimodem:LTC2951 U1
U 1 1 596D4336
P 5100 2500
F 0 "U1" H 4800 2850 60  0000 C CNN
F 1 "LTC2951-2" H 5100 2550 60  0000 C CNN
F 2 "TO_SOT_Packages_SMD:SOT-23-8" H 5100 1900 60  0001 C CNN
F 3 "http://cds.linear.com/docs/en/datasheet/295112fb.pdf" H 5050 2600 60  0001 C CNN
F 4 "LTC2951CTS8-2#TRMPBF" H 5100 2500 60  0001 C CNN "MFN"
F 5 "LTC2951CTS8-2#TRMPBFCT-ND" H 5100 2500 60  0001 C CNN "Digikey"
	1    5100 2500
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR01
U 1 1 597B55C9
P 850 2350
F 0 "#PWR01" H 850 2100 50  0001 C CNN
F 1 "GND" H 850 2200 50  0000 C CNN
F 2 "" H 850 2350 50  0000 C CNN
F 3 "" H 850 2350 50  0000 C CNN
	1    850  2350
	1    0    0    -1  
$EndComp
$Comp
L Transistor_BJT:BC857BS Q2
U 1 1 597E9C9D
P 2550 2850
F 0 "Q2" H 2350 2750 50  0000 L CNN
F 1 "BC857BS" H 2200 3000 50  0000 L CNN
F 2 "Package_TO_SOT_SMD:SOT-363_SC-70-6" H 2750 2950 50  0001 C CNN
F 3 "https://www.diodes.com/assets/Datasheets/ds30373.pdf" H 2550 2850 50  0001 C CNN
F 4 "BC857BS-7-F" H 2550 2850 60  0001 C CNN "MFN"
F 5 "BC857BS-7FDICT-ND" H 2550 2850 60  0001 C CNN "Digikey"
	1    2550 2850
	-1   0    0    1   
$EndComp
$Comp
L Transistor_BJT:BC857BS Q2
U 2 1 597E9CEC
P 3350 2850
F 0 "Q2" H 3100 2750 50  0000 L CNN
F 1 "BC857BS" H 3550 2800 50  0001 L CNN
F 2 "Package_TO_SOT_SMD:SOT-363_SC-70-6" H 3550 2950 50  0001 C CNN
F 3 "https://www.diodes.com/assets/Datasheets/ds30373.pdf" H 3350 2850 50  0001 C CNN
	2    3350 2850
	1    0    0    1   
$EndComp
$Comp
L power:GND #PWR02
U 1 1 597EB2AC
P 2100 2750
F 0 "#PWR02" H 2100 2500 50  0001 C CNN
F 1 "GND" H 2100 2600 50  0000 C CNN
F 2 "" H 2100 2750 50  0000 C CNN
F 3 "" H 2100 2750 50  0000 C CNN
	1    2100 2750
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C1
U 1 1 597EB89E
P 2100 2550
F 0 "C1" H 2192 2596 50  0000 L CNN
F 1 "1uF" H 2192 2505 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 2100 2550 50  0001 C CNN
F 3 "http://datasheets.avx.com/X7RDielectric.pdf" H 2100 2550 50  0001 C CNN
F 4 "0805ZC105KAT2A" H 2100 2550 60  0001 C CNN "MFN"
F 5 "478-1405-1-ND" H 2100 2550 60  0001 C CNN "Digikey"
	1    2100 2550
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR03
U 1 1 597EBC5B
P 4900 3050
F 0 "#PWR03" H 4900 2800 50  0001 C CNN
F 1 "GND" H 4900 2900 50  0000 C CNN
F 2 "" H 4900 3050 50  0000 C CNN
F 3 "" H 4900 3050 50  0000 C CNN
	1    4900 3050
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C3
U 1 1 597EBE18
P 5100 3200
F 0 "C3" H 5000 3300 50  0000 L CNN
F 1 "10 uF" V 5200 3100 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 5100 3200 50  0001 C CNN
F 3 "http://www.kemet.com/Lists/ProductCatalog/Attachments/19/KEM_C1006_X5R_SMD.pdf" H 5100 3200 50  0001 C CNN
F 4 "C0805C106K8PACTU" H 5100 3200 60  0001 C CNN "MFN"
F 5 "399-4925-1-ND" H 5100 3200 60  0001 C CNN "Digikey"
	1    5100 3200
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C4
U 1 1 597EBF6E
P 5300 3100
F 0 "C4" H 5200 3200 50  0000 L CNN
F 1 "0.033 uF" V 5400 3000 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 5300 3100 50  0001 C CNN
F 3 "http://katalog.we-online.de/pbs/datasheet/885012207014.pdf" H 5300 3100 50  0001 C CNN
F 4 "885012207014" H 5300 3100 60  0001 C CNN "MFN"
F 5 "732-8024-1-ND" H 5300 3100 60  0001 C CNN "Digikey"
	1    5300 3100
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR04
U 1 1 597EC278
P 5100 3400
F 0 "#PWR04" H 5100 3150 50  0001 C CNN
F 1 "GND" H 5100 3250 50  0000 C CNN
F 2 "" H 5100 3400 50  0000 C CNN
F 3 "" H 5100 3400 50  0000 C CNN
	1    5100 3400
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR05
U 1 1 597EC299
P 5300 3400
F 0 "#PWR05" H 5300 3150 50  0001 C CNN
F 1 "GND" H 5300 3250 50  0000 C CNN
F 2 "" H 5300 3400 50  0000 C CNN
F 3 "" H 5300 3400 50  0000 C CNN
	1    5300 3400
	1    0    0    -1  
$EndComp
$Comp
L Device:Polyfuse F1
U 1 1 597EC70F
P 1800 2300
F 0 "F1" V 2000 2250 50  0000 C CNN
F 1 "MF-MSMF200" V 1900 2250 50  0000 C CNN
F 2 "Resistors_SMD:R_1812" H 1850 2100 50  0001 L CNN
F 3 "http://www.bourns.com/docs/Product-Datasheets/mfmsmf.pdf" H 1800 2300 50  0001 C CNN
F 4 "MF-MSMF200-2" V 1800 2300 60  0001 C CNN "MFN"
F 5 "MF-MSMF200-2CT-ND" V 1800 2300 60  0001 C CNN "Digikey"
	1    1800 2300
	0    1    1    0   
$EndComp
Text GLabel 6400 2600 2    60   Input ~ 0
~KILL~
Text GLabel 6950 2450 2    60   Output ~ 0
~INT~
$Comp
L Device:R R1
U 1 1 597EE4F0
P 2450 3450
F 0 "R1" H 2520 3496 50  0000 L CNN
F 1 "10k(1%)" H 2520 3405 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 2380 3450 50  0001 C CNN
F 3 "http://www.yageo.com/documents/recent/PYu-RC_Group_51_RoHS_L_7.pdf" H 2450 3450 50  0001 C CNN
F 4 "RC0805FR-0710KL" H 2450 3450 60  0001 C CNN "MFN"
F 5 "311-10.0KCRCT-ND" H 2450 3450 60  0001 C CNN "Digikey"
	1    2450 3450
	1    0    0    -1  
$EndComp
$Comp
L pimodem:LTC4360 U2
U 1 1 597FC74A
P 6450 1800
F 0 "U2" H 6500 2150 60  0000 C CNN
F 1 "LTC4360-1" H 6400 1650 60  0000 C CNN
F 2 "TO_SOT_Packages_SMD:SC-70-8" H 6650 1900 60  0001 C CNN
F 3 "http://cds.linear.com/docs/en/datasheet/436012fa.pdf" H 6650 1900 60  0001 C CNN
F 4 "LTC4360ISC8-1#TRMPBF" H 6450 1800 60  0001 C CNN "MFN"
F 5 "LTC4360ISC8-1#TRMPBFCT-ND" H 6450 1800 60  0001 C CNN "Digikey"
	1    6450 1800
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_02x03_Odd_Even J1
U 1 1 597FD821
P 1200 2200
F 0 "J1" H 1250 1850 50  0000 C CNN
F 1 "PWR_HEADER" H 1250 1950 50  0000 C CNN
F 2 "Connectors_Molex:Molex_Microfit3_Header_02x03_Straight_43045-0628" H 1200 1000 50  0001 C CNN
F 3 "http://www.molex.com/pdm_docs/sd/430450628_sd.pdf" H 1200 1000 50  0001 C CNN
F 4 "43045-0628" H 1200 2200 60  0001 C CNN "MFN"
F 5 "WM10679-ND" H 1200 2200 60  0001 C CNN "Digikey"
	1    1200 2200
	-1   0    0    1   
$EndComp
$Comp
L power:+3V3 #PWR06
U 1 1 597FE163
P 6750 3200
F 0 "#PWR06" H 6750 3050 50  0001 C CNN
F 1 "+3V3" H 6765 3373 50  0000 C CNN
F 2 "" H 6750 3200 50  0001 C CNN
F 3 "" H 6750 3200 50  0001 C CNN
	1    6750 3200
	-1   0    0    1   
$EndComp
$Comp
L power:GND #PWR07
U 1 1 598009B2
P 2450 3700
F 0 "#PWR07" H 2450 3450 50  0001 C CNN
F 1 "GND" H 2450 3550 50  0000 C CNN
F 2 "" H 2450 3700 50  0000 C CNN
F 3 "" H 2450 3700 50  0000 C CNN
	1    2450 3700
	1    0    0    -1  
$EndComp
$Comp
L Device:R R2
U 1 1 598010E8
P 3450 3450
F 0 "R2" H 3300 3500 50  0000 L CNN
F 1 "47k(1%)" H 3100 3400 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 3380 3450 50  0001 C CNN
F 3 "http://www.yageo.com/documents/recent/PYu-RC_Group_51_RoHS_L_7.pdf" H 3450 3450 50  0001 C CNN
F 4 "RC0805FR-0747KL" H 3450 3450 60  0001 C CNN "MFN"
F 5 "311-47.0KCRCT-ND" H 3450 3450 60  0001 C CNN "Digikey"
	1    3450 3450
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR08
U 1 1 59801482
P 3450 3700
F 0 "#PWR08" H 3450 3450 50  0001 C CNN
F 1 "GND" H 3450 3550 50  0000 C CNN
F 2 "" H 3450 3700 50  0000 C CNN
F 3 "" H 3450 3700 50  0000 C CNN
	1    3450 3700
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR09
U 1 1 5980251F
P 7200 850
F 0 "#PWR09" H 7200 700 50  0001 C CNN
F 1 "+5V" H 7215 1023 50  0000 C CNN
F 2 "" H 7200 850 50  0001 C CNN
F 3 "" H 7200 850 50  0001 C CNN
	1    7200 850 
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C2
U 1 1 59802873
P 4050 2600
F 0 "C2" H 4142 2646 50  0000 L CNN
F 1 "47uf" H 4142 2555 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 4050 2600 50  0001 C CNN
F 3 "https://product.tdk.com/info/en/catalog/datasheets/mlcc_commercial_general_en.pdf" H 4050 2600 50  0001 C CNN
F 4 "C2012X5R1A476M125AC" H 4050 2600 60  0001 C CNN "MFN"
F 5 "445-8239-1-ND" H 4050 2600 60  0001 C CNN "Digikey"
	1    4050 2600
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR010
U 1 1 598029DE
P 4050 2900
F 0 "#PWR010" H 4050 2650 50  0001 C CNN
F 1 "GND" H 4050 2750 50  0000 C CNN
F 2 "" H 4050 2900 50  0000 C CNN
F 3 "" H 4050 2900 50  0000 C CNN
	1    4050 2900
	1    0    0    -1  
$EndComp
$Comp
L pimodem:TVS D1
U 1 1 598033DC
P 3700 2600
F 0 "D1" V 3654 2679 50  0000 L CNN
F 1 "TVS" V 3745 2679 50  0000 L CNN
F 2 "Diodes_SMD:D_SMB_Standard" H 3700 2600 50  0001 C CNN
F 3 "http://www.littelfuse.com/~/media/electronics/datasheets/tvs_diodes/littelfuse_tvs_diode_smbj_datasheet.pdf.pdf" H 3700 2600 50  0001 C CNN
F 4 "SMBJ5.0A" V 3700 2600 60  0001 C CNN "MFN"
F 5 "SMBJ5.0ALFCT-ND" V 3700 2600 60  0001 C CNN "Digikey"
	1    3700 2600
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR011
U 1 1 598034C3
P 3700 2900
F 0 "#PWR011" H 3700 2650 50  0001 C CNN
F 1 "GND" H 3700 2750 50  0000 C CNN
F 2 "" H 3700 2900 50  0000 C CNN
F 3 "" H 3700 2900 50  0000 C CNN
	1    3700 2900
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR012
U 1 1 5980D0EE
P 7150 2150
F 0 "#PWR012" H 7150 1900 50  0001 C CNN
F 1 "GND" H 7150 2000 50  0000 C CNN
F 2 "" H 7150 2150 50  0000 C CNN
F 3 "" H 7150 2150 50  0000 C CNN
	1    7150 2150
	0    -1   -1   0   
$EndComp
$Comp
L Device:R R3
U 1 1 5980D6DE
P 6750 2950
F 0 "R3" H 6600 3000 50  0000 L CNN
F 1 "10k(1%)" H 6400 2900 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 6680 2950 50  0001 C CNN
F 3 "" H 6750 2950 50  0001 C CNN
	1    6750 2950
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG02
U 1 1 5980EDA8
P 850 1800
F 0 "#FLG02" H 850 1875 50  0001 C CNN
F 1 "PWR_FLAG" H 850 1974 50  0000 C CNN
F 2 "" H 850 1800 50  0001 C CNN
F 3 "" H 850 1800 50  0001 C CNN
	1    850  1800
	1    0    0    -1  
$EndComp
$Comp
L pimodem:MAX3237 U4
U 1 1 59825CCA
P 1800 5400
F 0 "U4" H 1800 6765 50  0000 C CNN
F 1 "MAX3237" H 1800 6674 50  0000 C CNN
F 2 "Housings_SSOP:SSOP-28_5.3x10.2mm_Pitch0.65mm" H 1800 4050 50  0001 C CNN
F 3 "http://datasheets.maximintegrated.com/en/ds/MAX3222-MAX3241.pdf" H 1800 5400 60  0001 C CNN
F 4 "MAX3237EAI+" H 1800 5400 60  0001 C CNN "MFN"
F 5 "MAX3237EAI+-ND" H 1800 5400 60  0001 C CNN "Digikey"
	1    1800 5400
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C5
U 1 1 598260A6
P 850 4400
F 0 "C5" H 750 4500 50  0000 L CNN
F 1 "0.1uF" V 950 4300 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 850 4400 50  0001 C CNN
F 3 "http://www.kemet.com/Lists/ProductCatalog/Attachments/53/KEM_C1002_X7R_SMD.pdf" H 850 4400 50  0001 C CNN
F 4 "C0805C104K8RACTU" H 850 4400 60  0001 C CNN "MFN"
F 5 "399-7999-1-ND" H 850 4400 60  0001 C CNN "Digikey"
	1    850  4400
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C6
U 1 1 5982619E
P 850 4800
F 0 "C6" H 750 4900 50  0000 L CNN
F 1 "0.1uF" V 950 4700 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 850 4800 50  0001 C CNN
F 3 "" H 850 4800 50  0001 C CNN
	1    850  4800
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG03
U 1 1 598622E9
P 6850 850
F 0 "#FLG03" H 6850 925 50  0001 C CNN
F 1 "PWR_FLAG" H 6850 1024 50  0000 C CNN
F 2 "" H 6850 850 50  0001 C CNN
F 3 "" H 6850 850 50  0001 C CNN
	1    6850 850 
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C7
U 1 1 598633AF
P 3000 4400
F 0 "C7" V 2850 4400 50  0000 L CNN
F 1 "0.1uF" V 3100 4300 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 3000 4400 50  0001 C CNN
F 3 "" H 3000 4400 50  0001 C CNN
	1    3000 4400
	0    1    1    0   
$EndComp
$Comp
L Device:C_Small C8
U 1 1 59863435
P 2650 4600
F 0 "C8" V 2500 4650 50  0000 L CNN
F 1 "0.1uF" V 2750 4500 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 2650 4600 50  0001 C CNN
F 3 "" H 2650 4600 50  0001 C CNN
	1    2650 4600
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR014
U 1 1 59863ADA
P 2850 4600
F 0 "#PWR014" H 2850 4350 50  0001 C CNN
F 1 "GND" V 2855 4472 50  0000 R CNN
F 2 "" H 2850 4600 50  0001 C CNN
F 3 "" H 2850 4600 50  0001 C CNN
	1    2850 4600
	0    -1   -1   0   
$EndComp
$Comp
L Connector_Generic:Conn_02x20_Odd_Even J2
U 1 1 59867196
P 9600 2000
F 0 "J2" H 9650 3150 50  0000 C CNN
F 1 "PI_HEADER" H 9650 3050 50  0000 C CNN
F 2 "Socket_Strips:Socket_Strip_Straight_2x20_Pitch2.54mm" H 9600 1050 50  0001 C CNN
F 3 "http://sullinscorp.com/catalogs/146_PAGE119_.100_SFH11_SERIES_FEMALE_HDR_ST_RA.pdf" H 9600 1050 50  0001 C CNN
F 4 "SFH11-PBPC-D20-ST-BK" H 9600 2000 60  0001 C CNN "MFN"
F 5 "S9200-ND" H 9600 2000 60  0001 C CNN "Digikey"
	1    9600 2000
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR015
U 1 1 59867677
P 10200 1100
F 0 "#PWR015" H 10200 950 50  0001 C CNN
F 1 "+5V" V 10215 1228 50  0000 L CNN
F 2 "" H 10200 1100 50  0001 C CNN
F 3 "" H 10200 1100 50  0001 C CNN
	1    10200 1100
	0    1    1    0   
$EndComp
$Comp
L power:+3V3 #PWR016
U 1 1 59867A2F
P 9000 1100
F 0 "#PWR016" H 9000 950 50  0001 C CNN
F 1 "+3V3" V 9015 1228 50  0000 L CNN
F 2 "" H 9000 1100 50  0001 C CNN
F 3 "" H 9000 1100 50  0001 C CNN
	1    9000 1100
	0    -1   -1   0   
$EndComp
$Comp
L power:GND #PWR017
U 1 1 598680D9
P 9000 1500
F 0 "#PWR017" H 9000 1250 50  0001 C CNN
F 1 "GND" V 9005 1372 50  0000 R CNN
F 2 "" H 9000 1500 50  0001 C CNN
F 3 "" H 9000 1500 50  0001 C CNN
	1    9000 1500
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR018
U 1 1 59868188
P 10200 1300
F 0 "#PWR018" H 10200 1050 50  0001 C CNN
F 1 "GND" V 10205 1172 50  0000 R CNN
F 2 "" H 10200 1300 50  0001 C CNN
F 3 "" H 10200 1300 50  0001 C CNN
	1    10200 1300
	0    -1   -1   0   
$EndComp
$Comp
L power:+3V3 #PWR019
U 1 1 59869B5F
P 2650 6500
F 0 "#PWR019" H 2650 6350 50  0001 C CNN
F 1 "+3V3" V 2665 6628 50  0000 L CNN
F 2 "" H 2650 6500 50  0001 C CNN
F 3 "" H 2650 6500 50  0001 C CNN
	1    2650 6500
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR020
U 1 1 59869C78
P 1050 6300
F 0 "#PWR020" H 1050 6050 50  0001 C CNN
F 1 "GND" V 1055 6172 50  0000 R CNN
F 2 "" H 1050 6300 50  0001 C CNN
F 3 "" H 1050 6300 50  0001 C CNN
	1    1050 6300
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR021
U 1 1 59869CB6
P 2600 6300
F 0 "#PWR021" H 2600 6050 50  0001 C CNN
F 1 "GND" V 2605 6172 50  0000 R CNN
F 2 "" H 2600 6300 50  0001 C CNN
F 3 "" H 2600 6300 50  0001 C CNN
	1    2600 6300
	0    -1   -1   0   
$EndComp
$Comp
L power:PWR_FLAG #FLG04
U 1 1 5986B37B
P 9100 900
F 0 "#FLG04" H 9100 975 50  0001 C CNN
F 1 "PWR_FLAG" H 9100 1074 50  0000 C CNN
F 2 "" H 9100 900 50  0001 C CNN
F 3 "" H 9100 900 50  0001 C CNN
	1    9100 900 
	1    0    0    -1  
$EndComp
NoConn ~ 1200 5700
Text GLabel 10550 1400 2    60   Output ~ 0
TXD
Text GLabel 10250 1500 2    60   Input ~ 0
RXD
Text GLabel 1050 5900 0    60   Output ~ 0
RXD
Text GLabel 1000 5100 0    60   Input ~ 0
TXD
$Comp
L Connector_Generic:Conn_02x05_Odd_Even J3
U 1 1 5986E856
P 3750 5500
F 0 "J3" H 3800 5100 50  0000 C CNN
F 1 "SERIAL" H 3800 5200 50  0000 C CNN
F 2 "Pin_Headers:Pin_Header_Straight_2x05_Pitch2.54mm" H 3750 4300 50  0001 C CNN
F 3 "http://portal.fciconnect.com/Comergent//fci/drawing/67996.pdf" H 3750 4300 50  0001 C CNN
F 4 "67997-410HLF" H 3750 5500 60  0001 C CNN "MFN"
F 5 "609-3243-ND" H 3750 5500 60  0001 C CNN "Digikey"
	1    3750 5500
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR022
U 1 1 598718CF
P 1400 1550
F 0 "#PWR022" H 1400 1400 50  0001 C CNN
F 1 "+5V" H 1415 1723 50  0000 C CNN
F 2 "" H 1400 1550 50  0001 C CNN
F 3 "" H 1400 1550 50  0001 C CNN
	1    1400 1550
	1    0    0    -1  
$EndComp
$Comp
L Device:R R4
U 1 1 59872F74
P 5850 1950
F 0 "R4" H 5920 1996 50  0000 L CNN
F 1 "10k(1%)" H 5500 2000 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 5780 1950 50  0001 C CNN
F 3 "" H 5850 1950 50  0001 C CNN
	1    5850 1950
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x07 J4
U 1 1 59873DFC
P 10050 3600
F 0 "J4" H 10128 3641 50  0000 L CNN
F 1 "AUDIO_HEADER" H 10128 3550 50  0000 L CNN
F 2 "adafruit:MAX98357_I2S_DAC" H 10050 3600 50  0001 C CNN
F 3 "http://www.te.com/commerce/DocumentDelivery/DDEController?Action=srchrtrv&DocNm=644456&DocType=Customer+Drawing&DocLang=English" H 10050 3600 50  0001 C CNN
F 4 "3-644456-7" H 10050 3600 60  0001 C CNN "MFN"
F 5 "A31117-ND" H 10050 3600 60  0001 C CNN "Digikey"
	1    10050 3600
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR023
U 1 1 598741D1
P 9550 3900
F 0 "#PWR023" H 9550 3750 50  0001 C CNN
F 1 "+5V" V 9565 4028 50  0000 L CNN
F 2 "" H 9550 3900 50  0001 C CNN
F 3 "" H 9550 3900 50  0001 C CNN
	1    9550 3900
	0    -1   -1   0   
$EndComp
$Comp
L power:GND #PWR024
U 1 1 5987420E
P 9700 3800
F 0 "#PWR024" H 9700 3550 50  0001 C CNN
F 1 "GND" V 9705 3672 50  0000 R CNN
F 2 "" H 9700 3800 50  0001 C CNN
F 3 "" H 9700 3800 50  0001 C CNN
	1    9700 3800
	0    1    1    0   
$EndComp
Text GLabel 9550 3500 0    60   Input ~ 0
DIN
Text GLabel 10550 1600 2    60   Output ~ 0
BCLK
Text GLabel 9000 2800 0    60   Output ~ 0
LRCLK
Text GLabel 10200 3000 2    60   Output ~ 0
DIN
NoConn ~ 9900 2900
Text GLabel 9700 3300 0    60   Input ~ 0
LRCLK
Text GLabel 9250 3400 0    60   Input ~ 0
BCLK
NoConn ~ 9850 3700
Text GLabel 10200 2800 2    60   Input ~ 0
CTS
Text GLabel 8900 1600 0    60   Output ~ 0
RTS
Text GLabel 750  5200 0    60   Input ~ 0
RTS
Text GLabel 750  6000 0    60   Output ~ 0
CTS
Text GLabel 1050 6100 0    60   Output ~ 0
DTR
Text GLabel 1000 5300 0    60   Input ~ 0
DCD
Text GLabel 750  5400 0    60   Input ~ 0
DSR
Text GLabel 900  5500 0    60   Input ~ 0
RI
NoConn ~ 4050 5700
$Comp
L power:GND #PWR025
U 1 1 598805E5
P 3450 5500
F 0 "#PWR025" H 3450 5250 50  0001 C CNN
F 1 "GND" V 3455 5372 50  0000 R CNN
F 2 "" H 3450 5500 50  0001 C CNN
F 3 "" H 3450 5500 50  0001 C CNN
	1    3450 5500
	0    1    1    0   
$EndComp
Text GLabel 10200 2600 2    60   Input ~ 0
DTR
Text GLabel 9050 2500 0    60   Output ~ 0
DCD
Text GLabel 8800 2600 0    60   Output ~ 0
DSR
Text GLabel 9300 2700 0    60   Output ~ 0
RI
NoConn ~ 9400 1200
NoConn ~ 9400 1300
NoConn ~ 9400 1700
NoConn ~ 9400 1800
NoConn ~ 9400 2100
NoConn ~ 9900 1800
NoConn ~ 9900 1900
NoConn ~ 9900 2100
NoConn ~ 9900 2200
NoConn ~ 9400 2900
Text GLabel 8700 1400 0    60   Output ~ 0
~KILL~
Text GLabel 9000 2200 0    60   Input ~ 0
~INT~
$Comp
L Transistor_Array:ULN2803A U3
U 1 1 5988B5B5
P 7600 4450
F 0 "U3" H 7600 5017 50  0000 C CNN
F 1 "ULN2803A" H 7600 4926 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-18W_7.5x11.6mm_Pitch1.27mm" H 7650 3800 50  0001 L CNN
F 3 "http://www.ti.com/lit/ds/symlink/uln2803a.pdf" H 7700 4350 50  0001 C CNN
F 4 "ULN2803ADWR" H 7600 4450 60  0001 C CNN "MFN"
F 5 "296-15777-1-ND" H 7600 4450 60  0001 C CNN "Digikey"
	1    7600 4450
	1    0    0    -1  
$EndComp
NoConn ~ 8000 4150
$Comp
L power:GND #PWR026
U 1 1 5987C891
P 7600 5250
F 0 "#PWR026" H 7600 5000 50  0001 C CNN
F 1 "GND" H 7605 5077 50  0000 C CNN
F 2 "" H 7600 5250 50  0001 C CNN
F 3 "" H 7600 5250 50  0001 C CNN
	1    7600 5250
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_02x05_Odd_Even J5
U 1 1 5988A04C
P 8950 4700
F 0 "J5" H 9000 5000 50  0000 C CNN
F 1 "LED_PANEL" H 9000 4400 50  0000 C CNN
F 2 "Pin_Headers:Pin_Header_Straight_2x05_Pitch2.54mm" H 8950 3500 50  0001 C CNN
F 3 "http://portal.fciconnect.com/Comergent//fci/drawing/67996.pdf" H 8950 3500 50  0001 C CNN
F 4 "67997-410HLF" H 8950 4700 60  0001 C CNN "MFN"
F 5 "609-3243-ND" H 8950 4700 60  0001 C CNN "Digikey"
	1    8950 4700
	0    -1   -1   0   
$EndComp
$Comp
L power:GND #PWR027
U 1 1 5988A7B2
P 8750 5200
F 0 "#PWR027" H 8750 4950 50  0001 C CNN
F 1 "GND" V 8755 5072 50  0000 R CNN
F 2 "" H 8750 5200 50  0001 C CNN
F 3 "" H 8750 5200 50  0001 C CNN
	1    8750 5200
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR028
U 1 1 5988A7F7
P 9150 4250
F 0 "#PWR028" H 9150 4100 50  0001 C CNN
F 1 "+5V" V 9165 4378 50  0000 L CNN
F 2 "" H 9150 4250 50  0001 C CNN
F 3 "" H 9150 4250 50  0001 C CNN
	1    9150 4250
	1    0    0    -1  
$EndComp
Text GLabel 7050 4250 0    60   UnSpc ~ 0
LED1
Text GLabel 6750 4350 0    60   UnSpc ~ 0
LED2
Text GLabel 7050 4450 0    60   UnSpc ~ 0
LED3
Text GLabel 6750 4550 0    60   UnSpc ~ 0
LED4
Text GLabel 7050 4650 0    60   UnSpc ~ 0
LED5
Text GLabel 6750 4750 0    60   UnSpc ~ 0
LED6
Text GLabel 6500 6200 3    60   UnSpc ~ 0
LED7
Text GLabel 7050 6200 3    60   UnSpc ~ 0
LED8
NoConn ~ 10850 4500
NoConn ~ 10850 5000
NoConn ~ 10850 5500
NoConn ~ 10850 6000
$Comp
L power:GND #PWR029
U 1 1 59891422
P 9850 6250
F 0 "#PWR029" H 9850 6000 50  0001 C CNN
F 1 "GND" H 9855 6077 50  0000 C CNN
F 2 "" H 9850 6250 50  0001 C CNN
F 3 "" H 9850 6250 50  0001 C CNN
	1    9850 6250
	1    0    0    -1  
$EndComp
Text GLabel 8000 5400 0    60   UnSpc ~ 0
TXD
Text GLabel 8000 5550 0    60   UnSpc ~ 0
RXD
Text GLabel 8350 5400 2    60   UnSpc ~ 0
LED7
Text GLabel 8350 5550 2    60   UnSpc ~ 0
LED8
Text GLabel 8000 5700 0    60   UnSpc ~ 0
DSR
Text GLabel 8350 5700 2    60   UnSpc ~ 0
LED1
Text GLabel 9050 2000 0    60   UnSpc ~ 0
LED2
Text GLabel 8350 5850 2    60   UnSpc ~ 0
LED3
Text GLabel 8350 6000 2    60   UnSpc ~ 0
LED4
Text GLabel 8350 6150 2    60   UnSpc ~ 0
LED5
Text GLabel 8350 6300 2    60   UnSpc ~ 0
LED6
NoConn ~ 8350 6000
NoConn ~ 8350 6150
NoConn ~ 8350 6300
$Comp
L pimodem:74LS04 U5
U 1 1 59888443
P 6500 5550
F 0 "U5" V 6550 5400 50  0000 R CNN
F 1 "74LS04" V 6450 5400 50  0000 R CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 6500 5550 50  0001 C CNN
F 3 "http://www.ti.com/lit/ds/symlink/sn74ls04.pdf" H 6500 5550 50  0001 C CNN
F 4 "SN74LS04DR" V 6500 5550 60  0001 C CNN "MFN"
F 5 "296-14875-1-ND" V 6500 5550 60  0001 C CNN "Digikey"
	1    6500 5550
	0    -1   -1   0   
$EndComp
$Comp
L pimodem:74LS04 U5
U 2 1 5988894F
P 7050 5550
F 0 "U5" V 7096 5372 50  0000 R CNN
F 1 "74LS04" V 7005 5372 50  0000 R CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 7050 5550 50  0001 C CNN
F 3 "" H 7050 5550 50  0001 C CNN
	2    7050 5550
	0    -1   -1   0   
$EndComp
$Comp
L pimodem:74LS04 U5
U 3 1 59888F40
P 10400 4500
F 0 "U5" H 10400 4815 50  0000 C CNN
F 1 "74LS04" H 10400 4724 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 10400 4500 50  0001 C CNN
F 3 "" H 10400 4500 50  0001 C CNN
	3    10400 4500
	1    0    0    -1  
$EndComp
$Comp
L pimodem:74LS04 U5
U 4 1 5988939F
P 10400 5000
F 0 "U5" H 10400 5315 50  0000 C CNN
F 1 "74LS04" H 10400 5224 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 10400 5000 50  0001 C CNN
F 3 "" H 10400 5000 50  0001 C CNN
	4    10400 5000
	1    0    0    -1  
$EndComp
$Comp
L pimodem:74LS04 U5
U 5 1 59889407
P 10400 5500
F 0 "U5" H 10400 5815 50  0000 C CNN
F 1 "74LS04" H 10400 5724 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 10400 5500 50  0001 C CNN
F 3 "" H 10400 5500 50  0001 C CNN
	5    10400 5500
	1    0    0    -1  
$EndComp
$Comp
L pimodem:74LS04 U5
U 6 1 5988946B
P 10400 6000
F 0 "U5" H 10400 6315 50  0000 C CNN
F 1 "74LS04" H 10400 6224 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-14_3.9x8.7mm_Pitch1.27mm" H 10400 6000 50  0001 C CNN
F 3 "" H 10400 6000 50  0001 C CNN
	6    10400 6000
	1    0    0    -1  
$EndComp
$Comp
L Memory_EEPROM:24LC32 U6
U 1 1 598902DA
P 3750 7050
F 0 "U6" H 3600 6800 50  0000 C CNN
F 1 "24LC32" H 3900 6800 50  0000 C CNN
F 2 "Housings_SOIC:SOIC-8_3.9x4.9mm_Pitch1.27mm" H 3800 6800 50  0001 L CNN
F 3 "http://www.onsemi.com/pub/Collateral/CAT24C32-D.PDF" H 3750 6950 50  0001 C CNN
F 4 "CAT24C32WI-G" H 3750 7050 60  0001 C CNN "MFN"
F 5 "CAT24C32WI-GOS-ND" H 3750 7050 60  0001 C CNN "Digikey"
	1    3750 7050
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR030
U 1 1 5989095A
P 3750 7450
F 0 "#PWR030" H 3750 7200 50  0001 C CNN
F 1 "GND" V 3755 7322 50  0000 R CNN
F 2 "" H 3750 7450 50  0001 C CNN
F 3 "" H 3750 7450 50  0001 C CNN
	1    3750 7450
	1    0    0    -1  
$EndComp
$Comp
L Device:C_Small C9
U 1 1 59891D71
P 3550 6550
F 0 "C9" V 3400 6500 50  0000 L CNN
F 1 "0.1uF" V 3650 6450 50  0000 L CNN
F 2 "Capacitors_SMD:C_0805" H 3550 6550 50  0001 C CNN
F 3 "" H 3550 6550 50  0001 C CNN
	1    3550 6550
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR031
U 1 1 598922D6
P 3350 6550
F 0 "#PWR031" H 3350 6300 50  0001 C CNN
F 1 "GND" V 3355 6422 50  0000 R CNN
F 2 "" H 3350 6550 50  0001 C CNN
F 3 "" H 3350 6550 50  0001 C CNN
	1    3350 6550
	0    1    1    0   
$EndComp
$Comp
L power:+3V3 #PWR032
U 1 1 598928FC
P 3750 6500
F 0 "#PWR032" H 3750 6350 50  0001 C CNN
F 1 "+3V3" V 3765 6628 50  0000 L CNN
F 2 "" H 3750 6500 50  0001 C CNN
F 3 "" H 3750 6500 50  0001 C CNN
	1    3750 6500
	1    0    0    -1  
$EndComp
$Comp
L Device:R R5
U 1 1 59892F4A
P 5450 6700
F 0 "R5" H 5520 6746 50  0000 L CNN
F 1 "1k(1%)" H 5520 6655 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 5380 6700 50  0001 C CNN
F 3 "http://www.yageo.com/documents/recent/PYu-RC_Group_51_RoHS_L_7.pdf" H 5450 6700 50  0001 C CNN
F 4 "RC0805FR-071KL" H 5450 6700 60  0001 C CNN "MFN"
F 5 "311-1.00KCRCT-ND" H 5450 6700 60  0001 C CNN "Digikey"
	1    5450 6700
	1    0    0    -1  
$EndComp
$Comp
L Device:R R6
U 1 1 59893238
P 4450 6700
F 0 "R6" H 4520 6746 50  0000 L CNN
F 1 "3.9k(1%)" H 4520 6655 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 4380 6700 50  0001 C CNN
F 3 "https://www.seielect.com/Catalog/SEI-RMCF_RMCP.pdf" H 4450 6700 50  0001 C CNN
F 4 "RMCF0805FT3K90" H 4450 6700 60  0001 C CNN "MFN"
F 5 "RMCF0805FT3K90CT-ND" H 4450 6700 60  0001 C CNN "Digikey"
	1    4450 6700
	1    0    0    -1  
$EndComp
$Comp
L Device:R R7
U 1 1 598932AC
P 4950 6700
F 0 "R7" H 5020 6746 50  0000 L CNN
F 1 "3.9k(1%)" H 5020 6655 50  0000 L CNN
F 2 "Resistors_SMD:R_0805" V 4880 6700 50  0001 C CNN
F 3 "" H 4950 6700 50  0001 C CNN
	1    4950 6700
	1    0    0    -1  
$EndComp
Text GLabel 5950 7050 2    60   UnSpc ~ 0
ID_SC
Text GLabel 5650 6950 2    60   UnSpc ~ 0
ID_SD
$Comp
L Connector_Generic:Conn_01x02 J6
U 1 1 59896620
P 5650 7300
F 0 "J6" H 5750 7300 50  0000 L CNN
F 1 "WP" H 5750 7200 50  0000 L CNN
F 2 "Pin_Headers:Pin_Header_Straight_1x02_Pitch2.54mm" H 5650 7300 50  0001 C CNN
F 3 "https://cdn.harwin.com/pdfs/M20-999.pdf" H 5650 7300 50  0001 C CNN
F 4 "M20-9990246" H 5650 7300 60  0001 C CNN "MFN"
F 5 "952-2262-ND" H 5650 7300 60  0001 C CNN "Digikey"
	1    5650 7300
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR033
U 1 1 59897078
P 5450 7450
F 0 "#PWR033" H 5450 7200 50  0001 C CNN
F 1 "GND" V 5455 7322 50  0000 R CNN
F 2 "" H 5450 7450 50  0001 C CNN
F 3 "" H 5450 7450 50  0001 C CNN
	1    5450 7450
	1    0    0    -1  
$EndComp
Text GLabel 8800 2400 0    60   UnSpc ~ 0
ID_SD
Text GLabel 10450 2400 2    60   UnSpc ~ 0
ID_SC
$Comp
L Device:Q_PMOS_GSD Q1
U 1 1 5989C5E5
P 3050 2150
F 0 "Q1" V 3000 2300 50  0000 C CNN
F 1 "DMG2305UX" V 3250 2450 50  0000 C CNN
F 2 "TO_SOT_Packages_SMD:SOT-23" H 3250 2250 50  0001 C CNN
F 3 "https://www.diodes.com/assets/Datasheets/DMG2305UX.pdf" H 3050 2150 50  0001 C CNN
F 4 "DMG2305UX-13" V 3050 2150 60  0001 C CNN "MFN"
F 5 "DMG2305UX-13DICT-ND" V 3050 2150 60  0001 C CNN "Digikey"
	1    3050 2150
	0    -1   -1   0   
$EndComp
$Comp
L Device:Q_NMOS_GSD Q3
U 1 1 5989CF2E
P 6400 1100
F 0 "Q3" V 6743 1100 50  0000 C CNN
F 1 "DMG3414U" V 6652 1100 50  0000 C CNN
F 2 "TO_SOT_Packages_SMD:SOT-23" H 6600 1200 50  0001 C CNN
F 3 "https://www.diodes.com/assets/Datasheets/ds31739.pdf" H 6400 1100 50  0001 C CNN
F 4 "DMG3414U-7" V 6400 1100 60  0001 C CNN "MFN"
F 5 "DMG3414U-7DICT-ND" V 6400 1100 60  0001 C CNN "Digikey"
	1    6400 1100
	0    -1   -1   0   
$EndComp
Text GLabel 8000 5850 0    60   UnSpc ~ 0
DCD
$Comp
L power:GND #PWR013
U 1 1 598DB0E8
P 3250 4400
F 0 "#PWR013" H 3250 4150 50  0001 C CNN
F 1 "GND" V 3255 4272 50  0000 R CNN
F 2 "" H 3250 4400 50  0001 C CNN
F 3 "" H 3250 4400 50  0001 C CNN
	1    3250 4400
	0    -1   -1   0   
$EndComp
$Comp
L Device:R R8
U 1 1 598DB50F
P 8050 3600
F 0 "R8" V 7950 3600 50  0000 C CNN
F 1 "100k" V 7850 3600 50  0000 C CNN
F 2 "Resistors_SMD:R_0805" V 7980 3600 50  0001 C CNN
F 3 "http://www.yageo.com/documents/recent/PYu-RC_Group_51_RoHS_L_7.pdf" H 8050 3600 50  0001 C CNN
F 4 "RC0805FR-07100KL" V 8050 3600 60  0001 C CNN "MFN"
F 5 "311-100KCRCT-ND" V 8050 3600 60  0001 C CNN "Digikey"
	1    8050 3600
	0    1    1    0   
$EndComp
$Comp
L Device:R R9
U 1 1 598DB6CD
P 8050 3700
F 0 "R9" V 8150 3700 50  0000 C CNN
F 1 "100k" V 8250 3700 50  0000 C CNN
F 2 "Resistors_SMD:R_0805" V 7980 3700 50  0001 C CNN
F 3 "" H 8050 3700 50  0001 C CNN
	1    8050 3700
	0    1    1    0   
$EndComp
$Comp
L Connector_Generic:Conn_02x04_Odd_Even J7
U 1 1 598E0AB7
P 8600 3700
F 0 "J7" H 8650 3300 50  0000 C CNN
F 1 "GAIN" H 8650 3400 50  0000 C CNN
F 2 "Pin_Headers:Pin_Header_Straight_2x04_Pitch2.54mm" H 8600 2500 50  0001 C CNN
F 3 "http://portal.fciconnect.com/Comergent//fci/drawing/67996.pdf" H 8600 2500 50  0001 C CNN
F 4 "67996-408HLF" H 8600 3700 60  0001 C CNN "MFN"
F 5 "609-3391-ND" H 8600 3700 60  0001 C CNN "Digikey"
	1    8600 3700
	-1   0    0    1   
$EndComp
$Comp
L power:+5V #PWR034
U 1 1 598E1BF3
P 7800 3800
F 0 "#PWR034" H 7800 3650 50  0001 C CNN
F 1 "+5V" V 7815 3928 50  0000 L CNN
F 2 "" H 7800 3800 50  0001 C CNN
F 3 "" H 7800 3800 50  0001 C CNN
	1    7800 3800
	0    -1   -1   0   
$EndComp
$Comp
L power:GND #PWR035
U 1 1 598E1C52
P 7800 3500
F 0 "#PWR035" H 7800 3250 50  0001 C CNN
F 1 "GND" V 7805 3372 50  0000 R CNN
F 2 "" H 7800 3500 50  0001 C CNN
F 3 "" H 7800 3500 50  0001 C CNN
	1    7800 3500
	0    1    1    0   
$EndComp
NoConn ~ 9400 1100
NoConn ~ 9900 2300
Wire Wire Line
	850  2200 900  2200
Wire Wire Line
	850  1800 850  2100
Wire Wire Line
	900  2100 850  2100
Connection ~ 850  2200
Wire Wire Line
	4900 2950 4900 3050
Wire Wire Line
	5100 3100 5100 2950
Wire Wire Line
	5300 2950 5300 3000
Wire Wire Line
	5300 3400 5300 3200
Wire Wire Line
	5100 3300 5100 3400
Wire Wire Line
	5750 2600 6400 2600
Wire Wire Line
	900  2300 850  2300
Connection ~ 850  2300
Wire Wire Line
	1950 2300 2100 2300
Wire Wire Line
	2100 2300 2100 2450
Wire Wire Line
	2100 2750 2100 2650
Wire Wire Line
	1400 2300 1650 2300
Wire Wire Line
	2650 2050 2850 2050
Wire Wire Line
	2650 2300 2650 2050
Connection ~ 2100 2300
Wire Wire Line
	2450 2650 2450 2300
Connection ~ 2450 2300
Wire Wire Line
	2450 3050 2450 3200
Wire Wire Line
	2450 3200 2950 3200
Wire Wire Line
	2750 2850 2950 2850
Connection ~ 2450 3200
Wire Wire Line
	2450 3600 2450 3700
Wire Wire Line
	3450 3050 3450 3200
Wire Wire Line
	2950 3200 2950 2850
Connection ~ 2950 2850
Wire Wire Line
	3050 2350 3050 3200
Wire Wire Line
	3050 3200 3450 3200
Connection ~ 3450 3200
Wire Wire Line
	3450 3700 3450 3600
Wire Wire Line
	3250 2050 3450 2050
Wire Wire Line
	3450 2050 3450 2300
Wire Wire Line
	3450 2300 3700 2300
Wire Wire Line
	4050 2150 4050 2300
Connection ~ 3450 2300
Connection ~ 4050 2300
Wire Wire Line
	4050 2900 4050 2700
Wire Notes Line
	2350 1800 2350 4050
Wire Notes Line
	2350 4050 3550 4050
Wire Notes Line
	3550 4050 3550 1800
Wire Notes Line
	3550 1800 2350 1800
Wire Wire Line
	3700 2300 3700 2450
Connection ~ 3700 2300
Wire Wire Line
	3700 2900 3700 2750
Wire Wire Line
	4500 2500 4300 2500
Wire Wire Line
	4300 1750 4300 2500
Wire Wire Line
	1650 1750 4300 1750
Wire Wire Line
	1650 1750 1650 2200
Wire Wire Line
	1650 2200 1400 2200
Wire Wire Line
	6400 1400 6400 1300
Wire Wire Line
	5850 2300 5750 2300
Wire Wire Line
	5850 2100 5850 2150
Wire Wire Line
	5850 2150 5950 2150
Wire Wire Line
	4600 1750 5600 1750
Wire Wire Line
	4600 1750 4600 2150
Wire Wire Line
	4600 2150 4050 2150
Wire Wire Line
	6200 1000 5600 1000
Wire Wire Line
	5600 1000 5600 1750
Connection ~ 5600 1750
Wire Wire Line
	6600 1000 6850 1000
Wire Wire Line
	7200 850  7200 1000
Wire Wire Line
	7200 1750 6900 1750
Connection ~ 7200 1000
Wire Wire Line
	6900 2150 7150 2150
Wire Wire Line
	6750 2450 6750 2800
Connection ~ 6750 2450
Wire Wire Line
	6750 3200 6750 3100
Wire Wire Line
	5750 2450 6750 2450
Wire Wire Line
	850  4300 1200 4300
Wire Wire Line
	850  4500 1200 4500
Wire Wire Line
	850  4700 1200 4700
Wire Wire Line
	850  4900 1200 4900
Wire Wire Line
	6850 850  6850 1000
Connection ~ 6850 1000
Connection ~ 850  2100
Wire Wire Line
	2400 4600 2550 4600
Wire Wire Line
	2850 4600 2750 4600
Wire Wire Line
	3250 4400 3100 4400
Wire Wire Line
	2900 4400 2400 4400
Wire Wire Line
	9900 1100 10050 1100
Wire Wire Line
	9900 1200 10050 1200
Wire Wire Line
	10050 1200 10050 1100
Connection ~ 10050 1100
Wire Wire Line
	9100 1900 9400 1900
Wire Wire Line
	9100 900  9100 1100
Connection ~ 9100 1100
Wire Wire Line
	9900 1300 10150 1300
Wire Wire Line
	2400 6500 2650 6500
Wire Wire Line
	1050 6300 1200 6300
Wire Wire Line
	2600 6300 2400 6300
Wire Wire Line
	9900 1700 10150 1700
Wire Wire Line
	10150 1300 10150 1700
Connection ~ 10150 1300
Wire Wire Line
	10150 2000 9900 2000
Connection ~ 10150 1700
Wire Wire Line
	10150 2500 9900 2500
Connection ~ 10150 2000
Wire Wire Line
	10150 2700 9900 2700
Connection ~ 10150 2500
Wire Wire Line
	9000 1500 9150 1500
Wire Wire Line
	9400 2300 9150 2300
Wire Wire Line
	9150 1500 9150 2300
Connection ~ 9150 1500
Wire Wire Line
	9150 3000 9400 3000
Connection ~ 9150 2300
Wire Wire Line
	9900 1400 10550 1400
Wire Wire Line
	10250 1500 9900 1500
Wire Wire Line
	1000 5100 1200 5100
Wire Wire Line
	1050 5900 1200 5900
Wire Wire Line
	1400 2100 1400 1550
Connection ~ 5850 2150
Wire Wire Line
	5850 1800 5850 1750
Connection ~ 5850 1750
Wire Wire Line
	9550 3900 9850 3900
Wire Wire Line
	9850 3800 9700 3800
Wire Wire Line
	9550 3500 9850 3500
Wire Wire Line
	10200 3000 9900 3000
Wire Wire Line
	9900 1600 10550 1600
Wire Wire Line
	9000 2800 9400 2800
Wire Wire Line
	9700 3300 9850 3300
Wire Wire Line
	9250 3400 9850 3400
Wire Wire Line
	9900 2800 10200 2800
Wire Wire Line
	8900 1600 9400 1600
Wire Wire Line
	750  6000 1200 6000
Wire Wire Line
	750  5200 1200 5200
Wire Wire Line
	1050 6100 1200 6100
Wire Wire Line
	1000 5300 1200 5300
Wire Wire Line
	1200 5400 750  5400
Wire Wire Line
	900  5500 1200 5500
Wire Wire Line
	3450 5500 3550 5500
Wire Wire Line
	2400 5300 3550 5300
Wire Wire Line
	2400 5100 4050 5100
Wire Wire Line
	4050 5100 4050 5300
Wire Wire Line
	2400 5900 3150 5900
Wire Wire Line
	3150 5900 3150 5400
Wire Wire Line
	3150 5400 3550 5400
Wire Wire Line
	2400 6100 4350 6100
Wire Wire Line
	4350 6100 4350 5400
Wire Wire Line
	4350 5400 4050 5400
Wire Wire Line
	2400 5400 3050 5400
Wire Wire Line
	3050 5400 3050 5000
Wire Wire Line
	3050 5000 4150 5000
Wire Wire Line
	4150 5000 4150 5500
Wire Wire Line
	4150 5500 4050 5500
Wire Wire Line
	2400 6000 3400 6000
Wire Wire Line
	3400 6000 3400 5600
Wire Wire Line
	3400 5600 3550 5600
Wire Wire Line
	2400 5200 4250 5200
Wire Wire Line
	4250 5200 4250 5600
Wire Wire Line
	4250 5600 4050 5600
Wire Wire Line
	2400 5500 3000 5500
Wire Wire Line
	3000 5500 3000 5700
Wire Wire Line
	3000 5700 3550 5700
Wire Wire Line
	9900 2600 10200 2600
Wire Wire Line
	9050 2500 9400 2500
Wire Wire Line
	9300 2700 9400 2700
Wire Wire Line
	8800 2600 9400 2600
Wire Wire Line
	8700 1400 9400 1400
Wire Wire Line
	7600 5250 7600 5150
Wire Wire Line
	8000 4250 8750 4250
Wire Wire Line
	8750 4250 8750 4400
Wire Wire Line
	8000 4350 8600 4350
Wire Wire Line
	8600 4350 8600 5050
Wire Wire Line
	8600 5050 8850 5050
Wire Wire Line
	8850 5050 8850 4900
Wire Wire Line
	8000 4450 8500 4450
Wire Wire Line
	8500 4450 8500 4150
Wire Wire Line
	8500 4150 8850 4150
Wire Wire Line
	8850 4150 8850 4400
Wire Wire Line
	8000 4550 8500 4550
Wire Wire Line
	8500 4550 8500 4950
Wire Wire Line
	8500 4950 8950 4950
Wire Wire Line
	8950 4950 8950 4900
Wire Wire Line
	8000 4650 8400 4650
Wire Wire Line
	8400 4650 8400 4050
Wire Wire Line
	8400 4050 8950 4050
Wire Wire Line
	8950 4050 8950 4400
Wire Wire Line
	8000 4750 8450 4750
Wire Wire Line
	8450 4750 8450 5150
Wire Wire Line
	8450 5150 9050 5150
Wire Wire Line
	9050 5150 9050 4900
Wire Wire Line
	8000 4850 8350 4850
Wire Wire Line
	8350 4850 8350 4100
Wire Wire Line
	8350 4100 9050 4100
Wire Wire Line
	9050 4100 9050 4400
Wire Wire Line
	8000 4950 8350 4950
Wire Wire Line
	8350 4950 8350 5100
Wire Wire Line
	8350 5100 9150 5100
Wire Wire Line
	9150 5100 9150 4900
Wire Wire Line
	8750 4900 8750 5200
Wire Wire Line
	9150 4400 9150 4250
Wire Wire Line
	7050 4250 7200 4250
Wire Wire Line
	7200 4350 6750 4350
Wire Wire Line
	7050 4450 7200 4450
Wire Wire Line
	6750 4550 7200 4550
Wire Wire Line
	7200 4650 7050 4650
Wire Wire Line
	6750 4750 7200 4750
Wire Wire Line
	9850 4500 9850 5000
Wire Wire Line
	9850 4500 9950 4500
Wire Wire Line
	9950 5000 9850 5000
Connection ~ 9850 5000
Wire Wire Line
	9950 5500 9850 5500
Connection ~ 9850 5500
Wire Wire Line
	9850 6000 9950 6000
Connection ~ 9850 6000
Wire Wire Line
	7200 4850 6500 4850
Wire Wire Line
	6500 4850 6500 5100
Wire Wire Line
	7050 5100 7050 4950
Wire Wire Line
	7050 4950 7200 4950
Wire Wire Line
	7050 6200 7050 6000
Wire Wire Line
	6500 6000 6500 6200
Wire Wire Line
	8000 5400 8350 5400
Wire Wire Line
	8000 5550 8350 5550
Wire Wire Line
	8000 5700 8350 5700
Wire Wire Line
	9050 2000 9400 2000
Wire Wire Line
	3750 7350 3750 7400
Wire Wire Line
	3350 6950 3350 7050
Connection ~ 3350 7050
Wire Wire Line
	3350 7400 3750 7400
Connection ~ 3750 7400
Connection ~ 3350 7150
Wire Wire Line
	3750 6500 3750 6550
Wire Wire Line
	3450 6550 3350 6550
Wire Wire Line
	3650 6550 3750 6550
Connection ~ 3750 6550
Connection ~ 4450 6550
Connection ~ 4950 6550
Wire Wire Line
	4150 6950 4450 6950
Wire Wire Line
	4450 6950 4450 6850
Wire Wire Line
	4150 7050 4950 7050
Wire Wire Line
	4950 7050 4950 6850
Connection ~ 4950 7050
Connection ~ 4450 6950
Wire Wire Line
	4150 7150 5450 7150
Wire Wire Line
	5450 6850 5450 7150
Connection ~ 5450 7150
Wire Wire Line
	5450 7450 5450 7400
Wire Wire Line
	8800 2400 9400 2400
Wire Wire Line
	10450 2400 9900 2400
Wire Wire Line
	8000 5850 8350 5850
Wire Wire Line
	9850 3600 8800 3600
Wire Wire Line
	7800 3500 7850 3500
Wire Wire Line
	7800 3800 7850 3800
Wire Wire Line
	7850 3700 7850 3800
Connection ~ 7850 3800
Wire Wire Line
	7850 3600 7850 3500
Connection ~ 7850 3500
Wire Wire Line
	8200 3700 8300 3700
Wire Wire Line
	8300 3600 8200 3600
Wire Wire Line
	7850 3600 7900 3600
Wire Wire Line
	7900 3700 7850 3700
Wire Wire Line
	9100 1100 9000 1100
Wire Wire Line
	9000 2200 9400 2200
Wire Wire Line
	850  2200 850  2300
Wire Wire Line
	850  2300 850  2350
Wire Wire Line
	2100 2300 2450 2300
Wire Wire Line
	2450 2300 2650 2300
Wire Wire Line
	2450 3200 2450 3300
Wire Wire Line
	2950 2850 3150 2850
Wire Wire Line
	3450 3200 3450 3300
Wire Wire Line
	3450 2300 3450 2650
Wire Wire Line
	4050 2300 4500 2300
Wire Wire Line
	4050 2300 4050 2500
Wire Wire Line
	3700 2300 4050 2300
Wire Wire Line
	5600 1750 5850 1750
Wire Wire Line
	7200 1000 7200 1750
Wire Wire Line
	6750 2450 6950 2450
Wire Wire Line
	6850 1000 7200 1000
Wire Wire Line
	850  2100 850  2200
Wire Wire Line
	10050 1100 10200 1100
Wire Wire Line
	9100 1100 9100 1900
Wire Wire Line
	10150 1300 10200 1300
Wire Wire Line
	10150 1700 10150 2000
Wire Wire Line
	10150 2000 10150 2500
Wire Wire Line
	10150 2500 10150 2700
Wire Wire Line
	9150 1500 9400 1500
Wire Wire Line
	9150 2300 9150 3000
Wire Wire Line
	5850 2150 5850 2300
Wire Wire Line
	5850 1750 5950 1750
Wire Wire Line
	9850 5000 9850 5500
Wire Wire Line
	9850 5500 9850 6000
Wire Wire Line
	9850 6000 9850 6250
Wire Wire Line
	3350 7050 3350 7150
Wire Wire Line
	3750 7400 3750 7450
Wire Wire Line
	3350 7150 3350 7400
Wire Wire Line
	3750 6550 3750 6750
Wire Wire Line
	3750 6550 4450 6550
Wire Wire Line
	4450 6550 4950 6550
Wire Wire Line
	4950 6550 5450 6550
Wire Wire Line
	4950 7050 5950 7050
Wire Wire Line
	4450 6950 5650 6950
Wire Wire Line
	5450 7150 5450 7300
Wire Wire Line
	7850 3800 8300 3800
Wire Wire Line
	7850 3500 8300 3500
Wire Wire Line
	8800 3500 8800 3600
Connection ~ 8800 3600
Wire Wire Line
	8800 3700 8800 3600
Wire Wire Line
	8800 3800 8800 3700
Connection ~ 8800 3700
$EndSCHEMATC