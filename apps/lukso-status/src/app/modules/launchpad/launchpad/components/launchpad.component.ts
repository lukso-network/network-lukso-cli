import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { KeygenService } from '../services/keygen.service';

@Component({
  selector: 'lukso-launchpad',
  templateUrl: './launchpad.component.html',
  styleUrls: ['./launchpad.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LaunchpadComponent {
  keygenService: KeygenService;
  constructor(keygenService: KeygenService) {
    this.keygenService = keygenService;
  }

  genereateKeys() {
    this.keygenService.genereateKeys('12345678', 'l15-dev', '5').subscribe();
  }
}
