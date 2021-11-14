import { Component, Input } from '@angular/core';
import { DepositData } from '../../helpers/create-keys';
import { ethers } from 'ethers';

const addressTo = '0x000000000000000000000000000000000000cafE';
import depositABI from './deposit_contract_abi.json';

declare let window: any;

let contract: ethers.Contract;
let provider: any;
let signer: ethers.Signer;

@Component({
  selector: 'lukso-send-transactions',
  templateUrl: './send-transactions.component.html',
  styleUrls: ['./send-transactions.component.scss'],
})
export class SendTransactionsComponent {
  @Input() depositData: DepositData[] | null = null;

  async sendTransactions() {
    if (window.ethereum) {
      await window.ethereum.send('eth_requestAccounts');
      provider = new ethers.providers.Web3Provider(window.ethereum);
      signer = provider.getSigner();
      contract = new ethers.Contract(addressTo, depositABI as any, signer);
    }

    if (this.depositData === null) {
      console.error('Empty DepositData');
      throw new Error('Undefined');
    }

    await Promise.all(
      this.depositData.map(async (_depositData) => {
        console.log(_depositData);
        const { pubkey, withdrawal_credentials, signature, deposit_data_root } =
          _depositData;
        return contract.deposit(
          '0x' + pubkey,
          '0x' + withdrawal_credentials,
          '0x' + signature,
          '0x' + deposit_data_root,
          {
            from: await signer.getAddress(),
            value: ethers.utils.parseEther('32'),
          }
        );
      })
    );
    return true;
  }

  truncate(
    text: string,
    startChars: number,
    endChars: number,
    maxLength: number
  ) {
    if (text.length > maxLength) {
      const start = text.substring(0, startChars);
      const end = text.substring(text.length - endChars, text.length);
      return start + '...' + end;
    }
    return text;
  }
}
