import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { CURRENT_KEY_ACTION, NETWORKS } from '../../helpers/create-keys';
import { KeygenService } from '../../services/keygen.service';
import { saveAs } from 'file-saver';
import { BehaviorSubject, Observable } from 'rxjs';

interface KeyGenerationValues {
  network: string;
  amountOfValidators: number;
  password: string;
}

@Component({
  selector: 'lukso-launchpad',
  templateUrl: './launchpad.component.html',
  styleUrls: ['./launchpad.component.scss'],
})
export class LaunchpadComponent {
  keygenService: KeygenService;
  router: Router;
  showPasswordError = false;
  network$ = new BehaviorSubject<NETWORKS>(NETWORKS.L15_DEV);
  depositData$: Observable<any>;
  currentTask = {
    status: CURRENT_KEY_ACTION.IDLE,
  };

  constructor(keygenService: KeygenService, router: Router) {
    this.router = router;
    this.keygenService = keygenService;
    this.depositData$ = this.network$.pipe(
      switchMap((network: NETWORKS) => {
        return this.keygenService.getDepositData(network);
      })
    );
  }

  switchNetwork(event: any) {
    console.log(event);
    this.network$.next(event);
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
          console.log(response);
          this.currentTask.status = CURRENT_KEY_ACTION.COMPLETE;
          const blob: any = new Blob([response], {
            type: 'text/json; charset=utf-8',
          });
          saveAs(blob, 'validator_keys.zip');
        },
        error: (error: any) => console.log('Error downloading the file'),
      });
  }
}
