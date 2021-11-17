import { Component, ChangeDetectionStrategy } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';
import { ValidatorMetricsService } from '../../services/validator-metrics.service';
import { DataService } from '../../services/data.service';

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
  peersOverTime$: Observable<any>;
  validatorMetrics$: Observable<any>;
  network$: Observable<any>;

  constructor(
    softwareService: SoftwareService,
    vanguardService: VanguardService,
    validatorService: ValidatorMetricsService,
    pandoraService: PandoraMetricsService,
    dataService: DataService
  ) {
    this.softwareService = softwareService;
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.peersOverTime$ = pandoraService.getPeersOverTime$();
    this.vanguardMetrics$ = vanguardService.getMetrics$();
    this.validatorMetrics$ = validatorService.getMetrics$();
    this.network$ = dataService.getNetwork$();
  }

  stopClients() {
    this.softwareService.stopClients().subscribe();
  }
}
