import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { KeygenService } from '../services/keygen.service';

@Component({
  selector: 'lukso-launchpad',
  templateUrl: './launchpad.component.html',
  styleUrls: ['./launchpad.component.scss'],
})
export class LaunchpadComponent {
  keygenService: KeygenService;
  showPasswordError = false;
  isGeneratingKeys = false;
  constructor(keygenService: KeygenService) {
    this.keygenService = keygenService;
  }

  genereateKeys(pw1: string, pw2: string) {
    if (
      this.arePasswordsIdentical(pw1, pw2) &&
      this.isPasswordCriteriaMet(pw1)
    ) {
      this.isGeneratingKeys = true;
      this.keygenService
        .genereateKeys('12345678', 'l15-dev', '1')
        .subscribe(() => {
          this.isGeneratingKeys = false;
        });
    } else {
      this.showPasswordError = true;
    }
  }

  arePasswordsIdentical(pw1: string, pw2: string) {
    return pw1 === pw2;
  }

  isPasswordCriteriaMet(password: string) {
    // https://www.section.io/engineering-education/password-strength-checker-javascript/
    const pwStrength = new RegExp(
      '(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[^A-Za-z0-9])(?=.{8,})'
    );

    return pwStrength.test(password);
  }
}
