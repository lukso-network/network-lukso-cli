import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { Observable } from 'rxjs';
import { PandoraMetricsService } from '../pandora-metrics.service';

@Component({
  selector: 'lukso-pandora-status',
  templateUrl: './pandora-status.component.html',
  styleUrls: ['./pandora-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class PandoraStatusComponent implements OnInit {
  metrics$: Observable<any>;
  constructor(pandorsMetrics: PandoraMetricsService) {
    this.metrics$ = pandorsMetrics.getMetrics$();
   }

  ngOnInit(): void {
  }

}
