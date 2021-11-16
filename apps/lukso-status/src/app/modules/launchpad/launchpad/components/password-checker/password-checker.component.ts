import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';

@Component({
  selector: 'lukso-password-checker',
  templateUrl: './password-checker.component.html',
  styleUrls: ['./password-checker.component.scss'],
})
export class PasswordCheckerComponent implements OnChanges {
  @Input() password = '';

  minlength = false;
  hasNumber = false;
  hasCapitalCase = false;
  hasSmallCase = false;
  hasSpecialCharacters = false;
  ngOnChanges(changes: SimpleChanges) {
    if (changes.password) {
      this.minlength = new RegExp(/\d/).test(this.password);
      this.hasCapitalCase = new RegExp(/[A-Z]/).test(this.password);
      this.hasNumber = new RegExp(/[0-9]/).test(this.password);
      this.hasSmallCase = new RegExp(/[a-z]/).test(this.password);
      this.hasSpecialCharacters = new RegExp(
        /[ `!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?~]/
      ).test(this.password);
      this.minlength = this.password.length >= 8;
    }
  }
}
