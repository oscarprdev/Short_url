import { ShortUrlOutput } from '../../src/services/api/shortUrl';
import { mockedUrl } from './url-response';

export const shortUrlResponse: ShortUrlOutput = {
	status: 200,
	url: mockedUrl,
};
