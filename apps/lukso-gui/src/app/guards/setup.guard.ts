import { Inject, Injectable } from '@angular/core';
import { CanActivate, Router, UrlTree } from '@angular/router';
import { RxState } from '@rx-angular/state';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GLOBAL_RX_STATE, GlobalState } from '../shared/rx-state';

@Injectable({
  providedIn: 'root',
})
export class SetupGuard implements CanActivate {
  constructor(
    @Inject(GLOBAL_RX_STATE) private state: RxState<GlobalState>,
    private router: Router
  ) {}
  canActivate(): Observable<boolean | UrlTree> {
    return this.state.select('setupPerformed').pipe(
      map((setupPerformed) => {
        if (!setupPerformed) {
          return this.router.parseUrl('/setup');
        }
        return true;
      })
    );
  }
}
