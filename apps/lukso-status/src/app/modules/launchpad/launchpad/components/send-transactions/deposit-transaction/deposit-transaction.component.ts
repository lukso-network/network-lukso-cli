import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lukso-deposit-transaction',
  templateUrl: './deposit-transaction.component.html',
  styleUrls: ['./deposit-transaction.component.scss'],
})
export class DepositTransactionComponent implements OnInit {
  @Input() deposit: any;
  constructor() {}

  ngOnInit(): void {}

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
