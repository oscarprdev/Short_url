import { Locator, Page } from '@playwright/test';

export class AppPage {
	readonly homeDescription: Locator;
	readonly homeTitle: Locator;

	readonly urlForm: Locator;
	readonly urlInput: Locator;
	readonly urlButton: Locator;
	readonly urlCard: Locator;

	readonly loginModal: Locator;
	readonly loginButton: Locator;
	readonly defaultUserButton: Locator;
	readonly googleLoginButton: Locator;
	readonly githubLoginButton: Locator;

	readonly userScreen: Locator;
	readonly userScreenText: Locator;
	readonly userScreenNav: Locator;

	constructor(protected page: Page) {
		this.homeDescription = this.page.getByLabel('home description');
		this.homeTitle = this.page.getByLabel('home title');

		this.urlForm = this.page.getByTestId('url-form');
		this.urlInput = this.page.getByPlaceholder('Enter your link here...');
		this.urlButton = this.page.locator('button[type="submit"]');
		this.urlCard = this.page.getByRole('listitem');

		this.loginModal = this.page.getByLabel('login modal');
		this.loginButton = this.page.getByTestId('login-btn');
		this.defaultUserButton = this.page.getByRole('link').getByText('Default user');
		this.googleLoginButton = this.page.getByRole('link').getByText('Google');
		this.githubLoginButton = this.page.getByRole('link').getByText('Github');

		this.userScreen = this.page.getByRole('main');
		this.userScreenText = this.page.getByText('Short your favourite links!');
		this.userScreenNav = this.page.getByRole('navigation');
	}

	async getUrlLink(urlId: string) {
		return this.page.getByTestId(`url-link-${urlId}`);
	}

	async getLinkCounter(counterValue: number) {
		return this.page.getByText(`Total ${counterValue}`);
	}

	async getShortUrl(shortUrl: string) {
		return this.page.getByText(shortUrl);
	}
}
