import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'lukso-install',
  templateUrl: './install.component.html',
  styleUrls: ['./install.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class InstallComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
