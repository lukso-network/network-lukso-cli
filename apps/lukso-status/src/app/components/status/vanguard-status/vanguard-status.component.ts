import { Component, ChangeDetectionStrategy, Input } from '@angular/core';

@Component({
  selector: 'lukso-vanguard-status',
  templateUrl: './vanguard-status.component.html',
  styleUrls: ['./vanguard-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class VanguardStatusComponent {
  @Input() metrics: any = {};

  peersSelector = 'p2p_peer_count{state="Connected"}';
}
