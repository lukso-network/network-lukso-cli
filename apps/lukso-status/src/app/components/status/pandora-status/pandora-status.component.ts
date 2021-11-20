import {
  Component,
  ChangeDetectionStrategy,
  Input,
  OnChanges,
  SimpleChanges,
} from '@angular/core';
import { ClientVersion } from '../../../interfaces/client-versions';

@Component({
  selector: 'lukso-pandora-status',
  templateUrl: './pandora-status.component.html',
  styleUrls: ['./pandora-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PandoraStatusComponent implements OnChanges {
  @Input() metrics: any = {};
  @Input() peersOverTime: any = {};
  @Input() version: ClientVersion = {};
  legend = false;
  showLabels = false;
  animations = true;
  xAxis = false;
  yAxis = true;
  showYAxisLabel = false;
  showXAxisLabel = false;
  timeline = false;
  rangeFillOpacity = 1;
  customColors = [{ name: 'Peers', value: '#1CABE1' }];
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
