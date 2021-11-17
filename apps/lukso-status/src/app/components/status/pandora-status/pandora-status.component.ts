import { Component, ChangeDetectionStrategy, Input } from '@angular/core';
import * as shape from 'd3-shape';

@Component({
  selector: 'lukso-pandora-status',
  templateUrl: './pandora-status.component.html',
  styleUrls: ['./pandora-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PandoraStatusComponent {
  @Input() metrics: any = {};

  legend = false;
  showLabels = false;
  animations = true;
  xAxis = false;
  yAxis = true;
  showYAxisLabel = false;
  showXAxisLabel = false;
  timeline = false;
  curve = shape.curveBasis;
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
}
