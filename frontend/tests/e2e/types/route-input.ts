import { Page } from '@playwright/test';
import { FulfillResponse } from './fulfill-response';

export type RouteUrl = Parameters<Page['route']>[0];

export interface SetRouteInput {
	url: RouteUrl;
	response: FulfillResponse;
}
