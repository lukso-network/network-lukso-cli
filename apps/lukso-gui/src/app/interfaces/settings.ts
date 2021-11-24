import { ClientVersion } from './client-versions';

export interface Settings {
  hostName: string;
  coinbase: string;
  versions: ClientVersion;
}
