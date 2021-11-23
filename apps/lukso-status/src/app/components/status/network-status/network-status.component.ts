import { Component, Input, OnChanges } from '@angular/core';

@Component({
  selector: 'lukso-network-status',
  templateUrl: './network-status.component.html',
  styleUrls: ['./network-status.component.scss'],
})
export class NetworkStatusComponent implements OnChanges {
  @Input() networkData: { blockNumber: number; timestamp: number } | null =
    null;

  now: number = Date.now();
  ngOnChanges(): void {
    this.now = Date.now();
  }
}
