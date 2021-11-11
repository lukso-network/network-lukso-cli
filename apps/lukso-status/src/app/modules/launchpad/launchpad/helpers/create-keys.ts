export const enum CURRENT_KEY_ACTION {
  'IDLE' = 'Generate',
  'GENERATING' = 'Generating',
  'IMPORTING' = 'Importing',
  'COMPLETE' = 'Keys Downloaded',
}

// prettier-ignore
export interface DepositData {
  pubkey:                 string;
  withdrawal_credentials: string;
  amount:                 number;
  signature:              string;
  deposit_message_root:   string;
  deposit_data_root:      string;
  fork_version:           string;
  eth2_network_name:      string;
  deposit_cli_version:    string;
}
