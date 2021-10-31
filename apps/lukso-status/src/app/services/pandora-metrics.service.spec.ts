import { TestBed } from '@angular/core/testing';

import { PandoraMetricsService } from './pandora-metrics.service';

describe('PandoraMetricsService', () => {
  let service: PandoraMetricsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PandoraMetricsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
