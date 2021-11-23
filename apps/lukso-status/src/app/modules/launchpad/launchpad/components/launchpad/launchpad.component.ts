import { Component, Inject } from '@angular/core';
import { map, mapTo, switchMap, tap, withLatestFrom } from 'rxjs/operators';
import {
  CURRENT_KEY_ACTION,
  NETWORKS,
  KeyGenerationValues,
  DepositData,
} from '../../helpers/create-keys';
import { KeygenService } from '../../services/keygen.service';
import { saveAs } from 'file-saver';
import {
  GlobalState,
  GLOBAL_RX_STATE,
} from '../../../../../../app/shared/rx-state';
import { RxState } from '@rx-angular/state';
import { merge, Subject } from 'rxjs';

interface LaunchpadState {
  network: NETWORKS;
  depositData: DepositData[];
  status: CURRENT_KEY_ACTION;
}

@Component({
  selector: 'lukso-launchpad',
  templateUrl: './launchpad.component.html',
  styleUrls: ['./launchpad.component.scss'],
})
export class LaunchpadComponent extends RxState<LaunchpadState> {
  readonly network$ = this.select('network');
  readonly depositData$ = this.select('depositData');
  readonly status$ = this.select('status');

  state$ = this.select();
  createKeys$ = new Subject<KeyGenerationValues>();
  keygenService: KeygenService;

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,
    keygenService: KeygenService
  ) {
    super();
    this.keygenService = keygenService;

    this.hold(this.createKeys$.pipe(tap((values) => this.createKeys(values))));
    this.set({ status: CURRENT_KEY_ACTION.IDLE });
    this.connect('network', globalState.select('network'));
    this.connect(
      'depositData',
      merge(this.status$, this.network$).pipe(
        switchMap(() => {
          return this.keygenService.getDepositData('l15-dev');
        })
      )
    );
  }

  createKeys(values: KeyGenerationValues) {
    this.set({ status: CURRENT_KEY_ACTION.GENERATING });
    this.keygenService
      .genereateKeys(
        values.password,
        values.network,
        values.amountOfValidators.toString()
      )
      .pipe(
        switchMap(() => {
          this.set({ status: CURRENT_KEY_ACTION.IMPORTING });
          return this.keygenService.importKeys(values.password, values.network);
        })
      )
      .subscribe({
        next: (response: any) => {
          this.set({ status: CURRENT_KEY_ACTION.COMPLETE });
          const blob = new Blob([response], {
            type: 'text/json; charset=utf-8',
          });
          saveAs(blob, 'validator_keys.zip');
        },
        error: (error: Error) =>
          console.log('Error downloading the file', error),
      });
  }
}
