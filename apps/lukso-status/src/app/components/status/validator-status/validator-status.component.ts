import {
  Component,
  OnInit,
  ChangeDetectionStrategy,
  Input,
} from '@angular/core';

@Component({
  selector: 'lukso-validator-status',
  templateUrl: './validator-status.component.html',
  styleUrls: ['./validator-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ValidatorStatusComponent implements OnInit {
  @Input() metrics: any = {};

  constructor() {}

  ngOnInit(): void {}
}
