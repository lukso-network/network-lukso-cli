import { KeyValue } from '@angular/common';
import {
  Component,
  ChangeDetectionStrategy,
  Input,
  OnChanges,
  SimpleChanges,
} from '@angular/core';
import { ClientVersion } from '../../../interfaces/client-versions';
import { NETWORKS } from '../../../modules/launchpad/launchpad/helpers/create-keys';

const VALIDATOR_STATUSES: { [key: string]: string } = {
  '0': 'UNKNOWN',
  '1': 'DEPOSITED',
  '2': 'PENDING',
  '3': 'ACTIVE',
  '4': 'EXITING',
  '5': 'SLASHING',
  '6': 'EXITED',
};

@Component({
  selector: 'lukso-validator-status',
  templateUrl: './validator-status.component.html',
  styleUrls: ['./validator-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ValidatorStatusComponent implements OnChanges {
  validatorData: KeyValue<string, string>[] = [];
  @Input() metrics: any = {};
  @Input() network: NETWORKS | null = NETWORKS.L15_DEV;
  @Input() version: ClientVersion | undefined = {};

  env = '';
  VALIDATOR_STATUSES = VALIDATOR_STATUSES;

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.metrics?.currentValue) {
      this.validatorData = Object.entries(this.metrics)
        .filter(([key]) => {
          return key.includes('validator_statuses');
        })
        .map(([key, value]) => {
          const regex = new RegExp(/"([A-Za-z0-9]*)/);
          const match = key.match(regex) as RegExpMatchArray;
          return { key: match[1], value } as KeyValue<string, string>;
        });
    }

    if (changes?.env?.currentValue !== null) {
      this.env = this.getEnv(this.network as NETWORKS);
    }
  }

  getValidatorStatus(pubkey: string) {
    const statusNumber: string =
      this.metrics['validator_statuses{pubkey="' + pubkey + '"}'];

    return VALIDATOR_STATUSES[statusNumber];
  }

  getValidatorMetric(key: string, pubkey: string) {
    const statusNumber: string =
      this.metrics[key + '{pubkey="' + pubkey + '"}'];

    return statusNumber;
  }

  truncate(
    text: string,
    startChars: number,
    endChars: number,
    maxLength: number
  ) {
    if (text.length > maxLength) {
      const start = text.substring(0, startChars);
      const end = text.substring(text.length - endChars, text.length);
      return start + '...' + end;
    }
    return text;
  }

  private getEnv(network: string) {
    const namespace = network?.split('-')[1];

    if (namespace === 'dev') {
      return 'dev.';
    }

    if (namespace === 'staging') {
      return 'staging.';
    }

    return '';
  }
}
