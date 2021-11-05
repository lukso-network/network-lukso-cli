import { TestBed } from '@angular/core/testing';

import { ValidatorMetricsService } from './validator-metrics.service';

describe('ValidatorMetricsService', () => {
  let service: ValidatorMetricsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ValidatorMetricsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
