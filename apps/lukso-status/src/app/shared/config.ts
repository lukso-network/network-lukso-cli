export const DEFAULT_UPDATE_INTERVAL = 3000;
export const VALIDATOR_DEPOSIT_COST = 32;
export const DEPOSIT_CONTRACT_ADDRESS =
  '0x000000000000000000000000000000000000cafE';

export function getNamespacePrefix() {
  const network = localStorage.getItem('network');
  let statsPrefix = '';
  if (!(network == 'l15-prod')) {
    statsPrefix = network?.split('-')[1] + '.';
  }
  return statsPrefix;
}
