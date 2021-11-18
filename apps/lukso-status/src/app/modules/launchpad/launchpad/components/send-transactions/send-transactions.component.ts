import { Component, Input } from '@angular/core';
import { DepositData } from '../../helpers/create-keys';
import { ContractInterface, ContractTransaction, ethers } from 'ethers';

import DEPOSIT_ABI from './deposit_contract_abi.json';
import {
  DEPOSIT_CONTRACT_ADDRESS,
  VALIDATOR_DEPOSIT_COST,
} from '../../../../../shared/config';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare let window: any;

let contract: ethers.Contract;
let provider: ethers.providers.Web3Provider;
let signer: ethers.Signer;

@Component({
  selector: 'lukso-send-transactions',
  templateUrl: './send-transactions.component.html',
  styleUrls: ['./send-transactions.component.scss'],
})
export class SendTransactionsComponent {
  @Input() depositData: DepositData[] | null = null;

  constructor() {
    provider = new ethers.providers.Web3Provider(window.ethereum);
    signer = provider.getSigner();

    contract = new ethers.Contract(
      DEPOSIT_CONTRACT_ADDRESS,
      DEPOSIT_ABI as ContractInterface,
      signer
    );
  }

  async sendTransactions() {
    if (window.ethereum) {
      await window.ethereum.send('eth_requestAccounts');
    }

    if (this.depositData === null) {
      console.error('Empty DepositData');
      throw new Error('Undefined');
    }

    await Promise.all(
      this.depositData.map(async (_depositData) => {
        const { pubkey, withdrawal_credentials, signature, deposit_data_root } =
          _depositData;
        return contract
          .deposit(
            '0x' + pubkey,
            '0x' + withdrawal_credentials,
            '0x' + signature,
            '0x' + deposit_data_root,
            {
              from: await signer.getAddress(),
              value: ethers.utils.parseEther(VALIDATOR_DEPOSIT_COST.toString()),
            }
          )
          .then(async (transaction: ContractTransaction) => {
            _depositData.transaction = transaction;
            await transaction.wait();
            _depositData.transaction_confirmed = true;
          });
      })
    );

    return true;
  }
}
