import {
  Component,
  ChangeDetectionStrategy,
  Input,
  OnChanges,
  SimpleChanges,
} from '@angular/core';
import * as shape from 'd3-shape';

@Component({
  selector: 'lukso-pandora-status',
  templateUrl: './pandora-status.component.html',
  styleUrls: ['./pandora-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PandoraStatusComponent implements OnChanges {
  @Input() metrics: any = {};
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
  customColors = [{ name: 'Peers', value: 'rgb(0, 115, 192)' }];
  multi = [
    {
      name: 'Peers',
      series: [
        {
          name: '1',
          value: 3,
        },
        {
          name: '2',
          value: 4,
        },
        {
          name: '3',
          value: 3,
        },
        {
          name: '4',
          value: 3,
        },
        {
          name: '5',
          value: 4,
        },
        {
          name: '6',
          value: 3,
        },
        {
          name: '7',
          value: 3,
        },
        {
          name: '8',
          value: 4,
        },
      ],
    },
  ];

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
