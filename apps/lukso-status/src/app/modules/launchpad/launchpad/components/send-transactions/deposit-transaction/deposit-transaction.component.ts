import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { getNamespacePrefix } from '../../../../../../shared/config';
import { NETWORKS } from '../../../helpers/create-keys';

@Component({
  selector: 'lukso-deposit-transaction',
  templateUrl: './deposit-transaction.component.html',
  styleUrls: ['./deposit-transaction.component.scss'],
})
export class DepositTransactionComponent implements OnChanges {
  @Input() deposit: any;
  explorerLink: string | null = null;

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.deposit?.currentValue) {
      this.explorerLink = `http://${getNamespacePrefix(
        NETWORKS.L15_DEV
      )}explorer.vanguard.l15.lukso.network/validator/${
        changes.deposit?.currentValue.pubkey
      }`;
    }
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
