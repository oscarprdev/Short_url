import { FulfillResponse } from './fulfill-response';
import { Page } from '@playwright/test';

export type RouteUrl = Parameters<Page['route']>[0];

export interface SetRouteInput {
	url: RouteUrl;
	response: FulfillResponse;
}
