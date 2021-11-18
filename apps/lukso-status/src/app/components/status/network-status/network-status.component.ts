import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lukso-network-status',
  templateUrl: './network-status.component.html',
  styleUrls: ['./network-status.component.scss'],
})
export class NetworkStatusComponent {
  @Input() blockInfo: any;
  constructor() {}
}
