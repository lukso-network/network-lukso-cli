import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LaunchpadComponent } from './launchpad/components/launchpad.component';
import { TestComponent } from './test/test.component';

const routes: Routes = [
  {
    path: 'keys',
    component: LaunchpadComponent,
  },
  {
    path: 'test',
    component: TestComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class LaunchpadRoutingModule {}
