import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LaunchpadComponent } from './launchpad/components/launchpad.component';

const routes: Routes = [{ path: '', component: LaunchpadComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class LaunchpadRoutingModule {}
