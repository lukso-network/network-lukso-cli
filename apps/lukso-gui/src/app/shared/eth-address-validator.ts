import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

export function coinbaseValidator(nameRe: RegExp): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const isETHAddress = nameRe.test(control.value);
    return !isETHAddress ? { valid_eth1_address: true } : null;
  };
}
