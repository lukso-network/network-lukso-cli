import { TestBed } from '@angular/core/testing';

import { VanguardServiceService } from './vanguard-metrics.service';

describe('VanguardServiceService', () => {
  let service: VanguardServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(VanguardServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
