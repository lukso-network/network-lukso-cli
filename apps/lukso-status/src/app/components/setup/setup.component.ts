import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { RxState } from '@rx-angular/state';
import { Subject } from 'rxjs';
import { delay, tap } from 'rxjs/operators';
import { NETWORKS } from '../../modules/launchpad/launchpad/helpers/create-keys';

interface SetupState {
  inProgress: boolean;
}

@Component({
  selector: 'lukso-setup',
  templateUrl: './setup.component.html',
  styleUrls: ['./setup.component.scss'],
})
export class SetupComponent extends RxState<SetupState> implements OnInit {
  installBtn$ = new Subject();
  readonly inProgress$ = this.select('inProgress');
  readonly initialState: SetupState = { inProgress: false };
  constructor(private http: HttpClient) {
    super();
    this.set(this.initialState);
    this.connect(this.installBtn$, () => ({ inProgress: true }));
    this.hold(this.installBtn$, () =>
      this.http
        .post('/api/initial-setup', {
          network: NETWORKS.L15_DEV,
        })
        .subscribe(() => {
          this.set({ inProgress: false });
        })
    );
  }

  ngOnInit() {
    this.installBtn$.subscribe((result) => {
      console.log(result);
    });
  }
}
