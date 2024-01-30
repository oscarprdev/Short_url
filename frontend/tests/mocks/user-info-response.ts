import { GetUserInfoOutput } from '@/src/services/api/getUserInfo';

export const userInfoResponse: GetUserInfoOutput = {
	status: 200,
	user: {
		id: '615175183151032984302',
		email: 'email@email.com',
		picture: 'https:/image.com/ALV-UjVgV8w',
		createdAt: '2024-01-10T19:14:04.489431Z',
		updatedAt: '2024-01-10T19:14:04.489431Z',
		name: 'Tomm',
	},
	urls: [
		{
			id: 'f05a78f6-66e3-4a31-b718-6e94873b01f0',
			originalUrl: 'http:/short-it.dev/long/url',
			shortUrl: 'http:/short-it.dev/short/url',
			titleUrl: 'short_url_title',
			usage: 20,
			createdAt: '2022-01-31T12:34:56Z',
			updatedAt: '2022-01-31T12:34:56Z',
			expiresAt: '2022-01-31T23:59:59Z',
		},
	],
};
