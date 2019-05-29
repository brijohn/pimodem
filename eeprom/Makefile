
V ?= 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

CC ?= gcc
DTC ?= $(shell which dtc)
DTC_VER = $(shell $(DTC) -v 2>/dev/null | grep -Po "(?<=Version: DTC )([0-9.]+)")

BIN      = $(CURDIR)/bin
EEPMAKE = $(BIN)/eepmake

vercmp = $(word 1, $(sort $(1) $(2)))

.PHONY: eeprom
eeprom: check pimodem.dtb | $(EEPMAKE) ; $(info $(M) building eeprom…) @ ## Build eeprom
	$Q $(EEPMAKE) pimodem.txt pimodem.eep pimodem.dtb -c pins.yaml 2>&1 >/dev/null

.PHONY: check
check: ; $(info $(M) checking for device tree compiler…)
	$Q if [ -z "$(DTC_VER)" ]; then \
		echo "DTC not found"; \
		exit 1; \
	fi; \
	if [ $(call vercmp,1.4.3,$(DTC_VER)) != "1.4.3" ]; then \
		echo "DTC $(DTC_VER) < 1.4.3"; \
		exit 1; \
	fi

$(BIN)/eepmake: util/eeptypes.h util/eepmake.c ; $(info $(M) building eepmake…)
	$Q mkdir -p $(BIN)
	$Q $(CC) util/eepmake.c -Iutil -o $(EEPMAKE) -Wno-format

pimodem.dtb: pimodem.dts ; $(info $(M) building device tree blob…)
	$Q $(DTC) -@ -q -I dts -O dtb -o $@ $<

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf *.dtb
	@rm -rf pimodem.eep

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

