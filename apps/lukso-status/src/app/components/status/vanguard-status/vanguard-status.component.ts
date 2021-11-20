import {
  Component,
  ChangeDetectionStrategy,
  Input,
  SimpleChanges,
  OnChanges,
} from '@angular/core';
import { ClientVersion } from '../../../interfaces/client-versions';
import { GraphData } from '../../../interfaces/graph-data';

@Component({
  selector: 'lukso-vanguard-status',
  templateUrl: './vanguard-status.component.html',
  styleUrls: ['./vanguard-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class VanguardStatusComponent implements OnChanges {
  @Input() metrics: { peers?: number; headSlot?: number } = {};
  @Input() peersOverTime: any = {};
  @Input() version: ClientVersion = {};

  customColors = [{ name: 'Peers', value: '#1CABE1' }];
  graphData: GraphData[] = [];

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.peersOverTime?.currentValue) {
      this.graphData = this.setGraphData(changes.peersOverTime?.currentValue);
    }
  }

  private setGraphData(peersOverTime: any) {
    return [
      {
        name: 'Peers',
        series: Object.entries<number>(peersOverTime).map(([key, value]) => {
          return {
            name: key,
            value,
          };
        }),
      },
    ];
  }
}
