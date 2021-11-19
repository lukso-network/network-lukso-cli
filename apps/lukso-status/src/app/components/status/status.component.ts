import { Component, ChangeDetectionStrategy } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';
import { ValidatorMetricsService } from '../../services/validator-metrics.service';
import { DataService } from '../../services/data.service';
import { switchMap } from 'rxjs/operators';

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
  peersOverTimeVanguard$: Observable<any>;
  validatorMetrics$: Observable<any>;
  network$: Observable<any>;
  lastBlock$: Observable<any>;
  settings$: Observable<any>;

  constructor(
    softwareService: SoftwareService,
    vanguardService: VanguardService,
    validatorService: ValidatorMetricsService,
    pandoraService: PandoraMetricsService,
    dataService: DataService
  ) {
    this.softwareService = softwareService;
    this.pandoraMetrics$ = pandoraService.getMetrics$();
    this.peersOverTimePandora$ = pandoraService.getPeersOverTime$();
    this.peersOverTimeVanguard$ = vanguardService.getPeersOverTime$();
    this.lastBlock$ = pandoraService.myWSData$;
    this.vanguardMetrics$ = vanguardService.getMetrics$();
    this.validatorMetrics$ = validatorService.getMetrics$();
    this.network$ = dataService.getNetwork$();
    this.settings$ = this.network$.pipe(
      switchMap((network) => {
        return softwareService.getSettings(network);
      })
    );
  }

  stopClients() {
    this.softwareService.stopClients().subscribe();
  }
}
