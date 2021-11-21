import { Component, ChangeDetectionStrategy } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';
import { ValidatorMetricsService } from '../../services/validator-metrics.service';
import { DataService } from '../../services/data.service';
import { switchMap } from 'rxjs/operators';
import { DEFAULT_NETWORK } from '../../shared/config';

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
  peersOverTimePandora$: Observable<any>;
  vanguardPeersOverTime$: Observable<any>;
  validatorMetrics$: Observable<any>;
  network$: Observable<any>;
  lastBlock$: Observable<any>;
  settings$: Observable<any>;

  hasStopped = false;

  constructor(
    softwareService: SoftwareService,
    vanguardService: VanguardService,
    validatorService: ValidatorMetricsService,
    pandoraService: PandoraMetricsService,
    dataService: DataService
  ) {
    this.softwareService = softwareService;
    this.lastBlock$ = pandoraService.myWSData$;
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.peersOverTimePandora$ = pandoraService.getPeersOverTime$();
    this.vanguardPeersOverTime$ = vanguardService.getPeersOverTime$();
    this.vanguardMetrics$ = vanguardService.getMetrics$();
    this.validatorMetrics$ = validatorService.getMetrics$();
    this.network$ = dataService.getNetwork$();
    this.settings$ = this.network$.pipe(
      switchMap((network) => {
        return softwareService.getSettings(network);
      })
    );
  }

  stopClients(clients: string[]) {
    this.softwareService.stopClients(clients).subscribe(() => {});
  }

  startClients(options: any) {
    this.softwareService
      .startClients(
        localStorage.getItem('network') || DEFAULT_NETWORK,
        options.settings,
        options.clients
      )
      .subscribe();
  }
}
