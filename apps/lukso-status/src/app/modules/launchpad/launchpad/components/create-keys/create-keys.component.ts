import { Component, EventEmitter, Input, Output } from '@angular/core';
import {
  AbstractControlOptions,
  FormBuilder,
  FormGroup,
  Validators,
} from '@angular/forms';
import { CURRENT_KEY_ACTION, NETWORKS } from '../../helpers/create-keys';
import { CustomValidators } from '../../helpers/custom-validators';

@Component({
  selector: 'lukso-create-keys',
  templateUrl: './create-keys.component.html',
  styleUrls: ['./create-keys.component.scss'],
})
export class CreateKeysComponent {
  @Output() createKeys = new EventEmitter<any>();
  @Output() switchNetwork = new EventEmitter<NETWORKS>();
  @Input() currentTask = {
    status: CURRENT_KEY_ACTION.IDLE,
  };

  NETWORKS = NETWORKS;
  form: FormGroup = new FormGroup({});
  submitted = false;
  isGeneratingKeys = false;

  constructor(private fb: FormBuilder) {
    this.form = this.setupForm();
    this.form.controls.network.valueChanges.subscribe((network: NETWORKS) => {
      this.switchNetwork.emit(network);
    });
  }

  get f() {
    return this.form.controls;
  }

  onSubmit() {
    this.submitted = true;
    if (this.form.invalid) {
      return;
    }

    this.isGeneratingKeys = true;
    this.createKeys.emit(this.form.value);
  }

  private setupForm() {
    return this.fb.group(
      {
        network: [NETWORKS.L15_DEV, [Validators.required]],
        amountOfValidators: ['', [Validators.required]],
        password: [
          '',
          [
            Validators.required,
            CustomValidators.patternValidator(/\d/, { hasNumber: true }),
            CustomValidators.patternValidator(/[A-Z]/, {
              hasCapitalCase: true,
            }),
            CustomValidators.patternValidator(/[a-z]/, { hasSmallCase: true }),
            Validators.minLength(8),
          ],
        ],
        confirmPassword: ['', [Validators.required]],
      },
      {
        validator: CustomValidators.passwordMatchValidator,
      } as AbstractControlOptions
    );
  }
}
