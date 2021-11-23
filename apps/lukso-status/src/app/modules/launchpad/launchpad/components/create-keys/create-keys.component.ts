import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
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
export class CreateKeysComponent implements OnInit {
  @Output() createKeys = new EventEmitter<any>();
  @Output() switchNetwork = new EventEmitter<NETWORKS>();
  @Input() currentTask = {
    status: CURRENT_KEY_ACTION.IDLE,
  };

  form: FormGroup = new FormGroup({});
  submitted = false;
  isGeneratingKeys = false;

  constructor(private fb: FormBuilder) {
    this.form = this.setupForm();
  }
  ngOnInit(): void {
    this.form.controls.network.setValue(localStorage.getItem('network'));
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

  decrease() {
    const validators = parseInt(
      this.form.controls.amountOfValidators.value,
      10
    );
    if (validators > 1) {
      this.form.controls.amountOfValidators.setValue(validators - 1);
    }
  }
  increase() {
    const validators = parseInt(
      this.form.controls.amountOfValidators.value,
      10
    );
    this.form.controls.amountOfValidators.setValue(validators + 1);
  }

  filterNonNumeric(event: KeyboardEvent) {
    if (!(parseInt(event.key, 10) >= 0 && parseInt(event.key, 10) <= 9))
      return false;
    return true;
  }

  private setupForm() {
    return this.fb.group(
      {
        network: [localStorage.getItem('network'), [Validators.required]],
        amountOfValidators: ['1', [Validators.required]],
        password: [
          '',
          [
            Validators.required,
            CustomValidators.patternValidator(/\d/, { hasNumber: true }),
            CustomValidators.patternValidator(/[A-Z]/, {
              hasCapitalCase: true,
            }),
            CustomValidators.patternValidator(/[a-z]/, { hasSmallCase: true }),
            CustomValidators.patternValidator(
              /[ !@#$%^&*()_+\-=\\[\]{};':"\\|,.<>\\/?]/,
              { hasSpecialCharacters: true }
            ),
            Validators.minLength(8),
          ],
        ],
        confirmPassword: [
          '',
          [
            Validators.required,
            CustomValidators.patternValidator(/\d/, { hasNumber: true }),
            CustomValidators.patternValidator(/[A-Z]/, {
              hasCapitalCase: true,
            }),
            CustomValidators.patternValidator(/[a-z]/, { hasSmallCase: true }),
            CustomValidators.patternValidator(
              /[ !@#$%^&*()_+\-=\\[\]{};':"\\|,.<>\\/?]/,
              { hasSpecialCharacters: true }
            ),
            Validators.minLength(8),
          ],
        ],
      },
      {
        validator: CustomValidators.passwordMatchValidator,
      } as AbstractControlOptions
    );
  }
}
