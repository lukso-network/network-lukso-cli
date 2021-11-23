import {
  Component,
  ChangeDetectionStrategy,
  Inject,
  OnInit,
} from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { NETWORKS } from '../../modules/launchpad/launchpad/helpers/create-keys';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { ValidatorService } from '../../services/validator.service';
import { DEFAULT_NETWORK } from '../../shared/config';
import { coinbaseValidator } from '../../shared/eth-address-validator';
import { RxState } from '@rx-angular/state';
import { GlobalState, GLOBAL_RX_STATE } from '../../shared/rx-state';
import { Settings } from '../../interfaces/settings';

interface SettingsState {
  network: NETWORKS;
  settings: Settings;
}

@Component({
  selector: 'lukso-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SettingsComponent
  extends RxState<SettingsState>
  implements OnInit
{
  readonly network$ = this.select('network');
  readonly settings$ = this.select('settings');

  softwareService: SoftwareService;
  validatorService: ValidatorService;

  settingsForm: FormGroup;

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,
    softwareService: SoftwareService,
    validatorService: ValidatorService
  ) {
    super();

    this.settingsForm = this.initForm();
    this.validatorService = validatorService;
    this.softwareService = softwareService;

    this.connect('network', this.globalState.select('network'));
    this.connect('settings', this.globalState.select('settings'));
  }

  ngOnInit(): void {
    this.select('settings').subscribe((settings) => {
      this.settingsForm.patchValue(settings);
    });
  }

  get hostName() {
    return this.settingsForm.get('hostName') as FormControl;
  }
  get coinbase() {
    return this.settingsForm.get('coinbase') as FormControl;
  }
  get isValidatorEnabled() {
    return this.settingsForm.get('isValidatorEnabled') as FormControl;
  }
  get vanguard() {
    const versions = this.settingsForm.get('versions') as FormGroup;
    return versions.get('vanguard') as FormControl;
  }
  get pandora() {
    const versions = this.settingsForm.get('versions') as FormGroup;
    return versions.get('pandora') as FormControl;
  }
  get orchestrator() {
    const versions = this.settingsForm.get('versions') as FormGroup;
    return versions.get('orchestrator') as FormControl;
  }
  get validator() {
    const versions = this.settingsForm.get('versions') as FormGroup;
    return versions.get('validator') as FormControl;
  }

  onSubmit() {
    const network = localStorage.getItem('network') || DEFAULT_NETWORK;
    if (this.settingsForm.valid) {
      this.softwareService
        .setConfig(network, this.settingsForm.value)
        .subscribe();
    }
  }

  resetValidator() {
    const network = localStorage.getItem('network') || DEFAULT_NETWORK;
    this.validatorService.resetValidator(network).subscribe();
  }

  initForm() {
    return new FormGroup({
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
  }
}
