import { Component, Input } from '@angular/core';
import { RxState } from '@rx-angular/state';
import { Observable, Subject } from 'rxjs';

interface ButtonState {
  inProgress: boolean;
}

@Component({
  selector: 'lukso-button',
  templateUrl: './button.component.html',
  styleUrls: ['./button.component.scss'],
})
export class ButtonComponent extends RxState<ButtonState> {
  @Input() status$: Observable<any> | undefined;
  @Input() inProgressText = '';
  @Input() defaultText = '';
  btn$ = new Subject();
}
