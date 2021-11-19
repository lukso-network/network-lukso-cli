import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { SoftwareService } from '../../../services/available-versions/available-versions.service';
import { coinbaseValidator } from '../../../shared/eth-address-validator';

interface DownloadedSoftware {
  vanguard: string[];
  pandora: string[];
  'lukso-orchestrator': string[];
  'lukso-validator': string[];
}

@Component({
  selector: 'lukso-initial-setup',
  templateUrl: './initial-setup.component.html',
  styleUrls: ['./initial-setup.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InitialSetupComponent implements OnInit {
  softwareService: SoftwareService;
  router: Router;
  downloadedSoftware$: Observable<DownloadedSoftware>;
  vanguardVersions$: Observable<string[]>;
  pandoraVersions$: Observable<string[]>;
  orchestratorVersions$: Observable<string[]>;
  validatorVersions$: Observable<string[]>;

  setupForm = new FormGroup({
    hostName: new FormControl('', [Validators.required]),
    externalIp: new FormControl(''),
    isValidatorEnabled: new FormControl(0, [Validators.required]),
    versions: new FormGroup({
      vanguard: new FormControl(''),
      pandora: new FormControl(''),
      orchestrator: new FormControl(''),
      validator: new FormControl(''),
    }),
    coinbase: new FormControl('', [
      Validators.required,
      coinbaseValidator(/^0x[a-fA-F0-9]{40}$/i),
    ]),
  });

  constructor(softwareService: SoftwareService, router: Router) {
    this.softwareService = softwareService;
    this.router = router;

    this.downloadedSoftware$ = softwareService
      .getDownloadedVersions$()
      .pipe(shareReplay());

    this.vanguardVersions$ = this.downloadedSoftware$.pipe(
      map((result) => result['vanguard'])
    );
    this.pandoraVersions$ = this.downloadedSoftware$.pipe(
      map((result) => result['pandora'])
    );
    this.orchestratorVersions$ = this.downloadedSoftware$.pipe(
      map((result) => result['lukso-orchestrator'])
    );
    this.validatorVersions$ = this.downloadedSoftware$.pipe(
      map((result) => result['lukso-validator'])
    );
  }

  get hostName() {
    return this.setupForm.get('hostName') as FormControl;
  }
  get coinbase() {
    return this.setupForm.get('coinbase') as FormControl;
  }
  get isValidatorEnabled() {
    return this.setupForm.get('isValidatorEnabled') as FormControl;
  }
  get vanguard() {
    const versions = this.setupForm.get('versions') as FormGroup;
    return versions.get('vanguard') as FormControl;
  }
  get pandora() {
    const versions = this.setupForm.get('versions') as FormGroup;
    return versions.get('pandora') as FormControl;
  }
  get orchestrator() {
    const versions = this.setupForm.get('versions') as FormGroup;
    return versions.get('orchestrator') as FormControl;
  }
  get validator() {
    const versions = this.setupForm.get('versions') as FormGroup;
    return versions.get('validator') as FormControl;
  }

  ngOnInit(): void {
    const network = localStorage.getItem('network') as string;
    this.softwareService.getSettings(network).subscribe((result) => {
      this.setupForm.patchValue(result);
    });
  }

  onSubmit() {
    const network = localStorage.getItem('network') as string;
    if (this.setupForm.valid) {
      this.softwareService
        .setConfig(network, this.setupForm.value)
        .subscribe(() => {
          this.router.navigate(['/status']);
        });
    }
  }
}
