
.PHONY: all
all: eeprom pimodem

.PHONY: eeprom
eeprom:
	@make -s -C eeprom eeprom

.PHONY: pimodem
pimodem:
	@make -s -C src all

.PHONY: clean
clean:
	@make -s -C eeprom clean
	@make -s -C src clean
