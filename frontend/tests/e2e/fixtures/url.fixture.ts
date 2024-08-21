import { FulfillResponse } from '../types/fulfill-response';
import { test as base } from './app.fixture';
import { API_URL } from '@/src/constants/apiUrl';

interface PostsFixture {
	setShortUrlResponse(response: FulfillResponse): Promise<void>;
	setUserInfoResponse(response: FulfillResponse): Promise<void>;
}

export const test = base.extend<PostsFixture>({
	setShortUrlResponse: async ({ setRoute }, use) => {
		await use(async (response: FulfillResponse) => {
			await setRoute({
				url: new RegExp(`${API_URL}/url*`),
				response,
			});
		});
	},
	setUserInfoResponse: async ({ setRoute }, use) => {
		await use(async (response: FulfillResponse) => {
			await setRoute({
				url: new RegExp(`${API_URL}/users/describe/*`),
				response,
			});
		});
	},
});
