import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';

@Component({
  selector: 'lukso-network-status',
  templateUrl: './network-status.component.html',
  styleUrls: ['./network-status.component.scss'],
})
export class NetworkStatusComponent implements OnChanges {
  @Input() blockInfo: any;
  now: number = Date.now();
  constructor() {}
  ngOnChanges(changes: SimpleChanges): void {
    this.now = Date.now();
  }
}
