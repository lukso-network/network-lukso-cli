import { InjectionToken } from '@angular/core';
import { RxState } from '@rx-angular/state';
import { Settings } from '../interfaces/settings';
import { NETWORKS } from '../modules/launchpad/launchpad/helpers/create-keys';

export interface GlobalState {
  network: NETWORKS;
  settings: Settings;
  setupPerformed: boolean;
}

export const GLOBAL_RX_STATE = new InjectionToken<RxState<GlobalState>>(
  'GLOBAL_RX_STATE'
);
