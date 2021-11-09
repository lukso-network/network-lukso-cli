import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { SoftwareService } from '../../../services/available-versions/available-versions.service';

@Component({
  selector: 'lukso-initial-setup',
  templateUrl: './initial-setup.component.html',
  styleUrls: ['./initial-setup.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InitialSetupComponent implements OnInit {
  softwareService: SoftwareService;
  setupForm = new FormGroup({
    hostname: new FormControl('', [Validators.required]),
    // coinbase: new FormControl('', [Validators.required]),
  });

  constructor(softwareService: SoftwareService) {
    this.softwareService = softwareService;
  }

  ngOnInit(): void {
    this.softwareService.getConfig('l15-dev').subscribe((result) => {
      this.setupForm.controls.hostname.setValue(result.hostName);
    });
  }

  onSubmit() {
    this.softwareService
      .setConfig('l15-dev', this.setupForm.value)
      .subscribe((result) => {
        console.log(result);
      });
  }
}
