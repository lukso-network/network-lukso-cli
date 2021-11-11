import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';
import { switchMap } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';
import { Settings } from '../../interfaces/settings';
import { ValidatorMetricsService } from '../../services/validator-metrics.service';

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
  validatorMetrics$: Observable<any>;

  constructor(
    softwareService: SoftwareService,
    vanguardService: VanguardService,
    validatorService: ValidatorMetricsService,
    pandoraService: PandoraMetricsService
  ) {
    this.softwareService = softwareService;
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.vanguardMetrics$ = vanguardService.getMetrics$();
    this.validatorMetrics$ = validatorService.getMetrics$();
  }

  startClients(network: string) {
    this.softwareService
      .getConfig(network)
      .pipe(
        switchMap((settings: Settings) => {
          return this.softwareService.startClients(network, settings);
        })
      )
      .subscribe();
  }

  stopClients() {
    this.softwareService.stopClients().subscribe();
  }
}
