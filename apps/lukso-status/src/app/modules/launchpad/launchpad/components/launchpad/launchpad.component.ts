import { Component, Inject } from '@angular/core';
import { Router } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { CURRENT_KEY_ACTION, NETWORKS } from '../../helpers/create-keys';
import { KeygenService } from '../../services/keygen.service';
import { saveAs } from 'file-saver';
import { Observable } from 'rxjs';
import {
  GlobalState,
  GLOBAL_RX_STATE,
} from '../../../../../../app/shared/rx-state';
import { RxState } from '@rx-angular/state';

interface KeyGenerationValues {
  network: string;
  amountOfValidators: number;
  password: string;
}

interface LaunchpadState {
  network: NETWORKS;
}

@Component({
  selector: 'lukso-launchpad',
  templateUrl: './launchpad.component.html',
  styleUrls: ['./launchpad.component.scss'],
})
export class LaunchpadComponent extends RxState<LaunchpadState> {
  readonly network$ = this.select('network');

  keygenService: KeygenService;
  router: Router;
  showPasswordError = false;
  depositData$: Observable<any>;
  currentTask = {
    status: CURRENT_KEY_ACTION.IDLE,
  };

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,
    keygenService: KeygenService,
    router: Router
  ) {
    super();

    this.router = router;
    this.keygenService = keygenService;
    this.depositData$ = this.network$.pipe(
      switchMap((network: NETWORKS) => {
        return this.keygenService.getDepositData(network);
      })
    );
  }

  createKeys(values: KeyGenerationValues) {
    this.currentTask.status = CURRENT_KEY_ACTION.GENERATING;
    this.keygenService
      .genereateKeys(
        values.password,
        values.network,
        values.amountOfValidators.toString()
      )
      .pipe(
        switchMap(() => {
          this.currentTask.status = CURRENT_KEY_ACTION.IMPORTING;
          return this.keygenService.importKeys(values.password, values.network);
        })
      )
      .subscribe({
        next: (response: any) => {
          this.currentTask.status = CURRENT_KEY_ACTION.COMPLETE;
          const blob: any = new Blob([response], {
            type: 'text/json; charset=utf-8',
          });
          saveAs(blob, 'validator_keys.zip');
        },
        error: (error: Error) =>
          console.log('Error downloading the file', error),
      });
  }
}
