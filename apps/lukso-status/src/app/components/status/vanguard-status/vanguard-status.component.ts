import {
  Component,
  ChangeDetectionStrategy,
  Input,
  SimpleChanges,
} from '@angular/core';

@Component({
  selector: 'lukso-vanguard-status',
  templateUrl: './vanguard-status.component.html',
  styleUrls: ['./vanguard-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class VanguardStatusComponent {
  @Input() metrics: { peers?: number; headSlot?: number } = {};

  @Input() peersOverTime: any = {};

  legend = false;
  showLabels = false;
  animations = true;
  xAxis = false;
  yAxis = true;
  showYAxisLabel = false;
  showXAxisLabel = false;
  timeline = false;
  rangeFillOpacity = 1;
  customColors = [{ name: 'Peers', value: '#b62daf' }];
  multi: any = null;

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.peersOverTime?.currentValue) {
      const series = Object.entries<number>(
        changes.peersOverTime?.currentValue
      ).map(([key, value]) => {
        return {
          name: key,
          value,
        };
      });
      this.multi = [
        {
          name: 'Peers',
          series,
        },
      ];
    }
  }
}
