import { test as base } from 'playwright/test';
import { AppPage } from '../pages/app-page';
import { SetRouteInput } from '../types/route-input';

interface AppFixture {
	appPage: AppPage;
	navigateToApp(): Promise<void>;
	setRoute(input: SetRouteInput): Promise<void>;
}

export const test = base.extend<AppFixture>({
	appPage: async ({ page }, use) => {
		await use(new AppPage(page));
	},
	navigateToApp: async ({ page }, use) => {
		await use(async (): Promise<void> => {
			await page.goto(`http://localhost:5173/`);
		});
	},
	setRoute: async ({ page }, use) => {
		await use(async ({ url, response }): Promise<void> => {
			await page.route(url, async (route) => {
				if (route.request().method() === 'POST') {
					return await route.fulfill(response);
				}

				await route.fulfill(response);
			});
		});
	},
});
