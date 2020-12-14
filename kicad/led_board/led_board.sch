EESchema Schematic File Version 4
LIBS:led_board-cache
EELAYER 26 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "PiModem - LED board"
Date "2019-06-18"
Rev "1.0"
Comp ""
Comment1 ""
Comment2 "License: GPLv2"
Comment3 "Brian Johnson"
Comment4 "Copyright (c) 2019"
$EndDescr
$Comp
L Connector_Generic:Conn_02x05_Odd_Even J1
U 1 1 5D08EF40
P 7800 3600
F 0 "J1" H 7850 4017 50  0000 C CNN
F 1 "MAIN_BOARD" H 7850 3926 50  0000 C CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_2x05_P2.54mm_Vertical" H 7800 3600 50  0001 C CNN
F 3 "~" H 7800 3600 50  0001 C CNN
	1    7800 3600
	1    0    0    -1  
$EndComp
$Comp
L Device:LED D6
U 1 1 5D08F353
P 5650 4000
F 0 "D6" H 5641 4216 50  0000 C CNN
F 1 "SD" H 5641 4125 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 4000 50  0001 C CNN
F 3 "~" H 5650 4000 50  0001 C CNN
	1    5650 4000
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D5
U 1 1 5D08F397
P 5650 3650
F 0 "D5" H 5641 3866 50  0000 C CNN
F 1 "RD" H 5641 3775 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 3650 50  0001 C CNN
F 3 "~" H 5650 3650 50  0001 C CNN
	1    5650 3650
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D8
U 1 1 5D08F3C9
P 5650 4700
F 0 "D8" H 5641 4916 50  0000 C CNN
F 1 "TR" H 5641 4825 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 4700 50  0001 C CNN
F 3 "~" H 5650 4700 50  0001 C CNN
	1    5650 4700
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D4
U 1 1 5D08F3FD
P 5650 3300
F 0 "D4" H 5641 3516 50  0000 C CNN
F 1 "OH" H 5641 3425 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 3300 50  0001 C CNN
F 3 "~" H 5650 3300 50  0001 C CNN
	1    5650 3300
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D9
U 1 1 5D08F41F
P 5650 5050
F 0 "D9" H 5641 5266 50  0000 C CNN
F 1 "MR" H 5641 5175 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 5050 50  0001 C CNN
F 3 "~" H 5650 5050 50  0001 C CNN
	1    5650 5050
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D7
U 1 1 5D08F44F
P 5650 4350
F 0 "D7" H 5641 4566 50  0000 C CNN
F 1 "ACT" H 5641 4475 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 4350 50  0001 C CNN
F 3 "~" H 5650 4350 50  0001 C CNN
	1    5650 4350
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D3
U 1 1 5D08F47D
P 5650 2950
F 0 "D3" H 5641 3166 50  0000 C CNN
F 1 "CD" H 5641 3075 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 2950 50  0001 C CNN
F 3 "~" H 5650 2950 50  0001 C CNN
	1    5650 2950
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D2
U 1 1 5D08F4B9
P 5650 2600
F 0 "D2" H 5650 2800 50  0000 C CNN
F 1 "AA" H 5650 2700 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 2600 50  0001 C CNN
F 3 "~" H 5650 2600 50  0001 C CNN
	1    5650 2600
	-1   0    0    1   
$EndComp
$Comp
L Device:LED D1
U 1 1 5D08F4F3
P 5650 2250
F 0 "D1" H 5641 2466 50  0000 C CNN
F 1 "PWR" H 5641 2375 50  0000 C CNN
F 2 "LED_THT:LED_D5.0mm" H 5650 2250 50  0001 C CNN
F 3 "~" H 5650 2250 50  0001 C CNN
	1    5650 2250
	-1   0    0    1   
$EndComp
$Comp
L Device:R R3
U 1 1 5D08F7CB
P 5150 2950
F 0 "R3" V 5250 2950 50  0000 C CNN
F 1 "325" V 5034 2950 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 2950 50  0001 C CNN
F 3 "~" H 5150 2950 50  0001 C CNN
	1    5150 2950
	0    1    1    0   
$EndComp
$Comp
L Device:R R1
U 1 1 5D08FBE1
P 5150 2250
F 0 "R1" V 5250 2250 50  0000 C CNN
F 1 "325" V 5034 2250 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 2250 50  0001 C CNN
F 3 "~" H 5150 2250 50  0001 C CNN
	1    5150 2250
	0    1    1    0   
$EndComp
$Comp
L Device:R R2
U 1 1 5D08FC1D
P 5150 2600
F 0 "R2" V 5250 2600 50  0000 C CNN
F 1 "325" V 5034 2600 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 2600 50  0001 C CNN
F 3 "~" H 5150 2600 50  0001 C CNN
	1    5150 2600
	0    1    1    0   
$EndComp
$Comp
L Device:R R5
U 1 1 5D08FC5B
P 5150 3650
F 0 "R5" V 5250 3650 50  0000 C CNN
F 1 "325" V 5034 3650 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 3650 50  0001 C CNN
F 3 "~" H 5150 3650 50  0001 C CNN
	1    5150 3650
	0    1    1    0   
$EndComp
$Comp
L Device:R R4
U 1 1 5D08FC9B
P 5150 3300
F 0 "R4" V 5250 3300 50  0000 C CNN
F 1 "325" V 5034 3300 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 3300 50  0001 C CNN
F 3 "~" H 5150 3300 50  0001 C CNN
	1    5150 3300
	0    1    1    0   
$EndComp
$Comp
L Device:R R7
U 1 1 5D08FCD5
P 5150 4350
F 0 "R7" V 5250 4350 50  0000 C CNN
F 1 "325" V 5034 4350 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 4350 50  0001 C CNN
F 3 "~" H 5150 4350 50  0001 C CNN
	1    5150 4350
	0    1    1    0   
$EndComp
$Comp
L Device:R R6
U 1 1 5D08FD19
P 5150 4000
F 0 "R6" V 5250 4000 50  0000 C CNN
F 1 "325" V 5034 4000 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 4000 50  0001 C CNN
F 3 "~" H 5150 4000 50  0001 C CNN
	1    5150 4000
	0    1    1    0   
$EndComp
$Comp
L Device:R R9
U 1 1 5D08FD55
P 5150 5050
F 0 "R9" V 5250 5050 50  0000 C CNN
F 1 "325" V 5034 5050 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 5050 50  0001 C CNN
F 3 "~" H 5150 5050 50  0001 C CNN
	1    5150 5050
	0    1    1    0   
$EndComp
$Comp
L Device:R R8
U 1 1 5D08FD9F
P 5150 4700
F 0 "R8" V 5250 4700 50  0000 C CNN
F 1 "325" V 5034 4700 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0207_L6.3mm_D2.5mm_P10.16mm_Horizontal" V 5080 4700 50  0001 C CNN
F 3 "~" H 5150 4700 50  0001 C CNN
	1    5150 4700
	0    1    1    0   
$EndComp
$Comp
L power:+5V #PWR0101
U 1 1 5D090067
P 8250 4000
F 0 "#PWR0101" H 8250 3850 50  0001 C CNN
F 1 "+5V" H 8265 4173 50  0000 C CNN
F 2 "" H 8250 4000 50  0001 C CNN
F 3 "" H 8250 4000 50  0001 C CNN
	1    8250 4000
	-1   0    0    1   
$EndComp
$Comp
L power:GND #PWR0102
U 1 1 5D0900F1
P 7450 3200
F 0 "#PWR0102" H 7450 2950 50  0001 C CNN
F 1 "GND" H 7455 3027 50  0000 C CNN
F 2 "" H 7450 3200 50  0001 C CNN
F 3 "" H 7450 3200 50  0001 C CNN
	1    7450 3200
	-1   0    0    1   
$EndComp
Wire Wire Line
	8100 3800 8250 3800
Wire Wire Line
	8250 3800 8250 3950
Wire Wire Line
	7600 3400 7450 3400
Wire Wire Line
	7450 3400 7450 3250
Wire Wire Line
	5300 2250 5500 2250
Wire Wire Line
	5300 2600 5500 2600
Wire Wire Line
	5500 2950 5300 2950
Wire Wire Line
	5300 3300 5500 3300
Wire Wire Line
	5300 3650 5500 3650
Wire Wire Line
	5500 4000 5300 4000
Wire Wire Line
	5300 4350 5500 4350
Wire Wire Line
	5300 4700 5500 4700
Wire Wire Line
	5300 5050 5500 5050
Wire Wire Line
	5000 2250 4850 2250
Wire Wire Line
	4850 2250 4850 2600
Wire Wire Line
	4850 5050 5000 5050
Wire Wire Line
	5000 4700 4850 4700
Connection ~ 4850 4700
Wire Wire Line
	4850 4700 4850 5050
Wire Wire Line
	4850 4350 5000 4350
Connection ~ 4850 4350
Wire Wire Line
	4850 4350 4850 4700
Wire Wire Line
	5000 2600 4850 2600
Connection ~ 4850 2600
Wire Wire Line
	4850 2600 4850 2950
Wire Wire Line
	5000 2950 4850 2950
Connection ~ 4850 2950
Wire Wire Line
	4850 2950 4850 3300
Wire Wire Line
	4850 3300 5000 3300
Connection ~ 4850 3300
Wire Wire Line
	4850 3300 4850 3650
Wire Wire Line
	5000 3650 4850 3650
Connection ~ 4850 3650
Wire Wire Line
	4850 3650 4850 4000
Wire Wire Line
	4850 4000 5000 4000
Connection ~ 4850 4000
Wire Wire Line
	4850 4000 4850 4350
Wire Wire Line
	4850 2250 4850 1900
Connection ~ 4850 2250
$Comp
L power:+5V #PWR0103
U 1 1 5D09535D
P 4850 1900
F 0 "#PWR0103" H 4850 1750 50  0001 C CNN
F 1 "+5V" H 4865 2073 50  0000 C CNN
F 2 "" H 4850 1900 50  0001 C CNN
F 3 "" H 4850 1900 50  0001 C CNN
	1    4850 1900
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR0104
U 1 1 5D095DE7
P 6250 2250
F 0 "#PWR0104" H 6250 2000 50  0001 C CNN
F 1 "GND" V 6255 2122 50  0000 R CNN
F 2 "" H 6250 2250 50  0001 C CNN
F 3 "" H 6250 2250 50  0001 C CNN
	1    6250 2250
	0    -1   -1   0   
$EndComp
Wire Wire Line
	6250 2250 5800 2250
Wire Wire Line
	5800 2600 6300 2600
Wire Wire Line
	5800 2950 6300 2950
Wire Wire Line
	5800 3300 6300 3300
Wire Wire Line
	5800 3650 6300 3650
Wire Wire Line
	5800 4000 6300 4000
Wire Wire Line
	5800 4350 6300 4350
Wire Wire Line
	5800 4700 6300 4700
Wire Wire Line
	5800 5050 6300 5050
Text Label 6300 5050 2    50   ~ 0
MR
Text Label 6300 4700 2    50   ~ 0
TR
Text Label 6300 4350 2    50   ~ 0
ACT
Text Label 6300 4000 2    50   ~ 0
SD
Text Label 6300 3650 2    50   ~ 0
RD
Text Label 6300 3300 2    50   ~ 0
OH
Text Label 6300 2950 2    50   ~ 0
CD
Text Label 6300 2600 2    50   ~ 0
AA
Wire Wire Line
	7600 3600 7300 3600
Text Label 7300 3600 0    50   ~ 0
MR
Wire Wire Line
	8100 3500 8400 3500
Text Label 8400 3500 2    50   ~ 0
TR
Wire Wire Line
	8100 3400 8400 3400
Text Label 8400 3400 2    50   ~ 0
CD
Wire Wire Line
	8100 3700 8400 3700
Text Label 8400 3700 2    50   ~ 0
SD
Wire Wire Line
	7600 3800 7300 3800
Text Label 7300 3800 0    50   ~ 0
RD
Wire Wire Line
	7600 3500 7300 3500
Text Label 7300 3500 0    50   ~ 0
ACT
Wire Wire Line
	8100 3600 8400 3600
Wire Wire Line
	7600 3700 7300 3700
Text Label 8400 3600 2    50   ~ 0
OH
Text Label 7300 3700 0    50   ~ 0
AA
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5D0A0436
P 7450 3250
F 0 "#FLG0101" H 7450 3325 50  0001 C CNN
F 1 "PWR_FLAG" V 7450 3378 50  0000 L CNN
F 2 "" H 7450 3250 50  0001 C CNN
F 3 "~" H 7450 3250 50  0001 C CNN
	1    7450 3250
	0    -1   -1   0   
$EndComp
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5D0A046E
P 8250 3950
F 0 "#FLG0102" H 8250 4025 50  0001 C CNN
F 1 "PWR_FLAG" V 8250 4078 50  0000 L CNN
F 2 "" H 8250 3950 50  0001 C CNN
F 3 "~" H 8250 3950 50  0001 C CNN
	1    8250 3950
	0    1    1    0   
$EndComp
Connection ~ 7450 3250
Wire Wire Line
	7450 3250 7450 3200
Connection ~ 8250 3950
Wire Wire Line
	8250 3950 8250 4000
$EndSCHEMATC