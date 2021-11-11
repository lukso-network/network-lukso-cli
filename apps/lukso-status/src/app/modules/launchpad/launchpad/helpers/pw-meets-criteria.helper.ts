import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

export function pwOkValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const pwStrength = new RegExp(
      '(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[^A-Za-z0-9])(?=.{8,})'
    );
    const forbidden = pwStrength.test(control.value);
    return !forbidden ? { criteria_not_met: { value: control.value } } : null;
  };
}
