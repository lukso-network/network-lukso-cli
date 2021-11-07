import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';

@Component({
  selector: 'lukso-status',
  templateUrl: './status.component.html',
  styleUrls: ['./status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class StatusComponent {
  softwareService: SoftwareService;
  vanguardMetrics$: Observable<any>;
  pandoraMetrics$: Observable<any>;

  constructor(
    softwareService: SoftwareService,
    vanguardService: VanguardService,
    pandoraService: PandoraMetricsService
  ) {
    this.softwareService = softwareService;
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.vanguardMetrics$ = vanguardService.getMetrics$();
  }

  startClients() {
    this.softwareService.startClients().subscribe();
  }

  stopClients() {
    this.softwareService.stopClients().subscribe();
  }
}
