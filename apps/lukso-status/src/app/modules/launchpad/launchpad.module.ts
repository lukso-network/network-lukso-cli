import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LaunchpadRoutingModule } from './launchpad-routing.module';
import { LaunchpadComponent } from './launchpad/components/launchpad.component';
import { KeygenService } from './launchpad/services/keygen.service';
import { HttpClientModule } from '@angular/common/http';
import { TestComponent } from './test/test.component';

@NgModule({
  declarations: [LaunchpadComponent, TestComponent],
  providers: [KeygenService],
  imports: [CommonModule, LaunchpadRoutingModule, HttpClientModule],
})
export class LaunchpadModule {}
