import { Url } from '../../src/types/url';
import { shortUrlResponse } from '../mocks/short-url-response';
import { userInfoResponse } from '../mocks/user-info-response';
import { test } from './fixtures/url.fixture';
import { expect } from '@playwright/test';

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
		await expect(appPage.userScreenText).toBeVisible();
	});

	test('Should the user logged be able to short url and see it on the list', async ({
		appPage,
		setUserInfoResponse,
		setShortUrlResponse,
	}) => {
		const mockedNewUrl: Url = {
			...shortUrlResponse.url,
			originalUrl: 'https://random_url/4321',
			shortUrl: 'https://random_url/short-mocked-url',
			id: 'f05a78f6-66e3-4a31-b718-6e94873b0100',
		};

		await setUserInfoResponse({ json: userInfoResponse });
		await setShortUrlResponse({
			json: {
				status: shortUrlResponse.status,
				url: mockedNewUrl,
			},
		});

		await appPage.loginButton.click();

		await expect(appPage.loginModal).toBeVisible();
		await appPage.defaultUserButton.click();

		await expect(appPage.userScreenNav).toBeVisible();

		const counter = await appPage.getLinkCounter(userInfoResponse.urls.length);
		await expect(counter).toBeVisible();

		await expect(appPage.urlButton).toBeVisible();
		await expect(appPage.urlInput).toBeVisible();

		const urlsUpdated = userInfoResponse.urls.concat(mockedNewUrl);
		await setUserInfoResponse({
			json: {
				status: userInfoResponse.status,
				user: userInfoResponse.user,
				urls: urlsUpdated,
			},
		});

		await appPage.urlInput.fill(mockedNewUrl.originalUrl);
		await appPage.urlButton.click();

		const counterUpdated = await appPage.getLinkCounter(urlsUpdated.length);
		await expect(counterUpdated).toBeVisible();

		for (const url of urlsUpdated) {
			const urlCard = await appPage.getShortUrl(url.shortUrl);
			await expect(urlCard).toBeVisible();
		}
	});
});
