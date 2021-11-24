import {
  Component,
  Input,
  EventEmitter,
  Output,
  SimpleChanges,
  OnChanges,
  ChangeDetectionStrategy,
} from '@angular/core';

@Component({
  selector: 'lukso-status-box',
  templateUrl: './status-box.component.html',
  styleUrls: ['./status-box.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class StatusBoxComponent implements OnChanges {
  @Input() name: any;
  @Input() metrics: any = {};
  @Input() peersOverTime: any = {};
  @Input() settings: any = {};
  @Input() blockMetric = '';
  @Input() blockLabel = '';

  @Output() stopClient = new EventEmitter();
  @Output() startClient = new EventEmitter();

  customColors = [{ name: 'Peers', value: '#1CABE1' }];
  graphData: any = null;

  isStarting = false;
  isStopping = false;

  showOfflineOverlay = false;

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.metrics?.currentValue === undefined) {
      this.isStopping = false;
    }

    if (changes.metrics?.currentValue?.peers >= 0) {
      this.isStarting = false;
    }

    if (changes.peersOverTime?.currentValue) {
      this.graphData = [
        {
          name: 'Peers',
          series: Object.entries<number>(
            changes.peersOverTime?.currentValue
          ).map(([key, value]) => {
            return {
              name: key,
              value,
            };
          }),
        },
      ];
    }

    this.showOfflineOverlay =
      this.metrics?.chainData === undefined &&
      !(this.isStarting || this.isStopping);
  }

  toggleClient(startClient: boolean) {
    if (!startClient) {
      this.isStopping = true;
      this.stopClient.emit();
    } else {
      this.isStarting = true;
      this.startClient.emit({
        settings: this.settings,
        clients: [this.name.toLowerCase()],
      });
    }
  }
}
