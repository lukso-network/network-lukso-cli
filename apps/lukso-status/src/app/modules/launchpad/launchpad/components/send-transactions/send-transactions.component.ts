import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lukso-send-transactions',
  templateUrl: './send-transactions.component.html',
  styleUrls: ['./send-transactions.component.scss'],
})
export class SendTransactionsComponent implements OnInit {
  @Input() depositData: any;
  constructor() {}

  ngOnInit(): void {}
}
