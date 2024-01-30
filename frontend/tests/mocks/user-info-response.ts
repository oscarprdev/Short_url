import { GetUserInfoOutput } from '../../src/services/api/getUserInfo';
import { mockedUrl } from './url-response';

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
	urls: [mockedUrl],
};
