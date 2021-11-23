import { Component, ChangeDetectionStrategy, Inject } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';
import { VanguardService } from '../../services/vanguard-metrics.service';
import { PandoraMetricsService } from '../../services/pandora-metrics.service';
import { ValidatorMetricsService } from '../../services/validator-metrics.service';
import { DEFAULT_NETWORK } from '../../shared/config';
import { RxState } from '@rx-angular/state';
import { GlobalState, GLOBAL_RX_STATE } from '../../shared/rx-state';
import { NETWORKS } from '../../modules/launchpad/launchpad/helpers/create-keys';
import { Settings } from '../../interfaces/settings';

interface StatusState {
  network: NETWORKS;
  settings: Settings;
  networkData: any;
  pandoraMetrics: {
    lastBlock: number;
    peers: number;
  };
  vanguardMetrics: {
    lastSlot: number;
    peers: number;
  };
  pandoraPeersOverTime: { name: string; value: number }[];
  vanguardPeersOverTime: { name: string; value: number }[];
}

@Component({
  selector: 'lukso-status',
  templateUrl: './status.component.html',
  styleUrls: ['./status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class StatusComponent extends RxState<StatusState> {
  readonly network$ = this.select('network');
  readonly settings$ = this.select('settings');
  readonly pandoraPeersOverTime$ = this.select('pandoraPeersOverTime');
  readonly pandoraMetrics$ = this.select('pandoraMetrics');
  readonly vanguardPeersOverTime$ = this.select('vanguardPeersOverTime');
  readonly vanguardMetrics$ = this.select('vanguardMetrics');
  readonly networkData$ = this.select('networkData');

  softwareService: SoftwareService;
  validatorMetrics$: Observable<any>;

  hasStopped = false;

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,

    softwareService: SoftwareService,
    vanguardService: VanguardService,
    validatorService: ValidatorMetricsService,
    pandoraService: PandoraMetricsService
  ) {
    super();

    this.connect('network', this.globalState.select('network'));
    this.connect('settings', this.globalState.select('settings'));

    this.connect('pandoraPeersOverTime', pandoraService.getPeersOverTime$());
    this.connect('pandoraMetrics', pandoraService.getMetrics$());

    this.connect('vanguardPeersOverTime', vanguardService.getPeersOverTime$());
    this.connect('vanguardMetrics', vanguardService.getMetrics$());

    this.softwareService = softwareService;
    this.validatorMetrics$ = validatorService.getMetrics$();
  }

  stopClients(clients: string[]) {
    this.softwareService.stopClients(clients).subscribe();
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
