import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import {
  LetModule,
  PushModule,
  ViewportPrioModule,
} from '@rx-angular/template';
import { LaunchpadRoutingModule } from './launchpad-routing.module';
import { CreateKeysComponent } from './launchpad/components/create-keys/create-keys.component';
import { LaunchpadComponent } from './launchpad/components/launchpad/launchpad.component';
import { PasswordCheckerComponent } from './launchpad/components/password-checker/password-checker.component';
import { DepositTransactionComponent } from './launchpad/components/send-transactions/deposit-transaction/deposit-transaction.component';
import { SendTransactionsComponent } from './launchpad/components/send-transactions/send-transactions.component';
import { KeygenService } from './launchpad/services/keygen.service';

@NgModule({
  declarations: [
    LaunchpadComponent,
    CreateKeysComponent,
    PasswordCheckerComponent,
    SendTransactionsComponent,
    DepositTransactionComponent,
  ],
  providers: [KeygenService],
  imports: [
    CommonModule,
    LaunchpadRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    LetModule,
    PushModule,
    ViewportPrioModule,
  ],
})
export class LaunchpadModule {}
