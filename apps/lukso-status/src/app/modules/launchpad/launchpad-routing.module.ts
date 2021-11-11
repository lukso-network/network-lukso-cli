import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateKeysComponent } from './launchpad/components/create-keys/create-keys.component';
import { LaunchpadComponent } from './launchpad/components/launchpad/launchpad.component';
import { SendTransactionsComponent } from './launchpad/components/send-transactions/send-transactions.component';

const routes: Routes = [
  {
    path: '',
    component: LaunchpadComponent,
  },
  {
    path: 'keys',
    component: CreateKeysComponent,
  },
  {
    path: 'transactions',
    component: SendTransactionsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class LaunchpadRoutingModule {}
