import {
  Component,
  EventEmitter,
  Input,
  OnChanges,
  Output,
  SimpleChanges,
} from '@angular/core';
import {
  AbstractControlOptions,
  FormBuilder,
  FormGroup,
  Validators,
} from '@angular/forms';
import { CURRENT_KEY_ACTION } from '../../helpers/create-keys';
import { CustomValidators } from '../../helpers/custom-validators';

@Component({
  selector: 'lukso-create-keys',
  templateUrl: './create-keys.component.html',
  styleUrls: ['./create-keys.component.scss'],
})
export class CreateKeysComponent {
  @Output() createKeys = new EventEmitter<any>();
  @Input() currentTask = {
    status: CURRENT_KEY_ACTION.IDLE,
  };

  form: FormGroup = new FormGroup({});
  submitted = false;
  isGeneratingKeys = false;

  constructor(private fb: FormBuilder) {
    this.form = this.setupForm();
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
        network: ['l15-dev', [Validators.required]],
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