import {
  Component,
  ChangeDetectionStrategy,
  Inject,
  OnInit,
} from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { NETWORKS } from '../../modules/launchpad/launchpad/helpers/create-keys';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { ValidatorService } from '../../services/validator.service';
import { coinbaseValidator } from '../../shared/eth-address-validator';
import { RxState } from '@rx-angular/state';
import { GlobalState, GLOBAL_RX_STATE } from '../../shared/rx-state';
import { Settings } from '../../interfaces/settings';
import { switchMap } from 'rxjs/operators';
import { Subject } from 'rxjs';

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

  resetValidator$ = new Subject<NETWORKS>();
  saveSettings$ = new Subject<{ network: NETWORKS; settings: Settings }>();

  settingsForm: FormGroup;

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,
    fb: FormBuilder,
    softwareService: SoftwareService,
    validatorService: ValidatorService
  ) {
    super();

    this.settingsForm = this.initForm(fb);

    this.connect('network', this.globalState.select('network'));
    this.connect('settings', this.globalState.select('settings'));

    // Side Effects
    this.hold(this.resetValidator$, (network) =>
      validatorService.resetValidator(network)
    );
    this.hold(this.saveSettings$, (values) =>
      softwareService.setConfig(values.network, values.settings)
    );
  }

  ngOnInit(): void {
    this.select('settings').subscribe((settings) => {
      this.settingsForm.patchValue(settings);
    });
  }

  // TODO: check best practices, this certainly isn't it
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

  private initForm(fb: FormBuilder) {
    return fb.group({
      hostName: ['', [Validators.required]],
      externalIp: [''],
      isValidatorEnabled: [0, [Validators.required]],
      versions: fb.group({
        vanguard: [''],
        pandora: [''],
        orchestrator: [''],
        validator: [''],
      }),
      coinbase: [
        '',
        [Validators.required, coinbaseValidator(/^0x[a-fA-F0-9]{40}$/i)],
      ],
    });
  }
}
