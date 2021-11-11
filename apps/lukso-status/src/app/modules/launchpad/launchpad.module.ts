import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LaunchpadRoutingModule } from './launchpad-routing.module';
import { LaunchpadComponent } from './launchpad/components/launchpad/launchpad.component';
import { KeygenService } from './launchpad/services/keygen.service';
import { HttpClientModule } from '@angular/common/http';
import { TestComponent } from './test/test.component';
import { CreateKeysComponent } from './launchpad/components/create-keys/create-keys.component';
import { ReactiveFormsModule } from '@angular/forms';
import { PasswordCheckerComponent } from './launchpad/components/password-checker/password-checker.component';
import { SendTransactionsComponent } from './launchpad/components/send-transactions/send-transactions.component';

@NgModule({
  declarations: [
    LaunchpadComponent,
    TestComponent,
    CreateKeysComponent,
    PasswordCheckerComponent,
    SendTransactionsComponent,
  ],
  providers: [KeygenService],
  imports: [
    CommonModule,
    LaunchpadRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
  ],
})
export class LaunchpadModule {}
