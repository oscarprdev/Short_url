import { expect } from '@playwright/test';
import { test } from './fixtures/url.fixture';
import { shortUrlResponse } from '../mocks/short-url-response';
import { userInfoResponse } from '../mocks/user-info-response';

test.describe('Add posts', () => {
	test.beforeEach(async ({ navigateToApp }) => {
		await navigateToApp();
	});

	test('Should the user be able to short an url', async ({ appPage, setShortUrlResponse }) => {
		await setShortUrlResponse({ json: shortUrlResponse });

		await expect(appPage.homeTitle).toHaveText('Short - it');
		await expect(appPage.homeDescription).toHaveText('The simplest URL Shortner you were waiting for.');

		await expect(appPage.urlForm).toBeVisible();
		await expect(appPage.urlButton).toBeVisible();
		await expect(appPage.urlInput).toBeVisible();

		await appPage.urlInput.fill('https://random_url/1234');
		await appPage.urlButton.click();

		await expect(appPage.urlCard).toBeVisible();

		const urlLink = await appPage.getUrlLink(shortUrlResponse.url.id);
		await expect(urlLink).toBeVisible();
	});

	test('Should the user be able to go to user screen', async ({ appPage, setUserInfoResponse }) => {
		await setUserInfoResponse({ json: userInfoResponse });

		await expect(appPage.loginButton).toBeVisible();

		await appPage.loginButton.click();

		await expect(appPage.loginModal).toBeVisible();
		await expect(appPage.defaultUserButton).toBeVisible();
		await expect(appPage.googleLoginButton).toBeVisible();
		await expect(appPage.githubLoginButton).toBeVisible();

		await appPage.defaultUserButton.click();

		await expect(appPage.userScreen).toBeVisible();
		await expect(appPage.userScreenText).toHaveText('Short your favourite links!');
	});
});
