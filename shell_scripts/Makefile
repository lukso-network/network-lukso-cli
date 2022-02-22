help:
	@echo "use \"make FLAGNAME\" to do particular task";
	@echo "this makefile supports following FLAGNAMEs";
	@echo "   mnemonic        - Generate a mnemonic";
	@echo "   create-deposit  - create deposit-data.json"
	@echo "   create-wallet   - create wallet"

mnemonic:
	@echo "Please copy this mnemonic"; \
	./bin/eth2-val-tools mnemonic;

create-deposit:
	@read -p "Enter number of validators: " NUMBER_OF_VALIDATORS \
	&& read -p "Enter validator mnemonic: " VALIDATOR_MNEMONIC \
	&& read -p "Enter withdrawal mnemonic: " WITHDRAWAL_MNEMONIC \
	&& ./bin/eth2-val-tools deposit-data \
	--fork-version=0x62000069 \
	--source-max=$$NUMBER_OF_VALIDATORS \
	--source-min=0 \
	--validators-mnemonic="$$VALIDATOR_MNEMONIC" --withdrawals-mnemonic="$$WITHDRAWAL_MNEMONIC" \
	| sed 's/}/},/g' | sed '1 s/./[{/' > ./deposit_data.json; echo "]" >> ./deposit_data.json;


create-wallet:
	@read -p "Enter number of validators: " NUMBER_OF_VALIDATORS \
	&& read -p "Enter validator mnemonic: " VALIDATOR_MNEMONIC \
	&& read -p "Enter a valid keystore password: " PASSWORD \
	&& ./bin/eth2-val-tools keystores \
	--insecure \
    --out-loc "./keystore" \
    --prysm-pass $$PASSWORD \
    --source-max $$NUMBER_OF_VALIDATORS \
    --source-min 0 \
    --source-mnemonic "$$VALIDATOR_MNEMONIC" \
    && echo "$$PASSWORD" > ./keystore/password.txt;
