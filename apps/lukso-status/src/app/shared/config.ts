import { NETWORKS } from '../modules/launchpad/launchpad/helpers/create-keys';

export const DEFAULT_NETWORK = NETWORKS.L15_DEV;
export const DEFAULT_UPDATE_INTERVAL = 5000;
export const VALIDATOR_DEPOSIT_COST = 32;
export const DEPOSIT_CONTRACT_ADDRESS =
  '0x000000000000000000000000000000000000cafE';

export function getNamespacePrefix(network: string) {
  let statsPrefix = '';
  if (!(network == 'l15-prod')) {
    statsPrefix = network?.split('-')[1] + '.';
  }
  return statsPrefix;
}
